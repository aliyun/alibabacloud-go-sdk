// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOssUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizConfig(v string) *CreateOssUploadRequest
	GetBizConfig() *string
	SetCredentialType(v string) *CreateOssUploadRequest
	GetCredentialType() *string
	SetFileExt(v string) *CreateOssUploadRequest
	GetFileExt() *string
	SetFileType(v string) *CreateOssUploadRequest
	GetFileType() *string
	SetTargetBucket(v string) *CreateOssUploadRequest
	GetTargetBucket() *string
}

type CreateOssUploadRequest struct {
	// The **business configuration JSON**, which includes ProductionId and FolderId.
	//
	// example:
	//
	// {"ProductionId": "pd_***", "FolderId": "fd_***"}
	BizConfig *string `json:"BizConfig,omitempty" xml:"BizConfig,omitempty"`
	// The credential type. This is a reserved parameter. The current version always returns STS temporary credentials, and this parameter does not take effect.
	//
	// example:
	//
	// STS
	CredentialType *string `json:"CredentialType,omitempty" xml:"CredentialType,omitempty"`
	// The file name extension, such as mp4, wav, or png. Do not include the dot. If this parameter is missing, `MissingParameter` is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// mp4
	FileExt *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	// The file type. Valid values: `Media` (default), `StoryboardInput`, `ImageFusion`, `AIAppParams`, `Workflow`, and `CommonFile`. If an invalid value is specified, `InvalidParameter` is returned. This value is used as a path segment in the `ObjectKey` and is converted to lowercase.
	//
	// example:
	//
	// Media
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The explicitly specified target bucket.
	//
	// example:
	//
	// test-bucket
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
}

func (s CreateOssUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOssUploadRequest) GoString() string {
	return s.String()
}

func (s *CreateOssUploadRequest) GetBizConfig() *string {
	return s.BizConfig
}

func (s *CreateOssUploadRequest) GetCredentialType() *string {
	return s.CredentialType
}

func (s *CreateOssUploadRequest) GetFileExt() *string {
	return s.FileExt
}

func (s *CreateOssUploadRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateOssUploadRequest) GetTargetBucket() *string {
	return s.TargetBucket
}

func (s *CreateOssUploadRequest) SetBizConfig(v string) *CreateOssUploadRequest {
	s.BizConfig = &v
	return s
}

func (s *CreateOssUploadRequest) SetCredentialType(v string) *CreateOssUploadRequest {
	s.CredentialType = &v
	return s
}

func (s *CreateOssUploadRequest) SetFileExt(v string) *CreateOssUploadRequest {
	s.FileExt = &v
	return s
}

func (s *CreateOssUploadRequest) SetFileType(v string) *CreateOssUploadRequest {
	s.FileType = &v
	return s
}

func (s *CreateOssUploadRequest) SetTargetBucket(v string) *CreateOssUploadRequest {
	s.TargetBucket = &v
	return s
}

func (s *CreateOssUploadRequest) Validate() error {
	return dara.Validate(s)
}
