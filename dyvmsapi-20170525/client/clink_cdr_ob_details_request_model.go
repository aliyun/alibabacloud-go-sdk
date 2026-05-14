// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCdrObDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkCdrObDetailsRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkCdrObDetailsRequest
	GetHiddenType() *int64
	SetMainUniqueId(v string) *ClinkCdrObDetailsRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkCdrObDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkCdrObDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkCdrObDetailsRequest
	GetResourceOwnerId() *int64
}

type ClinkCdrObDetailsRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 是否隐藏号码。 0: 不隐藏，1: 中间四位，2: 最后八位，3: 全部号码，4: 最后四位。默认值为 0
	//
	// example:
	//
	// 1
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 通话记录唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkCdrObDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkCdrObDetailsRequest) GoString() string {
	return s.String()
}

func (s *ClinkCdrObDetailsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkCdrObDetailsRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkCdrObDetailsRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkCdrObDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkCdrObDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkCdrObDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkCdrObDetailsRequest) SetEnterpriseId(v int64) *ClinkCdrObDetailsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) SetHiddenType(v int64) *ClinkCdrObDetailsRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) SetMainUniqueId(v string) *ClinkCdrObDetailsRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) SetOwnerId(v int64) *ClinkCdrObDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) SetResourceOwnerAccount(v string) *ClinkCdrObDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) SetResourceOwnerId(v int64) *ClinkCdrObDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkCdrObDetailsRequest) Validate() error {
	return dara.Validate(s)
}
