// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProxyBrandUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandUserNick(v string) *CreateProxyBrandUserRequest
	GetBrandUserNick() *string
	SetChannelId(v string) *CreateProxyBrandUserRequest
	GetChannelId() *string
	SetClientToken(v string) *CreateProxyBrandUserRequest
	GetClientToken() *string
	SetCompany(v string) *CreateProxyBrandUserRequest
	GetCompany() *string
	SetContactName(v string) *CreateProxyBrandUserRequest
	GetContactName() *string
	SetContactPhone(v string) *CreateProxyBrandUserRequest
	GetContactPhone() *string
	SetIndustry(v string) *CreateProxyBrandUserRequest
	GetIndustry() *string
	SetProxyUserId(v int64) *CreateProxyBrandUserRequest
	GetProxyUserId() *int64
}

type CreateProxyBrandUserRequest struct {
	// This parameter is required.
	BrandUserNick *string `json:"BrandUserNick,omitempty" xml:"BrandUserNick,omitempty"`
	// This parameter is required.
	ChannelId   *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	Company      *string `json:"Company,omitempty" xml:"Company,omitempty"`
	ContactName  *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	ContactPhone *string `json:"ContactPhone,omitempty" xml:"ContactPhone,omitempty"`
	// This parameter is required.
	Industry *string `json:"Industry,omitempty" xml:"Industry,omitempty"`
	// This parameter is required.
	ProxyUserId *int64 `json:"ProxyUserId,omitempty" xml:"ProxyUserId,omitempty"`
}

func (s CreateProxyBrandUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProxyBrandUserRequest) GoString() string {
	return s.String()
}

func (s *CreateProxyBrandUserRequest) GetBrandUserNick() *string {
	return s.BrandUserNick
}

func (s *CreateProxyBrandUserRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *CreateProxyBrandUserRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProxyBrandUserRequest) GetCompany() *string {
	return s.Company
}

func (s *CreateProxyBrandUserRequest) GetContactName() *string {
	return s.ContactName
}

func (s *CreateProxyBrandUserRequest) GetContactPhone() *string {
	return s.ContactPhone
}

func (s *CreateProxyBrandUserRequest) GetIndustry() *string {
	return s.Industry
}

func (s *CreateProxyBrandUserRequest) GetProxyUserId() *int64 {
	return s.ProxyUserId
}

func (s *CreateProxyBrandUserRequest) SetBrandUserNick(v string) *CreateProxyBrandUserRequest {
	s.BrandUserNick = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetChannelId(v string) *CreateProxyBrandUserRequest {
	s.ChannelId = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetClientToken(v string) *CreateProxyBrandUserRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetCompany(v string) *CreateProxyBrandUserRequest {
	s.Company = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetContactName(v string) *CreateProxyBrandUserRequest {
	s.ContactName = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetContactPhone(v string) *CreateProxyBrandUserRequest {
	s.ContactPhone = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetIndustry(v string) *CreateProxyBrandUserRequest {
	s.Industry = &v
	return s
}

func (s *CreateProxyBrandUserRequest) SetProxyUserId(v int64) *CreateProxyBrandUserRequest {
	s.ProxyUserId = &v
	return s
}

func (s *CreateProxyBrandUserRequest) Validate() error {
	return dara.Validate(s)
}
