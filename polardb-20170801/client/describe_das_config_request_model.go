// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDasConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeDasConfigRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeDasConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDasConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeDasConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDasConfigRequest
	GetResourceOwnerId() *int64
}

type DescribeDasConfigRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/98094.html) operation to query the details of the clusters that belong to your Alibaba Cloud account, such as cluster IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-****************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDasConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDasConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDasConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeDasConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDasConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDasConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDasConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDasConfigRequest) SetDBClusterId(v string) *DescribeDasConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDasConfigRequest) SetOwnerAccount(v string) *DescribeDasConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDasConfigRequest) SetOwnerId(v int64) *DescribeDasConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDasConfigRequest) SetResourceOwnerAccount(v string) *DescribeDasConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDasConfigRequest) SetResourceOwnerId(v int64) *DescribeDasConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDasConfigRequest) Validate() error {
	return dara.Validate(s)
}
