// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeCDCVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCdcDbVersion(v string) *UpgradeCDCVersionRequest
	GetCdcDbVersion() *string
	SetCdcMinorVersion(v string) *UpgradeCDCVersionRequest
	GetCdcMinorVersion() *string
	SetDBInstanceName(v string) *UpgradeCDCVersionRequest
	GetDBInstanceName() *string
	SetInstanceName(v string) *UpgradeCDCVersionRequest
	GetInstanceName() *string
	SetRegionId(v string) *UpgradeCDCVersionRequest
	GetRegionId() *string
	SetSwitchMode(v string) *UpgradeCDCVersionRequest
	GetSwitchMode() *string
}

type UpgradeCDCVersionRequest struct {
	// The target database version for the upgrade. You can call the [DescribeDBClusterVersion](https://help.aliyun.com/document_detail/196830.html) operation to view the upgrade instructions for all database versions in the target region.
	//
	// example:
	//
	// 2.343
	CdcDbVersion *string `json:"CdcDbVersion,omitempty" xml:"CdcDbVersion,omitempty"`
	// The target minor version for the upgrade.
	//
	// example:
	//
	// polarx-cdc-kernel-5.4.19-20240928_17274884
	CdcMinorVersion *string `json:"CdcMinorVersion,omitempty" xml:"CdcMinorVersion,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in a specific region, including instance IDs.
	//
	// example:
	//
	// pxc-hzrp****3p72fi
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxc-shrvdc****wtf5uk-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region where the instance resides. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the details of regions supported by PolarDB-X, including region IDs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The switchover mode. Valid values:
	//
	// - 0: Switches immediately.
	//
	// - 1: Switches within the O&M window.
	//
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s UpgradeCDCVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeCDCVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeCDCVersionRequest) GetCdcDbVersion() *string {
	return s.CdcDbVersion
}

func (s *UpgradeCDCVersionRequest) GetCdcMinorVersion() *string {
	return s.CdcMinorVersion
}

func (s *UpgradeCDCVersionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeCDCVersionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpgradeCDCVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeCDCVersionRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *UpgradeCDCVersionRequest) SetCdcDbVersion(v string) *UpgradeCDCVersionRequest {
	s.CdcDbVersion = &v
	return s
}

func (s *UpgradeCDCVersionRequest) SetCdcMinorVersion(v string) *UpgradeCDCVersionRequest {
	s.CdcMinorVersion = &v
	return s
}

func (s *UpgradeCDCVersionRequest) SetDBInstanceName(v string) *UpgradeCDCVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeCDCVersionRequest) SetInstanceName(v string) *UpgradeCDCVersionRequest {
	s.InstanceName = &v
	return s
}

func (s *UpgradeCDCVersionRequest) SetRegionId(v string) *UpgradeCDCVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeCDCVersionRequest) SetSwitchMode(v string) *UpgradeCDCVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeCDCVersionRequest) Validate() error {
	return dara.Validate(s)
}
