// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPluginId(v string) *DeletePluginRequest
	GetPluginId() *string
	SetSecurityToken(v string) *DeletePluginRequest
	GetSecurityToken() *string
	SetTag(v []*DeletePluginRequestTag) *DeletePluginRequest
	GetTag() []*DeletePluginRequestTag
}

type DeletePluginRequest struct {
	// The ID of the plug-in to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a3f1a5279434f2ba74ccd91c295af9f
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*DeletePluginRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeletePluginRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginRequest) GoString() string {
	return s.String()
}

func (s *DeletePluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DeletePluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeletePluginRequest) GetTag() []*DeletePluginRequestTag {
	return s.Tag
}

func (s *DeletePluginRequest) SetPluginId(v string) *DeletePluginRequest {
	s.PluginId = &v
	return s
}

func (s *DeletePluginRequest) SetSecurityToken(v string) *DeletePluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeletePluginRequest) SetTag(v []*DeletePluginRequestTag) *DeletePluginRequest {
	s.Tag = v
	return s
}

func (s *DeletePluginRequest) Validate() error {
	return dara.Validate(s)
}

type DeletePluginRequestTag struct {
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

func (s DeletePluginRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginRequestTag) GoString() string {
	return s.String()
}

func (s *DeletePluginRequestTag) GetKey() *string {
	return s.Key
}

func (s *DeletePluginRequestTag) GetValue() *string {
	return s.Value
}

func (s *DeletePluginRequestTag) SetKey(v string) *DeletePluginRequestTag {
	s.Key = &v
	return s
}

func (s *DeletePluginRequestTag) SetValue(v string) *DeletePluginRequestTag {
	s.Value = &v
	return s
}

func (s *DeletePluginRequestTag) Validate() error {
	return dara.Validate(s)
}
