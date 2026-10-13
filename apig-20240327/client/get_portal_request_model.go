// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPortalRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPortalRequest struct {
}

func (s GetPortalRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPortalRequest) GoString() string {
	return s.String()
}

func (s *GetPortalRequest) Validate() error {
	return dara.Validate(s)
}
