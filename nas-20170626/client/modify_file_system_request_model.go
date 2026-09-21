// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoUpgradeConfig(v *ModifyFileSystemRequestAutoUpgradeConfig) *ModifyFileSystemRequest
	GetAutoUpgradeConfig() *ModifyFileSystemRequestAutoUpgradeConfig
	SetDescription(v string) *ModifyFileSystemRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyFileSystemRequest
	GetFileSystemId() *string
	SetOptions(v *ModifyFileSystemRequestOptions) *ModifyFileSystemRequest
	GetOptions() *ModifyFileSystemRequestOptions
}

type ModifyFileSystemRequest struct {
	// The auto-scaling configuration.
	AutoUpgradeConfig *ModifyFileSystemRequestAutoUpgradeConfig `json:"AutoUpgradeConfig,omitempty" xml:"AutoUpgradeConfig,omitempty" type:"Struct"`
	// The description of the file system.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter. It cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// NAS-test-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// - General-purpose NAS: `31a8e4****`.
	//
	// - Extreme NAS: The ID must start with `extreme-`, for example, `extreme-0015****`.
	//
	// - Cloud Parallel File Storage (CPFS): The ID must start with `cpfs-`, for example, `cpfs-125487****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The options.
	Options *ModifyFileSystemRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
}

func (s ModifyFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) GetAutoUpgradeConfig() *ModifyFileSystemRequestAutoUpgradeConfig {
	return s.AutoUpgradeConfig
}

func (s *ModifyFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFileSystemRequest) GetOptions() *ModifyFileSystemRequestOptions {
	return s.Options
}

func (s *ModifyFileSystemRequest) SetAutoUpgradeConfig(v *ModifyFileSystemRequestAutoUpgradeConfig) *ModifyFileSystemRequest {
	s.AutoUpgradeConfig = v
	return s
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetOptions(v *ModifyFileSystemRequestOptions) *ModifyFileSystemRequest {
	s.Options = v
	return s
}

func (s *ModifyFileSystemRequest) Validate() error {
	if s.AutoUpgradeConfig != nil {
		if err := s.AutoUpgradeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyFileSystemRequestAutoUpgradeConfig struct {
	// The capacity usage threshold.
	//
	// example:
	//
	// 80
	CapacityUsedRatio *int32 `json:"capacityUsedRatio,omitempty" xml:"capacityUsedRatio,omitempty"`
	// Specifies whether to enable auto-scaling.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The scaling increment.
	//
	// example:
	//
	// 100
	Step *int32 `json:"step,omitempty" xml:"step,omitempty"`
	// The duration.
	//
	// example:
	//
	// 30
	Time *int32 `json:"time,omitempty" xml:"time,omitempty"`
}

func (s ModifyFileSystemRequestAutoUpgradeConfig) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequestAutoUpgradeConfig) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) GetCapacityUsedRatio() *int32 {
	return s.CapacityUsedRatio
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) GetEnabled() *bool {
	return s.Enabled
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) GetStep() *int32 {
	return s.Step
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) GetTime() *int32 {
	return s.Time
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) SetCapacityUsedRatio(v int32) *ModifyFileSystemRequestAutoUpgradeConfig {
	s.CapacityUsedRatio = &v
	return s
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) SetEnabled(v bool) *ModifyFileSystemRequestAutoUpgradeConfig {
	s.Enabled = &v
	return s
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) SetStep(v int32) *ModifyFileSystemRequestAutoUpgradeConfig {
	s.Step = &v
	return s
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) SetTime(v int32) *ModifyFileSystemRequestAutoUpgradeConfig {
	s.Time = &v
	return s
}

func (s *ModifyFileSystemRequestAutoUpgradeConfig) Validate() error {
	return dara.Validate(s)
}

type ModifyFileSystemRequestOptions struct {
	// Specifies whether to enable the SMB Access-Based Enumeration (ABE) feature.
	//
	// example:
	//
	// false
	EnableABE *bool `json:"EnableABE,omitempty" xml:"EnableABE,omitempty"`
	// Specifies whether the OpLock feature is enabled.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Not enabled.
	//
	// > Only file systems of the SMB Protocol Type are supported.
	//
	// example:
	//
	// true
	EnableOplock *bool `json:"EnableOplock,omitempty" xml:"EnableOplock,omitempty"`
	// Specifies whether the Lingjun VSC mount target supports only access point-based access.
	//
	// example:
	//
	// false
	VscAccessPointAccessOnly *bool `json:"VscAccessPointAccessOnly,omitempty" xml:"VscAccessPointAccessOnly,omitempty"`
}

func (s ModifyFileSystemRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequestOptions) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequestOptions) GetEnableABE() *bool {
	return s.EnableABE
}

func (s *ModifyFileSystemRequestOptions) GetEnableOplock() *bool {
	return s.EnableOplock
}

func (s *ModifyFileSystemRequestOptions) GetVscAccessPointAccessOnly() *bool {
	return s.VscAccessPointAccessOnly
}

func (s *ModifyFileSystemRequestOptions) SetEnableABE(v bool) *ModifyFileSystemRequestOptions {
	s.EnableABE = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) SetEnableOplock(v bool) *ModifyFileSystemRequestOptions {
	s.EnableOplock = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) SetVscAccessPointAccessOnly(v bool) *ModifyFileSystemRequestOptions {
	s.VscAccessPointAccessOnly = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) Validate() error {
	return dara.Validate(s)
}
