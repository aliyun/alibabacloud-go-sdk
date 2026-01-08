// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindDmAccountShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCode(v string) *BindDmAccountShrinkRequest
	GetAccountCode() *string
	SetCustSpaceId(v string) *BindDmAccountShrinkRequest
	GetCustSpaceId() *string
	SetExtendAttrShrink(v string) *BindDmAccountShrinkRequest
	GetExtendAttrShrink() *string
	SetOwnerId(v int64) *BindDmAccountShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *BindDmAccountShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *BindDmAccountShrinkRequest
	GetResourceOwnerId() *int64
}

type BindDmAccountShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	AccountCode *string `json:"AccountCode,omitempty" xml:"AccountCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	CustSpaceId *string `json:"CustSpaceId,omitempty" xml:"CustSpaceId,omitempty"`
	// This parameter is required.
	ExtendAttrShrink     *string `json:"ExtendAttr,omitempty" xml:"ExtendAttr,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s BindDmAccountShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BindDmAccountShrinkRequest) GoString() string {
	return s.String()
}

func (s *BindDmAccountShrinkRequest) GetAccountCode() *string {
	return s.AccountCode
}

func (s *BindDmAccountShrinkRequest) GetCustSpaceId() *string {
	return s.CustSpaceId
}

func (s *BindDmAccountShrinkRequest) GetExtendAttrShrink() *string {
	return s.ExtendAttrShrink
}

func (s *BindDmAccountShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindDmAccountShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *BindDmAccountShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *BindDmAccountShrinkRequest) SetAccountCode(v string) *BindDmAccountShrinkRequest {
	s.AccountCode = &v
	return s
}

func (s *BindDmAccountShrinkRequest) SetCustSpaceId(v string) *BindDmAccountShrinkRequest {
	s.CustSpaceId = &v
	return s
}

func (s *BindDmAccountShrinkRequest) SetExtendAttrShrink(v string) *BindDmAccountShrinkRequest {
	s.ExtendAttrShrink = &v
	return s
}

func (s *BindDmAccountShrinkRequest) SetOwnerId(v int64) *BindDmAccountShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *BindDmAccountShrinkRequest) SetResourceOwnerAccount(v string) *BindDmAccountShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *BindDmAccountShrinkRequest) SetResourceOwnerId(v int64) *BindDmAccountShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *BindDmAccountShrinkRequest) Validate() error {
	return dara.Validate(s)
}
