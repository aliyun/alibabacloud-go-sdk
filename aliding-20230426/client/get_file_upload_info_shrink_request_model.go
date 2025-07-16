// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileUploadInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionShrink(v string) *GetFileUploadInfoShrinkRequest
	GetOptionShrink() *string
	SetParentDentryUuid(v string) *GetFileUploadInfoShrinkRequest
	GetParentDentryUuid() *string
	SetProtocol(v string) *GetFileUploadInfoShrinkRequest
	GetProtocol() *string
	SetTenantContextShrink(v string) *GetFileUploadInfoShrinkRequest
	GetTenantContextShrink() *string
}

type GetFileUploadInfoShrinkRequest struct {
	OptionShrink *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// example:
	//
	// dentryUuid
	ParentDentryUuid *string `json:"ParentDentryUuid,omitempty" xml:"ParentDentryUuid,omitempty"`
	// example:
	//
	// HEADER_SIGNATURE
	Protocol            *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s GetFileUploadInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFileUploadInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetFileUploadInfoShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetFileUploadInfoShrinkRequest) GetParentDentryUuid() *string {
	return s.ParentDentryUuid
}

func (s *GetFileUploadInfoShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *GetFileUploadInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetFileUploadInfoShrinkRequest) SetOptionShrink(v string) *GetFileUploadInfoShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetFileUploadInfoShrinkRequest) SetParentDentryUuid(v string) *GetFileUploadInfoShrinkRequest {
	s.ParentDentryUuid = &v
	return s
}

func (s *GetFileUploadInfoShrinkRequest) SetProtocol(v string) *GetFileUploadInfoShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *GetFileUploadInfoShrinkRequest) SetTenantContextShrink(v string) *GetFileUploadInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetFileUploadInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
