// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountMFAInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsMFAEnable(v bool) *GetAccountMFAInfoResponseBody
	GetIsMFAEnable() *bool
	SetRequestId(v string) *GetAccountMFAInfoResponseBody
	GetRequestId() *string
}

type GetAccountMFAInfoResponseBody struct {
	// Indicates whether MFA devices are enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsMFAEnable *bool `json:"IsMFAEnable,omitempty" xml:"IsMFAEnable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4BE83135-0B08-467C-B3A2-27B312FD0F57
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAccountMFAInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAccountMFAInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAccountMFAInfoResponseBody) GetIsMFAEnable() *bool {
	return s.IsMFAEnable
}

func (s *GetAccountMFAInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAccountMFAInfoResponseBody) SetIsMFAEnable(v bool) *GetAccountMFAInfoResponseBody {
	s.IsMFAEnable = &v
	return s
}

func (s *GetAccountMFAInfoResponseBody) SetRequestId(v string) *GetAccountMFAInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAccountMFAInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
