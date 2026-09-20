// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifyBackupPlanConfigRequest
	GetClusterId() *string
	SetFullBackupCycle(v string) *ModifyBackupPlanConfigRequest
	GetFullBackupCycle() *string
	SetMinHFileBackupCount(v string) *ModifyBackupPlanConfigRequest
	GetMinHFileBackupCount() *string
	SetNextFullBackupDate(v string) *ModifyBackupPlanConfigRequest
	GetNextFullBackupDate() *string
	SetTables(v string) *ModifyBackupPlanConfigRequest
	GetTables() *string
}

type ModifyBackupPlanConfigRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-m5eznlga4k5bcxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The full backup cycle, in days. Valid values: 3 to 9.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7
	FullBackupCycle *string `json:"FullBackupCycle,omitempty" xml:"FullBackupCycle,omitempty"`
	// The number of full backups to retain. Valid values: 3 to 8.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	MinHFileBackupCount *string `json:"MinHFileBackupCount,omitempty" xml:"MinHFileBackupCount,omitempty"`
	// The time of the next full backup. The specified time must be at least 6 minutes later than the current time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-11-09T18:00:00Z
	NextFullBackupDate *string `json:"NextFullBackupDate,omitempty" xml:"NextFullBackupDate,omitempty"`
	// The tables to back up. Specify one table name per line. Wildcards are supported. An asterisk (*) indicates all tables.
	//
	// This parameter is required.
	//
	// example:
	//
	// *
	Tables *string `json:"Tables,omitempty" xml:"Tables,omitempty"`
}

func (s ModifyBackupPlanConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifyBackupPlanConfigRequest) GetFullBackupCycle() *string {
	return s.FullBackupCycle
}

func (s *ModifyBackupPlanConfigRequest) GetMinHFileBackupCount() *string {
	return s.MinHFileBackupCount
}

func (s *ModifyBackupPlanConfigRequest) GetNextFullBackupDate() *string {
	return s.NextFullBackupDate
}

func (s *ModifyBackupPlanConfigRequest) GetTables() *string {
	return s.Tables
}

func (s *ModifyBackupPlanConfigRequest) SetClusterId(v string) *ModifyBackupPlanConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetFullBackupCycle(v string) *ModifyBackupPlanConfigRequest {
	s.FullBackupCycle = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetMinHFileBackupCount(v string) *ModifyBackupPlanConfigRequest {
	s.MinHFileBackupCount = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetNextFullBackupDate(v string) *ModifyBackupPlanConfigRequest {
	s.NextFullBackupDate = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) SetTables(v string) *ModifyBackupPlanConfigRequest {
	s.Tables = &v
	return s
}

func (s *ModifyBackupPlanConfigRequest) Validate() error {
	return dara.Validate(s)
}
