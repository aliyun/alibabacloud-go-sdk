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
	// example:
	//
	// Success
	Code *string     `json:"code,omitempty" xml:"code,omitempty"`
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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
