// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainEndpoints interface {
	dara.Model
	String() string
	GoString() string
}

type DomainEndpoints struct {
}

func (s DomainEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DomainEndpoints) GoString() string {
	return s.String()
}

func (s *DomainEndpoints) Validate() error {
	return dara.Validate(s)
}
