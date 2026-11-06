// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGlobalCloudAccountSystemRolePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGlobalCloudAccountSystemRolePolicyRequest struct {
}

func (s GetGlobalCloudAccountSystemRolePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGlobalCloudAccountSystemRolePolicyRequest) GoString() string {
	return s.String()
}

func (s *GetGlobalCloudAccountSystemRolePolicyRequest) Validate() error {
	return dara.Validate(s)
}
