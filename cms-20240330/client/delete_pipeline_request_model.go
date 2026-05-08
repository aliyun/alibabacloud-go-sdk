// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeletePipelineRequest struct {
}

func (s DeletePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelineRequest) Validate() error {
	return dara.Validate(s)
}
