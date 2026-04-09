// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBrowserRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetBrowserRequest struct {
}

func (s GetBrowserRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBrowserRequest) GoString() string {
	return s.String()
}

func (s *GetBrowserRequest) Validate() error {
	return dara.Validate(s)
}
