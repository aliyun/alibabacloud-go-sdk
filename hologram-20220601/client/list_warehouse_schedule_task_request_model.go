// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseScheduleTaskRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListWarehouseScheduleTaskRequest struct {
}

func (s ListWarehouseScheduleTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseScheduleTaskRequest) GoString() string {
	return s.String()
}

func (s *ListWarehouseScheduleTaskRequest) Validate() error {
	return dara.Validate(s)
}
