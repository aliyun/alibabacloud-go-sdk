// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*DescribeRenderingInstanceConfigurationRequestConfiguration) *DescribeRenderingInstanceConfigurationRequest
	GetConfiguration() []*DescribeRenderingInstanceConfigurationRequestConfiguration
	SetRenderingInstanceId(v string) *DescribeRenderingInstanceConfigurationRequest
	GetRenderingInstanceId() *string
}

type DescribeRenderingInstanceConfigurationRequest struct {
	// The configuration to query. By default, the configurations of all modules are returned.
	Configuration []*DescribeRenderingInstanceConfigurationRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// The ID of the cloud application service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationRequest) GetConfiguration() []*DescribeRenderingInstanceConfigurationRequestConfiguration {
	return s.Configuration
}

func (s *DescribeRenderingInstanceConfigurationRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DescribeRenderingInstanceConfigurationRequest) SetConfiguration(v []*DescribeRenderingInstanceConfigurationRequestConfiguration) *DescribeRenderingInstanceConfigurationRequest {
	s.Configuration = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationRequest) SetRenderingInstanceId(v string) *DescribeRenderingInstanceConfigurationRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationRequest) Validate() error {
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

type DescribeRenderingInstanceConfigurationRequestConfiguration struct {
	// The list of attribute names to query. By default, all attributes are returned.
	AttributeNames []*string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty" type:"Repeated"`
	// The name of the module.
	//
	// This parameter is required.
	//
	// example:
	//
	// location
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DescribeRenderingInstanceConfigurationRequestConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceConfigurationRequestConfiguration) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceConfigurationRequestConfiguration) GetAttributeNames() []*string {
	return s.AttributeNames
}

func (s *DescribeRenderingInstanceConfigurationRequestConfiguration) GetModuleName() *string {
	return s.ModuleName
}

func (s *DescribeRenderingInstanceConfigurationRequestConfiguration) SetAttributeNames(v []*string) *DescribeRenderingInstanceConfigurationRequestConfiguration {
	s.AttributeNames = v
	return s
}

func (s *DescribeRenderingInstanceConfigurationRequestConfiguration) SetModuleName(v string) *DescribeRenderingInstanceConfigurationRequestConfiguration {
	s.ModuleName = &v
	return s
}

func (s *DescribeRenderingInstanceConfigurationRequestConfiguration) Validate() error {
	return dara.Validate(s)
}
