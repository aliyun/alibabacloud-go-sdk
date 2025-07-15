// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePluginsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribePluginsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePluginsResponseBody
	GetPageSize() *int32
	SetPlugins(v *DescribePluginsResponseBodyPlugins) *DescribePluginsResponseBody
	GetPlugins() *DescribePluginsResponseBodyPlugins
	SetRequestId(v string) *DescribePluginsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribePluginsResponseBody
	GetTotalCount() *int32
}

type DescribePluginsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The returned information about the plug-in. It is an array consisting of PluginAttribute data.
	Plugins *DescribePluginsResponseBodyPlugins `json:"Plugins,omitempty" xml:"Plugins,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 46373DC4-19F1-4DC8-8C31-1107289BB5E0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribePluginsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePluginsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePluginsResponseBody) GetPlugins() *DescribePluginsResponseBodyPlugins {
	return s.Plugins
}

func (s *DescribePluginsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePluginsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribePluginsResponseBody) SetPageNumber(v int32) *DescribePluginsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePluginsResponseBody) SetPageSize(v int32) *DescribePluginsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribePluginsResponseBody) SetPlugins(v *DescribePluginsResponseBodyPlugins) *DescribePluginsResponseBody {
	s.Plugins = v
	return s
}

func (s *DescribePluginsResponseBody) SetRequestId(v string) *DescribePluginsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePluginsResponseBody) SetTotalCount(v int32) *DescribePluginsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribePluginsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePluginsResponseBodyPlugins struct {
	PluginAttribute []*DescribePluginsResponseBodyPluginsPluginAttribute `json:"PluginAttribute,omitempty" xml:"PluginAttribute,omitempty" type:"Repeated"`
}

func (s DescribePluginsResponseBodyPlugins) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponseBodyPlugins) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPlugins) GetPluginAttribute() []*DescribePluginsResponseBodyPluginsPluginAttribute {
	return s.PluginAttribute
}

func (s *DescribePluginsResponseBodyPlugins) SetPluginAttribute(v []*DescribePluginsResponseBodyPluginsPluginAttribute) *DescribePluginsResponseBodyPlugins {
	s.PluginAttribute = v
	return s
}

func (s *DescribePluginsResponseBodyPlugins) Validate() error {
	return dara.Validate(s)
}

type DescribePluginsResponseBodyPluginsPluginAttribute struct {
	// The creation time (UTC) of the plug-in.
	//
	// example:
	//
	// 2019-01-11T09:29:58Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The plug-in description.
	//
	// example:
	//
	// Throttling
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The last modification time (UTC) of the plug-in.
	//
	// example:
	//
	// 2019-01-11T09:29:58Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The definition statement of the plug-in.
	//
	// example:
	//
	// {"unit":"MINUTE","apiDefault":20}
	PluginData *string `json:"PluginData,omitempty" xml:"PluginData,omitempty"`
	// The ID of the plug-in.
	//
	// example:
	//
	// 9a3f1a5279434f2ba74ccd91c295af9f
	PluginId *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	// The name of the plug-in.
	//
	// example:
	//
	// firstPlugin
	PluginName *string `json:"PluginName,omitempty" xml:"PluginName,omitempty"`
	// The type of the plug-in.
	//
	// example:
	//
	// trafficControl
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The region where the plug-in is located.
	//
	// example:
	//
	// cn-qingdao
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The tags.
	Tags *DescribePluginsResponseBodyPluginsPluginAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttribute) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetPluginData() *string {
	return s.PluginData
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetPluginId() *string {
	return s.PluginId
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetPluginName() *string {
	return s.PluginName
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetPluginType() *string {
	return s.PluginType
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) GetTags() *DescribePluginsResponseBodyPluginsPluginAttributeTags {
	return s.Tags
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetCreatedTime(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetDescription(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.Description = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetModifiedTime(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginData(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginData = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginId(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginId = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginName(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginName = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetPluginType(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.PluginType = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetRegionId(v string) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) SetTags(v *DescribePluginsResponseBodyPluginsPluginAttributeTags) *DescribePluginsResponseBodyPluginsPluginAttribute {
	s.Tags = v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribePluginsResponseBodyPluginsPluginAttributeTags struct {
	TagInfo []*DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTags) GetTagInfo() []*DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo {
	return s.TagInfo
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTags) SetTagInfo(v []*DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) *DescribePluginsResponseBodyPluginsPluginAttributeTags {
	s.TagInfo = v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTags) Validate() error {
	return dara.Validate(s)
}

type DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo struct {
	// The key of the tag.
	//
	// example:
	//
	// testkey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// tetstvalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) SetKey(v string) *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) SetValue(v string) *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribePluginsResponseBodyPluginsPluginAttributeTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
