// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCardSmsDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCardId(v string) *GetCardSmsDetailsRequest
	GetBizCardId() *string
	SetBizDigitId(v string) *GetCardSmsDetailsRequest
	GetBizDigitId() *string
	SetBizSmsId(v string) *GetCardSmsDetailsRequest
	GetBizSmsId() *string
	SetCurrentPage(v int64) *GetCardSmsDetailsRequest
	GetCurrentPage() *int64
	SetOwnerId(v int64) *GetCardSmsDetailsRequest
	GetOwnerId() *int64
	SetPageSize(v int64) *GetCardSmsDetailsRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *GetCardSmsDetailsRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *GetCardSmsDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetCardSmsDetailsRequest
	GetResourceOwnerId() *int64
	SetSendDate(v string) *GetCardSmsDetailsRequest
	GetSendDate() *string
}

type GetCardSmsDetailsRequest struct {
	// Card SMS sending ID, which is the BizCardId field in the response when calling SendCardSms or SendBatchCardSms.
	//
	// example:
	//
	// 123456^0
	BizCardId *string `json:"BizCardId,omitempty" xml:"BizCardId,omitempty"`
	// Digital SMS sending ID, which is the BizDigitalId field in the response when calling SendCardSms or SendBatchCardSms.
	//
	// example:
	//
	// 12346^0
	BizDigitId *string `json:"BizDigitId,omitempty" xml:"BizDigitId,omitempty"`
	// Text SMS sending ID, which is the BizSmsId field in the response when calling SendCardSms or SendBatchCardSms.
	//
	// example:
	//
	// 1234576^0
	BizSmsId *string `json:"BizSmsId,omitempty" xml:"BizSmsId,omitempty"`
	// For paginated viewing of sending records, specify the current page number of the sending records.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// For paginated viewing of sending records, specify the number of card SMS records to display per page.
	//
	// The value range is 1~50.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Domestic phone number that received the SMS. Format: 11-digit phone number, for example, 1390000****.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Card SMS sending date, supports querying records from the last 30 days.
	//
	// Format: yyyyMMdd, for example, 20240112.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20240112
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
}

func (s GetCardSmsDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCardSmsDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetCardSmsDetailsRequest) GetBizCardId() *string {
	return s.BizCardId
}

func (s *GetCardSmsDetailsRequest) GetBizDigitId() *string {
	return s.BizDigitId
}

func (s *GetCardSmsDetailsRequest) GetBizSmsId() *string {
	return s.BizSmsId
}

func (s *GetCardSmsDetailsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *GetCardSmsDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCardSmsDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetCardSmsDetailsRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *GetCardSmsDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetCardSmsDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetCardSmsDetailsRequest) GetSendDate() *string {
	return s.SendDate
}

func (s *GetCardSmsDetailsRequest) SetBizCardId(v string) *GetCardSmsDetailsRequest {
	s.BizCardId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetBizDigitId(v string) *GetCardSmsDetailsRequest {
	s.BizDigitId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetBizSmsId(v string) *GetCardSmsDetailsRequest {
	s.BizSmsId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetCurrentPage(v int64) *GetCardSmsDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetOwnerId(v int64) *GetCardSmsDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetPageSize(v int64) *GetCardSmsDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetPhoneNumber(v string) *GetCardSmsDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetResourceOwnerAccount(v string) *GetCardSmsDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetResourceOwnerId(v int64) *GetCardSmsDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetCardSmsDetailsRequest) SetSendDate(v string) *GetCardSmsDetailsRequest {
	s.SendDate = &v
	return s
}

func (s *GetCardSmsDetailsRequest) Validate() error {
	return dara.Validate(s)
}
