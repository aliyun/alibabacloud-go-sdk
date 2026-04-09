// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteToolRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteToolRequest struct {
}

func (s DeleteToolRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteToolRequest) GoString() string {
	return s.String()
}

func (s *DeleteToolRequest) Validate() error {
	return dara.Validate(s)
}
