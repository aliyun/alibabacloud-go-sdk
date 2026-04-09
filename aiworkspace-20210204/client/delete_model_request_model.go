// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteModelRequest struct {
}

func (s DeleteModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) Validate() error {
	return dara.Validate(s)
}
