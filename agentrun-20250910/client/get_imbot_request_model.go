// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIMBotRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetIMBotRequest struct {
}

func (s GetIMBotRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIMBotRequest) GoString() string {
	return s.String()
}

func (s *GetIMBotRequest) Validate() error {
	return dara.Validate(s)
}
