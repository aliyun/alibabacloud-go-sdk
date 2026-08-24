// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetChatModelRequest struct {
}

func (s GetChatModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatModelRequest) GoString() string {
	return s.String()
}

func (s *GetChatModelRequest) Validate() error {
	return dara.Validate(s)
}
