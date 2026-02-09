// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartGatewayResponseBody
	GetCode() *string
	SetMessage(v string) *RestartGatewayResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartGatewayResponseBody
	GetRequestId() *string
}

type RestartGatewayResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RestartGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *RestartGatewayResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartGatewayResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartGatewayResponseBody) SetCode(v string) *RestartGatewayResponseBody {
	s.Code = &v
	return s
}

func (s *RestartGatewayResponseBody) SetMessage(v string) *RestartGatewayResponseBody {
	s.Message = &v
	return s
}

func (s *RestartGatewayResponseBody) SetRequestId(v string) *RestartGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
