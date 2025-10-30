// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnExtensionBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryAxnExtensionBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *QueryAxnExtensionBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest
	GetOwnerId() *int64
	SetQueryType(v string) *QueryAxnExtensionBindFixedLineRequest
	GetQueryType() *string
	SetResourceOwnerAccount(v string) *QueryAxnExtensionBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *QueryAxnExtensionBindFixedLineRequest
	GetSubId() *string
	SetTelA(v string) *QueryAxnExtensionBindFixedLineRequest
	GetTelA() *string
	SetTs(v string) *QueryAxnExtensionBindFixedLineRequest
	GetTs() *string
}

type QueryAxnExtensionBindFixedLineRequest struct {
	// 号池ID
	//
	// This parameter is required.
	//
	// example:
	//
	// alitest
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 消息请求唯一标识。
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345dkwkd99d
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// 查询类型。取值： 0：根据绑定ID查询。1：根据X和A/B号码查询
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定ID。QueryType参数为0时必传。
	//
	// example:
	//
	// A203**************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// A号码。QueryType参数为1时必传。
	//
	// example:
	//
	// 15500001111
	TelA *string `json:"TelA,omitempty" xml:"TelA,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetTelA() *string {
	return s.TelA
}

func (s *QueryAxnExtensionBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetAppId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetOrderId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetQueryType(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxnExtensionBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetSubId(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetTelA(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.TelA = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) SetTs(v string) *QueryAxnExtensionBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
