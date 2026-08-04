// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyContacterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContacterAddress(v string) *ModifyContacterRequest
	GetContacterAddress() *string
	SetContacterDingding(v string) *ModifyContacterRequest
	GetContacterDingding() *string
	SetContacterEmail(v string) *ModifyContacterRequest
	GetContacterEmail() *string
	SetContacterId(v int64) *ModifyContacterRequest
	GetContacterId() *int64
	SetContacterMobile(v string) *ModifyContacterRequest
	GetContacterMobile() *string
	SetContacterName(v string) *ModifyContacterRequest
	GetContacterName() *string
	SetContacterPosition(v string) *ModifyContacterRequest
	GetContacterPosition() *string
	SetContacterStaffNo(v string) *ModifyContacterRequest
	GetContacterStaffNo() *string
	SetContacterType(v string) *ModifyContacterRequest
	GetContacterType() *string
	SetContacterWangwang(v string) *ModifyContacterRequest
	GetContacterWangwang() *string
	SetEmailConfirmed(v bool) *ModifyContacterRequest
	GetEmailConfirmed() *bool
	SetMobileConfirmed(v bool) *ModifyContacterRequest
	GetMobileConfirmed() *bool
	SetUserId(v int64) *ModifyContacterRequest
	GetUserId() *int64
}

type ModifyContacterRequest struct {
	ContacterAddress  *string `json:"ContacterAddress,omitempty" xml:"ContacterAddress,omitempty"`
	ContacterDingding *string `json:"ContacterDingding,omitempty" xml:"ContacterDingding,omitempty"`
	ContacterEmail    *string `json:"ContacterEmail,omitempty" xml:"ContacterEmail,omitempty"`
	// This parameter is required.
	ContacterId     *int64  `json:"ContacterId,omitempty" xml:"ContacterId,omitempty"`
	ContacterMobile *string `json:"ContacterMobile,omitempty" xml:"ContacterMobile,omitempty"`
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

func (s ModifyContacterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyContacterRequest) GoString() string {
	return s.String()
}

func (s *ModifyContacterRequest) GetContacterAddress() *string {
	return s.ContacterAddress
}

func (s *ModifyContacterRequest) GetContacterDingding() *string {
	return s.ContacterDingding
}

func (s *ModifyContacterRequest) GetContacterEmail() *string {
	return s.ContacterEmail
}

func (s *ModifyContacterRequest) GetContacterId() *int64 {
	return s.ContacterId
}

func (s *ModifyContacterRequest) GetContacterMobile() *string {
	return s.ContacterMobile
}

func (s *ModifyContacterRequest) GetContacterName() *string {
	return s.ContacterName
}

func (s *ModifyContacterRequest) GetContacterPosition() *string {
	return s.ContacterPosition
}

func (s *ModifyContacterRequest) GetContacterStaffNo() *string {
	return s.ContacterStaffNo
}

func (s *ModifyContacterRequest) GetContacterType() *string {
	return s.ContacterType
}

func (s *ModifyContacterRequest) GetContacterWangwang() *string {
	return s.ContacterWangwang
}

func (s *ModifyContacterRequest) GetEmailConfirmed() *bool {
	return s.EmailConfirmed
}

func (s *ModifyContacterRequest) GetMobileConfirmed() *bool {
	return s.MobileConfirmed
}

func (s *ModifyContacterRequest) GetUserId() *int64 {
	return s.UserId
}

func (s *ModifyContacterRequest) SetContacterAddress(v string) *ModifyContacterRequest {
	s.ContacterAddress = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterDingding(v string) *ModifyContacterRequest {
	s.ContacterDingding = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterEmail(v string) *ModifyContacterRequest {
	s.ContacterEmail = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterId(v int64) *ModifyContacterRequest {
	s.ContacterId = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterMobile(v string) *ModifyContacterRequest {
	s.ContacterMobile = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterName(v string) *ModifyContacterRequest {
	s.ContacterName = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterPosition(v string) *ModifyContacterRequest {
	s.ContacterPosition = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterStaffNo(v string) *ModifyContacterRequest {
	s.ContacterStaffNo = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterType(v string) *ModifyContacterRequest {
	s.ContacterType = &v
	return s
}

func (s *ModifyContacterRequest) SetContacterWangwang(v string) *ModifyContacterRequest {
	s.ContacterWangwang = &v
	return s
}

func (s *ModifyContacterRequest) SetEmailConfirmed(v bool) *ModifyContacterRequest {
	s.EmailConfirmed = &v
	return s
}

func (s *ModifyContacterRequest) SetMobileConfirmed(v bool) *ModifyContacterRequest {
	s.MobileConfirmed = &v
	return s
}

func (s *ModifyContacterRequest) SetUserId(v int64) *ModifyContacterRequest {
	s.UserId = &v
	return s
}

func (s *ModifyContacterRequest) Validate() error {
	return dara.Validate(s)
}
