// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageDescription(v string) *ImportImageRequest
	GetImageDescription() *string
	SetImageFileURL(v string) *ImportImageRequest
	GetImageFileURL() *string
	SetImageName(v string) *ImportImageRequest
	GetImageName() *string
}

type ImportImageRequest struct {
	// example:
	//
	// android 12 custom image
	ImageDescription *string `json:"ImageDescription,omitempty" xml:"ImageDescription,omitempty"`
	// example:
	//
	// https://xxx.oss-xxx/xxxx.tgz
	ImageFileURL *string `json:"ImageFileURL,omitempty" xml:"ImageFileURL,omitempty"`
	// example:
	//
	// import custom image
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s ImportImageRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportImageRequest) GoString() string {
	return s.String()
}

func (s *ImportImageRequest) GetImageDescription() *string {
	return s.ImageDescription
}

func (s *ImportImageRequest) GetImageFileURL() *string {
	return s.ImageFileURL
}

func (s *ImportImageRequest) GetImageName() *string {
	return s.ImageName
}

func (s *ImportImageRequest) SetImageDescription(v string) *ImportImageRequest {
	s.ImageDescription = &v
	return s
}

func (s *ImportImageRequest) SetImageFileURL(v string) *ImportImageRequest {
	s.ImageFileURL = &v
	return s
}

func (s *ImportImageRequest) SetImageName(v string) *ImportImageRequest {
	s.ImageName = &v
	return s
}

func (s *ImportImageRequest) Validate() error {
	return dara.Validate(s)
}
