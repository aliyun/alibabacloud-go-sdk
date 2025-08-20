// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMAnswerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLLMAnswerRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeLLMAnswerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLLMAnswerRequest
	GetOwnerId() *int64
	SetQuery(v string) *DescribeLLMAnswerRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeLLMAnswerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLLMAnswerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLLMAnswerRequest
	GetResourceOwnerId() *int64
}

type DescribeLLMAnswerRequest struct {
	// The cluster ID.
	//
	// >  Enterprise Edition, Basic Edition, and Data Lakehouse Edition: You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// example:
	//
	// am-uf6g8w25jacm7****
	DBClusterId  *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The question proposed by a user.
	//
	// This parameter is required.
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// The region ID
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/612393.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLLMAnswerRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMAnswerRequest) GoString() string {
	return s.String()
}

func (s *DescribeLLMAnswerRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLLMAnswerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLLMAnswerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLLMAnswerRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeLLMAnswerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLLMAnswerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLLMAnswerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLLMAnswerRequest) SetDBClusterId(v string) *DescribeLLMAnswerRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetOwnerAccount(v string) *DescribeLLMAnswerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetOwnerId(v int64) *DescribeLLMAnswerRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetQuery(v string) *DescribeLLMAnswerRequest {
	s.Query = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetRegionId(v string) *DescribeLLMAnswerRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetResourceOwnerAccount(v string) *DescribeLLMAnswerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLLMAnswerRequest) SetResourceOwnerId(v int64) *DescribeLLMAnswerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLLMAnswerRequest) Validate() error {
	return dara.Validate(s)
}
