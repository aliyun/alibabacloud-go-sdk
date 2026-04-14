// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcBindingRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteVpcBindingRequest struct {
}

func (s DeleteVpcBindingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcBindingRequest) GoString() string {
	return s.String()
}

func (s *DeleteVpcBindingRequest) Validate() error {
	return dara.Validate(s)
}
