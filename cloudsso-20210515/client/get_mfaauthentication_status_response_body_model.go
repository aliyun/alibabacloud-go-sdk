// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMFAAuthenticationStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMFAAuthenticationStatus(v string) *GetMFAAuthenticationStatusResponseBody
	GetMFAAuthenticationStatus() *string
	SetRequestId(v string) *GetMFAAuthenticationStatusResponseBody
	GetRequestId() *string
}

type GetMFAAuthenticationStatusResponseBody struct {
	// Indicates whether MFA is enabled for users. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	MFAAuthenticationStatus *string `json:"MFAAuthenticationStatus,omitempty" xml:"MFAAuthenticationStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5E688346-DF1A-5537-9BFC-8A9974D29586
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMFAAuthenticationStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMFAAuthenticationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetMFAAuthenticationStatusResponseBody) GetMFAAuthenticationStatus() *string {
	return s.MFAAuthenticationStatus
}

func (s *GetMFAAuthenticationStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMFAAuthenticationStatusResponseBody) SetMFAAuthenticationStatus(v string) *GetMFAAuthenticationStatusResponseBody {
	s.MFAAuthenticationStatus = &v
	return s
}

func (s *GetMFAAuthenticationStatusResponseBody) SetRequestId(v string) *GetMFAAuthenticationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMFAAuthenticationStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
