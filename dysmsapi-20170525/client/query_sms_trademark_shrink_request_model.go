// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySmsTrademarkShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *QuerySmsTrademarkShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *QuerySmsTrademarkShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *QuerySmsTrademarkShrinkRequest
	GetResourceOwnerId() *int64
	SetTrademarkIdListShrink(v string) *QuerySmsTrademarkShrinkRequest
	GetTrademarkIdListShrink() *string
}

type QuerySmsTrademarkShrinkRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 商标实体id列表
	//
	// This parameter is required.
	TrademarkIdListShrink *string `json:"TrademarkIdList,omitempty" xml:"TrademarkIdList,omitempty"`
}

func (s QuerySmsTrademarkShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySmsTrademarkShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySmsTrademarkShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *QuerySmsTrademarkShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *QuerySmsTrademarkShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *QuerySmsTrademarkShrinkRequest) GetTrademarkIdListShrink() *string {
	return s.TrademarkIdListShrink
}

func (s *QuerySmsTrademarkShrinkRequest) SetOwnerId(v int64) *QuerySmsTrademarkShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *QuerySmsTrademarkShrinkRequest) SetResourceOwnerAccount(v string) *QuerySmsTrademarkShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *QuerySmsTrademarkShrinkRequest) SetResourceOwnerId(v int64) *QuerySmsTrademarkShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *QuerySmsTrademarkShrinkRequest) SetTrademarkIdListShrink(v string) *QuerySmsTrademarkShrinkRequest {
	s.TrademarkIdListShrink = &v
	return s
}

func (s *QuerySmsTrademarkShrinkRequest) Validate() error {
	return dara.Validate(s)
}
