// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatappSyncPhoneNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ChatappSyncPhoneNumberResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ChatappSyncPhoneNumberResponseBody
	GetCode() *string
	SetMessage(v string) *ChatappSyncPhoneNumberResponseBody
	GetMessage() *string
	SetPhoneNumbers(v []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers) *ChatappSyncPhoneNumberResponseBody
	GetPhoneNumbers() []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers
	SetRequestId(v string) *ChatappSyncPhoneNumberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChatappSyncPhoneNumberResponseBody
	GetSuccess() *bool
}

type ChatappSyncPhoneNumberResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The HTTP status code returned.
	//
	// 	- A value of OK indicates that the call is successful.
	//
	// 	- Other values indicate that the call fails. For more information, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// None.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The phone numbers.
	PhoneNumbers []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- **true**: The call was successful.
	//
	// 	- **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBody) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ChatappSyncPhoneNumberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ChatappSyncPhoneNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ChatappSyncPhoneNumberResponseBody) GetPhoneNumbers() []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	return s.PhoneNumbers
}

func (s *ChatappSyncPhoneNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChatappSyncPhoneNumberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChatappSyncPhoneNumberResponseBody) SetAccessDeniedDetail(v string) *ChatappSyncPhoneNumberResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetCode(v string) *ChatappSyncPhoneNumberResponseBody {
	s.Code = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetMessage(v string) *ChatappSyncPhoneNumberResponseBody {
	s.Message = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetPhoneNumbers(v []*ChatappSyncPhoneNumberResponseBodyPhoneNumbers) *ChatappSyncPhoneNumberResponseBody {
	s.PhoneNumbers = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetRequestId(v string) *ChatappSyncPhoneNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) SetSuccess(v bool) *ChatappSyncPhoneNumberResponseBody {
	s.Success = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBody) Validate() error {
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

type ChatappSyncPhoneNumberResponseBodyPhoneNumbers struct {
	CallingConfigure *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure `json:"CallingConfigure,omitempty" xml:"CallingConfigure,omitempty" type:"Struct"`
	// The verification state of the phone number.
	//
	// Valid values:
	//
	// 	- REVOKED: The review application is revoked.
	//
	// 	- MORE_INFORMATION_REQUESTED: More information needs to be provided.
	//
	// 	- VERIFIED: The phone number passes the verification.
	//
	// 	- REJECTED: The phone number fails to pass the verification.
	//
	// example:
	//
	// VERIFIED
	CodeVerificationStatus *string `json:"CodeVerificationStatus,omitempty" xml:"CodeVerificationStatus,omitempty"`
	// Indicates whether it is a WhatsApp Official Business Account (OBA).
	//
	// example:
	//
	// N
	IsOfficial *string `json:"IsOfficial,omitempty" xml:"IsOfficial,omitempty"`
	// The number of phone numbers to which messages can be sent in a day.
	//
	// example:
	//
	// TIER_10
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// The review status of the business display name.
	//
	// example:
	//
	// Approval
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// The review status of the new business display name.
	//
	// example:
	//
	// Approval
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800001234
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number.
	//
	// Valid values:
	//
	// 	- RED
	//
	// 	- YELLOW
	//
	// 	- GREEN
	//
	// example:
	//
	// GREEN
	QualityRating *string `json:"QualityRating,omitempty" xml:"QualityRating,omitempty"`
	// The state of the phone number.
	//
	// Valid values:
	//
	// 	- MIGRATED
	//
	// 	- FLAGGED
	//
	// 	- DISCONNECTED
	//
	// 	- UNVERIFIED
	//
	// 	- BANNED
	//
	// 	- RATE_LIMITED
	//
	// 	- PENDING
	//
	// 	- CONNECTED
	//
	// 	- UNKNOWN
	//
	// 	- DELETED
	//
	// 	- RESTRICTED
	//
	// example:
	//
	// PENDING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The callback URL to which status reports are sent by using HTTP callbacks.
	//
	// example:
	//
	// https://www.alibaba.com/status
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The status report queue.
	//
	// example:
	//
	// alicom-09399200-queue
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The URL that receives the MO messages.
	//
	// example:
	//
	// https://www.alibaba.com/inbound
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The mobile originated (MO) message queue.
	//
	// example:
	//
	// alicom-09399200-queue
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The display name of the business to which the phone number belongs.
	//
	// example:
	//
	// Alibaba
	VerifiedName *string `json:"VerifiedName,omitempty" xml:"VerifiedName,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbers) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetCallingConfigure() *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure {
	return s.CallingConfigure
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetCodeVerificationStatus() *string {
	return s.CodeVerificationStatus
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetIsOfficial() *string {
	return s.IsOfficial
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetMessagingLimitTier() *string {
	return s.MessagingLimitTier
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetNameStatus() *string {
	return s.NameStatus
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetNewNameStatus() *string {
	return s.NewNameStatus
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetQualityRating() *string {
	return s.QualityRating
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetStatus() *string {
	return s.Status
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetStatusCallbackUrl() *string {
	return s.StatusCallbackUrl
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetStatusQueue() *string {
	return s.StatusQueue
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetUpCallbackUrl() *string {
	return s.UpCallbackUrl
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetUpQueue() *string {
	return s.UpQueue
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) GetVerifiedName() *string {
	return s.VerifiedName
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetCallingConfigure(v *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.CallingConfigure = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetCodeVerificationStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.CodeVerificationStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetIsOfficial(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.IsOfficial = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetMessagingLimitTier(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.MessagingLimitTier = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetNameStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.NameStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetNewNameStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.NewNameStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetPhoneNumber(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.PhoneNumber = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetQualityRating(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.QualityRating = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.Status = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatusCallbackUrl(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.StatusCallbackUrl = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetStatusQueue(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.StatusQueue = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetUpCallbackUrl(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.UpCallbackUrl = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetUpQueue(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.UpQueue = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) SetVerifiedName(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbers {
	s.VerifiedName = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbers) Validate() error {
	if s.CallingConfigure != nil {
		if err := s.CallingConfigure.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure struct {
	Calling *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling `json:"Calling,omitempty" xml:"Calling,omitempty" type:"Struct"`
	// example:
	//
	// http://aliyun.com
	CallingCallbackUrl *string `json:"CallingCallbackUrl,omitempty" xml:"CallingCallbackUrl,omitempty"`
	// example:
	//
	// 100
	MaxTalkTime *int64 `json:"MaxTalkTime,omitempty" xml:"MaxTalkTime,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) GetCalling() *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling {
	return s.Calling
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) GetCallingCallbackUrl() *string {
	return s.CallingCallbackUrl
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) GetMaxTalkTime() *int64 {
	return s.MaxTalkTime
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) SetCalling(v *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure {
	s.Calling = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) SetCallingCallbackUrl(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure {
	s.CallingCallbackUrl = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) SetMaxTalkTime(v int64) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure {
	s.MaxTalkTime = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigure) Validate() error {
	if s.Calling != nil {
		if err := s.Calling.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling struct {
	CallHours *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours `json:"CallHours,omitempty" xml:"CallHours,omitempty" type:"Struct"`
	// example:
	//
	// DEFAULT
	CallIconVisibility *string `json:"CallIconVisibility,omitempty" xml:"CallIconVisibility,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	CallbackPermissionStatus *string `json:"CallbackPermissionStatus,omitempty" xml:"CallbackPermissionStatus,omitempty"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) GetCallHours() *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	return s.CallHours
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) GetCallIconVisibility() *string {
	return s.CallIconVisibility
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) GetCallbackPermissionStatus() *string {
	return s.CallbackPermissionStatus
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) GetStatus() *string {
	return s.Status
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) SetCallHours(v *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling {
	s.CallHours = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) SetCallIconVisibility(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling {
	s.CallIconVisibility = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) SetCallbackPermissionStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling {
	s.CallbackPermissionStatus = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) SetStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling {
	s.Status = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCalling) Validate() error {
	if s.CallHours != nil {
		if err := s.CallHours.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours struct {
	HolidaySchedule []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule `json:"HolidaySchedule,omitempty" xml:"HolidaySchedule,omitempty" type:"Repeated"`
	// example:
	//
	// ENABLED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimezoneId           *string                                                                                               `json:"TimezoneId,omitempty" xml:"TimezoneId,omitempty"`
	WeeklyOperatingHours []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours `json:"WeeklyOperatingHours,omitempty" xml:"WeeklyOperatingHours,omitempty" type:"Repeated"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetHolidaySchedule() []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	return s.HolidaySchedule
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetStatus() *string {
	return s.Status
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetTimezoneId() *string {
	return s.TimezoneId
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) GetWeeklyOperatingHours() []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	return s.WeeklyOperatingHours
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetHolidaySchedule(v []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.HolidaySchedule = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetStatus(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.Status = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetTimezoneId(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.TimezoneId = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) SetWeeklyOperatingHours(v []*ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours {
	s.WeeklyOperatingHours = v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHours) Validate() error {
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

type ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule struct {
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

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetDate() *string {
	return s.Date
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetEndTime() *string {
	return s.EndTime
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) GetStartTime() *string {
	return s.StartTime
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetDate(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.Date = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetEndTime(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.EndTime = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) SetStartTime(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule {
	s.StartTime = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursHolidaySchedule) Validate() error {
	return dara.Validate(s)
}

type ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours struct {
	// example:
	//
	// 2359
	CloseTime *string `json:"CloseTime,omitempty" xml:"CloseTime,omitempty"`
	// example:
	//
	// MONDAY
	DayOfWeek *string `json:"DayOfWeek,omitempty" xml:"DayOfWeek,omitempty"`
	// example:
	//
	// 0000
	OpenTime *string `json:"OpenTime,omitempty" xml:"OpenTime,omitempty"`
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) String() string {
	return dara.Prettify(s)
}

func (s ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GoString() string {
	return s.String()
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetCloseTime() *string {
	return s.CloseTime
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetDayOfWeek() *string {
	return s.DayOfWeek
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) GetOpenTime() *string {
	return s.OpenTime
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetCloseTime(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.CloseTime = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetDayOfWeek(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.DayOfWeek = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) SetOpenTime(v string) *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours {
	s.OpenTime = &v
	return s
}

func (s *ChatappSyncPhoneNumberResponseBodyPhoneNumbersCallingConfigureCallingCallHoursWeeklyOperatingHours) Validate() error {
	return dara.Validate(s)
}
