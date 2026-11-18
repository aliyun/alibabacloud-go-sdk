// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkflowGlobalEnvRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetWorkflowGlobalEnvRequest struct {
}

func (s GetWorkflowGlobalEnvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkflowGlobalEnvRequest) GoString() string {
	return s.String()
}

func (s *GetWorkflowGlobalEnvRequest) Validate() error {
	return dara.Validate(s)
}
