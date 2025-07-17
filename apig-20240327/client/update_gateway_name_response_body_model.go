// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateGatewayNameResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateGatewayNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateGatewayNameResponseBody
	GetRequestId() *string
}

type UpdateGatewayNameResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateGatewayNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayNameResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGatewayNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateGatewayNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateGatewayNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGatewayNameResponseBody) SetCode(v string) *UpdateGatewayNameResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetMessage(v string) *UpdateGatewayNameResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) SetRequestId(v string) *UpdateGatewayNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGatewayNameResponseBody) Validate() error {
	return dara.Validate(s)
}
