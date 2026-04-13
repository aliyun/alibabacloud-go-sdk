// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteChatRequest struct {
}

func (s DeleteChatRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatRequest) GoString() string {
	return s.String()
}

func (s *DeleteChatRequest) Validate() error {
	return dara.Validate(s)
}
