// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIMBotNoContent interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteIMBotNoContent struct {
}

func (s DeleteIMBotNoContent) String() string {
	return dara.Prettify(s)
}

func (s DeleteIMBotNoContent) GoString() string {
	return s.String()
}

func (s *DeleteIMBotNoContent) Validate() error {
	return dara.Validate(s)
}
