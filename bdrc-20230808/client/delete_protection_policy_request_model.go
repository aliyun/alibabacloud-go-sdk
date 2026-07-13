// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtectionPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteProtectionPolicyRequest struct {
}

func (s DeleteProtectionPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtectionPolicyRequest) Validate() error {
	return dara.Validate(s)
}
