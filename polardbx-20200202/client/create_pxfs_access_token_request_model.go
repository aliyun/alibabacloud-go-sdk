// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePxfsAccessTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreatePxfsAccessTokenRequest
	GetDBInstanceName() *string
	SetIdentityName(v string) *CreatePxfsAccessTokenRequest
	GetIdentityName() *string
	SetLabel(v string) *CreatePxfsAccessTokenRequest
	GetLabel() *string
	SetRegionId(v string) *CreatePxfsAccessTokenRequest
	GetRegionId() *string
	SetTenantId(v string) *CreatePxfsAccessTokenRequest
	GetTenantId() *string
	SetTtlSeconds(v int64) *CreatePxfsAccessTokenRequest
	GetTtlSeconds() *int64
}

type CreatePxfsAccessTokenRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The name of the identity.
	//
	// This parameter is required.
	//
	// example:
	//
	// agent-runtime
	IdentityName *string `json:"IdentityName,omitempty" xml:"IdentityName,omitempty"`
	// The label of the token.
	//
	// example:
	//
	// production-agent
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
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
	// The validity period of the token.
	//
	// example:
	//
	// 86400
	TtlSeconds *int64 `json:"TtlSeconds,omitempty" xml:"TtlSeconds,omitempty"`
}

func (s CreatePxfsAccessTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePxfsAccessTokenRequest) GoString() string {
	return s.String()
}

func (s *CreatePxfsAccessTokenRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreatePxfsAccessTokenRequest) GetIdentityName() *string {
	return s.IdentityName
}

func (s *CreatePxfsAccessTokenRequest) GetLabel() *string {
	return s.Label
}

func (s *CreatePxfsAccessTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePxfsAccessTokenRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePxfsAccessTokenRequest) GetTtlSeconds() *int64 {
	return s.TtlSeconds
}

func (s *CreatePxfsAccessTokenRequest) SetDBInstanceName(v string) *CreatePxfsAccessTokenRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) SetIdentityName(v string) *CreatePxfsAccessTokenRequest {
	s.IdentityName = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) SetLabel(v string) *CreatePxfsAccessTokenRequest {
	s.Label = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) SetRegionId(v string) *CreatePxfsAccessTokenRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) SetTenantId(v string) *CreatePxfsAccessTokenRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) SetTtlSeconds(v int64) *CreatePxfsAccessTokenRequest {
	s.TtlSeconds = &v
	return s
}

func (s *CreatePxfsAccessTokenRequest) Validate() error {
	return dara.Validate(s)
}
