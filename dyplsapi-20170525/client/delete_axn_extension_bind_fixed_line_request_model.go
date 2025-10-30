// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnExtensionBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAxnExtensionBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *DeleteAxnExtensionBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAxnExtensionBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *DeleteAxnExtensionBindFixedLineRequest
	GetSubId() *string
	SetTs(v string) *DeleteAxnExtensionBindFixedLineRequest
	GetTs() *string
}

type DeleteAxnExtensionBindFixedLineRequest struct {
	// 号池ID
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId              *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id
	//
	// This parameter is required.
	//
	// example:
	//
	// GHX0534X202504221531579290029
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *DeleteAxnExtensionBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetAppId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetOrderId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetSubId(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) SetTs(v string) *DeleteAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
