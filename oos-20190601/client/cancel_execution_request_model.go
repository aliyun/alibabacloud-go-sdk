// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *CancelExecutionRequest
	GetExecutionId() *string
	SetRegionId(v string) *CancelExecutionRequest
	GetRegionId() *string
}

type CancelExecutionRequest struct {
	// The ID of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-xxx
	ExecutionId *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionRequest) GoString() string {
	return s.String()
}

func (s *CancelExecutionRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *CancelExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelExecutionRequest) SetExecutionId(v string) *CancelExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *CancelExecutionRequest) SetRegionId(v string) *CancelExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *CancelExecutionRequest) Validate() error {
	return dara.Validate(s)
}
