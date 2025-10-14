// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeColumnarVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetColumnarVersion(v string) *UpgradeColumnarVersionRequest
	GetColumnarVersion() *string
	SetDBInstanceName(v string) *UpgradeColumnarVersionRequest
	GetDBInstanceName() *string
	SetInstanceName(v string) *UpgradeColumnarVersionRequest
	GetInstanceName() *string
	SetRegionId(v string) *UpgradeColumnarVersionRequest
	GetRegionId() *string
	SetSwitchMode(v string) *UpgradeColumnarVersionRequest
	GetSwitchMode() *string
}

type UpgradeColumnarVersionRequest struct {
	// example:
	//
	// polarx-col-kernel-5.4.20-20250819_17555906
	ColumnarVersion *string `json:"ColumnarVersion,omitempty" xml:"ColumnarVersion,omitempty"`
	// example:
	//
	// pxc-xxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// pxc-hzrh51fze****pon-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s UpgradeColumnarVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeColumnarVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeColumnarVersionRequest) GetColumnarVersion() *string {
	return s.ColumnarVersion
}

func (s *UpgradeColumnarVersionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradeColumnarVersionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpgradeColumnarVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeColumnarVersionRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *UpgradeColumnarVersionRequest) SetColumnarVersion(v string) *UpgradeColumnarVersionRequest {
	s.ColumnarVersion = &v
	return s
}

func (s *UpgradeColumnarVersionRequest) SetDBInstanceName(v string) *UpgradeColumnarVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeColumnarVersionRequest) SetInstanceName(v string) *UpgradeColumnarVersionRequest {
	s.InstanceName = &v
	return s
}

func (s *UpgradeColumnarVersionRequest) SetRegionId(v string) *UpgradeColumnarVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeColumnarVersionRequest) SetSwitchMode(v string) *UpgradeColumnarVersionRequest {
	s.SwitchMode = &v
	return s
}

func (s *UpgradeColumnarVersionRequest) Validate() error {
	return dara.Validate(s)
}
