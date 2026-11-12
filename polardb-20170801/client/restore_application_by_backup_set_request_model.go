// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreApplicationByBackupSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *RestoreApplicationByBackupSetRequest
	GetApplicationId() *string
	SetApplicationType(v string) *RestoreApplicationByBackupSetRequest
	GetApplicationType() *string
	SetBackupSetId(v string) *RestoreApplicationByBackupSetRequest
	GetBackupSetId() *string
	SetRegionId(v string) *RestoreApplicationByBackupSetRequest
	GetRegionId() *string
}

type RestoreApplicationByBackupSetRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The ID of the backup set used to restore the application. This parameter is required. If it is not provided, the API returns InvalidBackupSetID.NotFound (404).
	//
	// example:
	//
	// bs-123456
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
	// The region ID of the application.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestoreApplicationByBackupSetRequest) String() string {
	return dara.Prettify(s)
}

func (s RestoreApplicationByBackupSetRequest) GoString() string {
	return s.String()
}

func (s *RestoreApplicationByBackupSetRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *RestoreApplicationByBackupSetRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *RestoreApplicationByBackupSetRequest) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *RestoreApplicationByBackupSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestoreApplicationByBackupSetRequest) SetApplicationId(v string) *RestoreApplicationByBackupSetRequest {
	s.ApplicationId = &v
	return s
}

func (s *RestoreApplicationByBackupSetRequest) SetApplicationType(v string) *RestoreApplicationByBackupSetRequest {
	s.ApplicationType = &v
	return s
}

func (s *RestoreApplicationByBackupSetRequest) SetBackupSetId(v string) *RestoreApplicationByBackupSetRequest {
	s.BackupSetId = &v
	return s
}

func (s *RestoreApplicationByBackupSetRequest) SetRegionId(v string) *RestoreApplicationByBackupSetRequest {
	s.RegionId = &v
	return s
}

func (s *RestoreApplicationByBackupSetRequest) Validate() error {
	return dara.Validate(s)
}
