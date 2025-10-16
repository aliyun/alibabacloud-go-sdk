// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatFirewallPreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreCheckId(v string) *CreateNatFirewallPreCheckResponseBody
	GetPreCheckId() *string
	SetRequestId(v string) *CreateNatFirewallPreCheckResponseBody
	GetRequestId() *string
}

type CreateNatFirewallPreCheckResponseBody struct {
	// example:
	//
	// 2122
	PreCheckId *string `json:"PreCheckId,omitempty" xml:"PreCheckId,omitempty"`
	// example:
	//
	// A426611F-04FA-5205-8D04-4F6DCF09****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatFirewallPreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatFirewallPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatFirewallPreCheckResponseBody) GetPreCheckId() *string {
	return s.PreCheckId
}

func (s *CreateNatFirewallPreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatFirewallPreCheckResponseBody) SetPreCheckId(v string) *CreateNatFirewallPreCheckResponseBody {
	s.PreCheckId = &v
	return s
}

func (s *CreateNatFirewallPreCheckResponseBody) SetRequestId(v string) *CreateNatFirewallPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatFirewallPreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
