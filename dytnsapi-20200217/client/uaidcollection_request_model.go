// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUAIDCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *UAIDCollectionRequest
	GetAuthCode() *string
	SetCarrier(v string) *UAIDCollectionRequest
	GetCarrier() *string
	SetIp(v string) *UAIDCollectionRequest
	GetIp() *string
	SetOutId(v string) *UAIDCollectionRequest
	GetOutId() *string
	SetOwnerId(v int64) *UAIDCollectionRequest
	GetOwnerId() *int64
	SetProvince(v string) *UAIDCollectionRequest
	GetProvince() *string
	SetResourceOwnerAccount(v string) *UAIDCollectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UAIDCollectionRequest
	GetResourceOwnerId() *int64
	SetToken(v string) *UAIDCollectionRequest
	GetToken() *string
	SetUserGrantId(v string) *UAIDCollectionRequest
	GetUserGrantId() *string
}

type UAIDCollectionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// 示例值
	Carrier *string `json:"Carrier,omitempty" xml:"Carrier,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
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
	// 示例值示例值
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 示例值示例值
	UserGrantId *string `json:"UserGrantId,omitempty" xml:"UserGrantId,omitempty"`
}

func (s UAIDCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s UAIDCollectionRequest) GoString() string {
	return s.String()
}

func (s *UAIDCollectionRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *UAIDCollectionRequest) GetCarrier() *string {
	return s.Carrier
}

func (s *UAIDCollectionRequest) GetIp() *string {
	return s.Ip
}

func (s *UAIDCollectionRequest) GetOutId() *string {
	return s.OutId
}

func (s *UAIDCollectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UAIDCollectionRequest) GetProvince() *string {
	return s.Province
}

func (s *UAIDCollectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UAIDCollectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UAIDCollectionRequest) GetToken() *string {
	return s.Token
}

func (s *UAIDCollectionRequest) GetUserGrantId() *string {
	return s.UserGrantId
}

func (s *UAIDCollectionRequest) SetAuthCode(v string) *UAIDCollectionRequest {
	s.AuthCode = &v
	return s
}

func (s *UAIDCollectionRequest) SetCarrier(v string) *UAIDCollectionRequest {
	s.Carrier = &v
	return s
}

func (s *UAIDCollectionRequest) SetIp(v string) *UAIDCollectionRequest {
	s.Ip = &v
	return s
}

func (s *UAIDCollectionRequest) SetOutId(v string) *UAIDCollectionRequest {
	s.OutId = &v
	return s
}

func (s *UAIDCollectionRequest) SetOwnerId(v int64) *UAIDCollectionRequest {
	s.OwnerId = &v
	return s
}

func (s *UAIDCollectionRequest) SetProvince(v string) *UAIDCollectionRequest {
	s.Province = &v
	return s
}

func (s *UAIDCollectionRequest) SetResourceOwnerAccount(v string) *UAIDCollectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UAIDCollectionRequest) SetResourceOwnerId(v int64) *UAIDCollectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UAIDCollectionRequest) SetToken(v string) *UAIDCollectionRequest {
	s.Token = &v
	return s
}

func (s *UAIDCollectionRequest) SetUserGrantId(v string) *UAIDCollectionRequest {
	s.UserGrantId = &v
	return s
}

func (s *UAIDCollectionRequest) Validate() error {
	return dara.Validate(s)
}
