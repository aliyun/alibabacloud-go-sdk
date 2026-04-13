// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteVirtualResourceRequest struct {
}

func (s DeleteVirtualResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVirtualResourceRequest) Validate() error {
	return dara.Validate(s)
}
