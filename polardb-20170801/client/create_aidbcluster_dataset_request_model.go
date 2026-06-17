// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAIDBClusterDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateAIDBClusterDatasetRequest
	GetDBClusterId() *string
	SetDatasetName(v string) *CreateAIDBClusterDatasetRequest
	GetDatasetName() *string
	SetDatasetType(v string) *CreateAIDBClusterDatasetRequest
	GetDatasetType() *string
	SetImportMode(v string) *CreateAIDBClusterDatasetRequest
	GetImportMode() *string
	SetOwnerAccount(v string) *CreateAIDBClusterDatasetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateAIDBClusterDatasetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateAIDBClusterDatasetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateAIDBClusterDatasetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateAIDBClusterDatasetRequest
	GetResourceOwnerId() *int64
	SetTrainMode(v string) *CreateAIDBClusterDatasetRequest
	GetTrainMode() *string
}

type CreateAIDBClusterDatasetRequest struct {
	// The ID of the PolarDB database cluster.
	//
	// example:
	//
	// pc-2ze88***
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The dataset name.
	//
	// example:
	//
	// dataset01
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The type of the dataset. Valid values:
	//
	// - **train**: training set
	//
	// - **eval**: evaluation set
	//
	// example:
	//
	// train
	DatasetType *string `json:"DatasetType,omitempty" xml:"DatasetType,omitempty"`
	// The import method. Valid values:
	//
	// - **LocalImport**: local import
	//
	// example:
	//
	// LocalImport
	ImportMode   *string `json:"ImportMode,omitempty" xml:"ImportMode,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The training mode for the dataset. The dataset format depends on the mode. Valid values:
	//
	// - **sft**: supervised fine-tuning. For training sets only.
	//
	// - **grpo**: reinforcement learning optimization. For training sets only.
	//
	// - **text**: text generation. For validation sets only.
	//
	// example:
	//
	// sft
	TrainMode *string `json:"TrainMode,omitempty" xml:"TrainMode,omitempty"`
}

func (s CreateAIDBClusterDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAIDBClusterDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateAIDBClusterDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateAIDBClusterDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateAIDBClusterDatasetRequest) GetDatasetType() *string {
	return s.DatasetType
}

func (s *CreateAIDBClusterDatasetRequest) GetImportMode() *string {
	return s.ImportMode
}

func (s *CreateAIDBClusterDatasetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateAIDBClusterDatasetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateAIDBClusterDatasetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAIDBClusterDatasetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateAIDBClusterDatasetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateAIDBClusterDatasetRequest) GetTrainMode() *string {
	return s.TrainMode
}

func (s *CreateAIDBClusterDatasetRequest) SetDBClusterId(v string) *CreateAIDBClusterDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetDatasetName(v string) *CreateAIDBClusterDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetDatasetType(v string) *CreateAIDBClusterDatasetRequest {
	s.DatasetType = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetImportMode(v string) *CreateAIDBClusterDatasetRequest {
	s.ImportMode = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetOwnerAccount(v string) *CreateAIDBClusterDatasetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetOwnerId(v int64) *CreateAIDBClusterDatasetRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetRegionId(v string) *CreateAIDBClusterDatasetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetResourceOwnerAccount(v string) *CreateAIDBClusterDatasetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetResourceOwnerId(v int64) *CreateAIDBClusterDatasetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) SetTrainMode(v string) *CreateAIDBClusterDatasetRequest {
	s.TrainMode = &v
	return s
}

func (s *CreateAIDBClusterDatasetRequest) Validate() error {
	return dara.Validate(s)
}
