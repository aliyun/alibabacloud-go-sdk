// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnoseReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateDiagnoseReportRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *CreateDiagnoseReportRequest
	GetScalingGroupId() *string
}

type CreateDiagnoseReportRequest struct {
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1f2f6oxc2*******
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s CreateDiagnoseReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnoseReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnoseReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiagnoseReportRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *CreateDiagnoseReportRequest) SetRegionId(v string) *CreateDiagnoseReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiagnoseReportRequest) SetScalingGroupId(v string) *CreateDiagnoseReportRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *CreateDiagnoseReportRequest) Validate() error {
	return dara.Validate(s)
}
