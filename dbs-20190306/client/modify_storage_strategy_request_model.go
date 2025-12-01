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
	SetClientToken(v string) *ModifyStorageStrategyRequest
	GetClientToken() *string
	SetDuplicationArchivePeriod(v int32) *ModifyStorageStrategyRequest
	GetDuplicationArchivePeriod() *int32
	SetDuplicationInfrequentAccessPeriod(v int32) *ModifyStorageStrategyRequest
	GetDuplicationInfrequentAccessPeriod() *int32
	SetOwnerId(v string) *ModifyStorageStrategyRequest
	GetOwnerId() *string
}

type ModifyStorageStrategyRequest struct {
	// The ID of the backup schedule. You can call the [DescribeBackupPlanList](https://help.aliyun.com/document_detail/2869825.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dbsqdss5tmh****
	BackupPlanId *string `json:"BackupPlanId,omitempty" xml:"BackupPlanId,omitempty"`
	// The number of days for which the backup data is retained. Valid values: 0 to 1825.
	//
	// > Default value: 730.
	//
	// This parameter is required.
	//
	// example:
	//
	// 730
	BackupRetentionPeriod *int32 `json:"BackupRetentionPeriod,omitempty" xml:"BackupRetentionPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// dbstest
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The number of days after which the storage class of the backup data is changed to Archive. The value of this parameter must be smaller than the value of the BackupRetentionPeriod parameter. For more information about the Archive storage class, see [Storage class overview](https://help.aliyun.com/document_detail/51374.html).
	//
	// > Default value: 365.
	//
	// This parameter is required.
	//
	// example:
	//
	// 366
	DuplicationArchivePeriod *int32 `json:"DuplicationArchivePeriod,omitempty" xml:"DuplicationArchivePeriod,omitempty"`
	// The number of days after which the storage class of the backup data is changed to Infrequent Access (IA). The value of this parameter must be smaller than the value of the DuplicationArchivePeriod parameter. For more information about the IA storage class, see [Storage class overview](https://help.aliyun.com/document_detail/51374.html).
	//
	// > Default value: 180.
	//
	// This parameter is required.
	//
	// example:
	//
	// 190
	DuplicationInfrequentAccessPeriod *int32  `json:"DuplicationInfrequentAccessPeriod,omitempty" xml:"DuplicationInfrequentAccessPeriod,omitempty"`
	OwnerId                           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s *ModifyStorageStrategyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyStorageStrategyRequest) GetDuplicationArchivePeriod() *int32 {
	return s.DuplicationArchivePeriod
}

func (s *ModifyStorageStrategyRequest) GetDuplicationInfrequentAccessPeriod() *int32 {
	return s.DuplicationInfrequentAccessPeriod
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

func (s *ModifyStorageStrategyRequest) SetOwnerId(v string) *ModifyStorageStrategyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStorageStrategyRequest) Validate() error {
	return dara.Validate(s)
}
