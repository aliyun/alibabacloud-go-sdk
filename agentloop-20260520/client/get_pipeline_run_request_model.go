// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetPipelineRunRequest struct {
}

func (s GetPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *GetPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
