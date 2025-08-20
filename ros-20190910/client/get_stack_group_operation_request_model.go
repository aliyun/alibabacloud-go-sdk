// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackGroupOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *GetStackGroupOperationRequest
	GetOperationId() *string
	SetRegionId(v string) *GetStackGroupOperationRequest
	GetRegionId() *string
}

type GetStackGroupOperationRequest struct {
	// The operation ID. You can call the [ListStackGroupOperations](https://help.aliyun.com/document_detail/151342.html) operation to query the operation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6da106ca-1784-4a6f-a7e1-e723863d****
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The region ID of the stack group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetStackGroupOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackGroupOperationRequest) GoString() string {
	return s.String()
}

func (s *GetStackGroupOperationRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *GetStackGroupOperationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackGroupOperationRequest) SetOperationId(v string) *GetStackGroupOperationRequest {
	s.OperationId = &v
	return s
}

func (s *GetStackGroupOperationRequest) SetRegionId(v string) *GetStackGroupOperationRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackGroupOperationRequest) Validate() error {
	return dara.Validate(s)
}
