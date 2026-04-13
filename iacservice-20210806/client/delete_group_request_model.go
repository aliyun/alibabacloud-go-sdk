// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteGroupRequest struct {
}

func (s DeleteGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) Validate() error {
	return dara.Validate(s)
}
