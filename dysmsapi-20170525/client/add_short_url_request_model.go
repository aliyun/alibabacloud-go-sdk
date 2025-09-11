// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddShortUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveDays(v string) *AddShortUrlRequest
	GetEffectiveDays() *string
	SetOwnerId(v int64) *AddShortUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AddShortUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddShortUrlRequest
	GetResourceOwnerId() *int64
	SetShortUrlName(v string) *AddShortUrlRequest
	GetShortUrlName() *string
	SetSourceUrl(v string) *AddShortUrlRequest
	GetSourceUrl() *string
}

type AddShortUrlRequest struct {
	// The validity period of the short URL. Unit: days. The maximum validity period is 90 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	EffectiveDays        *string `json:"EffectiveDays,omitempty" xml:"EffectiveDays,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name of the short URL. The name cannot exceed 13 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// The Alibaba Cloud Short Link service.
	ShortUrlName *string `json:"ShortUrlName,omitempty" xml:"ShortUrlName,omitempty"`
	// The source URL. The URL cannot exceed 1,000 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s AddShortUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s AddShortUrlRequest) GoString() string {
	return s.String()
}

func (s *AddShortUrlRequest) GetEffectiveDays() *string {
	return s.EffectiveDays
}

func (s *AddShortUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddShortUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddShortUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddShortUrlRequest) GetShortUrlName() *string {
	return s.ShortUrlName
}

func (s *AddShortUrlRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *AddShortUrlRequest) SetEffectiveDays(v string) *AddShortUrlRequest {
	s.EffectiveDays = &v
	return s
}

func (s *AddShortUrlRequest) SetOwnerId(v int64) *AddShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *AddShortUrlRequest) SetResourceOwnerAccount(v string) *AddShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddShortUrlRequest) SetResourceOwnerId(v int64) *AddShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddShortUrlRequest) SetShortUrlName(v string) *AddShortUrlRequest {
	s.ShortUrlName = &v
	return s
}

func (s *AddShortUrlRequest) SetSourceUrl(v string) *AddShortUrlRequest {
	s.SourceUrl = &v
	return s
}

func (s *AddShortUrlRequest) Validate() error {
	return dara.Validate(s)
}
