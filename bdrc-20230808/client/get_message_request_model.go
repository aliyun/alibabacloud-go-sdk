// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMessageRequest struct {
}

func (s GetMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMessageRequest) GoString() string {
	return s.String()
}

func (s *GetMessageRequest) Validate() error {
	return dara.Validate(s)
}
