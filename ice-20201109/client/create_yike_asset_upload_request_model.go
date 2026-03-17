// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeAssetUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileExt(v string) *CreateYikeAssetUploadRequest
	GetFileExt() *string
	SetFileType(v string) *CreateYikeAssetUploadRequest
	GetFileType() *string
}

type CreateYikeAssetUploadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mp4
	FileExt *string `json:"FileExt,omitempty" xml:"FileExt,omitempty"`
	// example:
	//
	// StoryboardInput
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s CreateYikeAssetUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeAssetUploadRequest) GoString() string {
	return s.String()
}

func (s *CreateYikeAssetUploadRequest) GetFileExt() *string {
	return s.FileExt
}

func (s *CreateYikeAssetUploadRequest) GetFileType() *string {
	return s.FileType
}

func (s *CreateYikeAssetUploadRequest) SetFileExt(v string) *CreateYikeAssetUploadRequest {
	s.FileExt = &v
	return s
}

func (s *CreateYikeAssetUploadRequest) SetFileType(v string) *CreateYikeAssetUploadRequest {
	s.FileType = &v
	return s
}

func (s *CreateYikeAssetUploadRequest) Validate() error {
	return dara.Validate(s)
}
