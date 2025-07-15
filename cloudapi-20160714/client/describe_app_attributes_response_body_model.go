// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApps(v *DescribeAppAttributesResponseBodyApps) *DescribeAppAttributesResponseBody
	GetApps() *DescribeAppAttributesResponseBodyApps
	SetPageNumber(v int32) *DescribeAppAttributesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAppAttributesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAppAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAppAttributesResponseBody
	GetTotalCount() *int32
}

type DescribeAppAttributesResponseBody struct {
	// The returned app information. It is an array that consists of AppAttribute data.
	Apps *DescribeAppAttributesResponseBodyApps `json:"Apps,omitempty" xml:"Apps,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// 8883AC74-259D-4C0B-99FC-0B7F9A588B2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAppAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBody) GetApps() *DescribeAppAttributesResponseBodyApps {
	return s.Apps
}

func (s *DescribeAppAttributesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAppAttributesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAppAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppAttributesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAppAttributesResponseBody) SetApps(v *DescribeAppAttributesResponseBodyApps) *DescribeAppAttributesResponseBody {
	s.Apps = v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetPageNumber(v int32) *DescribeAppAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetPageSize(v int32) *DescribeAppAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetRequestId(v string) *DescribeAppAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) SetTotalCount(v int32) *DescribeAppAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAppAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAttributesResponseBodyApps struct {
	AppAttribute []*DescribeAppAttributesResponseBodyAppsAppAttribute `json:"AppAttribute,omitempty" xml:"AppAttribute,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesResponseBodyApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyApps) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyApps) GetAppAttribute() []*DescribeAppAttributesResponseBodyAppsAppAttribute {
	return s.AppAttribute
}

func (s *DescribeAppAttributesResponseBodyApps) SetAppAttribute(v []*DescribeAppAttributesResponseBodyAppsAppAttribute) *DescribeAppAttributesResponseBodyApps {
	s.AppAttribute = v
	return s
}

func (s *DescribeAppAttributesResponseBodyApps) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAttributesResponseBodyAppsAppAttribute struct {
	// The ID of the app.
	//
	// example:
	//
	// 20112314518278
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the app.
	//
	// example:
	//
	// CreateApptest
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The creation time (UTC) of the app.
	//
	// example:
	//
	// 2016-07-31T04:10:19Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the app.
	//
	// example:
	//
	// App test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Disabled    *bool   `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// 扩展信息
	//
	// example:
	//
	// 110461946884
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The modification time (UTC) of the app.
	//
	// example:
	//
	// 2016-07-31T04:10:19Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The tags.
	Tags *DescribeAppAttributesResponseBodyAppsAppAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttribute) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetAppName() *string {
	return s.AppName
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetExtend() *string {
	return s.Extend
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) GetTags() *DescribeAppAttributesResponseBodyAppsAppAttributeTags {
	return s.Tags
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetAppId(v int64) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.AppId = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetAppName(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.AppName = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetCreatedTime(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.CreatedTime = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetDescription(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Description = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetDisabled(v bool) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Disabled = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetExtend(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Extend = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetModifiedTime(v string) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) SetTags(v *DescribeAppAttributesResponseBodyAppsAppAttributeTags) *DescribeAppAttributesResponseBodyAppsAppAttribute {
	s.Tags = v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttribute) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAttributesResponseBodyAppsAppAttributeTags struct {
	TagInfo []*DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTags) GetTagInfo() []*DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo {
	return s.TagInfo
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTags) SetTagInfo(v []*DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) *DescribeAppAttributesResponseBodyAppsAppAttributeTags {
	s.TagInfo = v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTags) Validate() error {
	return dara.Validate(s)
}

type DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo struct {
	// The key of the tag.
	//
	// example:
	//
	// appid
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) SetKey(v string) *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) SetValue(v string) *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribeAppAttributesResponseBodyAppsAppAttributeTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
