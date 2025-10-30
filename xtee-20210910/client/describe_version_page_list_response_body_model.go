// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVersionPageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeVersionPageListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVersionPageListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeVersionPageListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeVersionPageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVersionPageListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeVersionPageListResponseBodyResultObject) *DescribeVersionPageListResponseBody
	GetResultObject() []*DescribeVersionPageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeVersionPageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeVersionPageListResponseBody
	GetTotalPage() *int32
}

type DescribeVersionPageListResponseBody struct {
	// The maximum amount of data to be read in this call, with a default value of 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 用来表示当前调用返回读取到的位置。空代表数据已经读取完毕。
	//
	// example:
	//
	// 1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	// Number of items per page, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeVersionPageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeVersionPageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionPageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVersionPageListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVersionPageListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVersionPageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVersionPageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVersionPageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVersionPageListResponseBody) GetResultObject() []*DescribeVersionPageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVersionPageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeVersionPageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVersionPageListResponseBody) SetMaxResults(v int32) *DescribeVersionPageListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetNextToken(v string) *DescribeVersionPageListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetRequestId(v string) *DescribeVersionPageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetCurrentPage(v int32) *DescribeVersionPageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetPageSize(v int32) *DescribeVersionPageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetResultObject(v []*DescribeVersionPageListResponseBodyResultObject) *DescribeVersionPageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetTotalItem(v int32) *DescribeVersionPageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) SetTotalPage(v int32) *DescribeVersionPageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVersionPageListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVersionPageListResponseBodyResultObject struct {
	// Change content.
	//
	// example:
	//
	// {\\"description\\":\\"自定义变量\\",\\"expression\\":\\"1\\",\\"expressionTitle\\":\\"1\\",\\"expressionVariable\\":\\"\\",\\"fieldRank\\":0,\\"id\\":392023,\\"name\\":\\"ex_OERlw0Zqfb23\\",\\"outlier\\":\\"SYS_ERROR\\",\\"outputs\\":\\"STRING\\",\\"refObjId\\":\\"de_auevox0336\\",\\"region\\":\\"SH\\",\\"title\\":\\"自定义变量\\",\\"version\\":4}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Creator.
	//
	// example:
	//
	// root
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Variable description.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Primary key ID of the version.
	//
	// example:
	//
	// 808
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Modifier.
	//
	// example:
	//
	// root
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// Name of the variable.
	//
	// example:
	//
	// ex_OERlw0Zqfb23
	ObjectCode *string `json:"objectCode,omitempty" xml:"objectCode,omitempty"`
	// Primary key ID of the variable.
	//
	// example:
	//
	// 392023
	ObjectId *int64 `json:"objectId,omitempty" xml:"objectId,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// Variable type.
	//
	// example:
	//
	// EXPRESSION
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User UID.
	//
	// example:
	//
	// 151222222222226
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeVersionPageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVersionPageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetContent() *string {
	return s.Content
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetModifier() *string {
	return s.Modifier
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetObjectCode() *string {
	return s.ObjectCode
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetObjectId() *int64 {
	return s.ObjectId
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetRegion() *string {
	return s.Region
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetType() *string {
	return s.Type
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeVersionPageListResponseBodyResultObject) GetVersion() *int64 {
	return s.Version
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetContent(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Content = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetCreator(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetDescription(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetId(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetModifier(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Modifier = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetObjectCode(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.ObjectCode = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetObjectId(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.ObjectId = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetRegion(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Region = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetType(v string) *DescribeVersionPageListResponseBodyResultObject {
	s.Type = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetUserId(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) SetVersion(v int64) *DescribeVersionPageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeVersionPageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
