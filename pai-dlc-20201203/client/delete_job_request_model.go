// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteJobRequest struct {
}

func (s DeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
