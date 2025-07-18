// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNacUserCertStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNacUserCertStatusResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateNacUserCertStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNacUserCertStatusResponseBody
	GetRequestId() *string
}

type UpdateNacUserCertStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNacUserCertStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNacUserCertStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNacUserCertStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNacUserCertStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNacUserCertStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNacUserCertStatusResponseBody) SetCode(v string) *UpdateNacUserCertStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNacUserCertStatusResponseBody) SetMessage(v string) *UpdateNacUserCertStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNacUserCertStatusResponseBody) SetRequestId(v string) *UpdateNacUserCertStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNacUserCertStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
