// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeInstanceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *UpgradeInstanceVersionRequest
	GetDrdsInstanceId() *string
	SetRegionId(v string) *UpgradeInstanceVersionRequest
	GetRegionId() *string
	SetRpm(v string) *UpgradeInstanceVersionRequest
	GetRpm() *string
}

type UpgradeInstanceVersionRequest struct {
	// The ID of the PolarDB-X 1.0 instance that you want to upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdshbgaen89****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version number of the PolarDB-X 1.0 instance. You can leave this parameter unspecified.
	//
	// example:
	//
	// t-drds-server-5.4.12-16348095.noarch.rpm
	Rpm *string `json:"Rpm,omitempty" xml:"Rpm,omitempty"`
}

func (s UpgradeInstanceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeInstanceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeInstanceVersionRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *UpgradeInstanceVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeInstanceVersionRequest) GetRpm() *string {
	return s.Rpm
}

func (s *UpgradeInstanceVersionRequest) SetDrdsInstanceId(v string) *UpgradeInstanceVersionRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRegionId(v string) *UpgradeInstanceVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) SetRpm(v string) *UpgradeInstanceVersionRequest {
	s.Rpm = &v
	return s
}

func (s *UpgradeInstanceVersionRequest) Validate() error {
	return dara.Validate(s)
}
