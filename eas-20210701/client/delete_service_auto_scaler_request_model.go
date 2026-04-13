// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAutoScalerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteServiceAutoScalerRequest struct {
}

func (s DeleteServiceAutoScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerRequest) Validate() error {
	return dara.Validate(s)
}
