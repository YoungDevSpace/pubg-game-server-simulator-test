# 🎮 PUBG Game Server Simulator  

배틀그라운드(PUBG) 게임 서버 시뮬레이터 - Kubernetes 인프라 학습용 프로젝트  

## 📖 프로젝트 소개  

실제 게임 로직 없이 **순수 인프라/DevOps** 기술만으로 대규모 멀티플레이어 게임의 백엔드 시스템을 시뮬레이션합니다.  

### 주요 기능  

- 🎯 100명 동시 플레이어 시뮬레이션  
- ⚔️ 실시간 전투 및 생존자 추적  
- 📊 페이즈별 게임 진행 (Early → Mid → Late)  
- 💻 CPU/메모리/네트워크 리소스 사용 패턴 재현  

## 🚀 실행 방법  

### 요구사항  

- Go 1.21 이상  
- 운영체제: Windows, macOS, Linux  

### 설치 및 실행  

- 저장소 클론  
git clone https://github.com/your-username/pubg-game-server-simulator-test.git  
cd pubg-game-server-simulator-test  

- 실행  
go run main.go  


## 📸 실행 결과  

━━━━━━━━━━━━━━━━━━━━━━━━━━━━  
PUBG Simulator v0.1  
━━━━━━━━━━━━━━━━━━━━━━━━━━━━  

✓ 100 players initialized  

[01s] Phase: Early Players Alive: 95 / 100  
[02s] Phase: Early Players Alive: 89 / 100  
[03s] Phase: Early Players Alive: 82 / 100  
...  
[90s] Phase: Late Players Alive: 1 / 100  

🏆 Game Over! Winner: Player_042  
⏱️ Total Time: 90 seconds  

=> 변경된 버전  

━━━━━━━━━━━━━━━━━━━━━━━━━━━━  
  🎮 PUBG Simulator v0.11  
━━━━━━━━━━━━━━━━━━━━━━━━━━━━  
  
Elapsed Time: 02:28  

┌────────────────────────────────────────────┐  
│ Players Alive:   2 / 100               │  
│ ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ │  
└────────────────────────────────────────────┘  

┌────────────────────────────────────────────┐  
│ 💻 CPU Usage:  1.6%                  │  
│ 🧠 Memory:     1060 MB                 │  
│ 🌐 Network:    0.3 MB/s               │  
└────────────────────────────────────────────┘  
  🔥 Final Circle!  
    - Player_071 (HP: 24)  

🏆 Game Over! Winner: Player_071  

<img width="406" height="367" alt="Image" src="https://github.com/user-attachments/assets/45b38354-cae6-4e2f-a764-07c92bf6918b" />  

## 🏗️ 프로젝트 구조  
pubg-game-server-simulator-test/  
├── main.go # 메인 시뮬레이터 코드  
├── go.mod # Go 모듈 설정  
├── go.sum # 의존성 체크섬  
└── README.md # 프로젝트 문서  

## 📅 개발 로드맵  

- [x] Week 1-2: 기본 시뮬레이터 개발  
- [ ] Week 3-4: Docker 컨테이너화  
- [ ] Week 5-6: Kubernetes 배포  
- [ ] Week 7-8: Agones 통합  
- [ ] Week 9-10: 멀티 리전 배포  
- [ ] Week 11-12: 관측성 및 최적화  

## 🎯 학습 목표  

1. ☸️ Kubernetes 게임 서버 오케스트레이션 (Agones)  
2. ☁️ AWS/GCP 멀티클라우드 인프라  
3. 📊 Prometheus/Grafana 관측성  
4. 🔄 CI/CD 파이프라인 (GitOps)  
5. 💰 비용 최적화 (Spot Instance)  

## 👥 팀원  

- **양소영** - 개발 / 인프라 엔지니어  

## 📚 참고 자료  

- [크래프톤 기술 블로그](https://careers.krafton.com/tech-blog)  
- [AWS GameTech](https://aws.amazon.com/gametech/)  

## 📄 라이선스  

MIT License  

## 📧 문의  

프로젝트 관련 문의: iamsoyoungdev@gmail.com  
