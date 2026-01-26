// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppApiByPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetAppApiByPageResponseBody
	GetCode() *int32
	SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody
	GetData() *GetAppApiByPageResponseBodyData
	SetMessage(v string) *GetAppApiByPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAppApiByPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppApiByPageResponseBody
	GetSuccess() *bool
}

type GetAppApiByPageResponseBody struct {
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
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *GetAppApiByPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B6A00968-82A8-4F14-9D1B-B53827DB****
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

func (s GetAppApiByPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAppApiByPageResponseBody) GetData() *GetAppApiByPageResponseBodyData {
	return s.Data
}

func (s *GetAppApiByPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAppApiByPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppApiByPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppApiByPageResponseBody) SetCode(v int32) *GetAppApiByPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetData(v *GetAppApiByPageResponseBodyData) *GetAppApiByPageResponseBody {
	s.Data = v
	return s
}

func (s *GetAppApiByPageResponseBody) SetMessage(v string) *GetAppApiByPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetRequestId(v string) *GetAppApiByPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppApiByPageResponseBody) SetSuccess(v bool) *GetAppApiByPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppApiByPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAppApiByPageResponseBodyData struct {
	// Is completed.
	//
	// example:
	//
	// false
	Completed *bool `json:"Completed,omitempty" xml:"Completed,omitempty"`
	// The data entries.
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 0
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetAppApiByPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAppApiByPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAppApiByPageResponseBodyData) GetCompleted() *bool {
	return s.Completed
}

func (s *GetAppApiByPageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetAppApiByPageResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *GetAppApiByPageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAppApiByPageResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *GetAppApiByPageResponseBodyData) SetCompleted(v bool) *GetAppApiByPageResponseBodyData {
	s.Completed = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetItems(v []map[string]interface{}) *GetAppApiByPageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPage(v int32) *GetAppApiByPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetPageSize(v int32) *GetAppApiByPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) SetTotal(v string) *GetAppApiByPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetAppApiByPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
