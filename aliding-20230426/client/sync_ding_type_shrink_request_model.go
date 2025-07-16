// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDingType(v string) *SyncDingTypeShrinkRequest
	GetDingType() *string
	SetIsDimission(v string) *SyncDingTypeShrinkRequest
	GetIsDimission() *string
	SetSource(v string) *SyncDingTypeShrinkRequest
	GetSource() *string
	SetTenantContextShrink(v string) *SyncDingTypeShrinkRequest
	GetTenantContextShrink() *string
	SetWorkNo(v string) *SyncDingTypeShrinkRequest
	GetWorkNo() *string
}

type SyncDingTypeShrinkRequest struct {
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
	Source              *string `json:"Source,omitempty" xml:"Source,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	WorkNo *string `json:"WorkNo,omitempty" xml:"WorkNo,omitempty"`
}

func (s SyncDingTypeShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeShrinkRequest) GoString() string {
	return s.String()
}

func (s *SyncDingTypeShrinkRequest) GetDingType() *string {
	return s.DingType
}

func (s *SyncDingTypeShrinkRequest) GetIsDimission() *string {
	return s.IsDimission
}

func (s *SyncDingTypeShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *SyncDingTypeShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *SyncDingTypeShrinkRequest) GetWorkNo() *string {
	return s.WorkNo
}

func (s *SyncDingTypeShrinkRequest) SetDingType(v string) *SyncDingTypeShrinkRequest {
	s.DingType = &v
	return s
}

func (s *SyncDingTypeShrinkRequest) SetIsDimission(v string) *SyncDingTypeShrinkRequest {
	s.IsDimission = &v
	return s
}

func (s *SyncDingTypeShrinkRequest) SetSource(v string) *SyncDingTypeShrinkRequest {
	s.Source = &v
	return s
}

func (s *SyncDingTypeShrinkRequest) SetTenantContextShrink(v string) *SyncDingTypeShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *SyncDingTypeShrinkRequest) SetWorkNo(v string) *SyncDingTypeShrinkRequest {
	s.WorkNo = &v
	return s
}

func (s *SyncDingTypeShrinkRequest) Validate() error {
	return dara.Validate(s)
}
