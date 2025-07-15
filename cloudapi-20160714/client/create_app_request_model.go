// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *CreateAppRequest
	GetAppCode() *string
	SetAppKey(v string) *CreateAppRequest
	GetAppKey() *string
	SetAppName(v string) *CreateAppRequest
	GetAppName() *string
	SetAppSecret(v string) *CreateAppRequest
	GetAppSecret() *string
	SetDescription(v string) *CreateAppRequest
	GetDescription() *string
	SetExtend(v string) *CreateAppRequest
	GetExtend() *string
	SetSecurityToken(v string) *CreateAppRequest
	GetSecurityToken() *string
	SetTag(v []*CreateAppRequestTag) *CreateAppRequest
	GetTag() []*CreateAppRequestTag
}

type CreateAppRequest struct {
	// The AppCode of the application.
	//
	// example:
	//
	// 3aaf905a0a1f4f0eabc6d891dfa08afc
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// The key of the application that is used to make an API call.
	//
	// example:
	//
	// 60030986
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The name of the application. The name must be 4 to 26 characters in length. The name can contain letters, digits, and underscores (_), and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateAppTest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The password of the application.
	//
	// example:
	//
	// nzyNqFkRWB2uLw86
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The description of the application. The description can be up to 180 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information.
	//
	// example:
	//
	// 110210264071
	Extend        *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*CreateAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateAppRequest) GetAppKey() *string {
	return s.AppKey
}

func (s *CreateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppRequest) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *CreateAppRequest) GetTag() []*CreateAppRequestTag {
	return s.Tag
}

func (s *CreateAppRequest) SetAppCode(v string) *CreateAppRequest {
	s.AppCode = &v
	return s
}

func (s *CreateAppRequest) SetAppKey(v string) *CreateAppRequest {
	s.AppKey = &v
	return s
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetAppSecret(v string) *CreateAppRequest {
	s.AppSecret = &v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetExtend(v string) *CreateAppRequest {
	s.Extend = &v
	return s
}

func (s *CreateAppRequest) SetSecurityToken(v string) *CreateAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *CreateAppRequest) SetTag(v []*CreateAppRequestTag) *CreateAppRequest {
	s.Tag = v
	return s
}

func (s *CreateAppRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAppRequestTag struct {
	// The key of the tag.
	//
	// Valid values of n: `[1, 20]`.
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// Valid values of n: `[1, 20]`. If the parameter has a value, you must specify a value for the tag key with the same N as tag.N.Key. Otherwise, an error is reported.
	//
	// example:
	//
	// \\" \\"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateAppRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestTag) GoString() string {
	return s.String()
}

func (s *CreateAppRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateAppRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateAppRequestTag) SetKey(v string) *CreateAppRequestTag {
	s.Key = &v
	return s
}

func (s *CreateAppRequestTag) SetValue(v string) *CreateAppRequestTag {
	s.Value = &v
	return s
}

func (s *CreateAppRequestTag) Validate() error {
	return dara.Validate(s)
}
