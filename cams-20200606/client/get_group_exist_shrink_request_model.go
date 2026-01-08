// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGroupExistShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *GetGroupExistShrinkRequest
	GetBizCode() *string
	SetBizExtendShrink(v string) *GetGroupExistShrinkRequest
	GetBizExtendShrink() *string
	SetGroupName(v string) *GetGroupExistShrinkRequest
	GetGroupName() *string
	SetOwnerId(v int64) *GetGroupExistShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetGroupExistShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetGroupExistShrinkRequest
	GetResourceOwnerId() *int64
}

type GetGroupExistShrinkRequest struct {
	// example:
	//
	// ALICOM_OPAAS
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// example:
	//
	// {}
	BizExtendShrink *string `json:"BizExtend,omitempty" xml:"BizExtend,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aaa
	GroupName            *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetGroupExistShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGroupExistShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetGroupExistShrinkRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *GetGroupExistShrinkRequest) GetBizExtendShrink() *string {
	return s.BizExtendShrink
}

func (s *GetGroupExistShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *GetGroupExistShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetGroupExistShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetGroupExistShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetGroupExistShrinkRequest) SetBizCode(v string) *GetGroupExistShrinkRequest {
	s.BizCode = &v
	return s
}

func (s *GetGroupExistShrinkRequest) SetBizExtendShrink(v string) *GetGroupExistShrinkRequest {
	s.BizExtendShrink = &v
	return s
}

func (s *GetGroupExistShrinkRequest) SetGroupName(v string) *GetGroupExistShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *GetGroupExistShrinkRequest) SetOwnerId(v int64) *GetGroupExistShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *GetGroupExistShrinkRequest) SetResourceOwnerAccount(v string) *GetGroupExistShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetGroupExistShrinkRequest) SetResourceOwnerId(v int64) *GetGroupExistShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetGroupExistShrinkRequest) Validate() error {
	return dara.Validate(s)
}
