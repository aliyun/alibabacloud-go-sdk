// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryShortUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QueryShortUrlRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QueryShortUrlRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryShortUrlRequest
	GetResourceOwnerId() *int64
	SetShortUrl(v string) *QueryShortUrlRequest
	GetShortUrl() *string
}

type QueryShortUrlRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The short URL. You can query the short URL by calling the [AddShortUrl](https://help.aliyun.com/document_detail/186774.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://****.cn/6y8uy7
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryShortUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryShortUrlRequest) GoString() string {
	return s.String()
}

func (s *QueryShortUrlRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryShortUrlRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryShortUrlRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryShortUrlRequest) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *QueryShortUrlRequest) SetOwnerId(v int64) *QueryShortUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryShortUrlRequest) SetResourceOwnerAccount(v string) *QueryShortUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryShortUrlRequest) SetResourceOwnerId(v int64) *QueryShortUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryShortUrlRequest) SetShortUrl(v string) *QueryShortUrlRequest {
	s.ShortUrl = &v
	return s
}

func (s *QueryShortUrlRequest) Validate() error {
	return dara.Validate(s)
}
