// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceCronScalerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteServiceCronScalerRequest struct {
}

func (s DeleteServiceCronScalerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerRequest) Validate() error {
	return dara.Validate(s)
}
