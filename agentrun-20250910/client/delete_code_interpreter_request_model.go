// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCodeInterpreterRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCodeInterpreterRequest struct {
}

func (s DeleteCodeInterpreterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCodeInterpreterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCodeInterpreterRequest) Validate() error {
	return dara.Validate(s)
}
