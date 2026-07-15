// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAiModelProviderRequest struct {
}

func (s GetAiModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAiModelProviderRequest) GoString() string {
	return s.String()
}

func (s *GetAiModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
