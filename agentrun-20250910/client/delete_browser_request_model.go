// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBrowserRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteBrowserRequest struct {
}

func (s DeleteBrowserRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteBrowserRequest) GoString() string {
	return s.String()
}

func (s *DeleteBrowserRequest) Validate() error {
	return dara.Validate(s)
}
