// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetBatchImportTaskRequest struct {
}

func (s GetBatchImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchImportTaskRequest) GoString() string {
	return s.String()
}

func (s *GetBatchImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
