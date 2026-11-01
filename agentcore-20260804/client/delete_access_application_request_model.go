// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAccessApplicationRequest struct {
}

func (s DeleteAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
