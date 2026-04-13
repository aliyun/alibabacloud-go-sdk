// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTerraformStateDetectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetTerraformStateDetectionRequest struct {
}

func (s GetTerraformStateDetectionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTerraformStateDetectionRequest) GoString() string {
	return s.String()
}

func (s *GetTerraformStateDetectionRequest) Validate() error {
	return dara.Validate(s)
}
