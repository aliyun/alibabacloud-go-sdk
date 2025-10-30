// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryAxnBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *QueryAxnBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *QueryAxnBindFixedLineRequest
	GetOwnerId() *int64
	SetPhone(v string) *QueryAxnBindFixedLineRequest
	GetPhone() *string
	SetQueryType(v string) *QueryAxnBindFixedLineRequest
	GetQueryType() *string
	SetResourceOwnerAccount(v string) *QueryAxnBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAxnBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *QueryAxnBindFixedLineRequest
	GetSubId() *string
	SetTelX(v string) *QueryAxnBindFixedLineRequest
	GetTelX() *string
	SetTs(v string) *QueryAxnBindFixedLineRequest
	GetTs() *string
}

type QueryAxnBindFixedLineRequest struct {
	// 号池id。
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
	// 示例值示例值
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// A号码，queryType=1时，必传
	//
	// example:
	//
	// 15500001111
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// 查询类型 0：根据绑定id查询 1：根据A号码查询 2：根据X查询
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 绑定id queryType=0时，必传
	//
	// example:
	//
	// GHX0534X202504221531579290029-2-1-aliaxn
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// 隐私号码
	//
	// example:
	//
	// 05718950
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// 格式yyyyMMddHHmmssSSS，时间采用北京时间，24小时制，精确至毫秒
	//
	// This parameter is required.
	//
	// example:
	//
	// 20250421141723098
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxnBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxnBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryAxnBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryAxnBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAxnBindFixedLineRequest) GetPhone() *string {
	return s.Phone
}

func (s *QueryAxnBindFixedLineRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryAxnBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAxnBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAxnBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *QueryAxnBindFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *QueryAxnBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *QueryAxnBindFixedLineRequest) SetAppId(v string) *QueryAxnBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetOrderId(v string) *QueryAxnBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetOwnerId(v int64) *QueryAxnBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetPhone(v string) *QueryAxnBindFixedLineRequest {
	s.Phone = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetQueryType(v string) *QueryAxnBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxnBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxnBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetSubId(v string) *QueryAxnBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetTelX(v string) *QueryAxnBindFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) SetTs(v string) *QueryAxnBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *QueryAxnBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
