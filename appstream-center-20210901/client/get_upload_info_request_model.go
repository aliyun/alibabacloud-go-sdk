// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizRegionId(v string) *GetUploadInfoRequest
	GetBizRegionId() *string
	SetFileType(v string) *GetUploadInfoRequest
	GetFileType() *string
}

type GetUploadInfoRequest struct {
	// The ID of the business region where the file is uploaded. If this parameter is not specified, the default OSS region configured for the service is used.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The type of the file to upload. Currently, only TASK is supported, which indicates an image task file.
	//
	// example:
	//
	// TASK
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GetUploadInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUploadInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUploadInfoRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *GetUploadInfoRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetUploadInfoRequest) SetBizRegionId(v string) *GetUploadInfoRequest {
	s.BizRegionId = &v
	return s
}

func (s *GetUploadInfoRequest) SetFileType(v string) *GetUploadInfoRequest {
	s.FileType = &v
	return s
}

func (s *GetUploadInfoRequest) Validate() error {
	return dara.Validate(s)
}
