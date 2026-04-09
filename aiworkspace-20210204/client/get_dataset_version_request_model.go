// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDatasetVersionRequest struct {
}

func (s GetDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
