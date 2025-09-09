// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrationId(v string) *RegisterServiceResponseBody
	GetRegistrationId() *string
	SetRequestId(v string) *RegisterServiceResponseBody
	GetRequestId() *string
}

type RegisterServiceResponseBody struct {
	// The registration ID.
	//
	// example:
	//
	// sr-72dd5071e90c40xxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A361BA9E-2713-52C8-AFFC-C26E5180456E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterServiceResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponseBody) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *RegisterServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterServiceResponseBody) SetRegistrationId(v string) *RegisterServiceResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterServiceResponseBody) SetRequestId(v string) *RegisterServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
