// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeBundleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBundleVersion(v string) *CreateCodeBundleResponseBody
	GetBundleVersion() *string
	SetCodeBundleId(v int64) *CreateCodeBundleResponseBody
	GetCodeBundleId() *int64
	SetCreatedAt(v string) *CreateCodeBundleResponseBody
	GetCreatedAt() *string
	SetFilename(v string) *CreateCodeBundleResponseBody
	GetFilename() *string
	SetProjectId(v int64) *CreateCodeBundleResponseBody
	GetProjectId() *int64
	SetRequestId(v string) *CreateCodeBundleResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateCodeBundleResponseBody
	GetStatus() *string
	SetUpdatedAt(v string) *CreateCodeBundleResponseBody
	GetUpdatedAt() *string
	SetUpload(v *CreateCodeBundleResponseBodyUpload) *CreateCodeBundleResponseBody
	GetUpload() *CreateCodeBundleResponseBodyUpload
}

type CreateCodeBundleResponseBody struct {
	// The version identifier of the function code package.
	//
	// example:
	//
	// 1
	BundleVersion *string `json:"bundleVersion,omitempty" xml:"bundleVersion,omitempty"`
	// The function code package ID.
	//
	// example:
	//
	// 111
	CodeBundleId *int64 `json:"codeBundleId,omitempty" xml:"codeBundleId,omitempty"`
	// The time when the function code package was created.
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
	// 123
	ProjectId *int64 `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A1F403F-0A85-5578-8B7C-55E3E9408659
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The status of the function code package.
	//
	// example:
	//
	// pending
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the function code package was last updated.
	//
	// example:
	//
	// 2026-08-27T00:53:46.774Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The upload credential. See the following fields for details.
	Upload *CreateCodeBundleResponseBodyUpload `json:"upload,omitempty" xml:"upload,omitempty" type:"Struct"`
}

func (s CreateCodeBundleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCodeBundleResponseBody) GetBundleVersion() *string {
	return s.BundleVersion
}

func (s *CreateCodeBundleResponseBody) GetCodeBundleId() *int64 {
	return s.CodeBundleId
}

func (s *CreateCodeBundleResponseBody) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateCodeBundleResponseBody) GetFilename() *string {
	return s.Filename
}

func (s *CreateCodeBundleResponseBody) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateCodeBundleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCodeBundleResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateCodeBundleResponseBody) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateCodeBundleResponseBody) GetUpload() *CreateCodeBundleResponseBodyUpload {
	return s.Upload
}

func (s *CreateCodeBundleResponseBody) SetBundleVersion(v string) *CreateCodeBundleResponseBody {
	s.BundleVersion = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetCodeBundleId(v int64) *CreateCodeBundleResponseBody {
	s.CodeBundleId = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetCreatedAt(v string) *CreateCodeBundleResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetFilename(v string) *CreateCodeBundleResponseBody {
	s.Filename = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetProjectId(v int64) *CreateCodeBundleResponseBody {
	s.ProjectId = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetRequestId(v string) *CreateCodeBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetStatus(v string) *CreateCodeBundleResponseBody {
	s.Status = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetUpdatedAt(v string) *CreateCodeBundleResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateCodeBundleResponseBody) SetUpload(v *CreateCodeBundleResponseBodyUpload) *CreateCodeBundleResponseBody {
	s.Upload = v
	return s
}

func (s *CreateCodeBundleResponseBody) Validate() error {
	if s.Upload != nil {
		if err := s.Upload.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCodeBundleResponseBodyUpload struct {
	// The expiration time of the credential in RFC 3339 format.
	//
	// example:
	//
	// 2026-09-03T05:58:47.88987539Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The HTTP method of the pre-signed URL. Valid values: PUT.
	//
	// example:
	//
	// PUT
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The pre-signed OSS PUT upload URL.
	//
	// example:
	//
	// https://codesec-beijing.oss-cn-beijing.aliyuncs.com/87766767%2F1001667%2F1004171.zip
	PutUrl *string `json:"putUrl,omitempty" xml:"putUrl,omitempty"`
	// The Content-Type header that the client must include when performing the PUT request. This field is returned when putUrl is present.
	//
	// example:
	//
	// application/octet-stream
	RequiredContentType *string `json:"requiredContentType,omitempty" xml:"requiredContentType,omitempty"`
}

func (s CreateCodeBundleResponseBodyUpload) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeBundleResponseBodyUpload) GoString() string {
	return s.String()
}

func (s *CreateCodeBundleResponseBodyUpload) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *CreateCodeBundleResponseBodyUpload) GetMethod() *string {
	return s.Method
}

func (s *CreateCodeBundleResponseBodyUpload) GetPutUrl() *string {
	return s.PutUrl
}

func (s *CreateCodeBundleResponseBodyUpload) GetRequiredContentType() *string {
	return s.RequiredContentType
}

func (s *CreateCodeBundleResponseBodyUpload) SetExpiresAt(v string) *CreateCodeBundleResponseBodyUpload {
	s.ExpiresAt = &v
	return s
}

func (s *CreateCodeBundleResponseBodyUpload) SetMethod(v string) *CreateCodeBundleResponseBodyUpload {
	s.Method = &v
	return s
}

func (s *CreateCodeBundleResponseBodyUpload) SetPutUrl(v string) *CreateCodeBundleResponseBodyUpload {
	s.PutUrl = &v
	return s
}

func (s *CreateCodeBundleResponseBodyUpload) SetRequiredContentType(v string) *CreateCodeBundleResponseBodyUpload {
	s.RequiredContentType = &v
	return s
}

func (s *CreateCodeBundleResponseBodyUpload) Validate() error {
	return dara.Validate(s)
}
