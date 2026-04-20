// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIDBClusterDatasetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContinuationToken(v string) *DescribeAIDBClusterDatasetsRequest
	GetContinuationToken() *string
	SetDBClusterId(v string) *DescribeAIDBClusterDatasetsRequest
	GetDBClusterId() *string
	SetDatasetId(v string) *DescribeAIDBClusterDatasetsRequest
	GetDatasetId() *string
	SetDatasetType(v string) *DescribeAIDBClusterDatasetsRequest
	GetDatasetType() *string
	SetOwnerAccount(v string) *DescribeAIDBClusterDatasetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAIDBClusterDatasetsRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeAIDBClusterDatasetsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAIDBClusterDatasetsRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeAIDBClusterDatasetsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAIDBClusterDatasetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAIDBClusterDatasetsRequest
	GetResourceOwnerId() *int64
	SetTrainMode(v string) *DescribeAIDBClusterDatasetsRequest
	GetTrainMode() *string
}

type DescribeAIDBClusterDatasetsRequest struct {
	// example:
	//
	// EFSDF-DF-***
	ContinuationToken *string `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	// example:
	//
	// pc-2ze88***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// pds-2ze88***
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// example:
	//
	// train
	DatasetType  *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// grpo
	TrainMode *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
}

func (s DescribeAIDBClusterDatasetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIDBClusterDatasetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAIDBClusterDatasetsRequest) GetContinuationToken() *string {
	return s.ContinuationToken
}

func (s *DescribeAIDBClusterDatasetsRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAIDBClusterDatasetsRequest) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DescribeAIDBClusterDatasetsRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *DescribeAIDBClusterDatasetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAIDBClusterDatasetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAIDBClusterDatasetsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAIDBClusterDatasetsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAIDBClusterDatasetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAIDBClusterDatasetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAIDBClusterDatasetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAIDBClusterDatasetsRequest) GetTrainMode() *string {
	return s.TrainMode
}

func (s *DescribeAIDBClusterDatasetsRequest) SetContinuationToken(v string) *DescribeAIDBClusterDatasetsRequest {
	s.ContinuationToken = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetDBClusterId(v string) *DescribeAIDBClusterDatasetsRequest {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetDatasetId(v string) *DescribeAIDBClusterDatasetsRequest {
	s.DatasetId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetDatasetType(v string) *DescribeAIDBClusterDatasetsRequest {
	s.DatasetType = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetOwnerAccount(v string) *DescribeAIDBClusterDatasetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetOwnerId(v int64) *DescribeAIDBClusterDatasetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetPageNumber(v string) *DescribeAIDBClusterDatasetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetPageSize(v string) *DescribeAIDBClusterDatasetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetRegionId(v string) *DescribeAIDBClusterDatasetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetResourceOwnerAccount(v string) *DescribeAIDBClusterDatasetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetResourceOwnerId(v int64) *DescribeAIDBClusterDatasetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) SetTrainMode(v string) *DescribeAIDBClusterDatasetsRequest {
	s.TrainMode = &v
	return s
}

func (s *DescribeAIDBClusterDatasetsRequest) Validate() error {
	return dara.Validate(s)
}
