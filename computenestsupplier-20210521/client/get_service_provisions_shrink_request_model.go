// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvisionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParametersShrink(v string) *GetServiceProvisionsShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *GetServiceProvisionsShrinkRequest
	GetRegionId() *string
	SetServiceId(v string) *GetServiceProvisionsShrinkRequest
	GetServiceId() *string
	SetServiceVersion(v string) *GetServiceProvisionsShrinkRequest
	GetServiceVersion() *string
	SetTemplateName(v string) *GetServiceProvisionsShrinkRequest
	GetTemplateName() *string
}

type GetServiceProvisionsShrinkRequest struct {
	// The parameters that are specified to deploy the service instance.
	//
	// example:
	//
	// {\\"RegionId\\":\\"cn-hangzhou\\"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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
	// service-20b8a396048346xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The template name.
	//
	// example:
	//
	// 模板1
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s GetServiceProvisionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvisionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetServiceProvisionsShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *GetServiceProvisionsShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceProvisionsShrinkRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceProvisionsShrinkRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetServiceProvisionsShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetServiceProvisionsShrinkRequest) SetParametersShrink(v string) *GetServiceProvisionsShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetRegionId(v string) *GetServiceProvisionsShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceId(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceId = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetServiceVersion(v string) *GetServiceProvisionsShrinkRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) SetTemplateName(v string) *GetServiceProvisionsShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetServiceProvisionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
