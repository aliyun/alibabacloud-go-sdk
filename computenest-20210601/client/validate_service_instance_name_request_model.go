// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateServiceInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ValidateServiceInstanceNameRequest
	GetClientToken() *string
	SetIsTrial(v bool) *ValidateServiceInstanceNameRequest
	GetIsTrial() *bool
	SetServiceId(v string) *ValidateServiceInstanceNameRequest
	GetServiceId() *string
	SetServiceInstanceName(v string) *ValidateServiceInstanceNameRequest
	GetServiceInstanceName() *string
	SetServiceVersion(v string) *ValidateServiceInstanceNameRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *ValidateServiceInstanceNameRequest
	GetTemplateName() *string
}

type ValidateServiceInstanceNameRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	IsTrial *bool `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// service-12xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testName
	ServiceInstanceName *string `json:"ServiceInstanceName,omitempty" xml:"ServiceInstanceName,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// 模板一
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ValidateServiceInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateServiceInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ValidateServiceInstanceNameRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateServiceInstanceNameRequest) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *ValidateServiceInstanceNameRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ValidateServiceInstanceNameRequest) GetServiceInstanceName() *string {
	return s.ServiceInstanceName
}

func (s *ValidateServiceInstanceNameRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *ValidateServiceInstanceNameRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ValidateServiceInstanceNameRequest) SetClientToken(v string) *ValidateServiceInstanceNameRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) SetIsTrial(v bool) *ValidateServiceInstanceNameRequest {
	s.IsTrial = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) SetServiceId(v string) *ValidateServiceInstanceNameRequest {
	s.ServiceId = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) SetServiceInstanceName(v string) *ValidateServiceInstanceNameRequest {
	s.ServiceInstanceName = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) SetServiceVersion(v string) *ValidateServiceInstanceNameRequest {
	s.ServiceVersion = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) SetTemplateName(v string) *ValidateServiceInstanceNameRequest {
	s.TemplateName = &v
	return s
}

func (s *ValidateServiceInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
