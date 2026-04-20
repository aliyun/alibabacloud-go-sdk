// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKubeType(v string) *DescribeAIDBClusterTasksRequest
	GetKubeType() *string
	SetOwnerAccount(v string) *DescribeAIDBClusterTasksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClusterTasksRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAIDBClusterTasksRequest
	GetRegionId() *string
	SetRelativeDBClusterId(v string) *DescribeAIDBClusterTasksRequest
	GetRelativeDBClusterId() *string
	SetResourceOwnerAccount(v string) *DescribeAIDBClusterTasksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClusterTasksRequest
	GetResourceOwnerId() *int64
}

type DescribeAIDBClusterTasksRequest struct {
	// example:
	//
	// aitrain
	KubeType     *string `json:"KubeType,omitempty" xml:"KubeType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// pc-2ze88***
	RelativeDBClusterId  *string `json:"RelativeDBClusterId,omitempty" xml:"RelativeDBClusterId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAIDBClusterTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterTasksRequest) GetKubeType() *string {
	return s.KubeType
}

func (s *DescribeAIDBClusterTasksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClusterTasksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClusterTasksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterTasksRequest) GetRelativeDBClusterId() *string {
	return s.RelativeDBClusterId
}

func (s *DescribeAIDBClusterTasksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClusterTasksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClusterTasksRequest) SetKubeType(v string) *DescribeAIDBClusterTasksRequest {
	s.KubeType = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetOwnerAccount(v string) *DescribeAIDBClusterTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetOwnerId(v int64) *DescribeAIDBClusterTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetRegionId(v string) *DescribeAIDBClusterTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetRelativeDBClusterId(v string) *DescribeAIDBClusterTasksRequest {
	s.RelativeDBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClusterTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) SetResourceOwnerId(v int64) *DescribeAIDBClusterTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClusterTasksRequest) Validate() error {
	return dara.Validate(s)
}
