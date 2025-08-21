// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*DeleteRenderingInstanceConfigurationRequestConfiguration) *DeleteRenderingInstanceConfigurationRequest
	GetConfiguration() []*DeleteRenderingInstanceConfigurationRequestConfiguration
	SetRenderingInstanceId(v string) *DeleteRenderingInstanceConfigurationRequest
	GetRenderingInstanceId() *string
}

type DeleteRenderingInstanceConfigurationRequest struct {
	Configuration []*DeleteRenderingInstanceConfigurationRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s DeleteRenderingInstanceConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceConfigurationRequest) GetConfiguration() []*DeleteRenderingInstanceConfigurationRequestConfiguration {
	return s.Configuration
}

func (s *DeleteRenderingInstanceConfigurationRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *DeleteRenderingInstanceConfigurationRequest) SetConfiguration(v []*DeleteRenderingInstanceConfigurationRequestConfiguration) *DeleteRenderingInstanceConfigurationRequest {
	s.Configuration = v
	return s
}

func (s *DeleteRenderingInstanceConfigurationRequest) SetRenderingInstanceId(v string) *DeleteRenderingInstanceConfigurationRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationRequest) Validate() error {
	return dara.Validate(s)
}

type DeleteRenderingInstanceConfigurationRequestConfiguration struct {
	AttributeNames []*string `json:"AttributeNames,omitempty" xml:"AttributeNames,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// location
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
}

func (s DeleteRenderingInstanceConfigurationRequestConfiguration) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceConfigurationRequestConfiguration) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceConfigurationRequestConfiguration) GetAttributeNames() []*string {
	return s.AttributeNames
}

func (s *DeleteRenderingInstanceConfigurationRequestConfiguration) GetModuleName() *string {
	return s.ModuleName
}

func (s *DeleteRenderingInstanceConfigurationRequestConfiguration) SetAttributeNames(v []*string) *DeleteRenderingInstanceConfigurationRequestConfiguration {
	s.AttributeNames = v
	return s
}

func (s *DeleteRenderingInstanceConfigurationRequestConfiguration) SetModuleName(v string) *DeleteRenderingInstanceConfigurationRequestConfiguration {
	s.ModuleName = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationRequestConfiguration) Validate() error {
	return dara.Validate(s)
}
