// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDataPipelineRequest struct {
}

func (s DeleteDataPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataPipelineRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataPipelineRequest) Validate() error {
	return dara.Validate(s)
}
