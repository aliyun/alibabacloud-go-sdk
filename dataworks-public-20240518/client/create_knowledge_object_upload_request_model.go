// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeObjectUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeObjectUploadRequest
	GetClientToken() *string
	SetContentMd5(v string) *CreateKnowledgeObjectUploadRequest
	GetContentMd5() *string
	SetContentSha256(v string) *CreateKnowledgeObjectUploadRequest
	GetContentSha256() *string
	SetContentType(v string) *CreateKnowledgeObjectUploadRequest
	GetContentType() *string
	SetFileName(v string) *CreateKnowledgeObjectUploadRequest
	GetFileName() *string
	SetKnowledgeBaseId(v string) *CreateKnowledgeObjectUploadRequest
	GetKnowledgeBaseId() *string
	SetSizeBytes(v int64) *CreateKnowledgeObjectUploadRequest
	GetSizeBytes() *int64
}

type CreateKnowledgeObjectUploadRequest struct {
	// The idempotency key for the upload request.
	//
	// example:
	//
	// upload-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The Base64-encoded MD5 hash value of the original file.
	//
	// This parameter is required.
	//
	// example:
	//
	// yaNM/IXZgmmMasifdgcavQ==
	ContentMd5 *string `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	// The SHA-256 checksum of the original file.
	//
	// example:
	//
	// ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
	ContentSha256 *string `json:"ContentSha256,omitempty" xml:"ContentSha256,omitempty"`
	// The HTTP media type for the upload.
	//
	// This parameter is required.
	//
	// example:
	//
	// text/markdown
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The original file name, which is not used as the object ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// policy.md
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The target knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The size of the original file in bytes. This parameter reuses the existing maximum size limit for single-file uploads, which defaults to 100 MiB. If the size exceeds the limit, the InputTooLarge error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1024
	SizeBytes *int64 `json:"SizeBytes,omitempty" xml:"SizeBytes,omitempty"`
}

func (s CreateKnowledgeObjectUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeObjectUploadRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeObjectUploadRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeObjectUploadRequest) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *CreateKnowledgeObjectUploadRequest) GetContentSha256() *string {
	return s.ContentSha256
}

func (s *CreateKnowledgeObjectUploadRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateKnowledgeObjectUploadRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateKnowledgeObjectUploadRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeObjectUploadRequest) GetSizeBytes() *int64 {
	return s.SizeBytes
}

func (s *CreateKnowledgeObjectUploadRequest) SetClientToken(v string) *CreateKnowledgeObjectUploadRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetContentMd5(v string) *CreateKnowledgeObjectUploadRequest {
	s.ContentMd5 = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetContentSha256(v string) *CreateKnowledgeObjectUploadRequest {
	s.ContentSha256 = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetContentType(v string) *CreateKnowledgeObjectUploadRequest {
	s.ContentType = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetFileName(v string) *CreateKnowledgeObjectUploadRequest {
	s.FileName = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeObjectUploadRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) SetSizeBytes(v int64) *CreateKnowledgeObjectUploadRequest {
	s.SizeBytes = &v
	return s
}

func (s *CreateKnowledgeObjectUploadRequest) Validate() error {
	return dara.Validate(s)
}
