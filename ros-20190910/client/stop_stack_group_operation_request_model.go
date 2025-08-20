// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopStackGroupOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *StopStackGroupOperationRequest
	GetOperationId() *string
	SetRegionId(v string) *StopStackGroupOperationRequest
	GetRegionId() *string
}

type StopStackGroupOperationRequest struct {
	// The ID of the operation.
	//
	// You can call the [ListStackGroupOperations](https://help.aliyun.com/document_detail/151342.html) operation to obtain the operation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopStackGroupOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s StopStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *StopStackGroupOperationRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *StopStackGroupOperationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopStackGroupOperationRequest) SetOperationId(v string) *StopStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

func (s *StopStackGroupOperationRequest) SetRegionId(v string) *StopStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

func (s *StopStackGroupOperationRequest) Validate() error {
	return dara.Validate(s)
}
