// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*UpdateRenderingInstanceConfigurationRequestConfiguration) *UpdateRenderingInstanceConfigurationRequest
	GetConfiguration() []*UpdateRenderingInstanceConfigurationRequestConfiguration
	SetRenderingInstanceId(v string) *UpdateRenderingInstanceConfigurationRequest
	GetRenderingInstanceId() *string
}

type UpdateRenderingInstanceConfigurationRequest struct {
	// This parameter is required.
	Configuration []*UpdateRenderingInstanceConfigurationRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationRequest) GetConfiguration() []*UpdateRenderingInstanceConfigurationRequestConfiguration {
	return s.Configuration
}

func (s *UpdateRenderingInstanceConfigurationRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *UpdateRenderingInstanceConfigurationRequest) SetConfiguration(v []*UpdateRenderingInstanceConfigurationRequestConfiguration) *UpdateRenderingInstanceConfigurationRequest {
	s.Configuration = v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequest) SetRenderingInstanceId(v string) *UpdateRenderingInstanceConfigurationRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateRenderingInstanceConfigurationRequestConfiguration struct {
	// This parameter is required.
	Attributes []*UpdateRenderingInstanceConfigurationRequestConfigurationAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// location
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationRequestConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationRequestConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationRequestConfiguration) GetAttributes() []*UpdateRenderingInstanceConfigurationRequestConfigurationAttributes {
	return s.Attributes
}

func (s *UpdateRenderingInstanceConfigurationRequestConfiguration) GetModuleName() *string {
	return s.ModuleName
}

func (s *UpdateRenderingInstanceConfigurationRequestConfiguration) SetAttributes(v []*UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) *UpdateRenderingInstanceConfigurationRequestConfiguration {
	s.Attributes = v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequestConfiguration) SetModuleName(v string) *UpdateRenderingInstanceConfigurationRequestConfiguration {
	s.ModuleName = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequestConfiguration) Validate() error {
	return dara.Validate(s)
}

type UpdateRenderingInstanceConfigurationRequestConfigurationAttributes struct {
	// This parameter is required.
	//
	// example:
	//
	// lon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) GetName() *string {
	return s.Name
}

func (s *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) GetValue() interface{} {
	return s.Value
}

func (s *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) SetName(v string) *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes {
	s.Name = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) SetValue(v interface{}) *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes {
	s.Value = v
	return s
}

func (s *UpdateRenderingInstanceConfigurationRequestConfigurationAttributes) Validate() error {
	return dara.Validate(s)
}
