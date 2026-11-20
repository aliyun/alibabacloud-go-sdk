// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxbBindFixedLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryAxbBindFixedLineRequest
	GetAppId() *string
	SetOrderId(v string) *QueryAxbBindFixedLineRequest
	GetOrderId() *string
	SetOwnerId(v int64) *QueryAxbBindFixedLineRequest
	GetOwnerId() *int64
	SetPhone(v string) *QueryAxbBindFixedLineRequest
	GetPhone() *string
	SetQueryType(v string) *QueryAxbBindFixedLineRequest
	GetQueryType() *string
	SetResourceOwnerAccount(v string) *QueryAxbBindFixedLineRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QueryAxbBindFixedLineRequest
	GetResourceOwnerId() *int64
	SetSubId(v string) *QueryAxbBindFixedLineRequest
	GetSubId() *string
	SetTelX(v string) *QueryAxbBindFixedLineRequest
	GetTelX() *string
	SetTs(v string) *QueryAxbBindFixedLineRequest
	GetTs() *string
}

type QueryAxbBindFixedLineRequest struct {
	// The number pool ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC10000016848****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The unique identifier of the message request.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123*******
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The A/B number. Required when QueryType is set to 1.
	//
	// example:
	//
	// 138****6534
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The query type. Valid values:
	//
	// - 0: Query by binding ID.
	//
	// - 1: Query by X and A/B numbers.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding ID. Required when QueryType is set to 0.
	//
	// example:
	//
	// A203**************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The X number. Required when QueryType is set to 1.
	//
	// example:
	//
	// 0571******23
	TelX *string `json:"TelX,omitempty" xml:"TelX,omitempty"`
	// The business timestamp. Format: yyyyMMddHHmmssSSS. The time is in UTC+8, in 24-hour format, accurate to milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20161114143116123
	Ts *string `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s QueryAxbBindFixedLineRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAxbBindFixedLineRequest) GoString() string {
	return s.String()
}

func (s *QueryAxbBindFixedLineRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryAxbBindFixedLineRequest) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryAxbBindFixedLineRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QueryAxbBindFixedLineRequest) GetPhone() *string {
	return s.Phone
}

func (s *QueryAxbBindFixedLineRequest) GetQueryType() *string {
	return s.QueryType
}

func (s *QueryAxbBindFixedLineRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QueryAxbBindFixedLineRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QueryAxbBindFixedLineRequest) GetSubId() *string {
	return s.SubId
}

func (s *QueryAxbBindFixedLineRequest) GetTelX() *string {
	return s.TelX
}

func (s *QueryAxbBindFixedLineRequest) GetTs() *string {
	return s.Ts
}

func (s *QueryAxbBindFixedLineRequest) SetAppId(v string) *QueryAxbBindFixedLineRequest {
	s.AppId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetOrderId(v string) *QueryAxbBindFixedLineRequest {
	s.OrderId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetOwnerId(v int64) *QueryAxbBindFixedLineRequest {
	s.OwnerId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetPhone(v string) *QueryAxbBindFixedLineRequest {
	s.Phone = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetQueryType(v string) *QueryAxbBindFixedLineRequest {
	s.QueryType = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetResourceOwnerAccount(v string) *QueryAxbBindFixedLineRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetResourceOwnerId(v int64) *QueryAxbBindFixedLineRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetSubId(v string) *QueryAxbBindFixedLineRequest {
	s.SubId = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetTelX(v string) *QueryAxbBindFixedLineRequest {
	s.TelX = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) SetTs(v string) *QueryAxbBindFixedLineRequest {
	s.Ts = &v
	return s
}

func (s *QueryAxbBindFixedLineRequest) Validate() error {
	return dara.Validate(s)
}
