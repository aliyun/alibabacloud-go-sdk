// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelStockOssCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CancelStockOssCheckTaskRequest
	GetRegionId() *string
	SetTaskId(v string) *CancelStockOssCheckTaskRequest
	GetTaskId() *string
}

type CancelStockOssCheckTaskRequest struct {
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// P_UNYVB
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelStockOssCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelStockOssCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelStockOssCheckTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelStockOssCheckTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelStockOssCheckTaskRequest) SetRegionId(v string) *CancelStockOssCheckTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CancelStockOssCheckTaskRequest) SetTaskId(v string) *CancelStockOssCheckTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelStockOssCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
