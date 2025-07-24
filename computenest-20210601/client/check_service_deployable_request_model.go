// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceDeployableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPostPaidAmount(v string) *CheckServiceDeployableRequest
	GetPostPaidAmount() *string
	SetPrePaidAmount(v string) *CheckServiceDeployableRequest
	GetPrePaidAmount() *string
	SetRegionId(v string) *CheckServiceDeployableRequest
	GetRegionId() *string
	SetServiceId(v string) *CheckServiceDeployableRequest
	GetServiceId() *string
	SetServiceVersion(v string) *CheckServiceDeployableRequest
	GetServiceVersion() *string
	SetTrialType(v string) *CheckServiceDeployableRequest
	GetTrialType() *string
}

type CheckServiceDeployableRequest struct {
	// Total amount of postpaid.
	//
	// example:
	//
	// 1.29
	PostPaidAmount *string `json:"PostPaidAmount,omitempty" xml:"PostPaidAmount,omitempty"`
	// Total amount of prepayment.
	//
	// example:
	//
	// 0.0
	PrePaidAmount *string `json:"PrePaidAmount,omitempty" xml:"PrePaidAmount,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The service version.
	//
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// The trial type of the service instance. Valid values:
	//
	// 	- **Trial**: Trials are supported.
	//
	// 	- **NotTrial**: Trials are not supported.
	//
	// example:
	//
	// NotTrial
	TrialType *string `json:"TrialType,omitempty" xml:"TrialType,omitempty"`
}

func (s CheckServiceDeployableRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceDeployableRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceDeployableRequest) GetPostPaidAmount() *string {
	return s.PostPaidAmount
}

func (s *CheckServiceDeployableRequest) GetPrePaidAmount() *string {
	return s.PrePaidAmount
}

func (s *CheckServiceDeployableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceDeployableRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *CheckServiceDeployableRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *CheckServiceDeployableRequest) GetTrialType() *string {
	return s.TrialType
}

func (s *CheckServiceDeployableRequest) SetPostPaidAmount(v string) *CheckServiceDeployableRequest {
	s.PostPaidAmount = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetPrePaidAmount(v string) *CheckServiceDeployableRequest {
	s.PrePaidAmount = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetRegionId(v string) *CheckServiceDeployableRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetServiceId(v string) *CheckServiceDeployableRequest {
	s.ServiceId = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetServiceVersion(v string) *CheckServiceDeployableRequest {
	s.ServiceVersion = &v
	return s
}

func (s *CheckServiceDeployableRequest) SetTrialType(v string) *CheckServiceDeployableRequest {
	s.TrialType = &v
	return s
}

func (s *CheckServiceDeployableRequest) Validate() error {
	return dara.Validate(s)
}
