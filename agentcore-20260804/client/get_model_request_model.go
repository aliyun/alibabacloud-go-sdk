// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetModelRequest struct {
}

func (s GetModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelRequest) GoString() string {
	return s.String()
}

func (s *GetModelRequest) Validate() error {
	return dara.Validate(s)
}
