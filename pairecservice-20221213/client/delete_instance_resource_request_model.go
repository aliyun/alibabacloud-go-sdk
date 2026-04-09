// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteInstanceResourceRequest struct {
}

func (s DeleteInstanceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResourceRequest) Validate() error {
	return dara.Validate(s)
}
