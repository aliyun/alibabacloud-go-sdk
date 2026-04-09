// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDatasetVersionRequest struct {
}

func (s DeleteDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
