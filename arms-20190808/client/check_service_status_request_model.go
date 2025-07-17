// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckServiceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckServiceStatusRequest
	GetRegionId() *string
	SetSvcCode(v string) *CheckServiceStatusRequest
	GetSvcCode() *string
}

type CheckServiceStatusRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service code of an Alibaba Cloud service. The service code of Managed Service for Prometheus is prometheus.
	//
	// This parameter is required.
	//
	// example:
	//
	// prometheus
	SvcCode *string `json:"SvcCode,omitempty" xml:"SvcCode,omitempty"`
}

func (s CheckServiceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckServiceStatusRequest) GetSvcCode() *string {
	return s.SvcCode
}

func (s *CheckServiceStatusRequest) SetRegionId(v string) *CheckServiceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckServiceStatusRequest) SetSvcCode(v string) *CheckServiceStatusRequest {
	s.SvcCode = &v
	return s
}

func (s *CheckServiceStatusRequest) Validate() error {
	return dara.Validate(s)
}
