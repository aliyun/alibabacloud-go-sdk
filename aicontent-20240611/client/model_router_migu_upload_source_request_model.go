// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterMiguUploadSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileType(v string) *ModelRouterMiguUploadSourceRequest
	GetFileType() *string
	SetServiceName(v string) *ModelRouterMiguUploadSourceRequest
	GetServiceName() *string
}

type ModelRouterMiguUploadSourceRequest struct {
	// The source file type. Valid values: VIDEO, IMAGE, AUDIO, and TEXT.
	//
	// This parameter is required.
	//
	// example:
	//
	// VIDEO
	FileType *string `json:"fileType,omitempty" xml:"fileType,omitempty"`
	// The business service name, such as kling, vidu, or wonder.
	//
	// This parameter is required.
	//
	// example:
	//
	// kling
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ModelRouterMiguUploadSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterMiguUploadSourceRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterMiguUploadSourceRequest) GetFileType() *string {
	return s.FileType
}

func (s *ModelRouterMiguUploadSourceRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ModelRouterMiguUploadSourceRequest) SetFileType(v string) *ModelRouterMiguUploadSourceRequest {
	s.FileType = &v
	return s
}

func (s *ModelRouterMiguUploadSourceRequest) SetServiceName(v string) *ModelRouterMiguUploadSourceRequest {
	s.ServiceName = &v
	return s
}

func (s *ModelRouterMiguUploadSourceRequest) Validate() error {
	return dara.Validate(s)
}
