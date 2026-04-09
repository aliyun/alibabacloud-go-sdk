// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteModelVersionRequest struct {
}

func (s DeleteModelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelVersionRequest) Validate() error {
	return dara.Validate(s)
}
