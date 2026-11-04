// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTimeTemplateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTimeTemplateDetailResponseBody
	GetCode() *string
	SetData(v *QueryTimeTemplateDetailResponseBodyData) *QueryTimeTemplateDetailResponseBody
	GetData() *QueryTimeTemplateDetailResponseBodyData
	SetErrorMessage(v string) *QueryTimeTemplateDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *QueryTimeTemplateDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTimeTemplateDetailResponseBody
	GetSuccess() *bool
}

type QueryTimeTemplateDetailResponseBody struct {
	// The response code:
	//
	// - **200**: The call is successful.
	//
	// - Other values: error codes. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The time template information returned when the call was successful.
	Data *QueryTimeTemplateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned when the call failed.
	//
	// example:
	//
	// template owner error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 06DC77A0-4622-42DB-9EE0-25FIOHS82JK1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful:
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

func (s QueryTimeTemplateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTimeTemplateDetailResponseBody) GetData() *QueryTimeTemplateDetailResponseBodyData {
	return s.Data
}

func (s *QueryTimeTemplateDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *QueryTimeTemplateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTimeTemplateDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTimeTemplateDetailResponseBody) SetCode(v string) *QueryTimeTemplateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetData(v *QueryTimeTemplateDetailResponseBodyData) *QueryTimeTemplateDetailResponseBody {
	s.Data = v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetErrorMessage(v string) *QueryTimeTemplateDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetRequestId(v string) *QueryTimeTemplateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) SetSuccess(v bool) *QueryTimeTemplateDetailResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTimeTemplateDetailResponseBodyData struct {
	// Indicates whether 24/7 recording is enabled:
	//
	// - **1**: Yes.
	//
	// - **0**: No. For detailed configuration, see the **TimeSectionList*	- parameter.
	//
	// example:
	//
	// 0
	AllDay *int32 `json:"AllDay,omitempty" xml:"AllDay,omitempty"`
	// Indicates whether the template is the default template:
	//
	// - **1**: Default template.
	//
	// - **0**: Custom template.
	//
	// example:
	//
	// 1
	Default *int32 `json:"Default,omitempty" xml:"Default,omitempty"`
	// The name of the time template.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the time template.
	//
	// example:
	//
	// 36f06322442e4e3f85a5ae8862c9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The list of time sections in the time template.
	TimeSectionList []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList `json:"TimeSectionList,omitempty" xml:"TimeSectionList,omitempty" type:"Repeated"`
}

func (s QueryTimeTemplateDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyData) GetAllDay() *int32 {
	return s.AllDay
}

func (s *QueryTimeTemplateDetailResponseBodyData) GetDefault() *int32 {
	return s.Default
}

func (s *QueryTimeTemplateDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *QueryTimeTemplateDetailResponseBodyData) GetTemplateId() *string {
	return s.TemplateId
}

func (s *QueryTimeTemplateDetailResponseBodyData) GetTimeSectionList() []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	return s.TimeSectionList
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetAllDay(v int32) *QueryTimeTemplateDetailResponseBodyData {
	s.AllDay = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetDefault(v int32) *QueryTimeTemplateDetailResponseBodyData {
	s.Default = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetName(v string) *QueryTimeTemplateDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetTemplateId(v string) *QueryTimeTemplateDetailResponseBodyData {
	s.TemplateId = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) SetTimeSectionList(v []*QueryTimeTemplateDetailResponseBodyDataTimeSectionList) *QueryTimeTemplateDetailResponseBodyData {
	s.TimeSectionList = v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyData) Validate() error {
	if s.TimeSectionList != nil {
		for _, item := range s.TimeSectionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryTimeTemplateDetailResponseBodyDataTimeSectionList struct {
	// The start time of each day, in seconds.
	//
	// example:
	//
	// 20000
	Begin *int32 `json:"Begin,omitempty" xml:"Begin,omitempty"`
	// The day of the week. Valid values: 0 to 6, where 0 indicates Sunday and 6 indicates Saturday.
	//
	// example:
	//
	// 0
	DayOfWeek *int32 `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// The end time of each day, in seconds.
	//
	// example:
	//
	// 60000
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) String() string {
	return dara.Prettify(s)
}

func (s QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GoString() string {
	return s.String()
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GetBegin() *int32 {
	return s.Begin
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GetDayOfWeek() *int32 {
	return s.DayOfWeek
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) GetEnd() *int32 {
	return s.End
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetBegin(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.Begin = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetDayOfWeek(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.DayOfWeek = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) SetEnd(v int32) *QueryTimeTemplateDetailResponseBodyDataTimeSectionList {
	s.End = &v
	return s
}

func (s *QueryTimeTemplateDetailResponseBodyDataTimeSectionList) Validate() error {
	return dara.Validate(s)
}
