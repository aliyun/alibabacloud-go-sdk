// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumDataForPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRumDataForPageResponseBody
	GetCode() *string
	SetData(v *GetRumDataForPageResponseBodyData) *GetRumDataForPageResponseBody
	GetData() *GetRumDataForPageResponseBodyData
	SetHttpStatusCode(v string) *GetRumDataForPageResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetRumDataForPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumDataForPageResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRumDataForPageResponseBody
	GetSuccess() *string
}

type GetRumDataForPageResponseBody struct {
	// The responses code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	Data *GetRumDataForPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// StartTime is mandatory for this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRumDataForPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumDataForPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumDataForPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRumDataForPageResponseBody) GetData() *GetRumDataForPageResponseBodyData {
	return s.Data
}

func (s *GetRumDataForPageResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetRumDataForPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumDataForPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumDataForPageResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRumDataForPageResponseBody) SetCode(v string) *GetRumDataForPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumDataForPageResponseBody) SetData(v *GetRumDataForPageResponseBodyData) *GetRumDataForPageResponseBody {
	s.Data = v
	return s
}

func (s *GetRumDataForPageResponseBody) SetHttpStatusCode(v string) *GetRumDataForPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRumDataForPageResponseBody) SetMessage(v string) *GetRumDataForPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumDataForPageResponseBody) SetRequestId(v string) *GetRumDataForPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumDataForPageResponseBody) SetSuccess(v string) *GetRumDataForPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetRumDataForPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRumDataForPageResponseBodyData struct {
	// A reserved parameter. Ignore this parameter.
	//
	// example:
	//
	// null
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
	// Indicates whether the query ends. Valid values: true and false.
	//
	// example:
	//
	// true
	Completion *string `json:"Completion,omitempty" xml:"Completion,omitempty"`
	// The queried data.
	Items []map[string]interface{} `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *string `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A reserved parameter. Ignore this parameter.
	//
	// example:
	//
	// null
	Preference *string `json:"Preference,omitempty" xml:"Preference,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 7
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetRumDataForPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRumDataForPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRumDataForPageResponseBodyData) GetAuthentication() *string {
	return s.Authentication
}

func (s *GetRumDataForPageResponseBodyData) GetCompletion() *string {
	return s.Completion
}

func (s *GetRumDataForPageResponseBodyData) GetItems() []map[string]interface{} {
	return s.Items
}

func (s *GetRumDataForPageResponseBodyData) GetPage() *string {
	return s.Page
}

func (s *GetRumDataForPageResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *GetRumDataForPageResponseBodyData) GetPreference() *string {
	return s.Preference
}

func (s *GetRumDataForPageResponseBodyData) GetTotal() *string {
	return s.Total
}

func (s *GetRumDataForPageResponseBodyData) SetAuthentication(v string) *GetRumDataForPageResponseBodyData {
	s.Authentication = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetCompletion(v string) *GetRumDataForPageResponseBodyData {
	s.Completion = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetItems(v []map[string]interface{}) *GetRumDataForPageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetPage(v string) *GetRumDataForPageResponseBodyData {
	s.Page = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetPageSize(v string) *GetRumDataForPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetPreference(v string) *GetRumDataForPageResponseBodyData {
	s.Preference = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) SetTotal(v string) *GetRumDataForPageResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetRumDataForPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
