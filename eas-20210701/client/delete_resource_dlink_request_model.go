// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceDLinkRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteResourceDLinkRequest struct {
}

func (s DeleteResourceDLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceDLinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkRequest) Validate() error {
	return dara.Validate(s)
}
