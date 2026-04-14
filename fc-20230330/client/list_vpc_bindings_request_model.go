// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVpcBindingsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListVpcBindingsRequest struct {
}

func (s ListVpcBindingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVpcBindingsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcBindingsRequest) Validate() error {
	return dara.Validate(s)
}
