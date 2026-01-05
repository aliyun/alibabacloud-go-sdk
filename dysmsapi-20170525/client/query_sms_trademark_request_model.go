// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTrademarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsTrademarkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsTrademarkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsTrademarkRequest
	GetResourceOwnerId() *int64
	SetTrademarkIdList(v []*int64) *QuerySmsTrademarkRequest
	GetTrademarkIdList() []*int64
}

type QuerySmsTrademarkRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 商标实体id列表
	//
	// This parameter is required.
	TrademarkIdList []*int64 `json:"TrademarkIdList,omitempty" xml:"TrademarkIdList,omitempty" type:"Repeated"`
}

func (s QuerySmsTrademarkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTrademarkRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTrademarkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsTrademarkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsTrademarkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsTrademarkRequest) GetTrademarkIdList() []*int64 {
	return s.TrademarkIdList
}

func (s *QuerySmsTrademarkRequest) SetOwnerId(v int64) *QuerySmsTrademarkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTrademarkRequest) SetResourceOwnerAccount(v string) *QuerySmsTrademarkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTrademarkRequest) SetResourceOwnerId(v int64) *QuerySmsTrademarkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsTrademarkRequest) SetTrademarkIdList(v []*int64) *QuerySmsTrademarkRequest {
	s.TrademarkIdList = v
	return s
}

func (s *QuerySmsTrademarkRequest) Validate() error {
	return dara.Validate(s)
}
