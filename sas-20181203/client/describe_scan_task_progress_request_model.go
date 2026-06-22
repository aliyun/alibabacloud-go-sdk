// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskProgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int64) *DescribeScanTaskProgressRequest
	GetTaskId() *int64
}

type DescribeScanTaskProgressRequest struct {
	// The ID of the virus scan task that you want to query.
	//
	// > You can call the [StartVirusScanTask](~~StartVirusScanTask~~) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 282832
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScanTaskProgressRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskProgressRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskProgressRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeScanTaskProgressRequest) SetTaskId(v int64) *DescribeScanTaskProgressRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeScanTaskProgressRequest) Validate() error {
	return dara.Validate(s)
}
