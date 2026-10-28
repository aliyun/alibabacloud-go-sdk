// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePxfsAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *RevokePxfsAccessTokenRequest
	GetDBInstanceName() *string
	SetRegionId(v string) *RevokePxfsAccessTokenRequest
	GetRegionId() *string
	SetTenantId(v string) *RevokePxfsAccessTokenRequest
	GetTenantId() *string
	SetTokenId(v string) *RevokePxfsAccessTokenRequest
	GetTokenId() *string
}

type RevokePxfsAccessTokenRequest struct {
	// The ID of the PolarDB-X instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
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
	// Token ID
	//
	// This parameter is required.
	//
	// example:
	//
	// token-001
	TokenId *string `json:"TokenId,omitempty" xml:"TokenId,omitempty"`
}

func (s RevokePxfsAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokePxfsAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *RevokePxfsAccessTokenRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *RevokePxfsAccessTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokePxfsAccessTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokePxfsAccessTokenRequest) GetTokenId() *string {
	return s.TokenId
}

func (s *RevokePxfsAccessTokenRequest) SetDBInstanceName(v string) *RevokePxfsAccessTokenRequest {
	s.DBInstanceName = &v
	return s
}

func (s *RevokePxfsAccessTokenRequest) SetRegionId(v string) *RevokePxfsAccessTokenRequest {
	s.RegionId = &v
	return s
}

func (s *RevokePxfsAccessTokenRequest) SetTenantId(v string) *RevokePxfsAccessTokenRequest {
	s.TenantId = &v
	return s
}

func (s *RevokePxfsAccessTokenRequest) SetTokenId(v string) *RevokePxfsAccessTokenRequest {
	s.TokenId = &v
	return s
}

func (s *RevokePxfsAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
