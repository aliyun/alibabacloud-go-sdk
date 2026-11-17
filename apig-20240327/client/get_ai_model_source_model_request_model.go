// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelSourceModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAiModelSourceModelRequest struct {
}

func (s GetAiModelSourceModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelSourceModelRequest) GoString() string {
	return s.String()
}

func (s *GetAiModelSourceModelRequest) Validate() error {
	return dara.Validate(s)
}
