// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMBotRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteIMBotRequest struct {
}

func (s DeleteIMBotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMBotRequest) GoString() string {
	return s.String()
}

func (s *DeleteIMBotRequest) Validate() error {
	return dara.Validate(s)
}
