// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationBackupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationBackupRequest
	GetApplicationId() *string
	SetApplicationType(v string) *CreateApplicationBackupRequest
	GetApplicationType() *string
	SetJobMode(v string) *CreateApplicationBackupRequest
	GetJobMode() *string
	SetRegionId(v string) *CreateApplicationBackupRequest
	GetRegionId() *string
}

type CreateApplicationBackupRequest struct {
	// The ID of the application to back up.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The application type, such as polarclaw.
	//
	// example:
	//
	// polarclaw
	ApplicationType *string `json:"ApplicationType,omitempty" xml:"ApplicationType,omitempty"`
	// The backup mode. Automated indicates an automatic backup, and Manual indicates a manual backup.
	//
	// This parameter is required.
	//
	// example:
	//
	// Manual
	JobMode *string `json:"JobMode,omitempty" xml:"JobMode,omitempty"`
	// The ID of the region where the application resides. The server uses the actual region of the application.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateApplicationBackupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationBackupRequest) GoString() string {
	return s.String()
}

func (s *CreateApplicationBackupRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationBackupRequest) GetApplicationType() *string {
	return s.ApplicationType
}

func (s *CreateApplicationBackupRequest) GetJobMode() *string {
	return s.JobMode
}

func (s *CreateApplicationBackupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateApplicationBackupRequest) SetApplicationId(v string) *CreateApplicationBackupRequest {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationBackupRequest) SetApplicationType(v string) *CreateApplicationBackupRequest {
	s.ApplicationType = &v
	return s
}

func (s *CreateApplicationBackupRequest) SetJobMode(v string) *CreateApplicationBackupRequest {
	s.JobMode = &v
	return s
}

func (s *CreateApplicationBackupRequest) SetRegionId(v string) *CreateApplicationBackupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateApplicationBackupRequest) Validate() error {
	return dara.Validate(s)
}
