// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterResourceTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegistrationId(v string) *RegisterResourceTypeResponseBody
	GetRegistrationId() *string
	SetRequestId(v string) *RegisterResourceTypeResponseBody
	GetRequestId() *string
}

type RegisterResourceTypeResponseBody struct {
	// The ID of the registration record. You can call the [ListResourceTypeRegistrations](https://help.aliyun.com/document_detail/2330740.html) operation to query registration records.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RegisterResourceTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RegisterResourceTypeResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeResponseBody) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *RegisterResourceTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RegisterResourceTypeResponseBody) SetRegistrationId(v string) *RegisterResourceTypeResponseBody {
	s.RegistrationId = &v
	return s
}

func (s *RegisterResourceTypeResponseBody) SetRequestId(v string) *RegisterResourceTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterResourceTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
