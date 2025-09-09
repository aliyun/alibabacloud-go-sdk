// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupDbNames(v string) *DescribeRestoreOrderRequest
	GetBackupDbNames() *string
	SetBackupId(v string) *DescribeRestoreOrderRequest
	GetBackupId() *string
	SetBackupLevel(v string) *DescribeRestoreOrderRequest
	GetBackupLevel() *string
	SetBackupMode(v string) *DescribeRestoreOrderRequest
	GetBackupMode() *string
	SetDrdsInstanceId(v string) *DescribeRestoreOrderRequest
	GetDrdsInstanceId() *string
	SetPreferredBackupTime(v string) *DescribeRestoreOrderRequest
	GetPreferredBackupTime() *string
}

type DescribeRestoreOrderRequest struct {
	// The name of the database involved in the backup.
	//
	// example:
	//
	// drds_flashback
	BackupDbNames *string `json:"BackupDbNames,omitempty" xml:"BackupDbNames,omitempty"`
	// The ID of the backup set.
	//
	// example:
	//
	// 1918df27-4563-11e9-8403-af4fbe******
	BackupId *string `json:"BackupId,omitempty" xml:"BackupId,omitempty"`
	// The level of the backup. Valid values:
	//
	// 	- **DB**: The database Level
	//
	// 	- **instance **: instance level
	//
	// example:
	//
	// db
	BackupLevel *string `json:"BackupLevel,omitempty" xml:"BackupLevel,omitempty"`
	// The backup mode. Valid values: **logic*	- or **phy**.
	//
	// example:
	//
	// phy
	BackupMode *string `json:"BackupMode,omitempty" xml:"BackupMode,omitempty"`
	// The ID of the instance for which to modify the backup policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds************
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The preferred backup time.
	//
	// example:
	//
	// 2019-09-16 15:12:53
	PreferredBackupTime *string `json:"PreferredBackupTime,omitempty" xml:"PreferredBackupTime,omitempty"`
}

func (s DescribeRestoreOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeRestoreOrderRequest) GetBackupDbNames() *string {
	return s.BackupDbNames
}

func (s *DescribeRestoreOrderRequest) GetBackupId() *string {
	return s.BackupId
}

func (s *DescribeRestoreOrderRequest) GetBackupLevel() *string {
	return s.BackupLevel
}

func (s *DescribeRestoreOrderRequest) GetBackupMode() *string {
	return s.BackupMode
}

func (s *DescribeRestoreOrderRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeRestoreOrderRequest) GetPreferredBackupTime() *string {
	return s.PreferredBackupTime
}

func (s *DescribeRestoreOrderRequest) SetBackupDbNames(v string) *DescribeRestoreOrderRequest {
	s.BackupDbNames = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupId(v string) *DescribeRestoreOrderRequest {
	s.BackupId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupLevel(v string) *DescribeRestoreOrderRequest {
	s.BackupLevel = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetBackupMode(v string) *DescribeRestoreOrderRequest {
	s.BackupMode = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetDrdsInstanceId(v string) *DescribeRestoreOrderRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeRestoreOrderRequest) SetPreferredBackupTime(v string) *DescribeRestoreOrderRequest {
	s.PreferredBackupTime = &v
	return s
}

func (s *DescribeRestoreOrderRequest) Validate() error {
	return dara.Validate(s)
}
