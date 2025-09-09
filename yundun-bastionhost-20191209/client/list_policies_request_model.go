// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListPoliciesRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListPoliciesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListPoliciesRequest
	GetPageSize() *string
	SetPolicyName(v string) *ListPoliciesRequest
	GetPolicyName() *string
	SetRegionId(v string) *ListPoliciesRequest
	GetRegionId() *string
}

type ListPoliciesRequest struct {
	// The bastion host ID.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-5yd393wzk08
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// > We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the control policy to query. Only exact match is supported.
	//
	// example:
	//
	// 123
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListPoliciesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPoliciesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListPoliciesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListPoliciesRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPoliciesRequest) SetInstanceId(v string) *ListPoliciesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPoliciesRequest) SetPageNumber(v string) *ListPoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesRequest) SetPageSize(v string) *ListPoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesRequest) SetPolicyName(v string) *ListPoliciesRequest {
	s.PolicyName = &v
	return s
}

func (s *ListPoliciesRequest) SetRegionId(v string) *ListPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
