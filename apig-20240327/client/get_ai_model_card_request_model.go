// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelCardRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAiModelCardRequest struct {
}

func (s GetAiModelCardRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelCardRequest) GoString() string {
	return s.String()
}

func (s *GetAiModelCardRequest) Validate() error {
	return dara.Validate(s)
}
