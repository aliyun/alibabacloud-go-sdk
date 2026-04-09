// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetApplicationProvisioningScopeRequest struct {
}

func (s GetApplicationProvisioningScopeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeRequest) Validate() error {
	return dara.Validate(s)
}
