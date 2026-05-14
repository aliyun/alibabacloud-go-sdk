// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDescribeCdrIbRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkDescribeCdrIbRequest
	GetHiddenType() *int64
	SetMainUniqueId(v string) *ClinkDescribeCdrIbRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDescribeCdrIbRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDescribeCdrIbRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeCdrIbRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeCdrIbRequest struct {
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
	// xxx
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeCdrIbRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeCdrIbRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkDescribeCdrIbRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrIbRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeCdrIbRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeCdrIbRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeCdrIbRequest) SetEnterpriseId(v int64) *ClinkDescribeCdrIbRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) SetHiddenType(v int64) *ClinkDescribeCdrIbRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) SetMainUniqueId(v string) *ClinkDescribeCdrIbRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) SetOwnerId(v int64) *ClinkDescribeCdrIbRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) SetResourceOwnerAccount(v string) *ClinkDescribeCdrIbRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) SetResourceOwnerId(v int64) *ClinkDescribeCdrIbRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeCdrIbRequest) Validate() error {
	return dara.Validate(s)
}
