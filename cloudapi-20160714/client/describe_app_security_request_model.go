// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecurityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DescribeAppSecurityRequest
	GetAppId() *int64
	SetSecurityToken(v string) *DescribeAppSecurityRequest
	GetSecurityToken() *string
	SetTag(v []*DescribeAppSecurityRequestTag) *DescribeAppSecurityRequest
	GetTag() []*DescribeAppSecurityRequestTag
}

type DescribeAppSecurityRequest struct {
	// The ID of the app.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20112314518278
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*DescribeAppSecurityRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAppSecurityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecurityRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppSecurityRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAppSecurityRequest) GetTag() []*DescribeAppSecurityRequestTag {
	return s.Tag
}

func (s *DescribeAppSecurityRequest) SetAppId(v int64) *DescribeAppSecurityRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppSecurityRequest) SetSecurityToken(v string) *DescribeAppSecurityRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppSecurityRequest) SetTag(v []*DescribeAppSecurityRequestTag) *DescribeAppSecurityRequest {
	s.Tag = v
	return s
}

func (s *DescribeAppSecurityRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeAppSecurityRequestTag struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppSecurityRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecurityRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAppSecurityRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAppSecurityRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAppSecurityRequestTag) SetKey(v string) *DescribeAppSecurityRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAppSecurityRequestTag) SetValue(v string) *DescribeAppSecurityRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAppSecurityRequestTag) Validate() error {
	return dara.Validate(s)
}
