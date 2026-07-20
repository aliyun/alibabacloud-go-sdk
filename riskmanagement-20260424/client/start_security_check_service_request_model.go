// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSecurityCheckServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StartSecurityCheckServiceRequest struct {
}

func (s StartSecurityCheckServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSecurityCheckServiceRequest) GoString() string {
	return s.String()
}

func (s *StartSecurityCheckServiceRequest) Validate() error {
	return dara.Validate(s)
}
