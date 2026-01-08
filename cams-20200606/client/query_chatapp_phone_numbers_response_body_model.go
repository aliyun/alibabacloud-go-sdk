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
	// The message returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The phone numbers.
	PhoneNumbers []*QueryChatappPhoneNumbersResponseBodyPhoneNumbers `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
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
	// The verification status of the phone number.
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
	// Valid values:
	//
	// 	- TIER_100K: 100,000
	//
	// 	- TIER_UNLIMITED: unlimited
	//
	// 	- TIER_250: 250
	//
	// 	- TIER_1K: 1,000
	//
	// 	- TIER_50: 50
	//
	// 	- TIER_10K: 10,000
	//
	// example:
	//
	// TIER_10
	MessagingLimitTier *string `json:"MessagingLimitTier,omitempty" xml:"MessagingLimitTier,omitempty"`
	// The review status of the name.
	//
	// example:
	//
	// Approval
	NameStatus *string `json:"NameStatus,omitempty" xml:"NameStatus,omitempty"`
	// The review status of the new display name of the enterprise.
	//
	// example:
	//
	// Approval
	NewNameStatus *string `json:"NewNameStatus,omitempty" xml:"NewNameStatus,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 8613800000
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The quality rating of the phone number.
	//
	// Valid values:
	//
	// 	- RED: low
	//
	// 	- YELLOW: medium
	//
	// 	- UNKNOWN: unknown
	//
	// 	- GREEN: high
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
	// CONNECTED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The URL that receives the status reports.
	//
	// example:
	//
	// https://ali.com/status
	StatusCallbackUrl *string `json:"StatusCallbackUrl,omitempty" xml:"StatusCallbackUrl,omitempty"`
	// The status report queue.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppStatus
	StatusQueue *string `json:"StatusQueue,omitempty" xml:"StatusQueue,omitempty"`
	// The URL that receives the MO messages.
	//
	// example:
	//
	// https://ali.com/inbound
	UpCallbackUrl *string `json:"UpCallbackUrl,omitempty" xml:"UpCallbackUrl,omitempty"`
	// The mobile originated (MO) message queue.
	//
	// example:
	//
	// Alicom-Queue-****-ChatAppInbound
	UpQueue *string `json:"UpQueue,omitempty" xml:"UpQueue,omitempty"`
	// The display name of the enterprise to which the phone number belongs.
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
	return dara.Validate(s)
}
