// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordPlansResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryRecordPlansResponseBody
	GetCode() *string
	SetData(v *QueryRecordPlansResponseBodyData) *QueryRecordPlansResponseBody
	GetData() *QueryRecordPlansResponseBodyData
	SetErrorMessage(v string) *QueryRecordPlansResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryRecordPlansResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryRecordPlansResponseBody
	GetSuccess() *bool
}

type QueryRecordPlansResponseBody struct {
	// The error code returned if the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query information.
	Data *QueryRecordPlansResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// tenant auth error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryRecordPlansResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlansResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryRecordPlansResponseBody) GetData() *QueryRecordPlansResponseBodyData {
	return s.Data
}

func (s *QueryRecordPlansResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryRecordPlansResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecordPlansResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryRecordPlansResponseBody) SetCode(v string) *QueryRecordPlansResponseBody {
	s.Code = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetData(v *QueryRecordPlansResponseBodyData) *QueryRecordPlansResponseBody {
	s.Data = v
	return s
}

func (s *QueryRecordPlansResponseBody) SetErrorMessage(v string) *QueryRecordPlansResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetRequestId(v string) *QueryRecordPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecordPlansResponseBody) SetSuccess(v bool) *QueryRecordPlansResponseBody {
	s.Success = &v
	return s
}

func (s *QueryRecordPlansResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryRecordPlansResponseBodyData struct {
	// The list of returned recording plan templates.
	List []*QueryRecordPlansResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The number of records returned on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of recording plans returned.
	//
	// example:
	//
	// 2
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s QueryRecordPlansResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlansResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBodyData) GetList() []*QueryRecordPlansResponseBodyDataList {
	return s.List
}

func (s *QueryRecordPlansResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *QueryRecordPlansResponseBodyData) GetPageCount() *int32 {
	return s.PageCount
}

func (s *QueryRecordPlansResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryRecordPlansResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *QueryRecordPlansResponseBodyData) SetList(v []*QueryRecordPlansResponseBodyDataList) *QueryRecordPlansResponseBodyData {
	s.List = v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPage(v int32) *QueryRecordPlansResponseBodyData {
	s.Page = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPageCount(v int32) *QueryRecordPlansResponseBodyData {
	s.PageCount = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetPageSize(v int32) *QueryRecordPlansResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) SetTotal(v int32) *QueryRecordPlansResponseBodyData {
	s.Total = &v
	return s
}

func (s *QueryRecordPlansResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryRecordPlansResponseBodyDataList struct {
	// The name of the recording plan.
	//
	// example:
	//
	// plan1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the recording plan.
	//
	// example:
	//
	// 6c5a397bc17c4b48ace86f0f0c30****
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the time template associated with the recording plan.
	//
	// example:
	//
	// 693e244a038f4255bb2f5db32874****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s QueryRecordPlansResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordPlansResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryRecordPlansResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *QueryRecordPlansResponseBodyDataList) GetPlanId() *string {
	return s.PlanId
}

func (s *QueryRecordPlansResponseBodyDataList) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryRecordPlansResponseBodyDataList) SetName(v string) *QueryRecordPlansResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetPlanId(v string) *QueryRecordPlansResponseBodyDataList {
	s.PlanId = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) SetTemplateId(v string) *QueryRecordPlansResponseBodyDataList {
	s.TemplateId = &v
	return s
}

func (s *QueryRecordPlansResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
