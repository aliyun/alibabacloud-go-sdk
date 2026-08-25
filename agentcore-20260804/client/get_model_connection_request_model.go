// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelConnectionRequest struct {
}

func (s GetModelConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelConnectionRequest) GoString() string {
	return s.String()
}

func (s *GetModelConnectionRequest) Validate() error {
	return dara.Validate(s)
}
