// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdcClassRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCDCNodeCount(v string) *ModifyCdcClassRequest
	GetCDCNodeCount() *string
	SetCdcClass(v string) *ModifyCdcClassRequest
	GetCdcClass() *string
	SetInstanceName(v string) *ModifyCdcClassRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyCdcClassRequest
	GetRegionId() *string
	SetSwitchMode(v string) *ModifyCdcClassRequest
	GetSwitchMode() *string
}

type ModifyCdcClassRequest struct {
	// example:
	//
	// 1
	CDCNodeCount *string `json:"CDCNodeCount,omitempty" xml:"CDCNodeCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// premium
	CdcClass *string `json:"CdcClass,omitempty" xml:"CdcClass,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pxc-hzrh51fze****pon-cdc
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 0
	SwitchMode *string `json:"SwitchMode,omitempty" xml:"SwitchMode,omitempty"`
}

func (s ModifyCdcClassRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdcClassRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdcClassRequest) GetCDCNodeCount() *string {
	return s.CDCNodeCount
}

func (s *ModifyCdcClassRequest) GetCdcClass() *string {
	return s.CdcClass
}

func (s *ModifyCdcClassRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyCdcClassRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCdcClassRequest) GetSwitchMode() *string {
	return s.SwitchMode
}

func (s *ModifyCdcClassRequest) SetCDCNodeCount(v string) *ModifyCdcClassRequest {
	s.CDCNodeCount = &v
	return s
}

func (s *ModifyCdcClassRequest) SetCdcClass(v string) *ModifyCdcClassRequest {
	s.CdcClass = &v
	return s
}

func (s *ModifyCdcClassRequest) SetInstanceName(v string) *ModifyCdcClassRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyCdcClassRequest) SetRegionId(v string) *ModifyCdcClassRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCdcClassRequest) SetSwitchMode(v string) *ModifyCdcClassRequest {
	s.SwitchMode = &v
	return s
}

func (s *ModifyCdcClassRequest) Validate() error {
	return dara.Validate(s)
}
