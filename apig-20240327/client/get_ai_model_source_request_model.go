// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAiModelSourceRequest struct {
}

func (s GetAiModelSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceRequest) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceRequest) Validate() error {
	return dara.Validate(s)
}
