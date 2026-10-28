// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradePxfsVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *UpgradePxfsVersionRequest
	GetDBInstanceName() *string
	SetMinorVersion(v string) *UpgradePxfsVersionRequest
	GetMinorVersion() *string
	SetRegionId(v string) *UpgradePxfsVersionRequest
	GetRegionId() *string
}

type UpgradePxfsVersionRequest struct {
	// The PolarDB-X instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-xxxxxxxx
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The target minor version.
	//
	// example:
	//
	// 1.1.0
	MinorVersion *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradePxfsVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradePxfsVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradePxfsVersionRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *UpgradePxfsVersionRequest) GetMinorVersion() *string {
	return s.MinorVersion
}

func (s *UpgradePxfsVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradePxfsVersionRequest) SetDBInstanceName(v string) *UpgradePxfsVersionRequest {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradePxfsVersionRequest) SetMinorVersion(v string) *UpgradePxfsVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradePxfsVersionRequest) SetRegionId(v string) *UpgradePxfsVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradePxfsVersionRequest) Validate() error {
	return dara.Validate(s)
}
