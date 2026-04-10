// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAddonReleaseRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAddonReleaseRequest struct {
}

func (s GetAddonReleaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAddonReleaseRequest) GoString() string {
	return s.String()
}

func (s *GetAddonReleaseRequest) Validate() error {
	return dara.Validate(s)
}
