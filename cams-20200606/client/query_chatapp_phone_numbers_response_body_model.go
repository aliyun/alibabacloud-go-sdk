// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryChatappPhoneNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryChatappPhoneNumbersResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryChatappPhoneNumbersResponseBody
	GetCode() *string
	SetData(v string) *QueryChatappPhoneNumbersResponseBody
	GetData() *string
	SetMessage(v string) *QueryChatappPhoneNumbersResponseBody
	GetMessage() *string
	SetPhoneNumbers(v []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers) *QueryChatappPhoneNumbersResponseBody
	GetPhoneNumbers() []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers
	SetRequestId(v string) *QueryChatappPhoneNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryChatappPhoneNumbersResponseBody
	GetSuccess() *bool
}

type QueryChatappPhoneNumbersResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of phone numbers.
	PhoneNumbers []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryChatappPhoneNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryChatappPhoneNumbersResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryChatappPhoneNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryChatappPhoneNumbersResponseBody) GetPhoneNumbers() []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	return s.PhoneNumbers
}

func (s *QueryChatappPhoneNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryChatappPhoneNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryChatappPhoneNumbersResponseBody) SetAccessDeniedDetail(v string) *QueryChatappPhoneNumbersResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetCode(v string) *QueryChatappPhoneNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetData(v string) *QueryChatappPhoneNumbersResponseBody {
	s.Data = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetMessage(v string) *QueryChatappPhoneNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetPhoneNumbers(v []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers) *QueryChatappPhoneNumbersResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetRequestId(v string) *QueryChatappPhoneNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) SetSuccess(v bool) *QueryChatappPhoneNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBody) Validate() error {
	if s.PhoneNumbers != nil {
		for _, item := range s.PhoneNumbers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbers struct {
	CallingConfigure *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure `json:"CallingConfigure,omitempty" xml:"CallingConfigure,omitempty" type:"Struct"`
	// The verification status of the phone number.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// example:
	//
	// N
	IsOfficial *string `json:"IsOfficial,omitempty" xml:"IsOfficial,omitempty"`
	// The messaging limit tier of the phone number.
	//
	// example:
	//
	// TIER_10K
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// The status of the name.
	//
	// example:
	//
	// Approval
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// The review status of the new name.
	//
	// example:
	//
	// Approval
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 861380000****
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number.
	//
	// example:
	//
	// GREEN
	QualityRating *string `json:"QualityRating,omitempty" xml:"QualityRating,omitempty"`
	// The status of the phone number.
	//
	// example:
	//
	// CONNECTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The callback URL for delivery receipts.
	//
	// example:
	//
	// https://ali.com/status
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The queue for delivery receipts of mobile terminated messages.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppStatus
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The callback URL for mobile originated messages.
	//
	// example:
	//
	// https://ali.com/inbound
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The queue for mobile originated messages.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppInbound
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The name used for the request.
	//
	// example:
	//
	// Alibaba
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbers) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetCallingConfigure() *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure {
	return s.CallingConfigure
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetCodeVerificationStatus() *string {
	return s.CodeVerificationStatus
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetIsOfficial() *string {
	return s.IsOfficial
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetMessagingLimitTier() *string {
	return s.MessagingLimitTier
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetNameStatus() *string {
	return s.NameStatus
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetNewNameStatus() *string {
	return s.NewNameStatus
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetQualityRating() *string {
	return s.QualityRating
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetStatus() *string {
	return s.Status
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetStatusCallbackUrl() *string {
	return s.StatusCallbackUrl
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetStatusQueue() *string {
	return s.StatusQueue
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetUpCallbackUrl() *string {
	return s.UpCallbackUrl
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetUpQueue() *string {
	return s.UpQueue
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) GetVerifiedName() *string {
	return s.VerifiedName
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetCallingConfigure(v *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.CallingConfigure = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetCodeVerificationStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.CodeVerificationStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetIsOfficial(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.IsOfficial = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetMessagingLimitTier(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.MessagingLimitTier = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetNameStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.NameStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetNewNameStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.NewNameStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetPhoneNumber(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.PhoneNumber = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetQualityRating(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.QualityRating = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.Status = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatusCallbackUrl(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.StatusCallbackUrl = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetStatusQueue(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.StatusQueue = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetUpCallbackUrl(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.UpCallbackUrl = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetUpQueue(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.UpQueue = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) SetVerifiedName(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbers {
	s.VerifiedName = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbers) Validate() error {
	if s.CallingConfigure != nil {
		if err := s.CallingConfigure.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure struct {
	Calling *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling `json:"Calling,omitempty" xml:"Calling,omitempty" type:"Struct"`
	// example:
	//
	// https://aliyun.com
	CallingCallbackUrl *string `json:"CallingCallbackUrl,omitempty" xml:"CallingCallbackUrl,omitempty"`
	// example:
	//
	// 1000
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) GetCalling() *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling {
	return s.Calling
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) GetCallingCallbackUrl() *string {
	return s.CallingCallbackUrl
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) SetCalling(v *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure {
	s.Calling = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) SetCallingCallbackUrl(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure {
	s.CallingCallbackUrl = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) SetMaxTalkTime(v int64) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure {
	s.MaxTalkTime = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigure) Validate() error {
	if s.Calling != nil {
		if err := s.Calling.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling struct {
	CallHours *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours `json:"CallHours,omitempty" xml:"CallHours,omitempty" type:"Struct"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) GetCallHours() *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	return s.CallHours
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) GetStatus() *string {
	return s.Status
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) SetCallHours(v *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling {
	s.CallHours = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) SetStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling {
	s.Status = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCalling) Validate() error {
	if s.CallHours != nil {
		if err := s.CallHours.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours struct {
	// example:
	//
	// DEFAULT
	CallIconVisibility *string `json:"CallIconVisibility,omitempty" xml:"CallIconVisibility,omitempty"`
	// example:
	//
	// ENABLED
	CallbackPermissionStatus *string                                                                                            `json:"CallbackPermissionStatus,omitempty" xml:"CallbackPermissionStatus,omitempty"`
	HolidaySchedule          []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule `json:"HolidaySchedule,omitempty" xml:"HolidaySchedule,omitempty" type:"Repeated"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimezoneId           *string                                                                                                 `json:"TimezoneId,omitempty" xml:"TimezoneId,omitempty"`
	WeeklyOperatingHours []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours `json:"WeeklyOperatingHours,omitempty" xml:"WeeklyOperatingHours,omitempty" type:"Repeated"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetCallIconVisibility() *string {
	return s.CallIconVisibility
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetCallbackPermissionStatus() *string {
	return s.CallbackPermissionStatus
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetHolidaySchedule() []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	return s.HolidaySchedule
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetStatus() *string {
	return s.Status
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetTimezoneId() *string {
	return s.TimezoneId
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetWeeklyOperatingHours() []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	return s.WeeklyOperatingHours
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetCallIconVisibility(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.CallIconVisibility = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetCallbackPermissionStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.CallbackPermissionStatus = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetHolidaySchedule(v []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.HolidaySchedule = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetStatus(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.Status = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetTimezoneId(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.TimezoneId = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetWeeklyOperatingHours(v []*QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.WeeklyOperatingHours = v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHours) Validate() error {
	if s.HolidaySchedule != nil {
		for _, item := range s.HolidaySchedule {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WeeklyOperatingHours != nil {
		for _, item := range s.WeeklyOperatingHours {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule struct {
	// example:
	//
	// 2026-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2359
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetDate() *string {
	return s.Date
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetDate(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.Date = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetEndTime(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.EndTime = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetStartTime(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.StartTime = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) Validate() error {
	return dara.Validate(s)
}

type QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours struct {
	// example:
	//
	// 2359
	CloseTime *string `json:"CloseTime,omitempty" xml:"CloseTime,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	DayOfWeek *string `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// example:
	//
	// 0000
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) String() string {
	return dara.Prettify(s)
}

func (s QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GoString() string {
	return s.String()
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetCloseTime() *string {
	return s.CloseTime
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetDayOfWeek() *string {
	return s.DayOfWeek
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetOpenTime() *string {
	return s.OpenTime
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetCloseTime(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.CloseTime = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetDayOfWeek(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.DayOfWeek = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetOpenTime(v string) *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.OpenTime = &v
	return s
}

func (s *QueryChatappPhoneNumbersResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) Validate() error {
	return dara.Validate(s)
}
