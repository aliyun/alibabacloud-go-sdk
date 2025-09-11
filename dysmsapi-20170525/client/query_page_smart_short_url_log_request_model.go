// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPageSmartShortUrlLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateDateEnd(v int64) *QueryPageSmartShortUrlLogRequest
	GetCreateDateEnd() *int64
	SetCreateDateStart(v int64) *QueryPageSmartShortUrlLogRequest
	GetCreateDateStart() *int64
	SetOwnerId(v int64) *QueryPageSmartShortUrlLogRequest
	GetOwnerId() *int64
	SetPageNo(v int64) *QueryPageSmartShortUrlLogRequest
	GetPageNo() *int64
	SetPageSize(v int64) *QueryPageSmartShortUrlLogRequest
	GetPageSize() *int64
	SetPhoneNumber(v string) *QueryPageSmartShortUrlLogRequest
	GetPhoneNumber() *string
	SetResourceOwnerAccount(v string) *QueryPageSmartShortUrlLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryPageSmartShortUrlLogRequest
	GetResourceOwnerId() *int64
	SetShortUrl(v string) *QueryPageSmartShortUrlLogRequest
	GetShortUrl() *string
}

type QueryPageSmartShortUrlLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 20181225
	CreateDateEnd *int64 `json:"CreateDateEnd,omitempty" xml:"CreateDateEnd,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20181225
	CreateDateStart *int64 `json:"CreateDateStart,omitempty" xml:"CreateDateStart,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1390000****
	PhoneNumber          *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// http://ays.cn/****
	ShortUrl *string `json:"ShortUrl,omitempty" xml:"ShortUrl,omitempty"`
}

func (s QueryPageSmartShortUrlLogRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPageSmartShortUrlLogRequest) GoString() string {
	return s.String()
}

func (s *QueryPageSmartShortUrlLogRequest) GetCreateDateEnd() *int64 {
	return s.CreateDateEnd
}

func (s *QueryPageSmartShortUrlLogRequest) GetCreateDateStart() *int64 {
	return s.CreateDateStart
}

func (s *QueryPageSmartShortUrlLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryPageSmartShortUrlLogRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryPageSmartShortUrlLogRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPageSmartShortUrlLogRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *QueryPageSmartShortUrlLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryPageSmartShortUrlLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryPageSmartShortUrlLogRequest) GetShortUrl() *string {
	return s.ShortUrl
}

func (s *QueryPageSmartShortUrlLogRequest) SetCreateDateEnd(v int64) *QueryPageSmartShortUrlLogRequest {
	s.CreateDateEnd = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetCreateDateStart(v int64) *QueryPageSmartShortUrlLogRequest {
	s.CreateDateStart = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetOwnerId(v int64) *QueryPageSmartShortUrlLogRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPageNo(v int64) *QueryPageSmartShortUrlLogRequest {
	s.PageNo = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPageSize(v int64) *QueryPageSmartShortUrlLogRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetPhoneNumber(v string) *QueryPageSmartShortUrlLogRequest {
	s.PhoneNumber = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetResourceOwnerAccount(v string) *QueryPageSmartShortUrlLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetResourceOwnerId(v int64) *QueryPageSmartShortUrlLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) SetShortUrl(v string) *QueryPageSmartShortUrlLogRequest {
	s.ShortUrl = &v
	return s
}

func (s *QueryPageSmartShortUrlLogRequest) Validate() error {
	return dara.Validate(s)
}
