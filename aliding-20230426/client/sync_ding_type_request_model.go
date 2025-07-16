// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingType(v string) *SyncDingTypeRequest
	GetDingType() *string
	SetIsDimission(v string) *SyncDingTypeRequest
	GetIsDimission() *string
	SetSource(v string) *SyncDingTypeRequest
	GetSource() *string
	SetTenantContext(v *SyncDingTypeRequestTenantContext) *SyncDingTypeRequest
	GetTenantContext() *SyncDingTypeRequestTenantContext
	SetWorkNo(v string) *SyncDingTypeRequest
	GetWorkNo() *string
}

type SyncDingTypeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ANT_DING
	DingType *string `json:"DingType,omitempty" xml:"DingType,omitempty"`
	// example:
	//
	// y
	IsDimission *string `json:"IsDimission,omitempty" xml:"IsDimission,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// antding
	Source        *string                           `json:"Source,omitempty" xml:"Source,omitempty"`
	TenantContext *SyncDingTypeRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
}

func (s SyncDingTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeRequest) GoString() string {
	return s.String()
}

func (s *SyncDingTypeRequest) GetDingType() *string {
	return s.DingType
}

func (s *SyncDingTypeRequest) GetIsDimission() *string {
	return s.IsDimission
}

func (s *SyncDingTypeRequest) GetSource() *string {
	return s.Source
}

func (s *SyncDingTypeRequest) GetTenantContext() *SyncDingTypeRequestTenantContext {
	return s.TenantContext
}

func (s *SyncDingTypeRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *SyncDingTypeRequest) SetDingType(v string) *SyncDingTypeRequest {
	s.DingType = &v
	return s
}

func (s *SyncDingTypeRequest) SetIsDimission(v string) *SyncDingTypeRequest {
	s.IsDimission = &v
	return s
}

func (s *SyncDingTypeRequest) SetSource(v string) *SyncDingTypeRequest {
	s.Source = &v
	return s
}

func (s *SyncDingTypeRequest) SetTenantContext(v *SyncDingTypeRequestTenantContext) *SyncDingTypeRequest {
	s.TenantContext = v
	return s
}

func (s *SyncDingTypeRequest) SetWorkNo(v string) *SyncDingTypeRequest {
	s.WorkNo = &v
	return s
}

func (s *SyncDingTypeRequest) Validate() error {
	return dara.Validate(s)
}

type SyncDingTypeRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s SyncDingTypeRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeRequestTenantContext) GoString() string {
	return s.String()
}

func (s *SyncDingTypeRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *SyncDingTypeRequestTenantContext) SetTenantId(v string) *SyncDingTypeRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *SyncDingTypeRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
