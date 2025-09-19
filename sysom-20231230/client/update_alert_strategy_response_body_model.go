// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAlertStrategyResponseBody
	GetCode() *string
	SetData(v interface{}) *UpdateAlertStrategyResponseBody
	GetData() interface{}
	SetMessage(v string) *UpdateAlertStrategyResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAlertStrategyResponseBody
	GetRequestId() *string
}

type UpdateAlertStrategyResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
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
	//             "enabled": True
	//
	// }
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Message *string     `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateAlertStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertStrategyResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAlertStrategyResponseBody) GetData() interface{} {
	return s.Data
}

func (s *UpdateAlertStrategyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAlertStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertStrategyResponseBody) SetCode(v string) *UpdateAlertStrategyResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAlertStrategyResponseBody) SetData(v interface{}) *UpdateAlertStrategyResponseBody {
	s.Data = v
	return s
}

func (s *UpdateAlertStrategyResponseBody) SetMessage(v string) *UpdateAlertStrategyResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAlertStrategyResponseBody) SetRequestId(v string) *UpdateAlertStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
