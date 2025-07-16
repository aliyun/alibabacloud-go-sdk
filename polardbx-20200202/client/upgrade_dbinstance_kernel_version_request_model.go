// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionRequest
	GetDBInstanceName() *string
	SetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest
	GetMinorVersion() *string
	SetRegionId(v string) *UpgradeDBInstanceKernelVersionRequest
	GetRegionId() *string
	SetSwitchMode(v string) *UpgradeDBInstanceKernelVersionRequest
	GetSwitchMode() *string
}

type UpgradeDBInstanceKernelVersionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// polarx-kernel_5.4.12-16349923_xcluster-20210926
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeDBInstanceKernelVersionRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetDBInstanceName(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetMinorVersion(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetRegionId(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) SetSwitchMode(v string) *UpgradeDBInstanceKernelVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionRequest) Validate() error {
	return dara.Validate(s)
}
