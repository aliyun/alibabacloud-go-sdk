// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJenkinsImageRegistryNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegistryId(v int64) *UpdateJenkinsImageRegistryNameRequest
	GetRegistryId() *int64
	SetRegistryName(v string) *UpdateJenkinsImageRegistryNameRequest
	GetRegistryName() *string
	SetSourceIp(v string) *UpdateJenkinsImageRegistryNameRequest
	GetSourceIp() *string
}

type UpdateJenkinsImageRegistryNameRequest struct {
	// The image repository ID.
	//
	// >You can call the [PageImageRegistry](~~PageImageRegistry~~) operation to obtain this parameter.
	//
	// example:
	//
	// 25090
	RegistryId *int64 `json:"RegistryId,omitempty" xml:"RegistryId,omitempty"`
	// The image repository name.
	//
	// example:
	//
	// a0603tk1
	RegistryName *string `json:"RegistryName,omitempty" xml:"RegistryName,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 121.33.XXX.XXX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s UpdateJenkinsImageRegistryNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJenkinsImageRegistryNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateJenkinsImageRegistryNameRequest) GetRegistryId() *int64 {
	return s.RegistryId
}

func (s *UpdateJenkinsImageRegistryNameRequest) GetRegistryName() *string {
	return s.RegistryName
}

func (s *UpdateJenkinsImageRegistryNameRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *UpdateJenkinsImageRegistryNameRequest) SetRegistryId(v int64) *UpdateJenkinsImageRegistryNameRequest {
	s.RegistryId = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameRequest) SetRegistryName(v string) *UpdateJenkinsImageRegistryNameRequest {
	s.RegistryName = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameRequest) SetSourceIp(v string) *UpdateJenkinsImageRegistryNameRequest {
	s.SourceIp = &v
	return s
}

func (s *UpdateJenkinsImageRegistryNameRequest) Validate() error {
	return dara.Validate(s)
}
