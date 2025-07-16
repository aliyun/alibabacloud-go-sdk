// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConvExtensionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobileUrl(v string) *UpdateConvExtensionRequest
	GetMobileUrl() *string
	SetPcUrl(v string) *UpdateConvExtensionRequest
	GetPcUrl() *string
	SetStaffIdList(v []*string) *UpdateConvExtensionRequest
	GetStaffIdList() []*string
	SetSystemUid(v string) *UpdateConvExtensionRequest
	GetSystemUid() *string
	SetTenantContext(v *UpdateConvExtensionRequestTenantContext) *UpdateConvExtensionRequest
	GetTenantContext() *UpdateConvExtensionRequestTenantContext
}

type UpdateConvExtensionRequest struct {
	// example:
	//
	// https://xxx
	MobileUrl *string `json:"MobileUrl,omitempty" xml:"MobileUrl,omitempty"`
	// example:
	//
	// https://xxx
	PcUrl       *string   `json:"PcUrl,omitempty" xml:"PcUrl,omitempty"`
	StaffIdList []*string `json:"StaffIdList,omitempty" xml:"StaffIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 546374856
	SystemUid     *string                                  `json:"SystemUid,omitempty" xml:"SystemUid,omitempty"`
	TenantContext *UpdateConvExtensionRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
}

func (s UpdateConvExtensionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionRequest) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionRequest) GetMobileUrl() *string {
	return s.MobileUrl
}

func (s *UpdateConvExtensionRequest) GetPcUrl() *string {
	return s.PcUrl
}

func (s *UpdateConvExtensionRequest) GetStaffIdList() []*string {
	return s.StaffIdList
}

func (s *UpdateConvExtensionRequest) GetSystemUid() *string {
	return s.SystemUid
}

func (s *UpdateConvExtensionRequest) GetTenantContext() *UpdateConvExtensionRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateConvExtensionRequest) SetMobileUrl(v string) *UpdateConvExtensionRequest {
	s.MobileUrl = &v
	return s
}

func (s *UpdateConvExtensionRequest) SetPcUrl(v string) *UpdateConvExtensionRequest {
	s.PcUrl = &v
	return s
}

func (s *UpdateConvExtensionRequest) SetStaffIdList(v []*string) *UpdateConvExtensionRequest {
	s.StaffIdList = v
	return s
}

func (s *UpdateConvExtensionRequest) SetSystemUid(v string) *UpdateConvExtensionRequest {
	s.SystemUid = &v
	return s
}

func (s *UpdateConvExtensionRequest) SetTenantContext(v *UpdateConvExtensionRequestTenantContext) *UpdateConvExtensionRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateConvExtensionRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateConvExtensionRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateConvExtensionRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateConvExtensionRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateConvExtensionRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateConvExtensionRequestTenantContext) SetTenantId(v string) *UpdateConvExtensionRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateConvExtensionRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
