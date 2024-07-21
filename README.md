### NAME:
**tg-notify** - A new cli application for sending message to TG channel

### USAGE:
```commandline
    tg-notify [global options] command [command options]
```
    

### EXAMPLE:
You can set just 2 ENVs (NOTIFYER_TELEGRAM_BOT_TOKEN and NOTIFYER_TELEGRAM_CHAT_ID) for sending messages.
```sh

    ./tg-notify --telegram_bot_token "1233453487583" --telegram_chat_id "-12423413" \
        --text "TEST message from Bot" deploy done
```

or

```sh
    export NOTIFYER_TELEGRAM_BOT_TOKEN="1224345456546"
    export NOTIFYER_TELEGRAM_CHAT_ID="-124335435"
    ./tg-notify --text "message from BOT" deploy done
```

### REMARK:
  It's much better to specify user in the end of text:
```sh
    ./tg-notify --text "message from BOT @SPECIFY_USER_TG_NAME" deploy done
```
        

### VERSION:
v1.0.0

### COMMANDS:
```
release, r Release [ready]
deploy, d Deploy [done]
help, h Shows a list of commands or help for one command
```

### GLOBAL OPTIONS:
```
   --telegram_bot_token value, --tbt value     Telegram Bot Token [$NOTIFYER_TELEGRAM_BOT_TOKEN]
   --telegram_chat_id value, --tci value       Telegram Chat ID (default: 0) [$NOTIFYER_TELEGRAM_CHAT_ID]
   --telegram_parse_mode value, --tpm value    Telegram Parse Mode (default: "HTML") [$NOTIFYER_TELEGRAM_PARSE_MODE]
   --ci_commit_author value, --author value    CI Commit Author - GitLab CI [$CI_COMMIT_AUTHOR]
   --ci_project_url value, --cpu value         CI Project URL - GitLab CI [$CI_PROJECT_URL]
   --ci_pipeline_id value, --cpi value         CI Pipeline ID - GitLab CI [$CI_PIPELINE_ID]
   --ci_commit_ref_name value, --branch value  CI Commit Branch or Tag name - GitLab CI [$CI_COMMIT_REF_NAME]
   --text value, -t value                      Text to send [$NOTIFYER_TEXT]
   --help, -h                                  show help
   --version, -v                               print the version
```
