// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *UploadMediaShrinkRequest
	GetTenantContextShrink() *string
	SetMediaName(v string) *UploadMediaShrinkRequest
	GetMediaName() *string
	SetMediaType(v string) *UploadMediaShrinkRequest
	GetMediaType() *string
	SetOrgId(v int64) *UploadMediaShrinkRequest
	GetOrgId() *int64
	SetUrl(v string) *UploadMediaShrinkRequest
	GetUrl() *string
}

type UploadMediaShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// test.jpg
	MediaName *string `json:"mediaName,omitempty" xml:"mediaName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// image
	MediaType *string `json:"mediaType,omitempty" xml:"mediaType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	OrgId *int64 `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://alicdn.com/xxx.jpg
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s UploadMediaShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *UploadMediaShrinkRequest) GetMediaName() *string {
	return s.MediaName
}

func (s *UploadMediaShrinkRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *UploadMediaShrinkRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *UploadMediaShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *UploadMediaShrinkRequest) SetTenantContextShrink(v string) *UploadMediaShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *UploadMediaShrinkRequest) SetMediaName(v string) *UploadMediaShrinkRequest {
	s.MediaName = &v
	return s
}

func (s *UploadMediaShrinkRequest) SetMediaType(v string) *UploadMediaShrinkRequest {
	s.MediaType = &v
	return s
}

func (s *UploadMediaShrinkRequest) SetOrgId(v int64) *UploadMediaShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *UploadMediaShrinkRequest) SetUrl(v string) *UploadMediaShrinkRequest {
	s.Url = &v
	return s
}

func (s *UploadMediaShrinkRequest) Validate() error {
	return dara.Validate(s)
}
