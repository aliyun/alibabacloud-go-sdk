// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEmailVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginCreateTime(v int64) *ListEmailVerificationRequest
	GetBeginCreateTime() *int64
	SetEmail(v string) *ListEmailVerificationRequest
	GetEmail() *string
	SetEndCreateTime(v int64) *ListEmailVerificationRequest
	GetEndCreateTime() *int64
	SetLang(v string) *ListEmailVerificationRequest
	GetLang() *string
	SetPageNum(v int32) *ListEmailVerificationRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListEmailVerificationRequest
	GetPageSize() *int32
	SetUserClientIp(v string) *ListEmailVerificationRequest
	GetUserClientIp() *string
	SetVerificationStatus(v int32) *ListEmailVerificationRequest
	GetVerificationStatus() *int32
}

type ListEmailVerificationRequest struct {
	// The start time for querying email verification creation, represented as the number of milliseconds since 00:00 on January 1, 1970, UTC.
	//
	// example:
	//
	// 1522080000000
	BeginCreateTime *int64 `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	// The email address to query. You can upload only one email address at a time.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The end time for querying the creation of email verification, calculated as the number of milliseconds since 00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1522080000000
	EndCreateTime *int64 `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	// Language of error messages returned by the API. Valid values:
	//
	// - **zh**: Chinese.
	//
	// - **en**: English.
	//
	// Default value is **en**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number for paging through the domain list. Default value is **1**. You can set this parameter based on your needs.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The page size for paging through the domain list. Default value is **500**, and the maximum value is **5000**. You can set this parameter based on your needs.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// User IP address. You can set it to **127.0.0.1**.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// Email verification status. Valid values:
	//
	// - **0**: Waiting for verification.
	//
	// - **1**: Verification succeeded.
	//
	// example:
	//
	// 1
	VerificationStatus *int32 `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
}

func (s ListEmailVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEmailVerificationRequest) GoString() string {
	return s.String()
}

func (s *ListEmailVerificationRequest) GetBeginCreateTime() *int64 {
	return s.BeginCreateTime
}

func (s *ListEmailVerificationRequest) GetEmail() *string {
	return s.Email
}

func (s *ListEmailVerificationRequest) GetEndCreateTime() *int64 {
	return s.EndCreateTime
}

func (s *ListEmailVerificationRequest) GetLang() *string {
	return s.Lang
}

func (s *ListEmailVerificationRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListEmailVerificationRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEmailVerificationRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *ListEmailVerificationRequest) GetVerificationStatus() *int32 {
	return s.VerificationStatus
}

func (s *ListEmailVerificationRequest) SetBeginCreateTime(v int64) *ListEmailVerificationRequest {
	s.BeginCreateTime = &v
	return s
}

func (s *ListEmailVerificationRequest) SetEmail(v string) *ListEmailVerificationRequest {
	s.Email = &v
	return s
}

func (s *ListEmailVerificationRequest) SetEndCreateTime(v int64) *ListEmailVerificationRequest {
	s.EndCreateTime = &v
	return s
}

func (s *ListEmailVerificationRequest) SetLang(v string) *ListEmailVerificationRequest {
	s.Lang = &v
	return s
}

func (s *ListEmailVerificationRequest) SetPageNum(v int32) *ListEmailVerificationRequest {
	s.PageNum = &v
	return s
}

func (s *ListEmailVerificationRequest) SetPageSize(v int32) *ListEmailVerificationRequest {
	s.PageSize = &v
	return s
}

func (s *ListEmailVerificationRequest) SetUserClientIp(v string) *ListEmailVerificationRequest {
	s.UserClientIp = &v
	return s
}

func (s *ListEmailVerificationRequest) SetVerificationStatus(v int32) *ListEmailVerificationRequest {
	s.VerificationStatus = &v
	return s
}

func (s *ListEmailVerificationRequest) Validate() error {
	return dara.Validate(s)
}
