// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPhoneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhone(v string) *GetUserPhoneResponseBody
	GetPhone() *string
	SetRequestId(v string) *GetUserPhoneResponseBody
	GetRequestId() *string
}

type GetUserPhoneResponseBody struct {
	// example:
	//
	// 18612345678
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CEADB586-51CB-1B6B-95BD-AB85A7A08E97
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserPhoneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPhoneResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPhoneResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *GetUserPhoneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPhoneResponseBody) SetPhone(v string) *GetUserPhoneResponseBody {
	s.Phone = &v
	return s
}

func (s *GetUserPhoneResponseBody) SetRequestId(v string) *GetUserPhoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPhoneResponseBody) Validate() error {
	return dara.Validate(s)
}
