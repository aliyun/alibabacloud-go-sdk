// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDatasetJobRequest struct {
}

func (s DeleteDatasetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetJobRequest) Validate() error {
	return dara.Validate(s)
}
