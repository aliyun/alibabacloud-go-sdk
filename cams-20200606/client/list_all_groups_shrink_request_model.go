// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllGroupsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *ListAllGroupsShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *ListAllGroupsShrinkRequest
	GetBizExtendShrink() *string
	SetOwnerId(v int64) *ListAllGroupsShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ListAllGroupsShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListAllGroupsShrinkRequest
	GetResourceOwnerId() *int64
}

type ListAllGroupsShrinkRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink      *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListAllGroupsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAllGroupsShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *ListAllGroupsShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *ListAllGroupsShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListAllGroupsShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListAllGroupsShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListAllGroupsShrinkRequest) SetBizCode(v string) *ListAllGroupsShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *ListAllGroupsShrinkRequest) SetBizExtendShrink(v string) *ListAllGroupsShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *ListAllGroupsShrinkRequest) SetOwnerId(v int64) *ListAllGroupsShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ListAllGroupsShrinkRequest) SetResourceOwnerAccount(v string) *ListAllGroupsShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListAllGroupsShrinkRequest) SetResourceOwnerId(v int64) *ListAllGroupsShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListAllGroupsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
