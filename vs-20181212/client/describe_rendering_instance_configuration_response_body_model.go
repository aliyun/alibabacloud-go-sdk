// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*DescribeRenderingInstanceConfigurationResponseBodyConfiguration) *DescribeRenderingInstanceConfigurationResponseBody
	GetConfiguration() []*DescribeRenderingInstanceConfigurationResponseBodyConfiguration
	SetRequestId(v string) *DescribeRenderingInstanceConfigurationResponseBody
	GetRequestId() *string
}

type DescribeRenderingInstanceConfigurationResponseBody struct {
	// The configuration.
	Configuration []*DescribeRenderingInstanceConfigurationResponseBodyConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationResponseBody) GetConfiguration() []*DescribeRenderingInstanceConfigurationResponseBodyConfiguration {
	return s.Configuration
}

func (s *DescribeRenderingInstanceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenderingInstanceConfigurationResponseBody) SetConfiguration(v []*DescribeRenderingInstanceConfigurationResponseBodyConfiguration) *DescribeRenderingInstanceConfigurationResponseBody {
	s.Configuration = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBody) SetRequestId(v string) *DescribeRenderingInstanceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBody) Validate() error {
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

type DescribeRenderingInstanceConfigurationResponseBodyConfiguration struct {
	// The list of properties for the module.
	Attributes []*DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes `json:"Attributes,omitempty" xml:"Attributes,omitempty" type:"Repeated"`
	// The name of the device simulation feature module. Valid values:
	//
	// 1. ctl: The control module.
	//
	// 2. prop: The property module.
	//
	// 3. location: The location module.
	//
	// 4. battery: The battery module.
	//
	// 5. network: The network module.
	//
	// 6. bluetooth: The Bluetooth module.
	//
	// 7. sim: The SIM card module.
	//
	// 8. display: The device module.
	//
	// 9. system: The basic module.
	//
	// example:
	//
	// location
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationResponseBodyConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationResponseBodyConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfiguration) GetAttributes() []*DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes {
	return s.Attributes
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfiguration) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfiguration) SetAttributes(v []*DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) *DescribeRenderingInstanceConfigurationResponseBodyConfiguration {
	s.Attributes = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfiguration) SetModuleName(v string) *DescribeRenderingInstanceConfigurationResponseBodyConfiguration {
	s.ModuleName = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfiguration) Validate() error {
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

type DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes struct {
	// The name of the attribute.
	//
	// example:
	//
	// lon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the attribute.
	//
	// example:
	//
	// 100
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) GetName() *string {
	return s.Name
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) GetValue() interface{} {
	return s.Value
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) SetName(v string) *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes {
	s.Name = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) SetValue(v interface{}) *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes {
	s.Value = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationResponseBodyConfigurationAttributes) Validate() error {
	return dara.Validate(s)
}
