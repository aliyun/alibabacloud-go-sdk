// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelGroupRequest struct {
}

func (s GetModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelGroupRequest) GoString() string {
	return s.String()
}

func (s *GetModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
