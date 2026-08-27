// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentType(v string) *GetSourceUploadSignatureRequest
	GetContentType() *string
	SetExpires(v int64) *GetSourceUploadSignatureRequest
	GetExpires() *int64
	SetFilename(v string) *GetSourceUploadSignatureRequest
	GetFilename() *string
	SetOperatingObjectName(v string) *GetSourceUploadSignatureRequest
	GetOperatingObjectName() *string
	SetScope(v string) *GetSourceUploadSignatureRequest
	GetScope() *string
	SetTenantId(v string) *GetSourceUploadSignatureRequest
	GetTenantId() *string
}

type GetSourceUploadSignatureRequest struct {
	// The content type. Valid values: Text and Markdown.
	//
	// example:
	//
	// string_value
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The expiration time of the signed URL, in seconds. Default value: 3600.
	//
	// example:
	//
	// 3600
	Expires *int64 `json:"expires,omitempty" xml:"expires,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.pdf
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// The name of the digital employee (operating object name). This parameter is optional.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// source
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The tenant ID to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetSourceUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *GetSourceUploadSignatureRequest) GetContentType() *string {
	return s.ContentType
}

func (s *GetSourceUploadSignatureRequest) GetExpires() *int64 {
	return s.Expires
}

func (s *GetSourceUploadSignatureRequest) GetFilename() *string {
	return s.Filename
}

func (s *GetSourceUploadSignatureRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *GetSourceUploadSignatureRequest) GetScope() *string {
	return s.Scope
}

func (s *GetSourceUploadSignatureRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetSourceUploadSignatureRequest) SetContentType(v string) *GetSourceUploadSignatureRequest {
	s.ContentType = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetExpires(v int64) *GetSourceUploadSignatureRequest {
	s.Expires = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetFilename(v string) *GetSourceUploadSignatureRequest {
	s.Filename = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetOperatingObjectName(v string) *GetSourceUploadSignatureRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetScope(v string) *GetSourceUploadSignatureRequest {
	s.Scope = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) SetTenantId(v string) *GetSourceUploadSignatureRequest {
	s.TenantId = &v
	return s
}

func (s *GetSourceUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
