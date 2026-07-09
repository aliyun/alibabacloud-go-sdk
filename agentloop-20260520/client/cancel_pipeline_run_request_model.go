// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPipelineRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CancelPipelineRunRequest struct {
}

func (s CancelPipelineRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPipelineRunRequest) GoString() string {
	return s.String()
}

func (s *CancelPipelineRunRequest) Validate() error {
	return dara.Validate(s)
}
