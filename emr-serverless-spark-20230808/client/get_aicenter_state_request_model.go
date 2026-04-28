// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICenterStateRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAICenterStateRequest struct {
}

func (s GetAICenterStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICenterStateRequest) GoString() string {
	return s.String()
}

func (s *GetAICenterStateRequest) Validate() error {
	return dara.Validate(s)
}
