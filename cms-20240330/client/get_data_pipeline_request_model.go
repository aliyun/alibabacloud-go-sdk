// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDataPipelineRequest struct {
}

func (s GetDataPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataPipelineRequest) GoString() string {
	return s.String()
}

func (s *GetDataPipelineRequest) Validate() error {
	return dara.Validate(s)
}
