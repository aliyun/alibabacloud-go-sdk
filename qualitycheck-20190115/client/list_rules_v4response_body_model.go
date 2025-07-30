// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRulesV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessType(v int32) *ListRulesV4ResponseBody
	GetBusinessType() *int32
	SetCode(v string) *ListRulesV4ResponseBody
	GetCode() *string
	SetCount(v int32) *ListRulesV4ResponseBody
	GetCount() *int32
	SetCurrentPage(v int32) *ListRulesV4ResponseBody
	GetCurrentPage() *int32
	SetData(v []*RuleCountInfo) *ListRulesV4ResponseBody
	GetData() []*RuleCountInfo
	SetHttpStatusCode(v int32) *ListRulesV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListRulesV4ResponseBody
	GetMessage() *string
	SetMessages(v []*string) *ListRulesV4ResponseBody
	GetMessages() []*string
	SetPageNumber(v int32) *ListRulesV4ResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRulesV4ResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRulesV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRulesV4ResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListRulesV4ResponseBody
	GetTotalCount() *int32
}

type ListRulesV4ResponseBody struct {
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 219
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 10
	CurrentPage *int32           `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	Data        []*RuleCountInfo `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
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
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 219
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRulesV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRulesV4ResponseBody) GoString() string {
	return s.String()
}

func (s *ListRulesV4ResponseBody) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *ListRulesV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRulesV4ResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListRulesV4ResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListRulesV4ResponseBody) GetData() []*RuleCountInfo {
	return s.Data
}

func (s *ListRulesV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListRulesV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRulesV4ResponseBody) GetMessages() []*string {
	return s.Messages
}

func (s *ListRulesV4ResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRulesV4ResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRulesV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRulesV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRulesV4ResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListRulesV4ResponseBody) SetBusinessType(v int32) *ListRulesV4ResponseBody {
	s.BusinessType = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCode(v string) *ListRulesV4ResponseBody {
	s.Code = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCount(v int32) *ListRulesV4ResponseBody {
	s.Count = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetCurrentPage(v int32) *ListRulesV4ResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetData(v []*RuleCountInfo) *ListRulesV4ResponseBody {
	s.Data = v
	return s
}

func (s *ListRulesV4ResponseBody) SetHttpStatusCode(v int32) *ListRulesV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetMessage(v string) *ListRulesV4ResponseBody {
	s.Message = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetMessages(v []*string) *ListRulesV4ResponseBody {
	s.Messages = v
	return s
}

func (s *ListRulesV4ResponseBody) SetPageNumber(v int32) *ListRulesV4ResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetPageSize(v int32) *ListRulesV4ResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetRequestId(v string) *ListRulesV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetSuccess(v bool) *ListRulesV4ResponseBody {
	s.Success = &v
	return s
}

func (s *ListRulesV4ResponseBody) SetTotalCount(v int32) *ListRulesV4ResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRulesV4ResponseBody) Validate() error {
	return dara.Validate(s)
}
