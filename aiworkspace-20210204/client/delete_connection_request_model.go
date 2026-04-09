// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteConnectionRequest struct {
}

func (s DeleteConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) Validate() error {
	return dara.Validate(s)
}
