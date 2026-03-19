// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupPlanId(v string) *ModifyStorageStrategyRequest
	GetBackupPlanId() *string
	SetBackupRetentionPeriod(v int32) *ModifyStorageStrategyRequest
	GetBackupRetentionPeriod() *int32
	SetBackupStorageEncryptMethod(v string) *ModifyStorageStrategyRequest
	GetBackupStorageEncryptMethod() *string
	SetClientToken(v string) *ModifyStorageStrategyRequest
	GetClientToken() *string
	SetDuplicationArchivePeriod(v int32) *ModifyStorageStrategyRequest
	GetDuplicationArchivePeriod() *int32
	SetDuplicationInfrequentAccessPeriod(v int32) *ModifyStorageStrategyRequest
	GetDuplicationInfrequentAccessPeriod() *int32
	SetIncrementBackupRetentionPeriod(v string) *ModifyStorageStrategyRequest
	GetIncrementBackupRetentionPeriod() *string
	SetIncrementDuplicationArchivePeriod(v string) *ModifyStorageStrategyRequest
	GetIncrementDuplicationArchivePeriod() *string
	SetIncrementDuplicationInfrequentAccessPeriod(v string) *ModifyStorageStrategyRequest
	GetIncrementDuplicationInfrequentAccessPeriod() *string
	SetLogBackupRetentionPeriod(v string) *ModifyStorageStrategyRequest
	GetLogBackupRetentionPeriod() *string
	SetLogDuplicationArchivePeriod(v string) *ModifyStorageStrategyRequest
	GetLogDuplicationArchivePeriod() *string
	SetLogDuplicationInfrequentAccessPeriod(v string) *ModifyStorageStrategyRequest
	GetLogDuplicationInfrequentAccessPeriod() *string
	SetOwnerId(v string) *ModifyStorageStrategyRequest
	GetOwnerId() *string
}

