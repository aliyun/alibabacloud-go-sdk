// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyntheticDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListSyntheticDetailResponseBody
	GetCode() *int64
	SetData(v *ListSyntheticDetailResponseBodyData) *ListSyntheticDetailResponseBody
	GetData() *ListSyntheticDetailResponseBodyData
	SetMessage(v string) *ListSyntheticDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSyntheticDetailResponseBody
	GetRequestId() *string
}

type ListSyntheticDetailResponseBody struct {
	// The HTTP status code returned. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *ListSyntheticDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 70675725-8F11-4817-8106-CFE0AD71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSyntheticDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListSyntheticDetailResponseBody) GetData() *ListSyntheticDetailResponseBodyData {
	return s.Data
}

func (s *ListSyntheticDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSyntheticDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSyntheticDetailResponseBody) SetCode(v int64) *ListSyntheticDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ListSyntheticDetailResponseBody) SetData(v *ListSyntheticDetailResponseBodyData) *ListSyntheticDetailResponseBody {
	s.Data = v
	return s
}

func (s *ListSyntheticDetailResponseBody) SetMessage(v string) *ListSyntheticDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ListSyntheticDetailResponseBody) SetRequestId(v string) *ListSyntheticDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSyntheticDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSyntheticDetailResponseBodyData struct {
	// The list of results.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A reserved field.
	//
	// example:
	//
	// null
	TaskCreateTime *int64 `json:"TaskCreateTime,omitempty" xml:"TaskCreateTime,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 12
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListSyntheticDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSyntheticDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSyntheticDetailResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *ListSyntheticDetailResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ListSyntheticDetailResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSyntheticDetailResponseBodyData) GetTaskCreateTime() *int64 {
	return s.TaskCreateTime
}

func (s *ListSyntheticDetailResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListSyntheticDetailResponseBodyData) SetItems(v []map[string]interface{}) *ListSyntheticDetailResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSyntheticDetailResponseBodyData) SetPage(v int32) *ListSyntheticDetailResponseBodyData {
	s.Page = &v
	return s
}

func (s *ListSyntheticDetailResponseBodyData) SetPageSize(v int32) *ListSyntheticDetailResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSyntheticDetailResponseBodyData) SetTaskCreateTime(v int64) *ListSyntheticDetailResponseBodyData {
	s.TaskCreateTime = &v
	return s
}

func (s *ListSyntheticDetailResponseBodyData) SetTotal(v int32) *ListSyntheticDetailResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListSyntheticDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
