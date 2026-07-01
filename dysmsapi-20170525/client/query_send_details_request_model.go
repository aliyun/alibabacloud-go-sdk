// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySendDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QuerySendDetailsRequest
	GetBizId() *string
	SetCurrentPage(v int64) *QuerySendDetailsRequest
	GetCurrentPage() *int64
	SetOwnerId(v int64) *QuerySendDetailsRequest
	GetOwnerId() *int64
	SetPageSize(v int64) *QuerySendDetailsRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *QuerySendDetailsRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *QuerySendDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySendDetailsRequest
	GetResourceOwnerId() *int64
	SetSendDate(v string) *QuerySendDetailsRequest
	GetSendDate() *string
}

type QuerySendDetailsRequest struct {
	// The delivery receipt ID. This is the `BizId` returned in the response when you call the [SendSms](https://help.aliyun.com/document_detail/419273.html) or [SendBatchSms](https://help.aliyun.com/document_detail/419274.html) operation.
	//
	// > You can specify only one `BizId`.
	//
	// example:
	//
	// 134523^435****
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The current page number for paginated results.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	OwnerId     *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of delivery records to return on each page.
	//
	// Valid values: 1 to 50.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The phone number to query. The format is as follows:
	//
	// - For messages to the Chinese mainland, use an 11-digit phone number, such as 1390000\\*\\*\\*\\*.
	//
	// - For international SMS, use the format: country/region code + phone number, such as 8520000\\*\\*\\*\\*.
	//
	// > You can specify only one phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The date when the SMS message was sent. You can query records from the past 30 days.
	//
	// Format: **yyyyMMdd**, for example, 20250601.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250601
	SendDate *string `json:"SendDate,omitempty" xml:"SendDate,omitempty"`
}

func (s QuerySendDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySendDetailsRequest) GoString() string {
	return s.String()
}

func (s *QuerySendDetailsRequest) GetBizId() *string {
	return s.BizId
}

func (s *QuerySendDetailsRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *QuerySendDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySendDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QuerySendDetailsRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QuerySendDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySendDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySendDetailsRequest) GetSendDate() *string {
	return s.SendDate
}

func (s *QuerySendDetailsRequest) SetBizId(v string) *QuerySendDetailsRequest {
	s.BizId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetCurrentPage(v int64) *QuerySendDetailsRequest {
	s.CurrentPage = &v
	return s
}

func (s *QuerySendDetailsRequest) SetOwnerId(v int64) *QuerySendDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPageSize(v int64) *QuerySendDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySendDetailsRequest) SetPhoneNumber(v string) *QuerySendDetailsRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerAccount(v string) *QuerySendDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySendDetailsRequest) SetResourceOwnerId(v int64) *QuerySendDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySendDetailsRequest) SetSendDate(v string) *QuerySendDetailsRequest {
	s.SendDate = &v
	return s
}

func (s *QuerySendDetailsRequest) Validate() error {
	return dara.Validate(s)
}
