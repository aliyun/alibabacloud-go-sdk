// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateApplyResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeFirewallTemplateApplyResultsRequest
	GetClientToken() *string
	SetFirewallTemplateId(v string) *DescribeFirewallTemplateApplyResultsRequest
	GetFirewallTemplateId() *string
	SetPageNumber(v int32) *DescribeFirewallTemplateApplyResultsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFirewallTemplateApplyResultsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeFirewallTemplateApplyResultsRequest
	GetRegionId() *string
	SetTaskId(v []*string) *DescribeFirewallTemplateApplyResultsRequest
	GetTaskId() []*string
}

type DescribeFirewallTemplateApplyResultsRequest struct {
	// The client token that is used to ensure the idempotence of the request. A client token is generated from your client to ensure uniqueness across different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The firewall template ID.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The page number of the firewall rule list.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the simple application server. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query the list of supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the template application task.
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s DescribeFirewallTemplateApplyResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateApplyResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFirewallTemplateApplyResultsRequest) GetTaskId() []*string {
	return s.TaskId
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetClientToken(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetFirewallTemplateId(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetPageNumber(v int32) *DescribeFirewallTemplateApplyResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetPageSize(v int32) *DescribeFirewallTemplateApplyResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetRegionId(v string) *DescribeFirewallTemplateApplyResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) SetTaskId(v []*string) *DescribeFirewallTemplateApplyResultsRequest {
	s.TaskId = v
	return s
}

func (s *DescribeFirewallTemplateApplyResultsRequest) Validate() error {
	return dara.Validate(s)
}
