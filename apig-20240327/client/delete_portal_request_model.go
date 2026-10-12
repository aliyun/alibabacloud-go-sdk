// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortalRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePortalRequest struct {
}

func (s DeletePortalRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePortalRequest) GoString() string {
	return s.String()
}

func (s *DeletePortalRequest) Validate() error {
	return dara.Validate(s)
}