type ModifyStorageStrategyRequest struct {
	// Backup plan ID. Obtain this parameter\\"s value by calling the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsqdss5tmh****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// Backup data retention period, in days. Valid values: 0 to 1825.
	//
	// > Default value: 730 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 730
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// example:
	//
	// encrypted
	BackupStorageEncryptMethod *string `json:"BackupStorageEncryptMethod,omitempty" xml:"BackupStorageEncryptMethod,omitempty"`
	// An arbitrary string used to ensure the idempotence of the request and prevent duplicate submissions.
	//
	// example:
	//
	// dbstest
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Time to convert to Archive Storage. This value must be less than the backup data retention period (BackupRetentionPeriod parameter). For more information about Archive Storage, see [Storage Type Overview](https://help.aliyun.com/document_detail/51374.html).
	//
	// > Default value: 365 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 366
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// Time to convert to Infrequent Access storage. This value must be less than the Archive Storage period (DuplicationArchivePeriod parameter). For more information about Infrequent Access storage, see [Storage Type Overview](https://help.aliyun.com/document_detail/51374.html).
	//
	// > Default value: 180 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 190
	DuplicationInfrequentAccessPeriod *int32 `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	// example:
	//
	// 365
	IncrementBackupRetentionPeriod *string `json:"IncrementBackupRetentionPeriod,omitempty" xml:"IncrementBackupRetentionPeriod,omitempty"`
	// example:
	//
	// 365
	IncrementDuplicationArchivePeriod *string `json:"IncrementDuplicationArchivePeriod,omitempty" xml:"IncrementDuplicationArchivePeriod,omitempty"`
	// example:
	//
	// 365
	IncrementDuplicationInfrequentAccessPeriod *string `json:"IncrementDuplicationInfrequentAccessPeriod,omitempty" xml:"IncrementDuplicationInfrequentAccessPeriod,omitempty"`
	// example:
	//
	// 365
	LogBackupRetentionPeriod *string `json:"LogBackupRetentionPeriod,omitempty" xml:"LogBackupRetentionPeriod,omitempty"`
	// example:
	//
	// 365
	LogDuplicationArchivePeriod *string `json:"LogDuplicationArchivePeriod,omitempty" xml:"LogDuplicationArchivePeriod,omitempty"`
	// example:
	//
	// 365
	LogDuplicationInfrequentAccessPeriod *string `json:"LogDuplicationInfrequentAccessPeriod,omitempty" xml:"LogDuplicationInfrequentAccessPeriod,omitempty"`
	OwnerId                              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyStorageStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyStorageStrategyRequest) GetBackupPlanId() *string {
	return s.BackupPlanId
}

func (s *ModifyStorageStrategyRequest) GetBackupRetentionPeriod() *int32 {
	return s.BackupRetentionPeriod
}

func (s *ModifyStorageStrategyRequest) GetBackupStorageEncryptMethod() *string {
	return s.BackupStorageEncryptMethod
}

func (s *ModifyStorageStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyStorageStrategyRequest) GetDuplicationArchivePeriod() *int32 {
	return s.DuplicationArchivePeriod
}

func (s *ModifyStorageStrategyRequest) GetDuplicationInfrequentAccessPeriod() *int32 {
	return s.DuplicationInfrequentAccessPeriod
}

func (s *ModifyStorageStrategyRequest) GetIncrementBackupRetentionPeriod() *string {
	return s.IncrementBackupRetentionPeriod
}

func (s *ModifyStorageStrategyRequest) GetIncrementDuplicationArchivePeriod() *string {
	return s.IncrementDuplicationArchivePeriod
}

func (s *ModifyStorageStrategyRequest) GetIncrementDuplicationInfrequentAccessPeriod() *string {
	return s.IncrementDuplicationInfrequentAccessPeriod
}

func (s *ModifyStorageStrategyRequest) GetLogBackupRetentionPeriod() *string {
	return s.LogBackupRetentionPeriod
}

func (s *ModifyStorageStrategyRequest) GetLogDuplicationArchivePeriod() *string {
	return s.LogDuplicationArchivePeriod
}

func (s *ModifyStorageStrategyRequest) GetLogDuplicationInfrequentAccessPeriod() *string {
	return s.LogDuplicationInfrequentAccessPeriod
}

func (s *ModifyStorageStrategyRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ModifyStorageStrategyRequest) SetBackupPlanId(v string) *ModifyStorageStrategyRequest {
	s.BackupPlanId = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetBackupRetentionPeriod(v int32) *ModifyStorageStrategyRequest {
	s.BackupRetentionPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetBackupStorageEncryptMethod(v string) *ModifyStorageStrategyRequest {
	s.BackupStorageEncryptMethod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetClientToken(v string) *ModifyStorageStrategyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetDuplicationArchivePeriod(v int32) *ModifyStorageStrategyRequest {
	s.DuplicationArchivePeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetDuplicationInfrequentAccessPeriod(v int32) *ModifyStorageStrategyRequest {
	s.DuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetIncrementBackupRetentionPeriod(v string) *ModifyStorageStrategyRequest {
	s.IncrementBackupRetentionPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetIncrementDuplicationArchivePeriod(v string) *ModifyStorageStrategyRequest {
	s.IncrementDuplicationArchivePeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetIncrementDuplicationInfrequentAccessPeriod(v string) *ModifyStorageStrategyRequest {
	s.IncrementDuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetLogBackupRetentionPeriod(v string) *ModifyStorageStrategyRequest {
	s.LogBackupRetentionPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetLogDuplicationArchivePeriod(v string) *ModifyStorageStrategyRequest {
	s.LogDuplicationArchivePeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetLogDuplicationInfrequentAccessPeriod(v string) *ModifyStorageStrategyRequest {
	s.LogDuplicationInfrequentAccessPeriod = &v
	return s
}

func (s *ModifyStorageStrategyRequest) SetOwnerId(v string) *ModifyStorageStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStorageStrategyRequest) Validate() error {
	return dara.Validate(s)
}
