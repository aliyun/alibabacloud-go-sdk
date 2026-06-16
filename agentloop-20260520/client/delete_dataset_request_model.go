// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDatasetRequest struct {
}

func (s DeleteDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetRequest) Validate() error {
	return dara.Validate(s)
}
