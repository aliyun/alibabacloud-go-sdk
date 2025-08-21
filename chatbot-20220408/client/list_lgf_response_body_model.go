// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLgfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLgfs(v []*ListLgfResponseBodyLgfs) *ListLgfResponseBody
	GetLgfs() []*ListLgfResponseBodyLgfs
	SetPageNumber(v int32) *ListLgfResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLgfResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLgfResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListLgfResponseBody
	GetTotalCount() *int32
}

type ListLgfResponseBody struct {
	Lgfs []*ListLgfResponseBodyLgfs `json:"Lgfs,omitempty" xml:"Lgfs,omitempty" type:"Repeated"`
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
	// 34fg57h2gh5783
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLgfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLgfResponseBody) GoString() string {
	return s.String()
}

func (s *ListLgfResponseBody) GetLgfs() []*ListLgfResponseBodyLgfs {
	return s.Lgfs
}

func (s *ListLgfResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLgfResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLgfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLgfResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLgfResponseBody) SetLgfs(v []*ListLgfResponseBodyLgfs) *ListLgfResponseBody {
	s.Lgfs = v
	return s
}

func (s *ListLgfResponseBody) SetPageNumber(v int32) *ListLgfResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLgfResponseBody) SetPageSize(v int32) *ListLgfResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLgfResponseBody) SetRequestId(v string) *ListLgfResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLgfResponseBody) SetTotalCount(v int32) *ListLgfResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLgfResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLgfResponseBodyLgfs struct {
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 256756734345
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// LGF ID
	//
	// example:
	//
	// 123
	LgfId *int64 `json:"LgfId,omitempty" xml:"LgfId,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	RuleText   *string `json:"RuleText,omitempty" xml:"RuleText,omitempty"`
}

func (s ListLgfResponseBodyLgfs) String() string {
	return dara.Prettify(s)
}

func (s ListLgfResponseBodyLgfs) GoString() string {
	return s.String()
}

func (s *ListLgfResponseBodyLgfs) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLgfResponseBodyLgfs) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListLgfResponseBodyLgfs) GetLgfId() *int64 {
	return s.LgfId
}

func (s *ListLgfResponseBodyLgfs) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListLgfResponseBodyLgfs) GetRuleText() *string {
	return s.RuleText
}

func (s *ListLgfResponseBodyLgfs) SetCreateTime(v string) *ListLgfResponseBodyLgfs {
	s.CreateTime = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetIntentId(v int64) *ListLgfResponseBodyLgfs {
	s.IntentId = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetLgfId(v int64) *ListLgfResponseBodyLgfs {
	s.LgfId = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetModifyTime(v string) *ListLgfResponseBodyLgfs {
	s.ModifyTime = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) SetRuleText(v string) *ListLgfResponseBodyLgfs {
	s.RuleText = &v
	return s
}

func (s *ListLgfResponseBodyLgfs) Validate() error {
	return dara.Validate(s)
}
