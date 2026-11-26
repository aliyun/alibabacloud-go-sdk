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
	// The business type. This parameter has no practical meaning and can be ignored.
	//
	// example:
	//
	// None
	BusinessType *int32 `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can determine the cause of the failure based on this parameter.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 219
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 10
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The returned results.
	Data []*RuleCountInfo `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error details when an error occurs, or **successful*	- when the request is successful.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The error details when an error occurs. This parameter is used when multiple messages are returned.
	Messages []*string `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Repeated"`
	// The current page number.
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
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. You can determine whether the request is successful based on this parameter. A value of **true*	- indicates success. A value of **false*	- or **null*	- indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of data entries.
	//
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
