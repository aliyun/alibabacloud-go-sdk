// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteCodeBundleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBundleVersion(v string) *CompleteCodeBundleResponseBody
	GetBundleVersion() *string
	SetByteSize(v int64) *CompleteCodeBundleResponseBody
	GetByteSize() *int64
	SetCodeBundleId(v int64) *CompleteCodeBundleResponseBody
	GetCodeBundleId() *int64
	SetContentType(v string) *CompleteCodeBundleResponseBody
	GetContentType() *string
	SetCreatedAt(v string) *CompleteCodeBundleResponseBody
	GetCreatedAt() *string
	SetFilename(v string) *CompleteCodeBundleResponseBody
	GetFilename() *string
	SetProjectId(v int64) *CompleteCodeBundleResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *CompleteCodeBundleResponseBody
	GetRequestId() *string
	SetStatus(v string) *CompleteCodeBundleResponseBody
	GetStatus() *string
	SetUpdatedAt(v string) *CompleteCodeBundleResponseBody
	GetUpdatedAt() *string
}

type CompleteCodeBundleResponseBody struct {
	// The code bundle version identifier.
	//
	// example:
	//
	// 1
	BundleVersion *string `json:"bundleVersion,omitempty" xml:"bundleVersion,omitempty"`
	// The object size.
	//
	// example:
	//
	// 111
	ByteSize *int64 `json:"byteSize,omitempty" xml:"byteSize,omitempty"`
	// The code bundle ID.
	//
	// example:
	//
	// 111
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The MIME type of the stored code bundle. This is typically application/octet-stream for pre-signed PUT operations.
	//
	// example:
	//
	// application/octet-stream
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The time when the record was created, in RFC 3339 format.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The file name.
	//
	// example:
	//
	// test-cases.zip
	Filename *string `json:"filename,omitempty" xml:"filename,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 76851f2b5bf0187fbc29e8bca4
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The code bundle status.
	//
	// example:
	//
	// ready
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the record was last updated.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CompleteCodeBundleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CompleteCodeBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteCodeBundleResponseBody) GetBundleVersion() *string {
	return s.BundleVersion
}

func (s *CompleteCodeBundleResponseBody) GetByteSize() *int64 {
	return s.ByteSize
}

func (s *CompleteCodeBundleResponseBody) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *CompleteCodeBundleResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *CompleteCodeBundleResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CompleteCodeBundleResponseBody) GetFilename() *string {
	return s.Filename
}

func (s *CompleteCodeBundleResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CompleteCodeBundleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CompleteCodeBundleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CompleteCodeBundleResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CompleteCodeBundleResponseBody) SetBundleVersion(v string) *CompleteCodeBundleResponseBody {
	s.BundleVersion = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetByteSize(v int64) *CompleteCodeBundleResponseBody {
	s.ByteSize = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetCodeBundleId(v int64) *CompleteCodeBundleResponseBody {
	s.CodeBundleId = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetContentType(v string) *CompleteCodeBundleResponseBody {
	s.ContentType = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetCreatedAt(v string) *CompleteCodeBundleResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetFilename(v string) *CompleteCodeBundleResponseBody {
	s.Filename = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetProjectId(v int64) *CompleteCodeBundleResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetRequestId(v string) *CompleteCodeBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetStatus(v string) *CompleteCodeBundleResponseBody {
	s.Status = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) SetUpdatedAt(v string) *CompleteCodeBundleResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CompleteCodeBundleResponseBody) Validate() error {
	return dara.Validate(s)
}
