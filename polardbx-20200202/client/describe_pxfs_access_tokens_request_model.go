// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsAccessTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribePxfsAccessTokensRequest
	GetDBInstanceName() *string
	SetIdentityName(v string) *DescribePxfsAccessTokensRequest
	GetIdentityName() *string
	SetRegionId(v string) *DescribePxfsAccessTokensRequest
	GetRegionId() *string
	SetTenantId(v string) *DescribePxfsAccessTokensRequest
	GetTenantId() *string
}

type DescribePxfsAccessTokensRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The identity name.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tenant-001
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribePxfsAccessTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsAccessTokensRequest) GoString() string {
	return s.String()
}

func (s *DescribePxfsAccessTokensRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribePxfsAccessTokensRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *DescribePxfsAccessTokensRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePxfsAccessTokensRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribePxfsAccessTokensRequest) SetDBInstanceName(v string) *DescribePxfsAccessTokensRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribePxfsAccessTokensRequest) SetIdentityName(v string) *DescribePxfsAccessTokensRequest {
	s.IdentityName = &v
	return s
}

func (s *DescribePxfsAccessTokensRequest) SetRegionId(v string) *DescribePxfsAccessTokensRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePxfsAccessTokensRequest) SetTenantId(v string) *DescribePxfsAccessTokensRequest {
	s.TenantId = &v
	return s
}

func (s *DescribePxfsAccessTokensRequest) Validate() error {
	return dara.Validate(s)
}
