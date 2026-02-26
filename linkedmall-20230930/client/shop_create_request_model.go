// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShopCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterSaleDingTalkId(v string) *ShopCreateRequest
	GetAfterSaleDingTalkId() *string
	SetDescription(v string) *ShopCreateRequest
	GetDescription() *string
	SetOperatorDingTalkId(v string) *ShopCreateRequest
	GetOperatorDingTalkId() *string
	SetPreSaleDingTalkId(v string) *ShopCreateRequest
	GetPreSaleDingTalkId() *string
	SetShopName(v string) *ShopCreateRequest
	GetShopName() *string
}

type ShopCreateRequest struct {
	// example:
	//
	// 12344335
	AfterSaleDingTalkId *string `json:"afterSaleDingTalkId,omitempty" xml:"afterSaleDingTalkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 店铺描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12344335
	OperatorDingTalkId *string `json:"operatorDingTalkId,omitempty" xml:"operatorDingTalkId,omitempty"`
	// example:
	//
	// 12344335
	PreSaleDingTalkId *string `json:"preSaleDingTalkId,omitempty" xml:"preSaleDingTalkId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 店铺名称
	ShopName *string `json:"shopName,omitempty" xml:"shopName,omitempty"`
}

func (s ShopCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ShopCreateRequest) GoString() string {
	return s.String()
}

func (s *ShopCreateRequest) GetAfterSaleDingTalkId() *string {
	return s.AfterSaleDingTalkId
}

func (s *ShopCreateRequest) GetDescription() *string {
	return s.Description
}

func (s *ShopCreateRequest) GetOperatorDingTalkId() *string {
	return s.OperatorDingTalkId
}

func (s *ShopCreateRequest) GetPreSaleDingTalkId() *string {
	return s.PreSaleDingTalkId
}

func (s *ShopCreateRequest) GetShopName() *string {
	return s.ShopName
}

func (s *ShopCreateRequest) SetAfterSaleDingTalkId(v string) *ShopCreateRequest {
	s.AfterSaleDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetDescription(v string) *ShopCreateRequest {
	s.Description = &v
	return s
}

func (s *ShopCreateRequest) SetOperatorDingTalkId(v string) *ShopCreateRequest {
	s.OperatorDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetPreSaleDingTalkId(v string) *ShopCreateRequest {
	s.PreSaleDingTalkId = &v
	return s
}

func (s *ShopCreateRequest) SetShopName(v string) *ShopCreateRequest {
	s.ShopName = &v
	return s
}

func (s *ShopCreateRequest) Validate() error {
	return dara.Validate(s)
}
