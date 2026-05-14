// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkDescribeCdrObDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnterpriseId(v int64) *ClinkDescribeCdrObDetailsRequest
	GetEnterpriseId() *int64
	SetMainUniqueId(v string) *ClinkDescribeCdrObDetailsRequest
	GetMainUniqueId() *string
	SetOwnerId(v int64) *ClinkDescribeCdrObDetailsRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ClinkDescribeCdrObDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ClinkDescribeCdrObDetailsRequest
	GetResourceOwnerId() *int64
}

type ClinkDescribeCdrObDetailsRequest struct {
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
	// AWS_DEV_MEDIA_SERVER_8-1536201058.19
	MainUniqueId         *string `json:"MainUniqueId,omitempty" xml:"MainUniqueId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ClinkDescribeCdrObDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ClinkDescribeCdrObDetailsRequest) GoString() string {
	return s.String()
}

func (s *ClinkDescribeCdrObDetailsRequest) GetEnterpriseId() *int64 {
	return s.EnterpriseId
}

func (s *ClinkDescribeCdrObDetailsRequest) GetMainUniqueId() *string {
	return s.MainUniqueId
}

func (s *ClinkDescribeCdrObDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ClinkDescribeCdrObDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ClinkDescribeCdrObDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ClinkDescribeCdrObDetailsRequest) SetEnterpriseId(v int64) *ClinkDescribeCdrObDetailsRequest {
	s.EnterpriseId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsRequest) SetMainUniqueId(v string) *ClinkDescribeCdrObDetailsRequest {
	s.MainUniqueId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsRequest) SetOwnerId(v int64) *ClinkDescribeCdrObDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsRequest) SetResourceOwnerAccount(v string) *ClinkDescribeCdrObDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsRequest) SetResourceOwnerId(v int64) *ClinkDescribeCdrObDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ClinkDescribeCdrObDetailsRequest) Validate() error {
	return dara.Validate(s)
}
