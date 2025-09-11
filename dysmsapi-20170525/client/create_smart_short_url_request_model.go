// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmartShortUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutId(v string) *CreateSmartShortUrlRequest
	GetOutId() *string
	SetOwnerId(v int64) *CreateSmartShortUrlRequest
	GetOwnerId() *int64
	SetPhoneNumbers(v string) *CreateSmartShortUrlRequest
	GetPhoneNumbers() *string
	SetResourceOwnerAccount(v string) *CreateSmartShortUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmartShortUrlRequest
	GetResourceOwnerId() *int64
	SetSourceUrl(v string) *CreateSmartShortUrlRequest
	GetSourceUrl() *string
}

type CreateSmartShortUrlRequest struct {
	// example:
	//
	// 示例值示例值
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 15900195***
	PhoneNumbers         *string `json:"PhoneNumbers,omitempty" xml:"PhoneNumbers,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s CreateSmartShortUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmartShortUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateSmartShortUrlRequest) GetOutId() *string {
	return s.OutId
}

func (s *CreateSmartShortUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmartShortUrlRequest) GetPhoneNumbers() *string {
	return s.PhoneNumbers
}

func (s *CreateSmartShortUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmartShortUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmartShortUrlRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *CreateSmartShortUrlRequest) SetOutId(v string) *CreateSmartShortUrlRequest {
	s.OutId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetOwnerId(v int64) *CreateSmartShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetPhoneNumbers(v string) *CreateSmartShortUrlRequest {
	s.PhoneNumbers = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetResourceOwnerAccount(v string) *CreateSmartShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetResourceOwnerId(v int64) *CreateSmartShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmartShortUrlRequest) SetSourceUrl(v string) *CreateSmartShortUrlRequest {
	s.SourceUrl = &v
	return s
}

func (s *CreateSmartShortUrlRequest) Validate() error {
	return dara.Validate(s)
}
