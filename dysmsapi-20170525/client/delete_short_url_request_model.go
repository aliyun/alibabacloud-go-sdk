// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteShortUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteShortUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteShortUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteShortUrlRequest
	GetResourceOwnerId() *int64
	SetSourceUrl(v string) *DeleteShortUrlRequest
	GetSourceUrl() *string
}

type DeleteShortUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The source address. The address can be up to 1,000 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.****.com/product/sms
	SourceUrl *string `json:"SourceUrl,omitempty" xml:"SourceUrl,omitempty"`
}

func (s DeleteShortUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteShortUrlRequest) GoString() string {
	return s.String()
}

func (s *DeleteShortUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteShortUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteShortUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteShortUrlRequest) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *DeleteShortUrlRequest) SetOwnerId(v int64) *DeleteShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteShortUrlRequest) SetResourceOwnerAccount(v string) *DeleteShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteShortUrlRequest) SetResourceOwnerId(v int64) *DeleteShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteShortUrlRequest) SetSourceUrl(v string) *DeleteShortUrlRequest {
	s.SourceUrl = &v
	return s
}

func (s *DeleteShortUrlRequest) Validate() error {
	return dara.Validate(s)
}
