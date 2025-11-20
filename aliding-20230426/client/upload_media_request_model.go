// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *UploadMediaRequestTenantContext) *UploadMediaRequest
	GetTenantContext() *UploadMediaRequestTenantContext
	SetMediaName(v string) *UploadMediaRequest
	GetMediaName() *string
	SetMediaType(v string) *UploadMediaRequest
	GetMediaType() *string
	SetOrgId(v int64) *UploadMediaRequest
	GetOrgId() *int64
	SetUrl(v string) *UploadMediaRequest
	GetUrl() *string
}

type UploadMediaRequest struct {
	TenantContext *UploadMediaRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s UploadMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaRequest) GoString() string {
	return s.String()
}

func (s *UploadMediaRequest) GetTenantContext() *UploadMediaRequestTenantContext {
	return s.TenantContext
}

func (s *UploadMediaRequest) GetMediaName() *string {
	return s.MediaName
}

func (s *UploadMediaRequest) GetMediaType() *string {
	return s.MediaType
}

func (s *UploadMediaRequest) GetOrgId() *int64 {
	return s.OrgId
}

func (s *UploadMediaRequest) GetUrl() *string {
	return s.Url
}

func (s *UploadMediaRequest) SetTenantContext(v *UploadMediaRequestTenantContext) *UploadMediaRequest {
	s.TenantContext = v
	return s
}

func (s *UploadMediaRequest) SetMediaName(v string) *UploadMediaRequest {
	s.MediaName = &v
	return s
}

func (s *UploadMediaRequest) SetMediaType(v string) *UploadMediaRequest {
	s.MediaType = &v
	return s
}

func (s *UploadMediaRequest) SetOrgId(v int64) *UploadMediaRequest {
	s.OrgId = &v
	return s
}

func (s *UploadMediaRequest) SetUrl(v string) *UploadMediaRequest {
	s.Url = &v
	return s
}

func (s *UploadMediaRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadMediaRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UploadMediaRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UploadMediaRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UploadMediaRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadMediaRequestTenantContext) SetTenantId(v string) *UploadMediaRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UploadMediaRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
