// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseVersionImage interface {
	dara.Model
	String() string
	GoString() string
	SetCpuArchitecture(v string) *ReleaseVersionImage
	GetCpuArchitecture() *string
	SetImageId(v string) *ReleaseVersionImage
	GetImageId() *string
	SetRuntimeEngineType(v string) *ReleaseVersionImage
	GetRuntimeEngineType() *string
}

type ReleaseVersionImage struct {
	CpuArchitecture   *string `json:"cpuArchitecture,omitempty" xml:"cpuArchitecture,omitempty"`
	ImageId           *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	RuntimeEngineType *string `json:"runtimeEngineType,omitempty" xml:"runtimeEngineType,omitempty"`
}

func (s ReleaseVersionImage) String() string {
	return dara.Prettify(s)
}

func (s ReleaseVersionImage) GoString() string {
	return s.String()
}

func (s *ReleaseVersionImage) GetCpuArchitecture() *string {
	return s.CpuArchitecture
}

func (s *ReleaseVersionImage) GetImageId() *string {
	return s.ImageId
}

func (s *ReleaseVersionImage) GetRuntimeEngineType() *string {
	return s.RuntimeEngineType
}

func (s *ReleaseVersionImage) SetCpuArchitecture(v string) *ReleaseVersionImage {
	s.CpuArchitecture = &v
	return s
}

func (s *ReleaseVersionImage) SetImageId(v string) *ReleaseVersionImage {
	s.ImageId = &v
	return s
}

func (s *ReleaseVersionImage) SetRuntimeEngineType(v string) *ReleaseVersionImage {
	s.RuntimeEngineType = &v
	return s
}

func (s *ReleaseVersionImage) Validate() error {
	return dara.Validate(s)
}
