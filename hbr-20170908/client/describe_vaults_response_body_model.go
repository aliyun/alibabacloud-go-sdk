// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeVaultsResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeVaultsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeVaultsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVaultsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVaultsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeVaultsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *DescribeVaultsResponseBody
	GetTotalCount() *int32
	SetVaults(v *DescribeVaultsResponseBodyVaults) *DescribeVaultsResponseBody
	GetVaults() *DescribeVaultsResponseBodyVaults
}

type DescribeVaultsResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned. If the call is successful, "successful" is returned. If the call fails, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number for pagination, starting from 1. The default value is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, with a minimum value of 1, a maximum value of 99, and a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful.
	//
	// - true: Success - false: Failure
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Returns the total number of backup repositories.
	//
	// example:
	//
	// 8
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vaults     *DescribeVaultsResponseBodyVaults `json:"Vaults,omitempty" xml:"Vaults,omitempty" type:"Struct"`
}

func (s DescribeVaultsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeVaultsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVaultsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVaultsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVaultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVaultsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeVaultsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVaultsResponseBody) GetVaults() *DescribeVaultsResponseBodyVaults {
	return s.Vaults
}

func (s *DescribeVaultsResponseBody) SetCode(v string) *DescribeVaultsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetMessage(v string) *DescribeVaultsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetPageNumber(v int32) *DescribeVaultsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetPageSize(v int32) *DescribeVaultsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetRequestId(v string) *DescribeVaultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetSuccess(v bool) *DescribeVaultsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetTotalCount(v int32) *DescribeVaultsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVaultsResponseBody) SetVaults(v *DescribeVaultsResponseBodyVaults) *DescribeVaultsResponseBody {
	s.Vaults = v
	return s
}

