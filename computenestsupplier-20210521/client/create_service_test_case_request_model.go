// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceTestCaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateServiceTestCaseRequest
	GetRegionId() *string
	SetServiceId(v string) *CreateServiceTestCaseRequest
	GetServiceId() *string
	SetServiceVersion(v string) *CreateServiceTestCaseRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *CreateServiceTestCaseRequest
	GetTemplateName() *string
	SetTestCaseName(v string) *CreateServiceTestCaseRequest
	GetTestCaseName() *string
	SetTestConfig(v string) *CreateServiceTestCaseRequest
	GetTestConfig() *string
}

type CreateServiceTestCaseRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_Image_Ecs
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Service Test case name.
	//
	// This parameter is required.
	//
	// example:
	//
	// case1
	TestCaseName *string `json:"TestCaseName,omitempty" xml:"TestCaseName,omitempty"`
	// The service test config
	//
	// This parameter is required.
	//
	// example:
	//
	// ---
	//
	// parameters:
	//
	//   PayType: "PostPaid"
	//
	//   EcsInstanceType: "$[iact3-auto]"
	//
	//   InstancePassword: "$[iact3-auto]"
	TestConfig *string `json:"TestConfig,omitempty" xml:"TestConfig,omitempty"`
}

func (s CreateServiceTestCaseRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceTestCaseRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceTestCaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateServiceTestCaseRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CreateServiceTestCaseRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CreateServiceTestCaseRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateServiceTestCaseRequest) GetTestCaseName() *string {
	return s.TestCaseName
}

func (s *CreateServiceTestCaseRequest) GetTestConfig() *string {
	return s.TestConfig
}

func (s *CreateServiceTestCaseRequest) SetRegionId(v string) *CreateServiceTestCaseRequest {
	s.RegionId = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetServiceId(v string) *CreateServiceTestCaseRequest {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetServiceVersion(v string) *CreateServiceTestCaseRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTemplateName(v string) *CreateServiceTestCaseRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTestCaseName(v string) *CreateServiceTestCaseRequest {
	s.TestCaseName = &v
	return s
}

func (s *CreateServiceTestCaseRequest) SetTestConfig(v string) *CreateServiceTestCaseRequest {
	s.TestConfig = &v
	return s
}

func (s *CreateServiceTestCaseRequest) Validate() error {
	return dara.Validate(s)
}
