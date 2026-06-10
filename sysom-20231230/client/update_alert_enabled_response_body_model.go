// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAlertEnabledResponseBody
	GetCode() *string
	SetData(v interface{}) *UpdateAlertEnabledResponseBody
	GetData() interface{}
	SetMessage(v string) *UpdateAlertEnabledResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAlertEnabledResponseBody
	GetRequestId() *string
}

type UpdateAlertEnabledResponseBody struct {
	// Status code
	//
	// - If `code == Success`, authorization succeeded.
	//
	// - Other status codes indicate authorization failed. When authorization fails, view the `message` field to obtain detailed error message.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned data
	//
	// example:
	//
	// {
	//
	//     "uid": "1808078950770264",
	//
	//     "name": "test-name",
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
	//         "cluster1"
	//
	//       ]
	//
	//     },
	//
	//     "enabled": true,
	//
	//     "id": 1,
	//
	//     "created_at": 1753172727,
	//
	//     "updated_at": 1753172727
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

func (s UpdateAlertEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertEnabledResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAlertEnabledResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateAlertEnabledResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAlertEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertEnabledResponseBody) SetCode(v string) *UpdateAlertEnabledResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlertEnabledResponseBody) SetData(v interface{}) *UpdateAlertEnabledResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAlertEnabledResponseBody) SetMessage(v string) *UpdateAlertEnabledResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlertEnabledResponseBody) SetRequestId(v string) *UpdateAlertEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
