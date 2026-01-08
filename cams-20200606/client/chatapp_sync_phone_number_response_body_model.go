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
	return dara.Validate(s)
}
