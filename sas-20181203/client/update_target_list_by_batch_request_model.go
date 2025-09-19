// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTargetListByBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchId(v int64) *UpdateTargetListByBatchRequest
	GetBatchId() *int64
	SetOperationList(v []*UpdateTargetListByBatchRequestOperationList) *UpdateTargetListByBatchRequest
	GetOperationList() []*UpdateTargetListByBatchRequestOperationList
}

type UpdateTargetListByBatchRequest struct {
	// The ID of the release batch.
	//
	// This parameter is required.
	//
	// example:
	//
	// 52370
	BatchId *int64 `json:"BatchId,omitempty" xml:"BatchId,omitempty"`
	// The operations on assets.
	//
	// This parameter is required.
	OperationList []*UpdateTargetListByBatchRequestOperationList `json:"OperationList,omitempty" xml:"OperationList,omitempty" type:"Repeated"`
}

func (s UpdateTargetListByBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTargetListByBatchRequest) GoString() string {
	return s.String()
}

func (s *UpdateTargetListByBatchRequest) GetBatchId() *int64 {
	return s.BatchId
}

func (s *UpdateTargetListByBatchRequest) GetOperationList() []*UpdateTargetListByBatchRequestOperationList {
	return s.OperationList
}

func (s *UpdateTargetListByBatchRequest) SetBatchId(v int64) *UpdateTargetListByBatchRequest {
	s.BatchId = &v
	return s
}

func (s *UpdateTargetListByBatchRequest) SetOperationList(v []*UpdateTargetListByBatchRequestOperationList) *UpdateTargetListByBatchRequest {
	s.OperationList = v
	return s
}

func (s *UpdateTargetListByBatchRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateTargetListByBatchRequestOperationList struct {
	// The ID of the server group.
	//
	// >  You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of server groups.
	//
	// example:
	//
	// 11883086
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The operation type. Valid values:
	//
	// 	- **add**: the add operation.
	//
	// 	- **del**: the remove operation.
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The UUID of the server.
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	//
	// example:
	//
	// de393767-6fe1-4a8d-837d-927a2b******
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the VPC-connected instance.
	//
	// example:
	//
	// vpc-bp1ow0rm9t92iza******
	VpcInstanceId *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty"`
}

func (s UpdateTargetListByBatchRequestOperationList) String() string {
	return dara.Prettify(s)
}

func (s UpdateTargetListByBatchRequestOperationList) GoString() string {
	return s.String()
}

func (s *UpdateTargetListByBatchRequestOperationList) GetGroupId() *string {
	return s.GroupId
}

func (s *UpdateTargetListByBatchRequestOperationList) GetOperation() *string {
	return s.Operation
}

func (s *UpdateTargetListByBatchRequestOperationList) GetUuid() *string {
	return s.Uuid
}

func (s *UpdateTargetListByBatchRequestOperationList) GetVpcInstanceId() *string {
	return s.VpcInstanceId
}

func (s *UpdateTargetListByBatchRequestOperationList) SetGroupId(v string) *UpdateTargetListByBatchRequestOperationList {
	s.GroupId = &v
	return s
}

func (s *UpdateTargetListByBatchRequestOperationList) SetOperation(v string) *UpdateTargetListByBatchRequestOperationList {
	s.Operation = &v
	return s
}

func (s *UpdateTargetListByBatchRequestOperationList) SetUuid(v string) *UpdateTargetListByBatchRequestOperationList {
	s.Uuid = &v
	return s
}

func (s *UpdateTargetListByBatchRequestOperationList) SetVpcInstanceId(v string) *UpdateTargetListByBatchRequestOperationList {
	s.VpcInstanceId = &v
	return s
}

func (s *UpdateTargetListByBatchRequestOperationList) Validate() error {
	return dara.Validate(s)
}
