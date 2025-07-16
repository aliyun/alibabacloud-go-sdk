// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileDownloadInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDentryId(v string) *GetFileDownloadInfoShrinkRequest
	GetDentryId() *string
	SetOptionShrink(v string) *GetFileDownloadInfoShrinkRequest
	GetOptionShrink() *string
	SetSpaceId(v string) *GetFileDownloadInfoShrinkRequest
	GetSpaceId() *string
	SetTenantContextShrink(v string) *GetFileDownloadInfoShrinkRequest
	GetTenantContextShrink() *string
}

type GetFileDownloadInfoShrinkRequest struct {
	// example:
	//
	// 798xxxxx
	DentryId     *string `json:"DentryId,omitempty" xml:"DentryId,omitempty"`
	OptionShrink *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// 854xxxx
	SpaceId             *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetFileDownloadInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileDownloadInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFileDownloadInfoShrinkRequest) GetDentryId() *string {
	return s.DentryId
}

func (s *GetFileDownloadInfoShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetFileDownloadInfoShrinkRequest) GetSpaceId() *string {
	return s.SpaceId
}

func (s *GetFileDownloadInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetFileDownloadInfoShrinkRequest) SetDentryId(v string) *GetFileDownloadInfoShrinkRequest {
	s.DentryId = &v
	return s
}

func (s *GetFileDownloadInfoShrinkRequest) SetOptionShrink(v string) *GetFileDownloadInfoShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetFileDownloadInfoShrinkRequest) SetSpaceId(v string) *GetFileDownloadInfoShrinkRequest {
	s.SpaceId = &v
	return s
}

func (s *GetFileDownloadInfoShrinkRequest) SetTenantContextShrink(v string) *GetFileDownloadInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetFileDownloadInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
