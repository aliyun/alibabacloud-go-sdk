// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDetailCdrIbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDetailCdrIbRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkDetailCdrIbRequest
	GetHiddenType() *int64
	SetMainUniqueId(v string) *ClinkDetailCdrIbRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDetailCdrIbRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDetailCdrIbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDetailCdrIbRequest
	GetResourceOwnerId() *int64
}

type ClinkDetailCdrIbRequest struct {
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
	// 0
	HiddenType *int64 `json:"HiddenType,omitempty" xml:"HiddenType,omitempty"`
	// 通话记录唯一标识
	//
	// This parameter is required.
	//
	// example:
	//
	// medias_1-162046xxxx.12
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDetailCdrIbRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDetailCdrIbRequest) GoString() string {
	return s.String()
}

func (s *ClinkDetailCdrIbRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDetailCdrIbRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkDetailCdrIbRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDetailCdrIbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDetailCdrIbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDetailCdrIbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDetailCdrIbRequest) SetEnterpriseId(v int64) *ClinkDetailCdrIbRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) SetHiddenType(v int64) *ClinkDetailCdrIbRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) SetMainUniqueId(v string) *ClinkDetailCdrIbRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) SetOwnerId(v int64) *ClinkDetailCdrIbRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) SetResourceOwnerAccount(v string) *ClinkDetailCdrIbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) SetResourceOwnerId(v int64) *ClinkDetailCdrIbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDetailCdrIbRequest) Validate() error {
	return dara.Validate(s)
}
