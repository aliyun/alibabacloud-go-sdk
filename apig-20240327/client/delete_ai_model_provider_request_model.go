// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAiModelProviderRequest struct {
}

func (s DeleteAiModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
