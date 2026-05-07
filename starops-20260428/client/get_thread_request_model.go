// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetThreadRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetThreadRequest struct {
}

func (s GetThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s GetThreadRequest) GoString() string {
	return s.String()
}

func (s *GetThreadRequest) Validate() error {
	return dara.Validate(s)
}
