// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPluginPipelineRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RunPluginPipelineRequest struct {
}

func (s RunPluginPipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPluginPipelineRequest) GoString() string {
	return s.String()
}

func (s *RunPluginPipelineRequest) Validate() error {
	return dara.Validate(s)
}
