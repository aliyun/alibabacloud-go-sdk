// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSoftwarelibVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMd5(v string) *CreateSoftwarelibVersionRequest
	GetMd5() *string
	SetOs(v string) *CreateSoftwarelibVersionRequest
	GetOs() *string
	SetPublisherType(v string) *CreateSoftwarelibVersionRequest
	GetPublisherType() *string
	SetSoftwareId(v string) *CreateSoftwarelibVersionRequest
	GetSoftwareId() *string
	SetSoftwareName(v string) *CreateSoftwarelibVersionRequest
	GetSoftwareName() *string
	SetSoftwarePkgName(v string) *CreateSoftwarelibVersionRequest
	GetSoftwarePkgName() *string
	SetSoftwarePkgSize(v int64) *CreateSoftwarelibVersionRequest
	GetSoftwarePkgSize() *int64
	SetSoftwareUrl(v string) *CreateSoftwarelibVersionRequest
	GetSoftwareUrl() *string
	SetSoftwareVersion(v string) *CreateSoftwarelibVersionRequest
	GetSoftwareVersion() *string
}

type CreateSoftwarelibVersionRequest struct {
	// The MD5 value of the software package. The value can be up to 64 characters in length.
	//
	// example:
	//
	// 0b5824cdd509d3ed560e2d20d29a1bcb
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The operating system to which the software package applies. Valid values:
	//
	// - **Windows**: Windows.
	//
	// - **Mac(Apple)**: macOS with Apple silicon.
	//
	// - **Mac(Intel)**: macOS with Intel processors.
	//
	// example:
	//
	// Windows
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// The software publisher type. Valid values:
	//
	// - **local**: local upload.
	//
	// - **thirdparty**: third-party link.
	//
	// example:
	//
	// local
	PublisherType *string `json:"PublisherType,omitempty" xml:"PublisherType,omitempty"`
	// The software ID in the software library. The value can be up to 64 characters in length. You can call [ListSoftwarelibSoftware](~~ListSoftwarelibSoftware~~) to obtain the value.
	//
	// This parameter is required.
	//
	// example:
	//
	// softwarelib-software-2c51808a3cc8****
	SoftwareId *string `json:"SoftwareId,omitempty" xml:"SoftwareId,omitempty"`
	// The software name. The value can be up to 128 characters in length.
	//
	// example:
	//
	// Thunder
	SoftwareName *string `json:"SoftwareName,omitempty" xml:"SoftwareName,omitempty"`
	// The file name of the software package. The value can be up to 128 characters in length.
	//
	// example:
	//
	// TestSoftware.exe
	SoftwarePkgName *string `json:"SoftwarePkgName,omitempty" xml:"SoftwarePkgName,omitempty"`
	// The size of the software package.
	//
	// example:
	//
	// 1000
	SoftwarePkgSize *int64 `json:"SoftwarePkgSize,omitempty" xml:"SoftwarePkgSize,omitempty"`
	// The download URL of the software package. If the publisher type is local, the value is the relative path of the software package in the OSS bucket. If the publisher type is thirdparty, the value is a third-party download URL.
	SoftwareUrl *string `json:"SoftwareUrl,omitempty" xml:"SoftwareUrl,omitempty"`
	// The software version number. The value can be up to 64 characters in length. The combination of operating system and version number must be unique within the same software. If a duplicate exists, a ResourceDuplicated error is returned.
	//
	// example:
	//
	// 1.0
	SoftwareVersion *string `json:"SoftwareVersion,omitempty" xml:"SoftwareVersion,omitempty"`
}

func (s CreateSoftwarelibVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSoftwarelibVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateSoftwarelibVersionRequest) GetMd5() *string {
	return s.Md5
}

func (s *CreateSoftwarelibVersionRequest) GetOs() *string {
	return s.Os
}

func (s *CreateSoftwarelibVersionRequest) GetPublisherType() *string {
	return s.PublisherType
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwareId() *string {
	return s.SoftwareId
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwareName() *string {
	return s.SoftwareName
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwarePkgName() *string {
	return s.SoftwarePkgName
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwarePkgSize() *int64 {
	return s.SoftwarePkgSize
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwareUrl() *string {
	return s.SoftwareUrl
}

func (s *CreateSoftwarelibVersionRequest) GetSoftwareVersion() *string {
	return s.SoftwareVersion
}

func (s *CreateSoftwarelibVersionRequest) SetMd5(v string) *CreateSoftwarelibVersionRequest {
	s.Md5 = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetOs(v string) *CreateSoftwarelibVersionRequest {
	s.Os = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetPublisherType(v string) *CreateSoftwarelibVersionRequest {
	s.PublisherType = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwareId(v string) *CreateSoftwarelibVersionRequest {
	s.SoftwareId = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwareName(v string) *CreateSoftwarelibVersionRequest {
	s.SoftwareName = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwarePkgName(v string) *CreateSoftwarelibVersionRequest {
	s.SoftwarePkgName = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwarePkgSize(v int64) *CreateSoftwarelibVersionRequest {
	s.SoftwarePkgSize = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwareUrl(v string) *CreateSoftwarelibVersionRequest {
	s.SoftwareUrl = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) SetSoftwareVersion(v string) *CreateSoftwarelibVersionRequest {
	s.SoftwareVersion = &v
	return s
}

func (s *CreateSoftwarelibVersionRequest) Validate() error {
	return dara.Validate(s)
}
