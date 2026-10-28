// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRunRequest struct {
}

func (s GetRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunRequest) GoString() string {
	return s.String()
}

func (s *GetRunRequest) Validate() error {
	return dara.Validate(s)
}
