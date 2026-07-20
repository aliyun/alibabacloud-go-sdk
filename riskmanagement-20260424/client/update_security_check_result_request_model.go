// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityCheckResultRequest interface {
	dara.Model
	String() string
	GoString() string
}

type UpdateSecurityCheckResultRequest struct {
}

func (s UpdateSecurityCheckResultRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityCheckResultRequest) GoString() string {
	return s.String()
}

func (s *UpdateSecurityCheckResultRequest) Validate() error {
	return dara.Validate(s)
}
