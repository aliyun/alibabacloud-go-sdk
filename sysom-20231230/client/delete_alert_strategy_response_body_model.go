// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAlertStrategyResponseBody
	GetCode() *string
	SetData(v interface{}) *DeleteAlertStrategyResponseBody
	GetData() interface{}
	SetMessage(v string) *DeleteAlertStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAlertStrategyResponseBody
	GetRequestId() *string
}

type DeleteAlertStrategyResponseBody struct {
	// Status code
	//
	// - If `code == Success`, authorization succeeded.
	//
	// - Other status codes indicate authorization failed. When authorization fails, check the `message` field for detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data.
	//
	// example:
	//
	// {
	//
	//     "uid": "1808078950770264",
	//
	//     "name": "test",
	//
	//     "strategy": {
	//
	//       "items": [
	//
	//         "节点CPU使用率检测"
	//
	//       ],
	//
	//       "clusters": [
	//
	//         "1"
	//
	//       ]
	//
	//     },
	//
	//     "enabled": false,
	//
	//     "id": 1,
	//
	//     "created_at": 1753170771,
	//
	//     "updated_at": 1753170811
	//
	//   }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlertStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAlertStrategyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteAlertStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAlertStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertStrategyResponseBody) SetCode(v string) *DeleteAlertStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAlertStrategyResponseBody) SetData(v interface{}) *DeleteAlertStrategyResponseBody {
	s.Data = v
	return s
}

func (s *DeleteAlertStrategyResponseBody) SetMessage(v string) *DeleteAlertStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlertStrategyResponseBody) SetRequestId(v string) *DeleteAlertStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
