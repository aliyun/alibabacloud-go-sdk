// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupNotifyPolicyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DescribeMonitorGroupNotifyPolicyListRequest
	GetGroupId() *string
	SetPageNumber(v int32) *DescribeMonitorGroupNotifyPolicyListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitorGroupNotifyPolicyListRequest
	GetPageSize() *int32
	SetPolicyType(v string) *DescribeMonitorGroupNotifyPolicyListRequest
	GetPolicyType() *string
	SetRegionId(v string) *DescribeMonitorGroupNotifyPolicyListRequest
	GetRegionId() *string
}

type DescribeMonitorGroupNotifyPolicyListRequest struct {
	// The ID of the application group.
	//
	// example:
	//
	// 6780****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the policy.
	//
	// Valid value: PauseNotify.
	//
	// This parameter is required.
	//
	// example:
	//
	// PauseNotify
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeMonitorGroupNotifyPolicyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupNotifyPolicyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetGroupId(v string) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPageNumber(v int32) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPageSize(v int32) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetPolicyType(v string) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.PolicyType = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) SetRegionId(v string) *DescribeMonitorGroupNotifyPolicyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorGroupNotifyPolicyListRequest) Validate() error {
	return dara.Validate(s)
}
