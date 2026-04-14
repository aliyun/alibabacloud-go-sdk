// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetStackRequest struct {
}

func (s GetStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackRequest) GoString() string {
	return s.String()
}

func (s *GetStackRequest) Validate() error {
	return dara.Validate(s)
}
