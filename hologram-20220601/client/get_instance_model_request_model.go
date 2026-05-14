// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetInstanceModelRequest struct {
}

func (s GetInstanceModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceModelRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceModelRequest) Validate() error {
	return dara.Validate(s)
}
