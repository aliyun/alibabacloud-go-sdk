// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumOcuStatisticDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetRumOcuStatisticDataResponseBody
	GetCode() *int64
	SetData(v *GetRumOcuStatisticDataResponseBodyData) *GetRumOcuStatisticDataResponseBody
	GetData() *GetRumOcuStatisticDataResponseBodyData
	SetMessage(v string) *GetRumOcuStatisticDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumOcuStatisticDataResponseBody
	GetRequestId() *string
}

type GetRumOcuStatisticDataResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *GetRumOcuStatisticDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// null
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 626037F5-FDEB-45B0-804C-B3C92797****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRumOcuStatisticDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetRumOcuStatisticDataResponseBody) GetData() *GetRumOcuStatisticDataResponseBodyData {
	return s.Data
}

func (s *GetRumOcuStatisticDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumOcuStatisticDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumOcuStatisticDataResponseBody) SetCode(v int64) *GetRumOcuStatisticDataResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBody) SetData(v *GetRumOcuStatisticDataResponseBodyData) *GetRumOcuStatisticDataResponseBody {
	s.Data = v
	return s
}

func (s *GetRumOcuStatisticDataResponseBody) SetMessage(v string) *GetRumOcuStatisticDataResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBody) SetRequestId(v string) *GetRumOcuStatisticDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRumOcuStatisticDataResponseBodyData struct {
	// Indicates whether the next page exists.
	//
	// example:
	//
	// true
	Complete *bool `json:"Complete,omitempty" xml:"Complete,omitempty"`
	// The queried data.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRumOcuStatisticDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRumOcuStatisticDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRumOcuStatisticDataResponseBodyData) GetComplete() *bool {
	return s.Complete
}

func (s *GetRumOcuStatisticDataResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetRumOcuStatisticDataResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetRumOcuStatisticDataResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRumOcuStatisticDataResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *GetRumOcuStatisticDataResponseBodyData) SetComplete(v bool) *GetRumOcuStatisticDataResponseBodyData {
	s.Complete = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBodyData) SetItems(v []map[string]interface{}) *GetRumOcuStatisticDataResponseBodyData {
	s.Items = v
	return s
}

func (s *GetRumOcuStatisticDataResponseBodyData) SetPage(v int32) *GetRumOcuStatisticDataResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBodyData) SetPageSize(v int32) *GetRumOcuStatisticDataResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBodyData) SetTotal(v int32) *GetRumOcuStatisticDataResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRumOcuStatisticDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
