// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *ModifyAppRequest
	GetAppId() *int64
	SetAppName(v string) *ModifyAppRequest
	GetAppName() *string
	SetDescription(v string) *ModifyAppRequest
	GetDescription() *string
	SetDisabled(v bool) *ModifyAppRequest
	GetDisabled() *bool
	SetExtend(v string) *ModifyAppRequest
	GetExtend() *string
	SetSecurityToken(v string) *ModifyAppRequest
	GetSecurityToken() *string
	SetTag(v []*ModifyAppRequestTag) *ModifyAppRequest
	GetTag() []*ModifyAppRequestTag
}

type ModifyAppRequest struct {
	// The ID of the app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The value must be 4 to 26 characters in length and can contain letters, digits, and underscores (_). It must start with a letter.
	//
	// This parameter is required only when you want to modify the value.
	//
	// example:
	//
	// jiedian_pord
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The description of the app. The description can contain a maximum of 180 characters in length.
	//
	// This parameter is required only when you want to modify the value.
	//
	// example:
	//
	// modidyTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// 扩展信息
	//
	// example:
	//
	// 110461946884
	Extend        *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*ModifyAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ModifyAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *ModifyAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyAppRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAppRequest) GetDisabled() *bool {
	return s.Disabled
}

func (s *ModifyAppRequest) GetExtend() *string {
	return s.Extend
}

func (s *ModifyAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyAppRequest) GetTag() []*ModifyAppRequestTag {
	return s.Tag
}

func (s *ModifyAppRequest) SetAppId(v int64) *ModifyAppRequest {
	s.AppId = &v
	return s
}

func (s *ModifyAppRequest) SetAppName(v string) *ModifyAppRequest {
	s.AppName = &v
	return s
}

func (s *ModifyAppRequest) SetDescription(v string) *ModifyAppRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppRequest) SetDisabled(v bool) *ModifyAppRequest {
	s.Disabled = &v
	return s
}

func (s *ModifyAppRequest) SetExtend(v string) *ModifyAppRequest {
	s.Extend = &v
	return s
}

func (s *ModifyAppRequest) SetSecurityToken(v string) *ModifyAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyAppRequest) SetTag(v []*ModifyAppRequestTag) *ModifyAppRequest {
	s.Tag = v
	return s
}

func (s *ModifyAppRequest) Validate() error {
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

type ModifyAppRequestTag struct {
	// The value of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// This parameter is required.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyAppRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppRequestTag) GoString() string {
	return s.String()
}

func (s *ModifyAppRequestTag) GetKey() *string {
	return s.Key
}

func (s *ModifyAppRequestTag) GetValue() *string {
	return s.Value
}

func (s *ModifyAppRequestTag) SetKey(v string) *ModifyAppRequestTag {
	s.Key = &v
	return s
}

func (s *ModifyAppRequestTag) SetValue(v string) *ModifyAppRequestTag {
	s.Value = &v
	return s
}

func (s *ModifyAppRequestTag) Validate() error {
	return dara.Validate(s)
}
