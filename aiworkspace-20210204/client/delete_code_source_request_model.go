// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCodeSourceRequest struct {
}

func (s DeleteCodeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceRequest) Validate() error {
	return dara.Validate(s)
}
