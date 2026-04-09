// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRunLabelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRunLabelRequest struct {
}

func (s DeleteRunLabelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRunLabelRequest) GoString() string {
	return s.String()
}

func (s *DeleteRunLabelRequest) Validate() error {
	return dara.Validate(s)
}
