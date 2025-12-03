// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClusterServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterServiceConfigRequest
	GetClusterId() *string
	SetOwnerId(v int64) *ListClusterServiceConfigRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *ListClusterServiceConfigRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListClusterServiceConfigRequest
	GetPageSize() *int32
	SetResourceOwnerAccount(v string) *ListClusterServiceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListClusterServiceConfigRequest
	GetResourceOwnerId() *int64
}

type ListClusterServiceConfigRequest struct {
	ClusterId            *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListClusterServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *ListClusterServiceConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterServiceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListClusterServiceConfigRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClusterServiceConfigRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClusterServiceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListClusterServiceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListClusterServiceConfigRequest) SetClusterId(v string) *ListClusterServiceConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetOwnerId(v int64) *ListClusterServiceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetPageNumber(v int32) *ListClusterServiceConfigRequest {
	s.PageNumber = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetPageSize(v int32) *ListClusterServiceConfigRequest {
	s.PageSize = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetResourceOwnerAccount(v string) *ListClusterServiceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListClusterServiceConfigRequest) SetResourceOwnerId(v int64) *ListClusterServiceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListClusterServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
