// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAiModelCardRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAiModelCardRequest struct {
}

func (s DeleteAiModelCardRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAiModelCardRequest) GoString() string {
	return s.String()
}

func (s *DeleteAiModelCardRequest) Validate() error {
	return dara.Validate(s)
}
