// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ResumePipelineRequest struct {
}

func (s ResumePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumePipelineRequest) GoString() string {
	return s.String()
}

func (s *ResumePipelineRequest) Validate() error {
	return dara.Validate(s)
}
