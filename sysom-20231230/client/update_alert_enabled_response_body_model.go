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
