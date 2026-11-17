// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicalModelDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetLogicalModelDeploymentRequest struct {
}

func (s GetLogicalModelDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogicalModelDeploymentRequest) GoString() string {
	return s.String()
}

func (s *GetLogicalModelDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
