// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelDetails(v *DescribeModelsResponseBodyModelDetails) *DescribeModelsResponseBody
	GetModelDetails() *DescribeModelsResponseBodyModelDetails
	SetPageNumber(v int32) *DescribeModelsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeModelsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeModelsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeModelsResponseBody
	GetTotalCount() *int32
}

type DescribeModelsResponseBody struct {
	// The returned information about models. It is an array consisting of ModelDetail data.
	ModelDetails *DescribeModelsResponseBodyModelDetails `json:"ModelDetails,omitempty" xml:"ModelDetails,omitempty" type:"Struct"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40306469-2FB5-417A-B723-AF1F4A4FA204
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBody) GetModelDetails() *DescribeModelsResponseBodyModelDetails {
	return s.ModelDetails
}

func (s *DescribeModelsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeModelsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeModelsResponseBody) SetModelDetails(v *DescribeModelsResponseBodyModelDetails) *DescribeModelsResponseBody {
	s.ModelDetails = v
	return s
}

func (s *DescribeModelsResponseBody) SetPageNumber(v int32) *DescribeModelsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeModelsResponseBody) SetPageSize(v int32) *DescribeModelsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeModelsResponseBody) SetRequestId(v string) *DescribeModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelsResponseBody) SetTotalCount(v int32) *DescribeModelsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeModelsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeModelsResponseBodyModelDetails struct {
	ModelDetail []*DescribeModelsResponseBodyModelDetailsModelDetail `json:"ModelDetail,omitempty" xml:"ModelDetail,omitempty" type:"Repeated"`
}

func (s DescribeModelsResponseBodyModelDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetails) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetails) GetModelDetail() []*DescribeModelsResponseBodyModelDetailsModelDetail {
	return s.ModelDetail
}

func (s *DescribeModelsResponseBodyModelDetails) SetModelDetail(v []*DescribeModelsResponseBodyModelDetailsModelDetail) *DescribeModelsResponseBodyModelDetails {
	s.ModelDetail = v
	return s
}

func (s *DescribeModelsResponseBodyModelDetails) Validate() error {
	return dara.Validate(s)
}

type DescribeModelsResponseBodyModelDetailsModelDetail struct {
	// The time when the model was created.
	//
	// example:
	//
	// 2019-01-29T11:07:48Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The description of the model definition.
	//
	// example:
	//
	// Model Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the API group to which the model belongs.
	//
	// example:
	//
	// 30e792398d6c4569b04c0e53a3494381
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the model.
	//
	// example:
	//
	// 766c0b9538a04bdf974953b5576783ba
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// Test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The URI of the model.
	//
	// example:
	//
	// https://apigateway.aliyun.com/models/30e792398d6c4569b04c0e53a3494381/766c0b9538a04bdf974953b5576783ba
	ModelRef *string `json:"ModelRef,omitempty" xml:"ModelRef,omitempty"`
	// The time when the model was last modified.
	//
	// example:
	//
	// 2019-01-29T11:07:48Z
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The definition of the model.
	//
	// example:
	//
	// {\\"type\\":\\"object\\",\\"properties\\":{\\"id\\":{\\"format\\":\\"int64\\",\\"maximum\\":100,\\"exclusiveMaximum\\":true,\\"type\\":\\"integer\\"},\\"name\\":{\\"maxLength\\":10,\\"type\\":\\"string\\"}}}
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The tags of the model.
	Tags *DescribeModelsResponseBodyModelDetailsModelDetailTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeModelsResponseBodyModelDetailsModelDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetailsModelDetail) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetDescription() *string {
	return s.Description
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetModelId() *string {
	return s.ModelId
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetModelRef() *string {
	return s.ModelRef
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetSchema() *string {
	return s.Schema
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) GetTags() *DescribeModelsResponseBodyModelDetailsModelDetailTags {
	return s.Tags
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetCreatedTime(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.CreatedTime = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetDescription(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.Description = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetGroupId(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.GroupId = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelId(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelId = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelName(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelName = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModelRef(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModelRef = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetModifiedTime(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.ModifiedTime = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetSchema(v string) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.Schema = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) SetTags(v *DescribeModelsResponseBodyModelDetailsModelDetailTags) *DescribeModelsResponseBodyModelDetailsModelDetail {
	s.Tags = v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeModelsResponseBodyModelDetailsModelDetailTags struct {
	TagInfo []*DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo `json:"TagInfo,omitempty" xml:"TagInfo,omitempty" type:"Repeated"`
}

func (s DescribeModelsResponseBodyModelDetailsModelDetailTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetailsModelDetailTags) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTags) GetTagInfo() []*DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo {
	return s.TagInfo
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTags) SetTagInfo(v []*DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) *DescribeModelsResponseBodyModelDetailsModelDetailTags {
	s.TagInfo = v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTags) Validate() error {
	return dara.Validate(s)
}

type DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo struct {
	// The tag key.
	//
	// example:
	//
	// ENV
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// ST4
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) GoString() string {
	return s.String()
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) GetValue() *string {
	return s.Value
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) SetKey(v string) *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo {
	s.Key = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) SetValue(v string) *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo {
	s.Value = &v
	return s
}

func (s *DescribeModelsResponseBodyModelDetailsModelDetailTagsTagInfo) Validate() error {
	return dara.Validate(s)
}
