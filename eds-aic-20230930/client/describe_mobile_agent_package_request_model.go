// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMobileAgentPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeMobileAgentPackageRequest
	GetInstanceIds() []*string
	SetPackageIds(v []*string) *DescribeMobileAgentPackageRequest
	GetPackageIds() []*string
	SetPackageSpec(v string) *DescribeMobileAgentPackageRequest
	GetPackageSpec() *string
	SetPackageStatus(v string) *DescribeMobileAgentPackageRequest
	GetPackageStatus() *string
	SetPageNum(v int32) *DescribeMobileAgentPackageRequest
	GetPageNum() *int32
	SetPageSize(v int32) *DescribeMobileAgentPackageRequest
	GetPageSize() *int32
}

type DescribeMobileAgentPackageRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	PackageIds  []*string `json:"PackageIds,omitempty" xml:"PackageIds,omitempty" type:"Repeated"`
	// example:
	//
	// advanced
	PackageSpec *string `json:"PackageSpec,omitempty" xml:"PackageSpec,omitempty"`
	// example:
	//
	// ACTIVE
	PackageStatus *string `json:"PackageStatus,omitempty" xml:"PackageStatus,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMobileAgentPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMobileAgentPackageRequest) GoString() string {
	return s.String()
}

func (s *DescribeMobileAgentPackageRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeMobileAgentPackageRequest) GetPackageIds() []*string {
	return s.PackageIds
}

func (s *DescribeMobileAgentPackageRequest) GetPackageSpec() *string {
	return s.PackageSpec
}

func (s *DescribeMobileAgentPackageRequest) GetPackageStatus() *string {
	return s.PackageStatus
}

func (s *DescribeMobileAgentPackageRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *DescribeMobileAgentPackageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMobileAgentPackageRequest) SetInstanceIds(v []*string) *DescribeMobileAgentPackageRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeMobileAgentPackageRequest) SetPackageIds(v []*string) *DescribeMobileAgentPackageRequest {
	s.PackageIds = v
	return s
}

func (s *DescribeMobileAgentPackageRequest) SetPackageSpec(v string) *DescribeMobileAgentPackageRequest {
	s.PackageSpec = &v
	return s
}

func (s *DescribeMobileAgentPackageRequest) SetPackageStatus(v string) *DescribeMobileAgentPackageRequest {
	s.PackageStatus = &v
	return s
}

func (s *DescribeMobileAgentPackageRequest) SetPageNum(v int32) *DescribeMobileAgentPackageRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMobileAgentPackageRequest) SetPageSize(v int32) *DescribeMobileAgentPackageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMobileAgentPackageRequest) Validate() error {
	return dara.Validate(s)
}
