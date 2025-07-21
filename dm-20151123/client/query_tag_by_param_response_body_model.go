// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagByParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *QueryTagByParamResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryTagByParamResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *QueryTagByParamResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *QueryTagByParamResponseBody
	GetTotalCount() *int32
	SetData(v *QueryTagByParamResponseBodyData) *QueryTagByParamResponseBody
	GetData() *QueryTagByParamResponseBodyData
}

type QueryTagByParamResponseBody struct {
	// Current page number
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 10A1AD70-E48E-476D-98D9-39BD92193837
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Data records
	Data *QueryTagByParamResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
}

func (s QueryTagByParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagByParamResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryTagByParamResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryTagByParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagByParamResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryTagByParamResponseBody) GetData() *QueryTagByParamResponseBodyData {
	return s.Data
}

func (s *QueryTagByParamResponseBody) SetPageNumber(v int32) *QueryTagByParamResponseBody {
	s.PageNumber = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetPageSize(v int32) *QueryTagByParamResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetRequestId(v string) *QueryTagByParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetTotalCount(v int32) *QueryTagByParamResponseBody {
	s.TotalCount = &v
	return s
}

func (s *QueryTagByParamResponseBody) SetData(v *QueryTagByParamResponseBodyData) *QueryTagByParamResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagByParamResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryTagByParamResponseBodyData struct {
	Tag []*QueryTagByParamResponseBodyDataTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s QueryTagByParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTagByParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBodyData) GetTag() []*QueryTagByParamResponseBodyDataTag {
	return s.Tag
}

func (s *QueryTagByParamResponseBodyData) SetTag(v []*QueryTagByParamResponseBodyDataTag) *QueryTagByParamResponseBodyData {
	s.Tag = v
	return s
}

func (s *QueryTagByParamResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryTagByParamResponseBodyDataTag struct {
	// Tag description
	//
	// example:
	//
	// test description
	TagDescription *string `json:"TagDescription,omitempty" xml:"TagDescription,omitempty"`
	// Tag ID
	//
	// example:
	//
	// 52366
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	// Tag name
	//
	// example:
	//
	// hellopal
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s QueryTagByParamResponseBodyDataTag) String() string {
	return dara.Prettify(s)
}

func (s QueryTagByParamResponseBodyDataTag) GoString() string {
	return s.String()
}

func (s *QueryTagByParamResponseBodyDataTag) GetTagDescription() *string {
	return s.TagDescription
}

func (s *QueryTagByParamResponseBodyDataTag) GetTagId() *string {
	return s.TagId
}

func (s *QueryTagByParamResponseBodyDataTag) GetTagName() *string {
	return s.TagName
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagDescription(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagDescription = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagId(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagId = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) SetTagName(v string) *QueryTagByParamResponseBodyDataTag {
	s.TagName = &v
	return s
}

func (s *QueryTagByParamResponseBodyDataTag) Validate() error {
	return dara.Validate(s)
}
