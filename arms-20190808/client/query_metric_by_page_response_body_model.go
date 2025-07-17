// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryMetricByPageResponseBody
	GetCode() *string
	SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody
	GetData() *QueryMetricByPageResponseBodyData
	SetMessage(v string) *QueryMetricByPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryMetricByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryMetricByPageResponseBody
	GetSuccess() *bool
}

type QueryMetricByPageResponseBody struct {
	// The HTTP status code returned for the request. Valid values:
	//
	// 	- 2XX: The request was successful.
	//
	// 	- 3XX: A redirection message was returned.
	//
	// 	- 4XX: The request was invalid.
	//
	// 	- 5XX: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the array object.
	Data *QueryMetricByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// StartTime is mandatory for this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 687F9CB7-4798-57BF-A6EE-E6CC76******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- `true`: The call was successful.
	//
	// 	- `false`: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryMetricByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryMetricByPageResponseBody) GetData() *QueryMetricByPageResponseBodyData {
	return s.Data
}

func (s *QueryMetricByPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryMetricByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMetricByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryMetricByPageResponseBody) SetCode(v string) *QueryMetricByPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetData(v *QueryMetricByPageResponseBodyData) *QueryMetricByPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryMetricByPageResponseBody) SetMessage(v string) *QueryMetricByPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetRequestId(v string) *QueryMetricByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricByPageResponseBody) SetSuccess(v bool) *QueryMetricByPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryMetricByPageResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryMetricByPageResponseBodyData struct {
	// Whether the paging query ends.
	//
	// true: end.
	//
	// false: Need to continue pagination (continue to query after CurrentPage+1).
	//
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The data entries returned.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryMetricByPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryMetricByPageResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *QueryMetricByPageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *QueryMetricByPageResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryMetricByPageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryMetricByPageResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryMetricByPageResponseBodyData) SetCompleted(v bool) *QueryMetricByPageResponseBodyData {
	s.Completed = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetItems(v []map[string]interface{}) *QueryMetricByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPage(v int32) *QueryMetricByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetPageSize(v int32) *QueryMetricByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) SetTotal(v int32) *QueryMetricByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryMetricByPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
