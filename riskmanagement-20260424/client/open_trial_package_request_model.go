// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTrialPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCloseSwitch(v int32) *OpenTrialPackageRequest
	GetAutoCloseSwitch() *int32
	SetRegionId(v string) *OpenTrialPackageRequest
	GetRegionId() *string
}

type OpenTrialPackageRequest struct {
	// example:
	//
	// 0
	AutoCloseSwitch *int32 `json:"AutoCloseSwitch,omitempty" xml:"AutoCloseSwitch,omitempty"`
	// example:
	//
	// cn-guangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenTrialPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenTrialPackageRequest) GoString() string {
	return s.String()
}

func (s *OpenTrialPackageRequest) GetAutoCloseSwitch() *int32 {
	return s.AutoCloseSwitch
}

func (s *OpenTrialPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenTrialPackageRequest) SetAutoCloseSwitch(v int32) *OpenTrialPackageRequest {
	s.AutoCloseSwitch = &v
	return s
}

func (s *OpenTrialPackageRequest) SetRegionId(v string) *OpenTrialPackageRequest {
	s.RegionId = &v
	return s
}

func (s *OpenTrialPackageRequest) Validate() error {
	return dara.Validate(s)
}
