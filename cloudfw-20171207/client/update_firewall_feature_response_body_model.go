// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFirewallFeatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFirewallFeatureResponseBody
	GetRequestId() *string
}

type UpdateFirewallFeatureResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFirewallFeatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFirewallFeatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFirewallFeatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFirewallFeatureResponseBody) SetRequestId(v string) *UpdateFirewallFeatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFirewallFeatureResponseBody) Validate() error {
	return dara.Validate(s)
}
