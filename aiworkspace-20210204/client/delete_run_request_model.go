// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRunRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRunRequest struct {
}

func (s DeleteRunRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRunRequest) GoString() string {
	return s.String()
}

func (s *DeleteRunRequest) Validate() error {
	return dara.Validate(s)
}
