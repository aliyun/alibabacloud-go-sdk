// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyPluginRequest
	GetDescription() *string
	SetPluginData(v string) *ModifyPluginRequest
	GetPluginData() *string
	SetPluginId(v string) *ModifyPluginRequest
	GetPluginId() *string
	SetPluginName(v string) *ModifyPluginRequest
	GetPluginName() *string
	SetSecurityToken(v string) *ModifyPluginRequest
	GetSecurityToken() *string
	SetTag(v []*ModifyPluginRequestTag) *ModifyPluginRequest
	GetTag() []*ModifyPluginRequestTag
}

type ModifyPluginRequest struct {
	// The description of the plug-in. The description can contain a maximum of 200 characters in length.
	//
	// example:
	//
	// modify plugin first
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The statement that is used to modify the plug-in definition.
	//
	// example:
	//
	// Plugin definition
	PluginData *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	// The ID of the plug-in whose information you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// a96926e82f994915a8da40a119374537
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The name of the plug-in. The name must be 4 to 50 characters in length and can contain letters, digits, and underscores (_). However, it cannot start with an underscore.
	//
	// example:
	//
	// modifyCors
	PluginName    *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*ModifyPluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ModifyPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPluginRequest) GoString() string {
	return s.String()
}

func (s *ModifyPluginRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPluginRequest) GetPluginData() *string {
	return s.PluginData
}

func (s *ModifyPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *ModifyPluginRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *ModifyPluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyPluginRequest) GetTag() []*ModifyPluginRequestTag {
	return s.Tag
}

func (s *ModifyPluginRequest) SetDescription(v string) *ModifyPluginRequest {
	s.Description = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginData(v string) *ModifyPluginRequest {
	s.PluginData = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginId(v string) *ModifyPluginRequest {
	s.PluginId = &v
	return s
}

func (s *ModifyPluginRequest) SetPluginName(v string) *ModifyPluginRequest {
	s.PluginName = &v
	return s
}

func (s *ModifyPluginRequest) SetSecurityToken(v string) *ModifyPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyPluginRequest) SetTag(v []*ModifyPluginRequestTag) *ModifyPluginRequest {
	s.Tag = v
	return s
}

func (s *ModifyPluginRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPluginRequestTag struct {
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

func (s ModifyPluginRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ModifyPluginRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyPluginRequestTag) GetKey() *string {
	return s.Key
}

func (s *ModifyPluginRequestTag) GetValue() *string {
	return s.Value
}

func (s *ModifyPluginRequestTag) SetKey(v string) *ModifyPluginRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyPluginRequestTag) SetValue(v string) *ModifyPluginRequestTag {
	s.Value = &v
	return s
}

func (s *ModifyPluginRequestTag) Validate() error {
	return dara.Validate(s)
}
