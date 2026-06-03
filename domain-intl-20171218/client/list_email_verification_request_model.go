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
	BeginCreateTime    *int64  `json:"BeginCreateTime,omitempty" xml:"BeginCreateTime,omitempty"`
	Email              *string `json:"Email,omitempty" xml:"Email,omitempty"`
	EndCreateTime      *int64  `json:"EndCreateTime,omitempty" xml:"EndCreateTime,omitempty"`
	Lang               *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	UserClientIp       *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	VerificationStatus *int32  `json:"VerificationStatus,omitempty" xml:"VerificationStatus,omitempty"`
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
