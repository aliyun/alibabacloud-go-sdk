// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacterAddress(v string) *CreateContacterRequest
	GetContacterAddress() *string
	SetContacterDingding(v string) *CreateContacterRequest
	GetContacterDingding() *string
	SetContacterEmail(v string) *CreateContacterRequest
	GetContacterEmail() *string
	SetContacterMobile(v string) *CreateContacterRequest
	GetContacterMobile() *string
	SetContacterName(v string) *CreateContacterRequest
	GetContacterName() *string
	SetContacterPosition(v string) *CreateContacterRequest
	GetContacterPosition() *string
	SetContacterStaffNo(v string) *CreateContacterRequest
	GetContacterStaffNo() *string
	SetContacterType(v string) *CreateContacterRequest
	GetContacterType() *string
	SetContacterWangwang(v string) *CreateContacterRequest
	GetContacterWangwang() *string
	SetEmailConfirmed(v bool) *CreateContacterRequest
	GetEmailConfirmed() *bool
	SetMobileConfirmed(v bool) *CreateContacterRequest
	GetMobileConfirmed() *bool
	SetUserId(v int64) *CreateContacterRequest
	GetUserId() *int64
}

type CreateContacterRequest struct {
	ContacterAddress  *string `json:"ContacterAddress,omitempty" xml:"ContacterAddress,omitempty"`
	ContacterDingding *string `json:"ContacterDingding,omitempty" xml:"ContacterDingding,omitempty"`
	ContacterEmail    *string `json:"ContacterEmail,omitempty" xml:"ContacterEmail,omitempty"`
	ContacterMobile   *string `json:"ContacterMobile,omitempty" xml:"ContacterMobile,omitempty"`
	// This parameter is required.
	ContacterName     *string `json:"ContacterName,omitempty" xml:"ContacterName,omitempty"`
	ContacterPosition *string `json:"ContacterPosition,omitempty" xml:"ContacterPosition,omitempty"`
	ContacterStaffNo  *string `json:"ContacterStaffNo,omitempty" xml:"ContacterStaffNo,omitempty"`
	ContacterType     *string `json:"ContacterType,omitempty" xml:"ContacterType,omitempty"`
	ContacterWangwang *string `json:"ContacterWangwang,omitempty" xml:"ContacterWangwang,omitempty"`
	EmailConfirmed    *bool   `json:"EmailConfirmed,omitempty" xml:"EmailConfirmed,omitempty"`
	MobileConfirmed   *bool   `json:"MobileConfirmed,omitempty" xml:"MobileConfirmed,omitempty"`
	// This parameter is required.
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CreateContacterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateContacterRequest) GoString() string {
	return s.String()
}

func (s *CreateContacterRequest) GetContacterAddress() *string {
	return s.ContacterAddress
}

func (s *CreateContacterRequest) GetContacterDingding() *string {
	return s.ContacterDingding
}

func (s *CreateContacterRequest) GetContacterEmail() *string {
	return s.ContacterEmail
}

func (s *CreateContacterRequest) GetContacterMobile() *string {
	return s.ContacterMobile
}

func (s *CreateContacterRequest) GetContacterName() *string {
	return s.ContacterName
}

func (s *CreateContacterRequest) GetContacterPosition() *string {
	return s.ContacterPosition
}

func (s *CreateContacterRequest) GetContacterStaffNo() *string {
	return s.ContacterStaffNo
}

func (s *CreateContacterRequest) GetContacterType() *string {
	return s.ContacterType
}

func (s *CreateContacterRequest) GetContacterWangwang() *string {
	return s.ContacterWangwang
}

func (s *CreateContacterRequest) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *CreateContacterRequest) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *CreateContacterRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *CreateContacterRequest) SetContacterAddress(v string) *CreateContacterRequest {
	s.ContacterAddress = &v
	return s
}

func (s *CreateContacterRequest) SetContacterDingding(v string) *CreateContacterRequest {
	s.ContacterDingding = &v
	return s
}

func (s *CreateContacterRequest) SetContacterEmail(v string) *CreateContacterRequest {
	s.ContacterEmail = &v
	return s
}

func (s *CreateContacterRequest) SetContacterMobile(v string) *CreateContacterRequest {
	s.ContacterMobile = &v
	return s
}

func (s *CreateContacterRequest) SetContacterName(v string) *CreateContacterRequest {
	s.ContacterName = &v
	return s
}

func (s *CreateContacterRequest) SetContacterPosition(v string) *CreateContacterRequest {
	s.ContacterPosition = &v
	return s
}

func (s *CreateContacterRequest) SetContacterStaffNo(v string) *CreateContacterRequest {
	s.ContacterStaffNo = &v
	return s
}

func (s *CreateContacterRequest) SetContacterType(v string) *CreateContacterRequest {
	s.ContacterType = &v
	return s
}

func (s *CreateContacterRequest) SetContacterWangwang(v string) *CreateContacterRequest {
	s.ContacterWangwang = &v
	return s
}

func (s *CreateContacterRequest) SetEmailConfirmed(v bool) *CreateContacterRequest {
	s.EmailConfirmed = &v
	return s
}

func (s *CreateContacterRequest) SetMobileConfirmed(v bool) *CreateContacterRequest {
	s.MobileConfirmed = &v
	return s
}

func (s *CreateContacterRequest) SetUserId(v int64) *CreateContacterRequest {
	s.UserId = &v
	return s
}

func (s *CreateContacterRequest) Validate() error {
	return dara.Validate(s)
}
