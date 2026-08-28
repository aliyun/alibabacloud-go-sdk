// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRiskNotificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRiskNotificationResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateRiskNotificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRiskNotificationResponseBody
	GetRequestId() *string
}

type UpdateRiskNotificationResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 01A00D7A-AA00-5BC0-9835-C7B15A3FE73A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateRiskNotificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRiskNotificationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRiskNotificationResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRiskNotificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRiskNotificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRiskNotificationResponseBody) SetCode(v string) *UpdateRiskNotificationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRiskNotificationResponseBody) SetMessage(v string) *UpdateRiskNotificationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRiskNotificationResponseBody) SetRequestId(v string) *UpdateRiskNotificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRiskNotificationResponseBody) Validate() error {
	return dara.Validate(s)
}
