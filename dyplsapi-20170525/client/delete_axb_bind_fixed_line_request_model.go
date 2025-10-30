// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxbBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAxbBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *DeleteAxbBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *DeleteAxbBindFixedLineRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteAxbBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteAxbBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *DeleteAxbBindFixedLineRequest
	GetSubId() *string
	SetTs(v string) *DeleteAxbBindFixedLineRequest
	GetTs() *string
}

type DeleteAxbBindFixedLineRequest struct {
	// 号池ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ALPT_1234
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
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
	// A93IOELD93
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 业务时间戳，格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116001
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DeleteAxbBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *DeleteAxbBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAxbBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *DeleteAxbBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteAxbBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteAxbBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteAxbBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *DeleteAxbBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *DeleteAxbBindFixedLineRequest) SetAppId(v string) *DeleteAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetOrderId(v string) *DeleteAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetOwnerId(v int64) *DeleteAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *DeleteAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *DeleteAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetSubId(v string) *DeleteAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) SetTs(v string) *DeleteAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *DeleteAxbBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
