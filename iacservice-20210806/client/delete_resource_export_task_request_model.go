// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceExportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteResourceExportTaskRequest struct {
}

func (s DeleteResourceExportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceExportTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceExportTaskRequest) Validate() error {
	return dara.Validate(s)
}
