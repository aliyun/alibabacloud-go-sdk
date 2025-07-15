// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfficeSiteSsoStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOfficeSiteSsoStatusResponseBody
	GetRequestId() *string
	SetSsoStatus(v bool) *GetOfficeSiteSsoStatusResponseBody
	GetSsoStatus() *bool
}

type GetOfficeSiteSsoStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether SSO is enabled.
	//
	// example:
	//
	// false
	SsoStatus *bool `json:"SsoStatus,omitempty" xml:"SsoStatus,omitempty"`
}

func (s GetOfficeSiteSsoStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOfficeSiteSsoStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOfficeSiteSsoStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOfficeSiteSsoStatusResponseBody) GetSsoStatus() *bool {
	return s.SsoStatus
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetRequestId(v string) *GetOfficeSiteSsoStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponseBody) SetSsoStatus(v bool) *GetOfficeSiteSsoStatusResponseBody {
	s.SsoStatus = &v
	return s
}

func (s *GetOfficeSiteSsoStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
