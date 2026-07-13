// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtectionPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetProtectionPolicyRequest struct {
}

func (s GetProtectionPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProtectionPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetProtectionPolicyRequest) Validate() error {
	return dara.Validate(s)
}
