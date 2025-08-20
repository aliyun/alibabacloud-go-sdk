// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLLMSimilarQuestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeLLMSimilarQuestionsRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DescribeLLMSimilarQuestionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLLMSimilarQuestionsRequest
	GetOwnerId() *int64
	SetQuery(v string) *DescribeLLMSimilarQuestionsRequest
	GetQuery() *string
	SetRegionId(v string) *DescribeLLMSimilarQuestionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLLMSimilarQuestionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLLMSimilarQuestionsRequest
	GetResourceOwnerId() *int64
}

type DescribeLLMSimilarQuestionsRequest struct {
	// The cluster ID.
	//
	// >  You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/454250.html) operation to query the information about all AnalyticDB for MySQL clusters within a region, including cluster IDs.
	//
	// example:
	//
	// am-bp1565u55p32****
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
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLLMSimilarQuestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLLMSimilarQuestionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLLMSimilarQuestionsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeLLMSimilarQuestionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLLMSimilarQuestionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLLMSimilarQuestionsRequest) GetQuery() *string {
	return s.Query
}

func (s *DescribeLLMSimilarQuestionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLLMSimilarQuestionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLLMSimilarQuestionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLLMSimilarQuestionsRequest) SetDBClusterId(v string) *DescribeLLMSimilarQuestionsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetOwnerAccount(v string) *DescribeLLMSimilarQuestionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetOwnerId(v int64) *DescribeLLMSimilarQuestionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetQuery(v string) *DescribeLLMSimilarQuestionsRequest {
	s.Query = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetRegionId(v string) *DescribeLLMSimilarQuestionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetResourceOwnerAccount(v string) *DescribeLLMSimilarQuestionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) SetResourceOwnerId(v int64) *DescribeLLMSimilarQuestionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLLMSimilarQuestionsRequest) Validate() error {
	return dara.Validate(s)
}
