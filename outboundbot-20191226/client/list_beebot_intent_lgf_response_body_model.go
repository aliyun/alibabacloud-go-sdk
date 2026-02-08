// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ListBeebotIntentLgfResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ListBeebotIntentLgfResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBeebotIntentLgfResponseBody
	GetHttpStatusCode() *int32
	SetLgfs(v []*ListBeebotIntentLgfResponseBodyLgfs) *ListBeebotIntentLgfResponseBody
	GetLgfs() []*ListBeebotIntentLgfResponseBodyLgfs
	SetMessage(v string) *ListBeebotIntentLgfResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListBeebotIntentLgfResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentLgfResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBeebotIntentLgfResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBeebotIntentLgfResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListBeebotIntentLgfResponseBody
	GetTotalCount() *int32
}

type ListBeebotIntentLgfResponseBody struct {
	// Internal request ID
	//
	// example:
	//
	// 497CFAFF-48CC-161A-AD2C-252DED569037
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// API status code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// List of LGF descriptions
	Lgfs []*ListBeebotIntentLgfResponseBodyLgfs `json:"Lgfs,omitempty" xml:"Lgfs,omitempty" type:"Repeated"`
	// API message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Count
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBeebotIntentLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentLgfResponseBody) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentLgfResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ListBeebotIntentLgfResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBeebotIntentLgfResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBeebotIntentLgfResponseBody) GetLgfs() []*ListBeebotIntentLgfResponseBodyLgfs {
	return s.Lgfs
}

func (s *ListBeebotIntentLgfResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBeebotIntentLgfResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentLgfResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBeebotIntentLgfResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBeebotIntentLgfResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBeebotIntentLgfResponseBody) SetBeebotRequestId(v string) *ListBeebotIntentLgfResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetCode(v string) *ListBeebotIntentLgfResponseBody {
	s.Code = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetHttpStatusCode(v int32) *ListBeebotIntentLgfResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetLgfs(v []*ListBeebotIntentLgfResponseBodyLgfs) *ListBeebotIntentLgfResponseBody {
	s.Lgfs = v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetMessage(v string) *ListBeebotIntentLgfResponseBody {
	s.Message = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetPageNumber(v int32) *ListBeebotIntentLgfResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetPageSize(v int32) *ListBeebotIntentLgfResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetRequestId(v string) *ListBeebotIntentLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetSuccess(v bool) *ListBeebotIntentLgfResponseBody {
	s.Success = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) SetTotalCount(v int32) *ListBeebotIntentLgfResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBody) Validate() error {
	if s.Lgfs != nil {
		for _, item := range s.Lgfs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBeebotIntentLgfResponseBodyLgfs struct {
	// Creation Time
	//
	// example:
	//
	// 2025-04-21 10:54:18.+0800
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Utterance template ID
	//
	// example:
	//
	// 5666117
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// Updated At
	//
	// example:
	//
	// 2025-04-21 10:54:18.+0800
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// Content of the utterance template
	//
	// example:
	//
	// 我是一个问法模版
	RuleText *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s ListBeebotIntentLgfResponseBodyLgfs) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentLgfResponseBodyLgfs) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) GetLgfId() *int64 {
	return s.LgfId
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) GetRuleText() *string {
	return s.RuleText
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) SetCreateTime(v string) *ListBeebotIntentLgfResponseBodyLgfs {
	s.CreateTime = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) SetIntentId(v int64) *ListBeebotIntentLgfResponseBodyLgfs {
	s.IntentId = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) SetLgfId(v int64) *ListBeebotIntentLgfResponseBodyLgfs {
	s.LgfId = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) SetModifyTime(v string) *ListBeebotIntentLgfResponseBodyLgfs {
	s.ModifyTime = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) SetRuleText(v string) *ListBeebotIntentLgfResponseBodyLgfs {
	s.RuleText = &v
	return s
}

func (s *ListBeebotIntentLgfResponseBodyLgfs) Validate() error {
	return dara.Validate(s)
}
