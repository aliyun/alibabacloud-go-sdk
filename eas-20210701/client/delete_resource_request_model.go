// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteResourceRequest struct {
}

func (s DeleteResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceRequest) Validate() error {
	return dara.Validate(s)
}
