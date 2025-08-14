// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafTagListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafTagListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSafTagListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSafTagListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSafTagListResponseBodyResultObject) *DescribeSafTagListResponseBody
	GetResultObject() []*DescribeSafTagListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSafTagListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSafTagListResponseBody
	GetTotalPage() *int32
}

type DescribeSafTagListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeSafTagListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSafTagListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafTagListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafTagListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafTagListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSafTagListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSafTagListResponseBody) GetResultObject() []*DescribeSafTagListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafTagListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSafTagListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSafTagListResponseBody) SetRequestId(v string) *DescribeSafTagListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafTagListResponseBody) SetCurrentPage(v int32) *DescribeSafTagListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSafTagListResponseBody) SetPageSize(v int32) *DescribeSafTagListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSafTagListResponseBody) SetResultObject(v []*DescribeSafTagListResponseBodyResultObject) *DescribeSafTagListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafTagListResponseBody) SetTotalItem(v int32) *DescribeSafTagListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSafTagListResponseBody) SetTotalPage(v int32) *DescribeSafTagListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSafTagListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSafTagListResponseBodyResultObject struct {
	// Tag description.
	//
	// example:
	//
	// 依赖IP画像模型识别输出风险高、中高、中的判定\\n数据来源主要基于IP的网络属性数据和对黑产行为轨迹的情报监控数据，涵盖可识别检测IP网络属性特征、恶意属性等特征\\n而模型评分的高中低主要是根据当前IP号段命中风险行为的恶意程度和种类来判定的\\n备注：以下rn0311、rn0312、rn0313、rn0314、rn0315属于IP画像模型针对当前请求IP识别出的风险类别
	TagDesc *string `json:"tagDesc,omitempty" xml:"tagDesc,omitempty"`
	// Tag meaning.
	//
	// example:
	//
	// IP风险评分高
	TagMean *string `json:"tagMean,omitempty" xml:"tagMean,omitempty"`
	// Tag name.
	//
	// example:
	//
	// rn0301
	TagName *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	// Tag identifier.
	//
	// example:
	//
	// rn0301
	TagState *string `json:"tagState,omitempty" xml:"tagState,omitempty"`
	// Tag type.
	//
	// example:
	//
	// IP风险类
	TagType *string `json:"tagType,omitempty" xml:"tagType,omitempty"`
	// Unique identifier of the tag key.
	//
	// example:
	//
	// rn0301
	TagUid *string `json:"tagUid,omitempty" xml:"tagUid,omitempty"`
	// Update time.
	//
	// example:
	//
	// 1684744034000
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DescribeSafTagListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafTagListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagDesc() *string {
	return s.TagDesc
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagMean() *string {
	return s.TagMean
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagName() *string {
	return s.TagName
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagState() *string {
	return s.TagState
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagType() *string {
	return s.TagType
}

func (s *DescribeSafTagListResponseBodyResultObject) GetTagUid() *string {
	return s.TagUid
}

func (s *DescribeSafTagListResponseBodyResultObject) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagDesc(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagDesc = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagMean(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagMean = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagName(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagName = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagState(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagState = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagType(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagType = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetTagUid(v string) *DescribeSafTagListResponseBodyResultObject {
	s.TagUid = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) SetUpdateTime(v string) *DescribeSafTagListResponseBodyResultObject {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSafTagListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
