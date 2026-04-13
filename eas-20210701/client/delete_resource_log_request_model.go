// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceLogRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteResourceLogRequest struct {
}

func (s DeleteResourceLogRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceLogRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogRequest) Validate() error {
	return dara.Validate(s)
}
