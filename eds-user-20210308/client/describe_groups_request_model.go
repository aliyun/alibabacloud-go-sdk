// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeGroupsRequest
	GetBizType() *string
	SetExcludeAttachedLoginPolicyGroups(v bool) *DescribeGroupsRequest
	GetExcludeAttachedLoginPolicyGroups() *bool
	SetGroupId(v string) *DescribeGroupsRequest
	GetGroupId() *string
	SetGroupName(v string) *DescribeGroupsRequest
	GetGroupName() *string
	SetLoginPolicyId(v string) *DescribeGroupsRequest
	GetLoginPolicyId() *string
	SetPageNumber(v int32) *DescribeGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGroupsRequest
	GetPageSize() *int32
	SetSolutionId(v string) *DescribeGroupsRequest
	GetSolutionId() *string
	SetTransferFileNeedApproval(v bool) *DescribeGroupsRequest
	GetTransferFileNeedApproval() *bool
}

type DescribeGroupsRequest struct {
	// example:
	//
	// ENTERPRISE
	BizType                          *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExcludeAttachedLoginPolicyGroups *bool   `json:"ExcludeAttachedLoginPolicyGroups,omitempty" xml:"ExcludeAttachedLoginPolicyGroups,omitempty"`
	// example:
	//
	// ug-12341234****
	GroupId       *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	GroupName     *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	LoginPolicyId *string `json:"LoginPolicyId,omitempty" xml:"LoginPolicyId,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// co-0esnf80jab***
	SolutionId               *string `json:"SolutionId,omitempty" xml:"SolutionId,omitempty"`
	TransferFileNeedApproval *bool   `json:"TransferFileNeedApproval,omitempty" xml:"TransferFileNeedApproval,omitempty"`
}

func (s DescribeGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeGroupsRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeGroupsRequest) GetExcludeAttachedLoginPolicyGroups() *bool {
	return s.ExcludeAttachedLoginPolicyGroups
}

func (s *DescribeGroupsRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeGroupsRequest) GetLoginPolicyId() *string {
	return s.LoginPolicyId
}

func (s *DescribeGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGroupsRequest) GetSolutionId() *string {
	return s.SolutionId
}

func (s *DescribeGroupsRequest) GetTransferFileNeedApproval() *bool {
	return s.TransferFileNeedApproval
}

func (s *DescribeGroupsRequest) SetBizType(v string) *DescribeGroupsRequest {
	s.BizType = &v
	return s
}

func (s *DescribeGroupsRequest) SetExcludeAttachedLoginPolicyGroups(v bool) *DescribeGroupsRequest {
	s.ExcludeAttachedLoginPolicyGroups = &v
	return s
}

func (s *DescribeGroupsRequest) SetGroupId(v string) *DescribeGroupsRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeGroupsRequest) SetGroupName(v string) *DescribeGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *DescribeGroupsRequest) SetLoginPolicyId(v string) *DescribeGroupsRequest {
	s.LoginPolicyId = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageNumber(v int32) *DescribeGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeGroupsRequest) SetPageSize(v int32) *DescribeGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeGroupsRequest) SetSolutionId(v string) *DescribeGroupsRequest {
	s.SolutionId = &v
	return s
}

func (s *DescribeGroupsRequest) SetTransferFileNeedApproval(v bool) *DescribeGroupsRequest {
	s.TransferFileNeedApproval = &v
	return s
}

func (s *DescribeGroupsRequest) Validate() error {
	return dara.Validate(s)
}
