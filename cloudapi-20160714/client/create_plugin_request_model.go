// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePluginRequest
	GetDescription() *string
	SetPluginData(v string) *CreatePluginRequest
	GetPluginData() *string
	SetPluginName(v string) *CreatePluginRequest
	GetPluginName() *string
	SetPluginType(v string) *CreatePluginRequest
	GetPluginType() *string
	SetSecurityToken(v string) *CreatePluginRequest
	GetSecurityToken() *string
	SetTag(v []*CreatePluginRequestTag) *CreatePluginRequest
	GetTag() []*CreatePluginRequestTag
}

type CreatePluginRequest struct {
	// The description of the plug-in. The description can contain a maximum of 200 characters in length.
	//
	// example:
	//
	// createPlugin
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The plug-in definition. Supported formats: JSON and YAML.
	//
	// This parameter is required.
	//
	// example:
	//
	// Plugin definition
	PluginData *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	// The name of the plug-in. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). However, it cannot start with an underscore.
	//
	// This parameter is required.
	//
	// example:
	//
	// NewCors
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The type of the plug-in. Valid values:
	//
	// 	- **ipControl: IP address-based access control**
	//
	// 	- **trafficControl: throttling**
	//
	// 	- **backendSignature: backend signature**
	//
	// 	- **jwtAuth*	- :JWT (OpenId Connect) authentication
	//
	// 	- **cors*	- :cross-origin resource sharing (CORS)
	//
	// 	- **caching**
	//
	// This parameter is required.
	//
	// example:
	//
	// cors
	PluginType    *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*CreatePluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePluginRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePluginRequest) GetPluginData() *string {
	return s.PluginData
}

func (s *CreatePluginRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *CreatePluginRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreatePluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreatePluginRequest) GetTag() []*CreatePluginRequestTag {
	return s.Tag
}

func (s *CreatePluginRequest) SetDescription(v string) *CreatePluginRequest {
	s.Description = &v
	return s
}

func (s *CreatePluginRequest) SetPluginData(v string) *CreatePluginRequest {
	s.PluginData = &v
	return s
}

func (s *CreatePluginRequest) SetPluginName(v string) *CreatePluginRequest {
	s.PluginName = &v
	return s
}

func (s *CreatePluginRequest) SetPluginType(v string) *CreatePluginRequest {
	s.PluginType = &v
	return s
}

func (s *CreatePluginRequest) SetSecurityToken(v string) *CreatePluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreatePluginRequest) SetTag(v []*CreatePluginRequestTag) *CreatePluginRequest {
	s.Tag = v
	return s
}

func (s *CreatePluginRequest) Validate() error {
	return dara.Validate(s)
}

type CreatePluginRequestTag struct {
	// The key of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// This parameter is required.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// This parameter is required.
	//
	// example:
	//
	// \\" \\"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePluginRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePluginRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePluginRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePluginRequestTag) SetKey(v string) *CreatePluginRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePluginRequestTag) SetValue(v string) *CreatePluginRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePluginRequestTag) Validate() error {
	return dara.Validate(s)
}
