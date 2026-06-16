// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPipelineRequest struct {
}

func (s GetPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRequest) Validate() error {
	return dara.Validate(s)
}
