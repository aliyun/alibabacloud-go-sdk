// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDVerificationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *UAIDVerificationRequest
	GetAuthCode() *string
	SetCarrier(v string) *UAIDVerificationRequest
	GetCarrier() *string
	SetIp(v string) *UAIDVerificationRequest
	GetIp() *string
	SetOutId(v string) *UAIDVerificationRequest
	GetOutId() *string
	SetOwnerId(v int64) *UAIDVerificationRequest
	GetOwnerId() *int64
	SetProvince(v string) *UAIDVerificationRequest
	GetProvince() *string
	SetResourceOwnerAccount(v string) *UAIDVerificationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UAIDVerificationRequest
	GetResourceOwnerId() *int64
	SetToken(v string) *UAIDVerificationRequest
	GetToken() *string
	SetUserGrantId(v string) *UAIDVerificationRequest
	GetUserGrantId() *string
}

type UAIDVerificationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// HwD97InG
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// CM
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// b8b5b3a*******0b9893484fdf412c99
	OutId   *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	Province             *string `json:"Province,omitempty" xml:"Province,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// MIGfMA0********3DQEBAQUAA4GNADCB
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// Md****a3Em
	UserGrantId *string `json:"UserGrantId,omitempty" xml:"UserGrantId,omitempty"`
}

func (s UAIDVerificationRequest) String() string {
	return dara.Prettify(s)
}

func (s UAIDVerificationRequest) GoString() string {
	return s.String()
}

func (s *UAIDVerificationRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *UAIDVerificationRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *UAIDVerificationRequest) GetIp() *string {
	return s.Ip
}

func (s *UAIDVerificationRequest) GetOutId() *string {
	return s.OutId
}

func (s *UAIDVerificationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UAIDVerificationRequest) GetProvince() *string {
	return s.Province
}

func (s *UAIDVerificationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UAIDVerificationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UAIDVerificationRequest) GetToken() *string {
	return s.Token
}

func (s *UAIDVerificationRequest) GetUserGrantId() *string {
	return s.UserGrantId
}

func (s *UAIDVerificationRequest) SetAuthCode(v string) *UAIDVerificationRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDVerificationRequest) SetCarrier(v string) *UAIDVerificationRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDVerificationRequest) SetIp(v string) *UAIDVerificationRequest {
	s.Ip = &v
	return s
}

func (s *UAIDVerificationRequest) SetOutId(v string) *UAIDVerificationRequest {
	s.OutId = &v
	return s
}

func (s *UAIDVerificationRequest) SetOwnerId(v int64) *UAIDVerificationRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDVerificationRequest) SetProvince(v string) *UAIDVerificationRequest {
	s.Province = &v
	return s
}

func (s *UAIDVerificationRequest) SetResourceOwnerAccount(v string) *UAIDVerificationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDVerificationRequest) SetResourceOwnerId(v int64) *UAIDVerificationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDVerificationRequest) SetToken(v string) *UAIDVerificationRequest {
	s.Token = &v
	return s
}

func (s *UAIDVerificationRequest) SetUserGrantId(v string) *UAIDVerificationRequest {
	s.UserGrantId = &v
	return s
}

func (s *UAIDVerificationRequest) Validate() error {
	return dara.Validate(s)
}
