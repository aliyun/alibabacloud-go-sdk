// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBTenantApiKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeAgenticDBTenantApiKeysRequest
	GetDBClusterId() *string
	SetPageNumber(v int32) *DescribeAgenticDBTenantApiKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAgenticDBTenantApiKeysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAgenticDBTenantApiKeysRequest
	GetRegionId() *string
	SetTenantName(v string) *DescribeAgenticDBTenantApiKeysRequest
	GetTenantName() *string
}

type DescribeAgenticDBTenantApiKeysRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pagc-bp1abcdef1234567
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// my-saas-app
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeAgenticDBTenantApiKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBTenantApiKeysRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBTenantApiKeysRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAgenticDBTenantApiKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAgenticDBTenantApiKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAgenticDBTenantApiKeysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAgenticDBTenantApiKeysRequest) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeAgenticDBTenantApiKeysRequest) SetDBClusterId(v string) *DescribeAgenticDBTenantApiKeysRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysRequest) SetPageNumber(v int32) *DescribeAgenticDBTenantApiKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysRequest) SetPageSize(v int32) *DescribeAgenticDBTenantApiKeysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysRequest) SetRegionId(v string) *DescribeAgenticDBTenantApiKeysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysRequest) SetTenantName(v string) *DescribeAgenticDBTenantApiKeysRequest {
	s.TenantName = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysRequest) Validate() error {
	return dara.Validate(s)
}
