// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListTemplateResponseBody
	GetCode() *int32
	SetData(v []*ListTemplateResponseBodyData) *ListTemplateResponseBody
	GetData() []*ListTemplateResponseBodyData
	SetMessage(v string) *ListTemplateResponseBody
	GetMessage() *string
	SetNextToken(v int32) *ListTemplateResponseBody
	GetNextToken() *int32
	SetRequestId(v string) *ListTemplateResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTemplateResponseBody
	GetTotalCount() *int32
}

type ListTemplateResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details about templates.
	Data []*ListTemplateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	NextToken *int32 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListTemplateResponseBody) GetData() []*ListTemplateResponseBodyData {
	return s.Data
}

func (s *ListTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListTemplateResponseBody) GetNextToken() *int32 {
	return s.NextToken
}

func (s *ListTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTemplateResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTemplateResponseBody) SetCode(v int32) *ListTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *ListTemplateResponseBody) SetData(v []*ListTemplateResponseBodyData) *ListTemplateResponseBody {
	s.Data = v
	return s
}

func (s *ListTemplateResponseBody) SetMessage(v string) *ListTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *ListTemplateResponseBody) SetNextToken(v int32) *ListTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTemplateResponseBody) SetRequestId(v string) *ListTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTemplateResponseBody) SetTotalCount(v int32) *ListTemplateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTemplateResponseBodyData struct {
	// The time when the template was created.
	//
	// example:
	//
	// 2021-03-18 16:41:31
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DocumentUrl *string `json:"DocumentUrl,omitempty" xml:"DocumentUrl,omitempty"`
	// The URL of the architecture image.
	//
	// example:
	//
	// bp-studio-template/sr-U37UD2YQCRJ75X5V.png
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// cadt-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmyjt3c5om3hi
	ResourceGroupId *string                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListTemplateResponseBodyDataTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the tag that is added to the template.
	//
	// example:
	//
	// 1
	TagId *int32 `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// Deprecated
	//
	// The name of the tag that is added to the template.
	//
	// example:
	//
	// Official template
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// CJQ6H0XUEQ20IYJQ
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s ListTemplateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListTemplateResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListTemplateResponseBodyData) GetDocumentUrl() *string {
	return s.DocumentUrl
}

func (s *ListTemplateResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *ListTemplateResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListTemplateResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListTemplateResponseBodyData) GetTag() []*ListTemplateResponseBodyDataTag {
	return s.Tag
}

func (s *ListTemplateResponseBodyData) GetTagId() *int32 {
	return s.TagId
}

func (s *ListTemplateResponseBodyData) GetTagName() *string {
	return s.TagName
}

func (s *ListTemplateResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListTemplateResponseBodyData) SetCreateTime(v string) *ListTemplateResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetDescription(v string) *ListTemplateResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetDocumentUrl(v string) *ListTemplateResponseBodyData {
	s.DocumentUrl = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetImageURL(v string) *ListTemplateResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetName(v string) *ListTemplateResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetResourceGroupId(v string) *ListTemplateResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTag(v []*ListTemplateResponseBodyDataTag) *ListTemplateResponseBodyData {
	s.Tag = v
	return s
}

func (s *ListTemplateResponseBodyData) SetTagId(v int32) *ListTemplateResponseBodyData {
	s.TagId = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTagName(v string) *ListTemplateResponseBodyData {
	s.TagName = &v
	return s
}

func (s *ListTemplateResponseBodyData) SetTemplateId(v string) *ListTemplateResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *ListTemplateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListTemplateResponseBodyDataTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTemplateResponseBodyDataTag) String() string {
	return dara.Prettify(s)
}

func (s ListTemplateResponseBodyDataTag) GoString() string {
	return s.String()
}

func (s *ListTemplateResponseBodyDataTag) GetKey() *string {
	return s.Key
}

func (s *ListTemplateResponseBodyDataTag) GetValue() *string {
	return s.Value
}

func (s *ListTemplateResponseBodyDataTag) SetKey(v string) *ListTemplateResponseBodyDataTag {
	s.Key = &v
	return s
}

func (s *ListTemplateResponseBodyDataTag) SetValue(v string) *ListTemplateResponseBodyDataTag {
	s.Value = &v
	return s
}

func (s *ListTemplateResponseBodyDataTag) Validate() error {
	return dara.Validate(s)
}
