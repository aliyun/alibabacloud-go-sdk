// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAlertStrategyResponseBody
	GetCode() *string
	SetData(v interface{}) *CreateAlertStrategyResponseBody
	GetData() interface{}
	SetMessage(v string) *CreateAlertStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAlertStrategyResponseBody
	GetRequestId() *string
}

type CreateAlertStrategyResponseBody struct {
	// The status code.
	//
	// - If `code == Success`, the request is successful.
	//
	// - Other status codes indicate a failed request. If the request fails, check the `message` field for detailed error information.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {
	//
	//             "uid": "uid-1",
	//
	//             "name": "test-name",
	//
	//             "strategy": {
	//
	//                 "clusters": ["test-cluster-1","test-cluster-2"],
	//
	//                 "items": ["test-item1","test-item2"]
	//
	//             },
	//
	//             "enabled": True,
	//
	//
	//
	//         }
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The error message.
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the error information of the request.
	//
	// example:
	//
	// 创建策略失败
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAlertStrategyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *CreateAlertStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAlertStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertStrategyResponseBody) SetCode(v string) *CreateAlertStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAlertStrategyResponseBody) SetData(v interface{}) *CreateAlertStrategyResponseBody {
	s.Data = v
	return s
}

func (s *CreateAlertStrategyResponseBody) SetMessage(v string) *CreateAlertStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlertStrategyResponseBody) SetRequestId(v string) *CreateAlertStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
