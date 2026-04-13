// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGroupRequest struct {
}

func (s GetGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) Validate() error {
	return dara.Validate(s)
}
