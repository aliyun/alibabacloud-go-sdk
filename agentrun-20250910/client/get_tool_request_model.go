// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetToolRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetToolRequest struct {
}

func (s GetToolRequest) String() string {
	return dara.Prettify(s)
}

func (s GetToolRequest) GoString() string {
	return s.String()
}

func (s *GetToolRequest) Validate() error {
	return dara.Validate(s)
}
