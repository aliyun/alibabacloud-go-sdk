// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetBatchExportTaskRequest struct {
}

func (s GetBatchExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetBatchExportTaskRequest) GoString() string {
	return s.String()
}

func (s *GetBatchExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
