// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFirewallFeatureRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetFirewallFeatureRequest struct {
}

func (s GetFirewallFeatureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFirewallFeatureRequest) GoString() string {
	return s.String()
}

func (s *GetFirewallFeatureRequest) Validate() error {
	return dara.Validate(s)
}
