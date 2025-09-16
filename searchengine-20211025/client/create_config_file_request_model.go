// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *CreateConfigFileRequest
	GetFileName() *string
	SetOssPath(v string) *CreateConfigFileRequest
	GetOssPath() *string
	SetParentFullPath(v string) *CreateConfigFileRequest
	GetParentFullPath() *string
}

type CreateConfigFileRequest struct {
	// The name of the directory.
	//
	// example:
	//
	// /schemas/device_event_xt_schema.json
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The Object Storage Service (OSS) URL of the file.
	//
	// example:
	//
	// oss://xxx/xxxx/xxx
	OssPath *string `json:"ossPath,omitempty" xml:"ossPath,omitempty"`
	// The path of the parent directory.
	//
	// example:
	//
	// /
	ParentFullPath *string `json:"parentFullPath,omitempty" xml:"parentFullPath,omitempty"`
}

func (s CreateConfigFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigFileRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateConfigFileRequest) GetOssPath() *string {
	return s.OssPath
}

func (s *CreateConfigFileRequest) GetParentFullPath() *string {
	return s.ParentFullPath
}

func (s *CreateConfigFileRequest) SetFileName(v string) *CreateConfigFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateConfigFileRequest) SetOssPath(v string) *CreateConfigFileRequest {
	s.OssPath = &v
	return s
}

func (s *CreateConfigFileRequest) SetParentFullPath(v string) *CreateConfigFileRequest {
	s.ParentFullPath = &v
	return s
}

func (s *CreateConfigFileRequest) Validate() error {
	return dara.Validate(s)
}
