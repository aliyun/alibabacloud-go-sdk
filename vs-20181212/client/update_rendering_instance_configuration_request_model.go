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
	// The configuration content.
	//
	// This parameter is required.
	Configuration []*UpdateRenderingInstanceConfigurationRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// The instance ID of the cloud application service.
	//
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
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRenderingInstanceConfigurationRequestConfiguration struct {
	// The list of attributes corresponding to the module.
	//
	// This parameter is required.
	Attributes []*UpdateRenderingInstanceConfigurationRequestConfigurationAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the real-device simulation feature module. Valid values:
	//
	// - ctl: Control module.
	//
	// - prop: Property module.
	//
	// - location: Location module.
	//
	// - battery: Battery module.
	//
	// - network: Network module.
	//
	// - bluetooth: Bluetooth module.
	//
	// - sim: SIM card module.
	//
	// - display: Device module.
	//
	// - system: Basic module.
	//
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
	if s.Attributes != nil {
		for _, item := range s.Attributes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateRenderingInstanceConfigurationRequestConfigurationAttributes struct {
	// The name of the attribute.
	//
	// This parameter is required.
	//
	// example:
	//
	// lon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
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
