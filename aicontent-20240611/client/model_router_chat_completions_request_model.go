// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterChatCompletionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *ModelRouterChatCompletionsRequest
	GetBody() interface{}
}

type ModelRouterChatCompletionsRequest struct {
	// example:
	//
	// {
	//
	//     "stream": true,
	//
	//     "messages": [
	//
	//       {
	//
	//         "role": "user",
	//
	//         "content": "1+1"
	//
	//       }
	//
	//     ],
	//
	//     "model_id": 15,
	//
	//     "stream_options": {
	//
	//       "include_usage": true
	//
	//     }
	//
	//   }
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelRouterChatCompletionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterChatCompletionsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterChatCompletionsRequest) GetBody() interface{} {
	return s.Body
}

func (s *ModelRouterChatCompletionsRequest) SetBody(v interface{}) *ModelRouterChatCompletionsRequest {
	s.Body = v
	return s
}

func (s *ModelRouterChatCompletionsRequest) Validate() error {
	return dara.Validate(s)
}
