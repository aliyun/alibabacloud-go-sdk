// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListRegionsRequest struct {
}

func (s ListRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) Validate() error {
	return dara.Validate(s)
}
