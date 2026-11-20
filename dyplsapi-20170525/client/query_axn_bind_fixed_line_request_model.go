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
	// The A/B phone number. This parameter is required when QueryType is set to 1.
	//
	// example:
	//
	// 138****6534
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The query type. Valid values:
	//
	// - 0: query by binding ID.
	//
	// - 1: query by X and A/B phone number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	QueryType            *string `json:"QueryType,omitempty" xml:"QueryType,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The binding ID. This parameter is required when QueryType is set to 1.
	//
	// example:
	//
	// A203**************
	SubId *string `json:"SubId,omitempty" xml:"SubId,omitempty"`
	// The X phone number. This parameter is required when QueryType is set to 1.
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
