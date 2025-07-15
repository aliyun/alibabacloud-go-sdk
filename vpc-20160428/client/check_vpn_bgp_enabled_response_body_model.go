// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVpnBgpEnabledResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBgpEnabled(v bool) *CheckVpnBgpEnabledResponseBody
	GetBgpEnabled() *bool
	SetRequestId(v string) *CheckVpnBgpEnabledResponseBody
	GetRequestId() *string
}

type CheckVpnBgpEnabledResponseBody struct {
	// Indicates whether the region supports BGP.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	BgpEnabled *bool `json:"BgpEnabled,omitempty" xml:"BgpEnabled,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DF448D2E-215D-334F-B10D-7C3B087EF996
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckVpnBgpEnabledResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckVpnBgpEnabledResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVpnBgpEnabledResponseBody) GetBgpEnabled() *bool {
	return s.BgpEnabled
}

func (s *CheckVpnBgpEnabledResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckVpnBgpEnabledResponseBody) SetBgpEnabled(v bool) *CheckVpnBgpEnabledResponseBody {
	s.BgpEnabled = &v
	return s
}

func (s *CheckVpnBgpEnabledResponseBody) SetRequestId(v string) *CheckVpnBgpEnabledResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVpnBgpEnabledResponseBody) Validate() error {
	return dara.Validate(s)
}
