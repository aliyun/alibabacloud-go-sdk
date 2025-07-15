// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePluginsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsRequest
	GetPageSize() *int32
	SetPluginId(v string) *DescribePluginsRequest
	GetPluginId() *string
	SetPluginName(v string) *DescribePluginsRequest
	GetPluginName() *string
	SetPluginType(v string) *DescribePluginsRequest
	GetPluginType() *string
	SetSecurityToken(v string) *DescribePluginsRequest
	GetSecurityToken() *string
	SetTag(v []*DescribePluginsRequestTag) *DescribePluginsRequest
	GetTag() []*DescribePluginsRequestTag
}

type DescribePluginsRequest struct {
	// The number of the page to return. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the plug-in.
	//
	// example:
	//
	// a96926e82f994915a8da40a119374537
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// testPlugin
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The business type of the plug-in.
	//
	// example:
	//
	// cors
	PluginType    *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the lifecycle rule. You can specify multiple tags.
	//
	// example:
	//
	// Key， Value
	Tag []*DescribePluginsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePluginsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsRequest) GoString() string {
	return s.String()
}

func (s *DescribePluginsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginsRequest) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginsRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *DescribePluginsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribePluginsRequest) GetTag() []*DescribePluginsRequestTag {
	return s.Tag
}

func (s *DescribePluginsRequest) SetPageNumber(v int32) *DescribePluginsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsRequest) SetPageSize(v int32) *DescribePluginsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginId(v string) *DescribePluginsRequest {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginName(v string) *DescribePluginsRequest {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsRequest) SetPluginType(v string) *DescribePluginsRequest {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsRequest) SetSecurityToken(v string) *DescribePluginsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribePluginsRequest) SetTag(v []*DescribePluginsRequestTag) *DescribePluginsRequest {
	s.Tag = v
	return s
}

func (s *DescribePluginsRequest) Validate() error {
	return dara.Validate(s)
}

type DescribePluginsRequestTag struct {
	// The key of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// example:
	//
	// env
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// N can be an integer from 1 to 20.``
	//
	// example:
	//
	// \\" \\"
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePluginsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePluginsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribePluginsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribePluginsRequestTag) SetKey(v string) *DescribePluginsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePluginsRequestTag) SetValue(v string) *DescribePluginsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribePluginsRequestTag) Validate() error {
	return dara.Validate(s)
}
