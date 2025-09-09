// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeHiStoreInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *UpgradeHiStoreInstanceRequest
	GetDrdsInstanceId() *string
	SetHistoreInstanceId(v string) *UpgradeHiStoreInstanceRequest
	GetHistoreInstanceId() *string
	SetRegionId(v string) *UpgradeHiStoreInstanceRequest
	GetRegionId() *string
}

type UpgradeHiStoreInstanceRequest struct {
	// The ID of the PolarDB-X 1.0 instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdssad23sdfc
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The ID of the column-oriented storage instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// hi-sesex2e
	HistoreInstanceId *string `json:"HistoreInstanceId,omitempty" xml:"HistoreInstanceId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeHiStoreInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeHiStoreInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeHiStoreInstanceRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *UpgradeHiStoreInstanceRequest) GetHistoreInstanceId() *string {
	return s.HistoreInstanceId
}

func (s *UpgradeHiStoreInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpgradeHiStoreInstanceRequest) SetDrdsInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetHistoreInstanceId(v string) *UpgradeHiStoreInstanceRequest {
	s.HistoreInstanceId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) SetRegionId(v string) *UpgradeHiStoreInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeHiStoreInstanceRequest) Validate() error {
	return dara.Validate(s)
}
