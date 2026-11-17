// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomDomainRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCustomDomainRequest struct {
}

func (s DeleteCustomDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomDomainRequest) Validate() error {
	return dara.Validate(s)
}
