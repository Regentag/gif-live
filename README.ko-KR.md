# gif-live
GIF 파일을 터미널에서 [curl]() 애니메이션으로 재생합니다.
[hugomd/ascii-live](https://github.com/hugomd/ascii-live)을 참고하여 개발하였으나, 웹 프레임워크로 Echo를 사용합니다.

![데모 영상](demo.webp)
  ※ WebP 데모 영상이 재생되지 않으면 다음 링크(GIF)로 확인하세요: [https://imgur.com/hMs9WNL](https://imgur.com/hMs9WNL)

# 로컬에서 실행
로컬에서 `:1323` 포트로 서버를 실행하려면 다음과 같이 하십시오:
```bash
go run main.go
```

서버 실행 후 다른 터미널에서 다음과 같이 실행을 확인합니다:
```bash
curl http://localhost:1323/[gifname]
```

예시로서 다음 3개의 GIF 파일을 포함하고 있습니다.
 * chirno
 * reimu
 * cat

# 온라인 데모
Go 언어 개발환경이 없거나, 실행 결과만 보고 싶다면 다음 주소로 확인하세요. Heroku에서 실행 중이므로 끊김이 발생하거나 속도가 느릴 수 있습니다.
```bash
curl https://hamkong.herokuapp.com/gifanime?g=[gifname]
```

gifname은 동일하게 `chirno`, `reimu`, `cat`이 제공됩니다.

# aisimage
gif-live는 GIF 파일을 터미널에 그리기 위하여 [eliukblau/pixterm](https://github.com/eliukblau/pixterm)의 `ansimage` 모듈 코드를 수정하여 사용합니다.
`ansimage`의 원본 소스코드에 대한 권리는 Eliuk Blau에게 있습니다.
