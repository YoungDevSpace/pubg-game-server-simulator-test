package main

import (
    "fmt"
    "math/rand"
    "time"

    "github.com/fatih/color"  // 추가
)

// 플레이어 구조체
type Player struct {
    ID       string
    Health   int
    Alive    bool
    Position [2]float64  // X, Y 좌표
}

func main() {

    // 컬러 정의
    green := color.New(color.FgGreen, color.Bold).SprintFunc()
    red := color.New(color.FgRed, color.Bold).SprintFunc()
    yellow := color.New(color.FgYellow, color.Bold).SprintFunc()
    cyan := color.New(color.FgCyan, color.Bold).SprintFunc()
    
    // fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    // fmt.Printf("  %s\n", cyan("🎮 PUBG Simulator v0.3"))
    // fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
    // fmt.Println()
    
    rand.Seed(time.Now().UnixNano())
    
    // 100명 플레이어 생성
    players := make([]*Player, 100)
    for i := 0; i < 100; i++ {
        players[i] = &Player{
            ID:       fmt.Sprintf("Player_%03d", i+1),
            Health:   100,
            Alive:    true,
            Position: [2]float64{rand.Float64() * 8000, rand.Float64() * 8000},
        }
    }
    
    fmt.Printf("%s\n", green("✓ 100 players initialized"))
    fmt.Println()
    
    startTime := time.Now()
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()
    
    for {
        <-ticker.C
        
        alive := countAlive(players)
        if alive <= 1 {
            break
        }

        simulateBattle(players)

        // 화면 클리어 (선택)
        clearScreen()

        // ━━━━━━━━━━━━━━━━━━━━━━━━━━━━
        // 대시보드 출력
        // ━━━━━━━━━━━━━━━━━━━━━━━━━━━━
        elapsed := time.Since(startTime)
        
        fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
        fmt.Printf("┃  %s  ┃\n", cyan("🎮 PUBG Game Server Simulator"))
        fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
        fmt.Println()
        
        // 게임 정보
        fmt.Printf("Elapsed Time: %s\n", yellow(fmt.Sprintf("%02d:%02d", 
            int(elapsed.Minutes()), int(elapsed.Seconds())%60)))
        fmt.Println()
        
        // 플레이어 상태
        fmt.Println("┌────────────────────────────────────────────┐")
        fmt.Printf("│ Players Alive: %s / %s               │\n", 
            green(fmt.Sprintf("%3d", alive)), 
            yellow("100"))
        fmt.Printf("│ %s │\n", makeProgressBar(alive, 100, 30))
        fmt.Println("└────────────────────────────────────────────┘")
        fmt.Println()
        
        // CPU 사용 시뮬레이션
        cpuPercent := float64(alive) * 0.8
        fmt.Println("┌────────────────────────────────────────────┐")
        fmt.Printf("│ 💻 CPU Usage:  %s%%                  │\n", 
            getCPUColor(cpuPercent))
        fmt.Printf("│ 🧠 Memory:     %s MB                 │\n", 
            green(fmt.Sprintf("%4d", 1000+(alive*30))))
        fmt.Printf("│ 🌐 Network:    %s MB/s               │\n", 
            cyan(fmt.Sprintf("%.1f", float64(alive)*0.15)))
        fmt.Println("└────────────────────────────────────────────┘")
        
        // 10명 이하일 때 상세 정보
        if alive <= 10 {
            fmt.Println(yellow("  🔥 Final Circle!"))
            for _, p := range players {
                if p.Alive {
                    hpColor := green
                    if p.Health < 50 {
                        hpColor = red
                    }
                    fmt.Printf("    - %s (HP: %s)\n", p.ID, hpColor(fmt.Sprintf("%d", p.Health)))
                }
            }
        }
    }
    
    fmt.Println()
    winner := getWinner(players)
    fmt.Printf("%s\n", green(fmt.Sprintf("🏆 Game Over! Winner: %s", winner)))
}

// main 함수 시작 부분에 추가
func clearScreen() {
    fmt.Print("\033[H\033[2J")
}

func countAlive(players []*Player) int {
    count := 0
    for _, p := range players {
        if p.Alive {
            count++
        }
    }
    return count
}

func simulateBattle(players []*Player) {
    // 살아있는 플레이어 중 랜덤하게 데미지
    for _, p := range players {
        if !p.Alive {
            continue
        }
        
        // 1% 확률로 데미지
        if rand.Float64() < 0.05 {
            damage := rand.Intn(50) + 50  // 50-100 데미지
            p.Health -= damage
            
            if p.Health <= 0 {
                p.Alive = false
            }
        }
    }
}

func getWinner(players []*Player) string {
    for _, p := range players {
        if p.Alive {
            return p.ID
        }
    }
    return "Unknown"
}

func makeProgressBar(current, total, width int) string {
    filled := int(float64(current) / float64(total) * float64(width))
    bar := ""
    
    green := color.New(color.FgGreen).SprintFunc()
    
    for i := 0; i < width; i++ {
        if i < filled {
            bar += green("█")
        } else {
            bar += "░"
        }
    }
    
    return bar
}

// 헬퍼 함수
func getCPUColor(percent float64) string {
    s := fmt.Sprintf("%.1f", percent)
    if percent < 50 {
        return color.GreenString(s)
    } else if percent < 80 {
        return color.YellowString(s)
    } else {
        return color.RedString(s)
    }
}


