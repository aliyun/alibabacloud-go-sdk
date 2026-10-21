// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLifecyclePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeLifecyclePoliciesRequest
	GetDescription() *string
	SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest
	GetFileSystemId() *string
	SetLifecyclePolicyId(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyId() *string
	SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyName() *string
	SetLifecyclePolicyType(v string) *DescribeLifecyclePoliciesRequest
	GetLifecyclePolicyType() *string
	SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLifecyclePoliciesRequest
	GetPageSize() *int32
	SetPath(v string) *DescribeLifecyclePoliciesRequest
	GetPath() *string
	SetStorageType(v string) *DescribeLifecyclePoliciesRequest
	GetStorageType() *string
}

type DescribeLifecyclePoliciesRequest struct {
	// The description of the policy.
	//
	// >This parameter is supported only by CPFS for Lingjun.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the lifecycle policy.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// Policy Name of the lifecycle policy. Policy Name must meet the following rules:
	//
	//
	// Policy Name must be 3 to 64 characters in length. It must start with a letter and can contain letters, digits, underscores (_), and hyphens (-).
	//
	// >For General-purpose NAS file systems, this parameter is optional. If you specify this parameter, the specified value is used. If you do not specify this parameter, the value of LifecyclePolicyId is used.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The policy type. Valid values:
	//
	// - Auto: Automatic execution.
	//
	// - OnDemand: On-demand execution.
	//
	// >This parameter is supported only by CPFS for Lingjun.
	//
	// example:
	//
	// Auto
	LifecyclePolicyType *string `json:"LifecyclePolicyType,omitempty" xml:"LifecyclePolicyType,omitempty"`
	// The page number of the list.
	//
	// Starting value (default value): 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of lifecycle management policies to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters lifecycle policies by path.
	//
	// >This parameter is supported only by CPFS for Lingjun.
	//
	// example:
	//
	// /abc/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The tiered storage type. Valid values:
	//
	// - InfrequentAccess: IA storage class.
	//
	// - Archive: Archive storage class.
	//
	// >If you do not specify StorageType, all lifecycle policies are returned.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DescribeLifecyclePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLifecyclePoliciesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLifecyclePoliciesRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeLifecyclePoliciesRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *DescribeLifecyclePoliciesRequest) GetLifecyclePolicyType() *string {
	return s.LifecyclePolicyType
}

func (s *DescribeLifecyclePoliciesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLifecyclePoliciesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLifecyclePoliciesRequest) GetPath() *string {
	return s.Path
}

func (s *DescribeLifecyclePoliciesRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeLifecyclePoliciesRequest) SetDescription(v string) *DescribeLifecyclePoliciesRequest {
	s.Description = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetFileSystemId(v string) *DescribeLifecyclePoliciesRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyId(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyId = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyName(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetLifecyclePolicyType(v string) *DescribeLifecyclePoliciesRequest {
	s.LifecyclePolicyType = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageNumber(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPageSize(v int32) *DescribeLifecyclePoliciesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetPath(v string) *DescribeLifecyclePoliciesRequest {
	s.Path = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) SetStorageType(v string) *DescribeLifecyclePoliciesRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeLifecyclePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
