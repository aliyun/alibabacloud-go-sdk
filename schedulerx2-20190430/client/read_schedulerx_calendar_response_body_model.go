// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) *ReadSchedulerxCalendarResponseBody
	GetAccessDeniedDetail() *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail
	SetCode(v int32) *ReadSchedulerxCalendarResponseBody
	GetCode() *int32
	SetData(v *ReadSchedulerxCalendarResponseBodyData) *ReadSchedulerxCalendarResponseBody
	GetData() *ReadSchedulerxCalendarResponseBodyData
	SetMessage(v string) *ReadSchedulerxCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadSchedulerxCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadSchedulerxCalendarResponseBody
	GetSuccess() *bool
}

type ReadSchedulerxCalendarResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// *
	Data *ReadSchedulerxCalendarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// unknown exception occurred
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// C8E5FB4A-6D8D-424D-9AAA-4FE06BB74FF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadSchedulerxCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarResponseBody) GetAccessDeniedDetail() *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReadSchedulerxCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReadSchedulerxCalendarResponseBody) GetData() *ReadSchedulerxCalendarResponseBodyData {
	return s.Data
}

func (s *ReadSchedulerxCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadSchedulerxCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadSchedulerxCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadSchedulerxCalendarResponseBody) SetAccessDeniedDetail(v *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) *ReadSchedulerxCalendarResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) SetCode(v int32) *ReadSchedulerxCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) SetData(v *ReadSchedulerxCalendarResponseBodyData) *ReadSchedulerxCalendarResponseBody {
	s.Data = v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) SetMessage(v string) *ReadSchedulerxCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) SetRequestId(v string) *ReadSchedulerxCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) SetSuccess(v bool) *ReadSchedulerxCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadSchedulerxCalendarResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// edas:ReadSchedulerxCalendar
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The account of the principal.
	//
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaDEssEE4MDg4NTQyLTVGMTYtNTFEQy1CODJCLUFFMDY4NUVDQ0ZBQQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The permission denial type.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReadSchedulerxCalendarResponseBodyData struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 15
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// When there is more data to retrieve, the server returns a nextToken. You can use this token in a subsequent request to continue reading from where you left off.
	//
	// example:
	//
	// O39nXKu5XafATl3/cJjSJw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// *
	Records []*ReadSchedulerxCalendarResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ReadSchedulerxCalendarResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ReadSchedulerxCalendarResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ReadSchedulerxCalendarResponseBodyData) GetRecords() []*ReadSchedulerxCalendarResponseBodyDataRecords {
	return s.Records
}

func (s *ReadSchedulerxCalendarResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ReadSchedulerxCalendarResponseBodyData) SetMaxResults(v int32) *ReadSchedulerxCalendarResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyData) SetNextToken(v string) *ReadSchedulerxCalendarResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyData) SetRecords(v []*ReadSchedulerxCalendarResponseBodyDataRecords) *ReadSchedulerxCalendarResponseBodyData {
	s.Records = v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyData) SetTotal(v int64) *ReadSchedulerxCalendarResponseBodyData {
	s.Total = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadSchedulerxCalendarResponseBodyDataRecords struct {
	// The calendar name.
	//
	// example:
	//
	// 2025workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// The creator.
	//
	// example:
	//
	// 1827811800526xxx
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// The months and days.
	//
	// example:
	//
	// [
	//
	//   {"month":1,"days":[3,4,5,6,9,10,11,12,13,16,17,18,19,20,28,29,30,31]},
	//
	//   {"month":2,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28]},
	//
	//   {"month":3,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},
	//
	//   {"month":4,"days":[3,4,6,7,10,11,12,13,14,17,18,19,20,21,23,24,25,26,27,28]},
	//
	//   {"month":5,"days":[4,5,6,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]},
	//
	//   {"month":6,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,25,26,27,28,29,30]},
	//
	//   {"month":7,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},
	//
	//   {"month":8,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},
	//
	//   {"month":9,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28]},
	//
	//   {"month":10,"days":[7,8,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30,31]},
	//
	//   {"month":11,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30]},
	//
	//   {"month":12,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]}
	//
	// ]
	MonthDaysContent *string `json:"MonthDaysContent,omitempty" xml:"MonthDaysContent,omitempty"`
	// Indicates whether it is a system calendar.
	//
	// example:
	//
	// false
	SystemCalendar *bool `json:"SystemCalendar,omitempty" xml:"SystemCalendar,omitempty"`
	// The year.
	//
	// example:
	//
	// 2025
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ReadSchedulerxCalendarResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxCalendarResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) GetCalendarName() *string {
	return s.CalendarName
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) GetCreator() *string {
	return s.Creator
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) GetMonthDaysContent() *string {
	return s.MonthDaysContent
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) GetSystemCalendar() *bool {
	return s.SystemCalendar
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) GetYear() *int32 {
	return s.Year
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) SetCalendarName(v string) *ReadSchedulerxCalendarResponseBodyDataRecords {
	s.CalendarName = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) SetCreator(v string) *ReadSchedulerxCalendarResponseBodyDataRecords {
	s.Creator = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) SetMonthDaysContent(v string) *ReadSchedulerxCalendarResponseBodyDataRecords {
	s.MonthDaysContent = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) SetSystemCalendar(v bool) *ReadSchedulerxCalendarResponseBodyDataRecords {
	s.SystemCalendar = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) SetYear(v int32) *ReadSchedulerxCalendarResponseBodyDataRecords {
	s.Year = &v
	return s
}

func (s *ReadSchedulerxCalendarResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
