// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliDingMinutesContentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMinutesId(v string) *GetAliDingMinutesContentRequest
	GetMinutesId() *string
	SetTenantId(v string) *GetAliDingMinutesContentRequest
	GetTenantId() *string
}

type GetAliDingMinutesContentRequest struct {
	// The DingTalk minutes ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76327569643231383535353939365f3436383537393431335f32
	MinutesId *string `json:"minutesId,omitempty" xml:"minutesId,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetAliDingMinutesContentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliDingMinutesContentRequest) GoString() string {
	return s.String()
}

func (s *GetAliDingMinutesContentRequest) GetMinutesId() *string {
	return s.MinutesId
}

func (s *GetAliDingMinutesContentRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetAliDingMinutesContentRequest) SetMinutesId(v string) *GetAliDingMinutesContentRequest {
	s.MinutesId = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) SetTenantId(v string) *GetAliDingMinutesContentRequest {
	s.TenantId = &v
	return s
}

func (s *GetAliDingMinutesContentRequest) Validate() error {
	return dara.Validate(s)
}
