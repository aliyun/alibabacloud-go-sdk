// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVaultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateVaultResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateVaultResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVaultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateVaultResponseBody
	GetSuccess() *bool
}

type UpdateVaultResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, a success message is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateVaultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVaultResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVaultResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateVaultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVaultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVaultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateVaultResponseBody) SetCode(v string) *UpdateVaultResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVaultResponseBody) SetMessage(v string) *UpdateVaultResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVaultResponseBody) SetRequestId(v string) *UpdateVaultResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVaultResponseBody) SetSuccess(v bool) *UpdateVaultResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVaultResponseBody) Validate() error {
	return dara.Validate(s)
}
