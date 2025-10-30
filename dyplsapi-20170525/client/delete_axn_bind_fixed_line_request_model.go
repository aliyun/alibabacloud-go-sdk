// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAxnBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *DeleteAxnBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *DeleteAxnBindFixedLineRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAxnBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAxnBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *DeleteAxnBindFixedLineRequest
	GetSubId() *string
	SetTs(v string) *DeleteAxnBindFixedLineRequest
	GetTs() *string
}

type DeleteAxnBindFixedLineRequest struct {
	// 号池ID。
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
	// GHX0534X202504221531579290029-2-1-aliaxn
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

func (s DeleteAxnBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxnBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAxnBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteAxnBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAxnBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAxnBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAxnBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *DeleteAxnBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *DeleteAxnBindFixedLineRequest) SetAppId(v string) *DeleteAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetOrderId(v string) *DeleteAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetSubId(v string) *DeleteAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) SetTs(v string) *DeleteAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *DeleteAxnBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
