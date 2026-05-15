// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetWorkspaceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *SetWorkspaceQuotaRequest
	GetAutoPay() *bool
	SetClientToken(v string) *SetWorkspaceQuotaRequest
	GetClientToken() *string
	SetCuQuota(v int32) *SetWorkspaceQuotaRequest
	GetCuQuota() *int32
	SetRegion(v string) *SetWorkspaceQuotaRequest
	GetRegion() *string
	SetWorkspaceId(v string) *SetWorkspaceQuotaRequest
	GetWorkspaceId() *string
}

type SetWorkspaceQuotaRequest struct {
	// example:
	//
	// false
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// example:
	//
	// acdxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 200
	CuQuota *int32 `json:"CuQuota,omitempty" xml:"CuQuota,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20923*****7291
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SetWorkspaceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetWorkspaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetWorkspaceQuotaRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *SetWorkspaceQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetWorkspaceQuotaRequest) GetCuQuota() *int32 {
	return s.CuQuota
}

func (s *SetWorkspaceQuotaRequest) GetRegion() *string {
	return s.Region
}

func (s *SetWorkspaceQuotaRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SetWorkspaceQuotaRequest) SetAutoPay(v bool) *SetWorkspaceQuotaRequest {
	s.AutoPay = &v
	return s
}

func (s *SetWorkspaceQuotaRequest) SetClientToken(v string) *SetWorkspaceQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *SetWorkspaceQuotaRequest) SetCuQuota(v int32) *SetWorkspaceQuotaRequest {
	s.CuQuota = &v
	return s
}

func (s *SetWorkspaceQuotaRequest) SetRegion(v string) *SetWorkspaceQuotaRequest {
	s.Region = &v
	return s
}

func (s *SetWorkspaceQuotaRequest) SetWorkspaceId(v string) *SetWorkspaceQuotaRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SetWorkspaceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
