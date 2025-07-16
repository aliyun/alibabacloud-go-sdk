// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultipartFileUploadInfosShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionShrink(v string) *GetMultipartFileUploadInfosShrinkRequest
	GetOptionShrink() *string
	SetPartNumbersShrink(v string) *GetMultipartFileUploadInfosShrinkRequest
	GetPartNumbersShrink() *string
	SetTenantContextShrink(v string) *GetMultipartFileUploadInfosShrinkRequest
	GetTenantContextShrink() *string
	SetUploadKey(v string) *GetMultipartFileUploadInfosShrinkRequest
	GetUploadKey() *string
}

type GetMultipartFileUploadInfosShrinkRequest struct {
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	PartNumbersShrink   *string `json:"PartNumbers,omitempty" xml:"PartNumbers,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// hwHPAAAAAipHxxxxx
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s GetMultipartFileUploadInfosShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMultipartFileUploadInfosShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetMultipartFileUploadInfosShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetMultipartFileUploadInfosShrinkRequest) GetPartNumbersShrink() *string {
	return s.PartNumbersShrink
}

func (s *GetMultipartFileUploadInfosShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetMultipartFileUploadInfosShrinkRequest) GetUploadKey() *string {
	return s.UploadKey
}

func (s *GetMultipartFileUploadInfosShrinkRequest) SetOptionShrink(v string) *GetMultipartFileUploadInfosShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkRequest) SetPartNumbersShrink(v string) *GetMultipartFileUploadInfosShrinkRequest {
	s.PartNumbersShrink = &v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkRequest) SetTenantContextShrink(v string) *GetMultipartFileUploadInfosShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkRequest) SetUploadKey(v string) *GetMultipartFileUploadInfosShrinkRequest {
	s.UploadKey = &v
	return s
}

func (s *GetMultipartFileUploadInfosShrinkRequest) Validate() error {
	return dara.Validate(s)
}
