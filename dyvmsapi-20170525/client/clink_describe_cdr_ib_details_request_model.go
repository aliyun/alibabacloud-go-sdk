// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrIbDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDescribeCdrIbDetailsRequest
	GetEnterpriseId() *int64
	SetMainUniqueId(v string) *ClinkDescribeCdrIbDetailsRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDescribeCdrIbDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDescribeCdrIbDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeCdrIbDetailsRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeCdrIbDetailsRequest struct {
	// 呼叫中心 id
	//
	// This parameter is required.
	//
	// example:
	//
	// 8004970
	EnterpriseId *int64 `json:"EnterpriseId,omitempty" xml:"EnterpriseId,omitempty"`
	// 通话记录唯一标识
	//
	// example:
	//
	// uniq_1-162046xxxx.12
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeCdrIbDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrIbDetailsRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrIbDetailsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeCdrIbDetailsRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrIbDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeCdrIbDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeCdrIbDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeCdrIbDetailsRequest) SetEnterpriseId(v int64) *ClinkDescribeCdrIbDetailsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsRequest) SetMainUniqueId(v string) *ClinkDescribeCdrIbDetailsRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsRequest) SetOwnerId(v int64) *ClinkDescribeCdrIbDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsRequest) SetResourceOwnerAccount(v string) *ClinkDescribeCdrIbDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsRequest) SetResourceOwnerId(v int64) *ClinkDescribeCdrIbDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeCdrIbDetailsRequest) Validate() error {
	return dara.Validate(s)
}
