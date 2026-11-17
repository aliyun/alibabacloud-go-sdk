// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceEditImpactRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAiModelSourceEditImpactRequest struct {
}

func (s GetAiModelSourceEditImpactRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceEditImpactRequest) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceEditImpactRequest) Validate() error {
	return dara.Validate(s)
}