func (s *DescribeVaultsResponseBody) Validate() error {
	if s.Vaults != nil {
		if err := s.Vaults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVaultsResponseBodyVaults struct {
	Vault []*DescribeVaultsResponseBodyVaultsVault `json:"Vault,omitempty" xml:"Vault,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaults) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaults) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaults) GetVault() []*DescribeVaultsResponseBodyVaultsVault {
	return s.Vault
}

func (s *DescribeVaultsResponseBodyVaults) SetVault(v []*DescribeVaultsResponseBodyVaultsVault) *DescribeVaultsResponseBodyVaults {
	s.Vault = v
	return s
}

func (s *DescribeVaultsResponseBodyVaults) Validate() error {
	if s.Vault != nil {
		for _, item := range s.Vault {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVaultsResponseBodyVaultsVault struct {
	ArchiveBytesDone          *int64                                                     `json:"ArchiveBytesDone,omitempty" xml:"ArchiveBytesDone,omitempty"`
	ArchiveStorageSize        *int64                                                     `json:"ArchiveStorageSize,omitempty" xml:"ArchiveStorageSize,omitempty"`
	BackupPlanStatistics      *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics `json:"BackupPlanStatistics,omitempty" xml:"BackupPlanStatistics,omitempty" type:"Struct"`
	BucketName                *string                                                    `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	BytesDone                 *int64                                                     `json:"BytesDone,omitempty" xml:"BytesDone,omitempty"`
	ChargeType                *string                                                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ChargedStorageSize        *int64                                                     `json:"ChargedStorageSize,omitempty" xml:"ChargedStorageSize,omitempty"`
	CompressionAlgorithm      *string                                                    `json:"CompressionAlgorithm,omitempty" xml:"CompressionAlgorithm,omitempty"`
	CreatedTime               *int64                                                     `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	Dedup                     *bool                                                      `json:"Dedup,omitempty" xml:"Dedup,omitempty"`
	Description               *string                                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	EncryptType               *string                                                    `json:"EncryptType,omitempty" xml:"EncryptType,omitempty"`
	IndexAvailable            *bool                                                      `json:"IndexAvailable,omitempty" xml:"IndexAvailable,omitempty"`
	IndexLevel                *string                                                    `json:"IndexLevel,omitempty" xml:"IndexLevel,omitempty"`
	IndexUpdateTime           *int64                                                     `json:"IndexUpdateTime,omitempty" xml:"IndexUpdateTime,omitempty"`
	KmsKeyId                  *string                                                    `json:"KmsKeyId,omitempty" xml:"KmsKeyId,omitempty"`
	LatestReplicationTime     *int64                                                     `json:"LatestReplicationTime,omitempty" xml:"LatestReplicationTime,omitempty"`
	RedundancyType            *string                                                    `json:"RedundancyType,omitempty" xml:"RedundancyType,omitempty"`
	Replication               *bool                                                      `json:"Replication,omitempty" xml:"Replication,omitempty"`
	ReplicationProgress       *DescribeVaultsResponseBodyVaultsVaultReplicationProgress  `json:"ReplicationProgress,omitempty" xml:"ReplicationProgress,omitempty" type:"Struct"`
	ReplicationSourceOwnerId  *int64                                                     `json:"ReplicationSourceOwnerId,omitempty" xml:"ReplicationSourceOwnerId,omitempty"`
	ReplicationSourceRegionId *string                                                    `json:"ReplicationSourceRegionId,omitempty" xml:"ReplicationSourceRegionId,omitempty"`
	ReplicationSourceVault    *bool                                                      `json:"ReplicationSourceVault,omitempty" xml:"ReplicationSourceVault,omitempty"`
	ReplicationSourceVaultId  *string                                                    `json:"ReplicationSourceVaultId,omitempty" xml:"ReplicationSourceVaultId,omitempty"`
	ReplicationStatus         *string                                                    `json:"ReplicationStatus,omitempty" xml:"ReplicationStatus,omitempty"`
	ReplicationTargetOwnerId  *int64                                                     `json:"ReplicationTargetOwnerId,omitempty" xml:"ReplicationTargetOwnerId,omitempty"`
	ReplicationTargetRegionId *string                                                    `json:"ReplicationTargetRegionId,omitempty" xml:"ReplicationTargetRegionId,omitempty"`
	ReplicationTargetVaultId  *string                                                    `json:"ReplicationTargetVaultId,omitempty" xml:"ReplicationTargetVaultId,omitempty"`
	ResourceGroupId           *string                                                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Retention                 *int64                                                     `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RsTargetAccountIds        *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds   `json:"RsTargetAccountIds,omitempty" xml:"RsTargetAccountIds,omitempty" type:"Struct"`
	SearchEnabled             *bool                                                      `json:"SearchEnabled,omitempty" xml:"SearchEnabled,omitempty"`
	SnapshotCount             *int64                                                     `json:"SnapshotCount,omitempty" xml:"SnapshotCount,omitempty"`
	SourceTypes               *DescribeVaultsResponseBodyVaultsVaultSourceTypes          `json:"SourceTypes,omitempty" xml:"SourceTypes,omitempty" type:"Struct"`
	Status                    *string                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageSize               *int64                                                     `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	Tags                      *DescribeVaultsResponseBodyVaultsVaultTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	TrialInfo                 *DescribeVaultsResponseBodyVaultsVaultTrialInfo            `json:"TrialInfo,omitempty" xml:"TrialInfo,omitempty" type:"Struct"`
	UpdatedTime               *int64                                                     `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
	VaultId                   *string                                                    `json:"VaultId,omitempty" xml:"VaultId,omitempty"`
	VaultName                 *string                                                    `json:"VaultName,omitempty" xml:"VaultName,omitempty"`
	VaultOwnerId              *int64                                                     `json:"VaultOwnerId,omitempty" xml:"VaultOwnerId,omitempty"`
	VaultRegionId             *string                                                    `json:"VaultRegionId,omitempty" xml:"VaultRegionId,omitempty"`
	VaultStatusMessage        *string                                                    `json:"VaultStatusMessage,omitempty" xml:"VaultStatusMessage,omitempty"`
	VaultStorageClass         *string                                                    `json:"VaultStorageClass,omitempty" xml:"VaultStorageClass,omitempty"`
	VaultType                 *string                                                    `json:"VaultType,omitempty" xml:"VaultType,omitempty"`
	WormEnabled               *bool                                                      `json:"WormEnabled,omitempty" xml:"WormEnabled,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVault) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVault) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetArchiveBytesDone() *int64 {
	return s.ArchiveBytesDone
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetArchiveStorageSize() *int64 {
	return s.ArchiveStorageSize
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetBackupPlanStatistics() *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	return s.BackupPlanStatistics
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetBucketName() *string {
	return s.BucketName
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetBytesDone() *int64 {
	return s.BytesDone
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetChargedStorageSize() *int64 {
	return s.ChargedStorageSize
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetCompressionAlgorithm() *string {
	return s.CompressionAlgorithm
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetCreatedTime() *int64 {
	return s.CreatedTime
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetDedup() *bool {
	return s.Dedup
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetDescription() *string {
	return s.Description
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetEncryptType() *string {
	return s.EncryptType
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetIndexAvailable() *bool {
	return s.IndexAvailable
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetIndexLevel() *string {
	return s.IndexLevel
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetIndexUpdateTime() *int64 {
	return s.IndexUpdateTime
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetKmsKeyId() *string {
	return s.KmsKeyId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetLatestReplicationTime() *int64 {
	return s.LatestReplicationTime
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetRedundancyType() *string {
	return s.RedundancyType
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplication() *bool {
	return s.Replication
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationProgress() *DescribeVaultsResponseBodyVaultsVaultReplicationProgress {
	return s.ReplicationProgress
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationSourceOwnerId() *int64 {
	return s.ReplicationSourceOwnerId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationSourceRegionId() *string {
	return s.ReplicationSourceRegionId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationSourceVault() *bool {
	return s.ReplicationSourceVault
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationSourceVaultId() *string {
	return s.ReplicationSourceVaultId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationStatus() *string {
	return s.ReplicationStatus
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationTargetOwnerId() *int64 {
	return s.ReplicationTargetOwnerId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationTargetRegionId() *string {
	return s.ReplicationTargetRegionId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetReplicationTargetVaultId() *string {
	return s.ReplicationTargetVaultId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetRetention() *int64 {
	return s.Retention
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetRsTargetAccountIds() *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds {
	return s.RsTargetAccountIds
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetSearchEnabled() *bool {
	return s.SearchEnabled
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetSnapshotCount() *int64 {
	return s.SnapshotCount
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetSourceTypes() *DescribeVaultsResponseBodyVaultsVaultSourceTypes {
	return s.SourceTypes
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetStatus() *string {
	return s.Status
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetStorageSize() *int64 {
	return s.StorageSize
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetTags() *DescribeVaultsResponseBodyVaultsVaultTags {
	return s.Tags
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetTrialInfo() *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	return s.TrialInfo
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetUpdatedTime() *int64 {
	return s.UpdatedTime
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultId() *string {
	return s.VaultId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultName() *string {
	return s.VaultName
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultOwnerId() *int64 {
	return s.VaultOwnerId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultRegionId() *string {
	return s.VaultRegionId
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultStatusMessage() *string {
	return s.VaultStatusMessage
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultStorageClass() *string {
	return s.VaultStorageClass
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetVaultType() *string {
	return s.VaultType
}

func (s *DescribeVaultsResponseBodyVaultsVault) GetWormEnabled() *bool {
	return s.WormEnabled
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetArchiveBytesDone(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ArchiveBytesDone = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetArchiveStorageSize(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ArchiveStorageSize = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBackupPlanStatistics(v *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) *DescribeVaultsResponseBodyVaultsVault {
	s.BackupPlanStatistics = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBucketName(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.BucketName = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetBytesDone(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.BytesDone = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetChargeType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ChargeType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetChargedStorageSize(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ChargedStorageSize = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetCompressionAlgorithm(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.CompressionAlgorithm = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetCreatedTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.CreatedTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetDedup(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.Dedup = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetDescription(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.Description = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetEncryptType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.EncryptType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexAvailable(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexAvailable = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexLevel(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexLevel = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetIndexUpdateTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.IndexUpdateTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetKmsKeyId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.KmsKeyId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetLatestReplicationTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.LatestReplicationTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetRedundancyType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.RedundancyType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplication(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.Replication = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationProgress(v *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationProgress = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceOwnerId(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceOwnerId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceRegionId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceRegionId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceVault(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceVault = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationSourceVaultId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationSourceVaultId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationStatus(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationStatus = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationTargetOwnerId(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationTargetOwnerId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationTargetRegionId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationTargetRegionId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetReplicationTargetVaultId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ReplicationTargetVaultId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetResourceGroupId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetRetention(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.Retention = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetRsTargetAccountIds(v *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) *DescribeVaultsResponseBodyVaultsVault {
	s.RsTargetAccountIds = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSearchEnabled(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.SearchEnabled = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSnapshotCount(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.SnapshotCount = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetSourceTypes(v *DescribeVaultsResponseBodyVaultsVaultSourceTypes) *DescribeVaultsResponseBodyVaultsVault {
	s.SourceTypes = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetStatus(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.Status = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetStorageSize(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.StorageSize = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetTags(v *DescribeVaultsResponseBodyVaultsVaultTags) *DescribeVaultsResponseBodyVaultsVault {
	s.Tags = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetTrialInfo(v *DescribeVaultsResponseBodyVaultsVaultTrialInfo) *DescribeVaultsResponseBodyVaultsVault {
	s.TrialInfo = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetUpdatedTime(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.UpdatedTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultName(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultName = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultOwnerId(v int64) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultOwnerId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultRegionId(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultRegionId = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultStatusMessage(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultStatusMessage = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultStorageClass(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultStorageClass = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetVaultType(v string) *DescribeVaultsResponseBodyVaultsVault {
	s.VaultType = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) SetWormEnabled(v bool) *DescribeVaultsResponseBodyVaultsVault {
	s.WormEnabled = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVault) Validate() error {
	if s.BackupPlanStatistics != nil {
		if err := s.BackupPlanStatistics.Validate(); err != nil {
			return err
		}
	}
	if s.ReplicationProgress != nil {
		if err := s.ReplicationProgress.Validate(); err != nil {
			return err
		}
	}
	if s.RsTargetAccountIds != nil {
		if err := s.RsTargetAccountIds.Validate(); err != nil {
			return err
		}
	}
	if s.SourceTypes != nil {
		if err := s.SourceTypes.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	if s.TrialInfo != nil {
		if err := s.TrialInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics struct {
	Archive          *int32 `json:"Archive,omitempty" xml:"Archive,omitempty"`
	CommonFileSystem *int32 `json:"CommonFileSystem,omitempty" xml:"CommonFileSystem,omitempty"`
	CommonNas        *int32 `json:"CommonNas,omitempty" xml:"CommonNas,omitempty"`
	Csg              *int32 `json:"Csg,omitempty" xml:"Csg,omitempty"`
	EcsFile          *int32 `json:"EcsFile,omitempty" xml:"EcsFile,omitempty"`
	EcsHana          *int32 `json:"EcsHana,omitempty" xml:"EcsHana,omitempty"`
	Isilon           *int32 `json:"Isilon,omitempty" xml:"Isilon,omitempty"`
	LocalFile        *int32 `json:"LocalFile,omitempty" xml:"LocalFile,omitempty"`
	LocalVm          *int32 `json:"LocalVm,omitempty" xml:"LocalVm,omitempty"`
	MySql            *int32 `json:"MySql,omitempty" xml:"MySql,omitempty"`
	Nas              *int32 `json:"Nas,omitempty" xml:"Nas,omitempty"`
	Oracle           *int32 `json:"Oracle,omitempty" xml:"Oracle,omitempty"`
	Oss              *int32 `json:"Oss,omitempty" xml:"Oss,omitempty"`
	Ots              *int32 `json:"Ots,omitempty" xml:"Ots,omitempty"`
	SqlServer        *int32 `json:"SqlServer,omitempty" xml:"SqlServer,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetArchive() *int32 {
	return s.Archive
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetCommonFileSystem() *int32 {
	return s.CommonFileSystem
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetCommonNas() *int32 {
	return s.CommonNas
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetCsg() *int32 {
	return s.Csg
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetEcsFile() *int32 {
	return s.EcsFile
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetEcsHana() *int32 {
	return s.EcsHana
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetIsilon() *int32 {
	return s.Isilon
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetLocalFile() *int32 {
	return s.LocalFile
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetLocalVm() *int32 {
	return s.LocalVm
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetMySql() *int32 {
	return s.MySql
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetNas() *int32 {
	return s.Nas
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetOracle() *int32 {
	return s.Oracle
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetOss() *int32 {
	return s.Oss
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetOts() *int32 {
	return s.Ots
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) GetSqlServer() *int32 {
	return s.SqlServer
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetArchive(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Archive = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetCommonFileSystem(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.CommonFileSystem = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetCommonNas(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.CommonNas = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetCsg(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Csg = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetEcsFile(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.EcsFile = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetEcsHana(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.EcsHana = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetIsilon(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Isilon = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetLocalFile(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.LocalFile = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetLocalVm(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.LocalVm = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetMySql(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.MySql = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetNas(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Nas = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOracle(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Oracle = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOss(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Oss = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetOts(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.Ots = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) SetSqlServer(v int32) *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics {
	s.SqlServer = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultBackupPlanStatistics) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsResponseBodyVaultsVaultReplicationProgress struct {
	HistoricalReplicationProgress *int32 `json:"HistoricalReplicationProgress,omitempty" xml:"HistoricalReplicationProgress,omitempty"`
	NewReplicationProgress        *int64 `json:"NewReplicationProgress,omitempty" xml:"NewReplicationProgress,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultReplicationProgress) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultReplicationProgress) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) GetHistoricalReplicationProgress() *int32 {
	return s.HistoricalReplicationProgress
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) GetNewReplicationProgress() *int64 {
	return s.NewReplicationProgress
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) SetHistoricalReplicationProgress(v int32) *DescribeVaultsResponseBodyVaultsVaultReplicationProgress {
	s.HistoricalReplicationProgress = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) SetNewReplicationProgress(v int64) *DescribeVaultsResponseBodyVaultsVaultReplicationProgress {
	s.NewReplicationProgress = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultReplicationProgress) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds struct {
	RsTargetAccountId []*int64 `json:"RsTargetAccountId,omitempty" xml:"RsTargetAccountId,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) GetRsTargetAccountId() []*int64 {
	return s.RsTargetAccountId
}

func (s *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) SetRsTargetAccountId(v []*int64) *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds {
	s.RsTargetAccountId = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultRsTargetAccountIds) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsResponseBodyVaultsVaultSourceTypes struct {
	SourceType []*string `json:"SourceType,omitempty" xml:"SourceType,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaultsVaultSourceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultSourceTypes) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultSourceTypes) GetSourceType() []*string {
	return s.SourceType
}

func (s *DescribeVaultsResponseBodyVaultsVaultSourceTypes) SetSourceType(v []*string) *DescribeVaultsResponseBodyVaultsVaultSourceTypes {
	s.SourceType = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultSourceTypes) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsResponseBodyVaultsVaultTags struct {
	Tag []*DescribeVaultsResponseBodyVaultsVaultTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTags) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTags) GetTag() []*DescribeVaultsResponseBodyVaultsVaultTagsTag {
	return s.Tag
}

func (s *DescribeVaultsResponseBodyVaultsVaultTags) SetTag(v []*DescribeVaultsResponseBodyVaultsVaultTagsTag) *DescribeVaultsResponseBodyVaultsVaultTags {
	s.Tag = v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVaultsResponseBodyVaultsVaultTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) SetKey(v string) *DescribeVaultsResponseBodyVaultsVaultTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) SetValue(v string) *DescribeVaultsResponseBodyVaultsVaultTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTagsTag) Validate() error {
	return dara.Validate(s)
}

type DescribeVaultsResponseBodyVaultsVaultTrialInfo struct {
	KeepAfterTrialExpiration *bool  `json:"KeepAfterTrialExpiration,omitempty" xml:"KeepAfterTrialExpiration,omitempty"`
	TrialExpireTime          *int64 `json:"TrialExpireTime,omitempty" xml:"TrialExpireTime,omitempty"`
	TrialStartTime           *int64 `json:"TrialStartTime,omitempty" xml:"TrialStartTime,omitempty"`
	TrialVaultReleaseTime    *int64 `json:"TrialVaultReleaseTime,omitempty" xml:"TrialVaultReleaseTime,omitempty"`
}

func (s DescribeVaultsResponseBodyVaultsVaultTrialInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultsResponseBodyVaultsVaultTrialInfo) GoString() string {
	return s.String()
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) GetKeepAfterTrialExpiration() *bool {
	return s.KeepAfterTrialExpiration
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) GetTrialExpireTime() *int64 {
	return s.TrialExpireTime
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) GetTrialStartTime() *int64 {
	return s.TrialStartTime
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) GetTrialVaultReleaseTime() *int64 {
	return s.TrialVaultReleaseTime
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetKeepAfterTrialExpiration(v bool) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.KeepAfterTrialExpiration = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialExpireTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialExpireTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialStartTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialStartTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) SetTrialVaultReleaseTime(v int64) *DescribeVaultsResponseBodyVaultsVaultTrialInfo {
	s.TrialVaultReleaseTime = &v
	return s
}

func (s *DescribeVaultsResponseBodyVaultsVaultTrialInfo) Validate() error {
	return dara.Validate(s)
}
