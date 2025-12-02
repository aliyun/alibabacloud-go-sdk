// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBackupConfig(v string) *UpdateBackupConfigRequest
	GetBackupConfig() *string
	SetRegionId(v string) *UpdateBackupConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *UpdateBackupConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *UpdateBackupConfigRequest
	GetServiceCode() *string
}

type UpdateBackupConfigRequest struct {
	// Evidence backup configuration.
	//
	// example:
	//
	// {}
	BackupConfig *string `json:"BackupConfig,omitempty" xml:"BackupConfig,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource type.
	//
	// example:
	//
	// video
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Service code.
	//
	// example:
	//
	// videoDetection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s UpdateBackupConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigRequest) GetBackupConfig() *string {
	return s.BackupConfig
}

func (s *UpdateBackupConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateBackupConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UpdateBackupConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *UpdateBackupConfigRequest) SetBackupConfig(v string) *UpdateBackupConfigRequest {
	s.BackupConfig = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetRegionId(v string) *UpdateBackupConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetResourceType(v string) *UpdateBackupConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *UpdateBackupConfigRequest) SetServiceCode(v string) *UpdateBackupConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *UpdateBackupConfigRequest) Validate() error {
	return dara.Validate(s)
}
