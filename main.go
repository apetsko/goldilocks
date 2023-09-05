package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"gopkg.in/yaml.v3"
)

func main() {
	//mock listen for fl0.com deploy
	port := os.Getenv("PORT")
	go func() {
		http.HandleFunc("/", getHello)
		err := http.ListenAndServe(":"+port, nil)
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	token := os.Getenv("tgtoken")
	var data map[int]int

	err := yaml.Unmarshal(mapping, &data)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(token)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s:%d] %s\n", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, remapString(update.Message.Text, data))

		bot.Send(msg)
	}

}

// remapString сhanges characters according to the data specified in the mapping.map file
func remapString(s string, m map[int]int) string {
	var sb strings.Builder
	for _, v := range s {
		if val, ok := m[int(v)]; ok {
			sb.WriteString(string(rune(val)))
		} else {
			sb.WriteString(string(v))
		}
	}
	return sb.String()
}

var mapping = []byte(`#eng start
65: 5573
66: 5108
67: 264
68: 5598
69: 7960
70: 401
71: 1292
72: 7977
73: 8153
74: 1032
75: 1050
76: 315
77: 5047
78: 7750
79: 1256
80: 1056
81: 490
82: 11364
83: 1029
84: 1196
85: 360
86: 5178
87: 11378
88: 1061
89: 7934
90: 548
97: 7937
98: 1100
99: 7428
100: 7429
101: 6513
102: 402
103: 1409
104: 7721
105: 7984
106: 1112
107: 409
108: 621
109: 7747
110: 1352
111: 1257
112: 7465
113: 1382
114: 7775
115: 1109
116: 648
117: 1405
118: 709
119: 1309
120: 1203
121: 655
122: 9761
#eng end

#rus start
1040: 256
1041: 386
1042: 385
1043: 1270
1044: 916
1045: 1212
1025: 7868
1046: 1217
1047: 540
1048: 886
1049: 1250
1050: 1180
1051: 581
1052: 1018
1053: 1186
1054: 927
1055: 1316
1056: 420
1057: 1017
1058: 882
1059: 1038
1060: 569
1061: 1276
1062: 1039
1063: 1206
1064: 412
1065: 994
1066: 1122
1067: 1192
1068: 1164
1069: 1260
1070: 984
1071: 5101
1072: 551
1073: 948
1074: 665
1075: 1271
1076: 609
1077: 1213
1105: 7869
1078: 1218
1079: 541
1080: 887
1081: 1251
1082: 954
1083: 652
1084: 653
1085: 668
1086: 959
1087: 627
1088: 961
1089: 1010
1090: 1197
1091: 611
1092: 981
1093: 1277
1094: 1119
1095: 1207
1096: 623
1097: 995
1098: 384
1099: 1193
1100: 1165
1101: 1014
1102: 985
1103: 1126
#rus end`)

func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("got / request")
	io.WriteString(w, "Hello, I'm https://t.me/zolotovoloska_bot!\n")
}
