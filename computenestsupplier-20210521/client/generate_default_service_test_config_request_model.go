// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateDefaultServiceTestConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceId(v string) *GenerateDefaultServiceTestConfigRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GenerateDefaultServiceTestConfigRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GenerateDefaultServiceTestConfigRequest
	GetTemplateName() *string
}

type GenerateDefaultServiceTestConfigRequest struct {
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-fd1aec438c974828bb7b
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// ECS standalone
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GenerateDefaultServiceTestConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateDefaultServiceTestConfigRequest) GoString() string {
	return s.String()
}

func (s *GenerateDefaultServiceTestConfigRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GenerateDefaultServiceTestConfigRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GenerateDefaultServiceTestConfigRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GenerateDefaultServiceTestConfigRequest) SetServiceId(v string) *GenerateDefaultServiceTestConfigRequest {
	s.ServiceId = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigRequest) SetServiceVersion(v string) *GenerateDefaultServiceTestConfigRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigRequest) SetTemplateName(v string) *GenerateDefaultServiceTestConfigRequest {
	s.TemplateName = &v
	return s
}

func (s *GenerateDefaultServiceTestConfigRequest) Validate() error {
	return dara.Validate(s)
}
