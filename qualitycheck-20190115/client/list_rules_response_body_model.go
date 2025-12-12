// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRulesResponseBody
	GetCode() *string
	SetCount(v int32) *ListRulesResponseBody
	GetCount() *int32
	SetData(v []*ListRulesResponseBodyData) *ListRulesResponseBody
	GetData() []*ListRulesResponseBodyData
	SetMessage(v string) *ListRulesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListRulesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRulesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRulesResponseBody
	GetSuccess() *bool
}

type ListRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 20
	Count *int32                       `json:"Count,omitempty" xml:"Count,omitempty"`
	Data  []*ListRulesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRulesResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListRulesResponseBody) GetData() []*ListRulesResponseBodyData {
	return s.Data
}

func (s *ListRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRulesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRulesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRulesResponseBody) SetCode(v string) *ListRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesResponseBody) SetCount(v int32) *ListRulesResponseBody {
	s.Count = &v
	return s
}

func (s *ListRulesResponseBody) SetData(v []*ListRulesResponseBodyData) *ListRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesResponseBody) SetMessage(v string) *ListRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesResponseBody) SetPageNumber(v int32) *ListRulesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRulesResponseBody) SetPageSize(v int32) *ListRulesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRulesResponseBody) SetRequestId(v string) *ListRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesResponseBody) SetSuccess(v bool) *ListRulesResponseBody {
	s.Success = &v
	return s
}

func (s *ListRulesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRulesResponseBodyData struct {
	BusinessCategoryNameList []*string `json:"BusinessCategoryNameList,omitempty" xml:"BusinessCategoryNameList,omitempty" type:"Repeated"`
	Comments                 *string   `json:"Comments,omitempty" xml:"Comments,omitempty"`
	// example:
	//
	// 2020-04-20T20:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1234567
	Rid *int64 `json:"Rid,omitempty" xml:"Rid,omitempty"`
	// example:
	//
	// 1
	RuleType *int32 `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	// example:
	//
	// 1
	Type     *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRulesResponseBodyData) GetBusinessCategoryNameList() []*string {
	return s.BusinessCategoryNameList
}

func (s *ListRulesResponseBodyData) GetComments() *string {
	return s.Comments
}

func (s *ListRulesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRulesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListRulesResponseBodyData) GetRid() *int64 {
	return s.Rid
}

func (s *ListRulesResponseBodyData) GetRuleType() *int32 {
	return s.RuleType
}

func (s *ListRulesResponseBodyData) GetType() *int32 {
	return s.Type
}

func (s *ListRulesResponseBodyData) GetTypeName() *string {
	return s.TypeName
}

func (s *ListRulesResponseBodyData) SetBusinessCategoryNameList(v []*string) *ListRulesResponseBodyData {
	s.BusinessCategoryNameList = v
	return s
}

func (s *ListRulesResponseBodyData) SetComments(v string) *ListRulesResponseBodyData {
	s.Comments = &v
	return s
}

func (s *ListRulesResponseBodyData) SetCreateTime(v string) *ListRulesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListRulesResponseBodyData) SetName(v string) *ListRulesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRid(v int64) *ListRulesResponseBodyData {
	s.Rid = &v
	return s
}

func (s *ListRulesResponseBodyData) SetRuleType(v int32) *ListRulesResponseBodyData {
	s.RuleType = &v
	return s
}

func (s *ListRulesResponseBodyData) SetType(v int32) *ListRulesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListRulesResponseBodyData) SetTypeName(v string) *ListRulesResponseBodyData {
	s.TypeName = &v
	return s
}

func (s *ListRulesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
