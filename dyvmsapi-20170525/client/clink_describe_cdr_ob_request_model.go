// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDescribeCdrObRequest
	GetEnterpriseId() *int64
	SetHiddenType(v int64) *ClinkDescribeCdrObRequest
	GetHiddenType() *int64
	SetMainUniqueId(v string) *ClinkDescribeCdrObRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDescribeCdrObRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDescribeCdrObRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeCdrObRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeCdrObRequest struct {
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
	// AWS_DEV_MEDIA_SERVER_8-1536892698.2
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeCdrObRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeCdrObRequest) GetHiddenType() *int64 {
	return s.HiddenType
}

func (s *ClinkDescribeCdrObRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrObRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeCdrObRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeCdrObRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeCdrObRequest) SetEnterpriseId(v int64) *ClinkDescribeCdrObRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) SetHiddenType(v int64) *ClinkDescribeCdrObRequest {
	s.HiddenType = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) SetMainUniqueId(v string) *ClinkDescribeCdrObRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) SetOwnerId(v int64) *ClinkDescribeCdrObRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) SetResourceOwnerAccount(v string) *ClinkDescribeCdrObRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) SetResourceOwnerId(v int64) *ClinkDescribeCdrObRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeCdrObRequest) Validate() error {
	return dara.Validate(s)
}
