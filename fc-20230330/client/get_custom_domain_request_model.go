// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCustomDomainRequest struct {
}

func (s GetCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *GetCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
