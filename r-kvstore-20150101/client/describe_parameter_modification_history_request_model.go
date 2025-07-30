// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterModificationHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeParameterModificationHistoryRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeParameterModificationHistoryRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeParameterModificationHistoryRequest
	GetNodeId() *string
	SetOwnerAccount(v string) *DescribeParameterModificationHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterModificationHistoryRequest
	GetOwnerId() *int64
	SetParameterName(v string) *DescribeParameterModificationHistoryRequest
	GetParameterName() *string
	SetResourceOwnerAccount(v string) *DescribeParameterModificationHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterModificationHistoryRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *DescribeParameterModificationHistoryRequest
	GetSecurityToken() *string
	SetStartTime(v string) *DescribeParameterModificationHistoryRequest
	GetStartTime() *string
}

type DescribeParameterModificationHistoryRequest struct {
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-05T09:49:27Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the node.
	//
	// > You can set this parameter to query the parameter settings of the specified node in a cluster instance.
	//
	// example:
	//
	// r-bp1xxxxxxxxxxxxx-db-0
	NodeId       *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// script_check_enable
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-05T08:49:27Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeParameterModificationHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterModificationHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterModificationHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeParameterModificationHistoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeParameterModificationHistoryRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeParameterModificationHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterModificationHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterModificationHistoryRequest) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeParameterModificationHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterModificationHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterModificationHistoryRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeParameterModificationHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeParameterModificationHistoryRequest) SetEndTime(v string) *DescribeParameterModificationHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetInstanceId(v string) *DescribeParameterModificationHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetNodeId(v string) *DescribeParameterModificationHistoryRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetParameterName(v string) *DescribeParameterModificationHistoryRequest {
	s.ParameterName = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerAccount(v string) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetResourceOwnerId(v int64) *DescribeParameterModificationHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetSecurityToken(v string) *DescribeParameterModificationHistoryRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) SetStartTime(v string) *DescribeParameterModificationHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeParameterModificationHistoryRequest) Validate() error {
	return dara.Validate(s)
}
