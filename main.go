package main

import (
	"github.com/go-playground/webhooks/v6/gitlab"
	"net/http"
)

const (
	path = "/webhooks"
)

func main() {

	println("开始监听 Gitlab Hooks消息")

	hook, _ := gitlab.New(gitlab.Options.Secret("docs-dev-platform"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		println("收到 Gitlab Hooks消息")

		payload, err := hook.Parse(r, gitlab.PushEvents)
		if err != nil {
			if err == gitlab.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}

		switch payload.(type) {

		case gitlab.JobEventPayload:
			println("JobEventPayload")
			//release := payload.(gitlab.JobEventPayload)
			//// Do whatever you want from here...
			//fmt.Printf("%+v", release)
			break
		case gitlab.MergeRequest:
			println("MergeRequest")
			//pullRequest := payload.(gitlab.MergeRequest)
			//// Do whatever you want from here...
			//fmt.Printf("%+v", pullRequest)
			break
		case gitlab.PushEventPayload:
			{
				println("PushEventPayload")
				PushPayload := payload.(gitlab.PushEventPayload)

				PullEvent(PushPayload)

			}
			break
		default:
			println("命令不在里面")
		}

	})
	http.ListenAndServe(":3000", nil)
}
