// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenVbrHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeCenVbrHealthCheckResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenVbrHealthCheckResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCenVbrHealthCheckResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCenVbrHealthCheckResponseBody
	GetTotalCount() *int32
	SetVbrHealthChecks(v *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) *DescribeCenVbrHealthCheckResponseBody
	GetVbrHealthChecks() *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks
}

type DescribeCenVbrHealthCheckResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B8C9702E-304A-4E18-AC89-BE2D91C2C176
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount      *int32                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VbrHealthChecks *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks `json:"VbrHealthChecks,omitempty" xml:"VbrHealthChecks,omitempty" type:"Struct"`
}

func (s DescribeCenVbrHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenVbrHealthCheckResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenVbrHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCenVbrHealthCheckResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCenVbrHealthCheckResponseBody) GetVbrHealthChecks() *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	return s.VbrHealthChecks
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetPageNumber(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetPageSize(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetRequestId(v string) *DescribeCenVbrHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetTotalCount(v int32) *DescribeCenVbrHealthCheckResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) SetVbrHealthChecks(v *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) *DescribeCenVbrHealthCheckResponseBody {
	s.VbrHealthChecks = v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBody) Validate() error {
	if s.VbrHealthChecks != nil {
		if err := s.VbrHealthChecks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks struct {
	VbrHealthCheck []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck `json:"VbrHealthCheck,omitempty" xml:"VbrHealthCheck,omitempty" type:"Repeated"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) GetVbrHealthCheck() []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	return s.VbrHealthCheck
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) SetVbrHealthCheck(v []*DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks {
	s.VbrHealthCheck = v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecks) Validate() error {
	if s.VbrHealthCheck != nil {
		for _, item := range s.VbrHealthCheck {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck struct {
	CenId               *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	HealthCheckInterval *int32  `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	HealthCheckOnly     *bool   `json:"HealthCheckOnly,omitempty" xml:"HealthCheckOnly,omitempty"`
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	HealthyThreshold    *int32  `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	VbrInstanceId       *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	VbrInstanceRegionId *string `json:"VbrInstanceRegionId,omitempty" xml:"VbrInstanceRegionId,omitempty"`
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GoString() string {
	return s.String()
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetCenId() *string {
	return s.CenId
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetDescription() *string {
	return s.Description
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetHealthCheckInterval() *int32 {
	return s.HealthCheckInterval
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetHealthCheckOnly() *bool {
	return s.HealthCheckOnly
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetHealthyThreshold() *int32 {
	return s.HealthyThreshold
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) GetVbrInstanceRegionId() *string {
	return s.VbrInstanceRegionId
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetCenId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.CenId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetDescription(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.Description = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckInterval(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckInterval = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckOnly(v bool) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckOnly = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckSourceIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthCheckTargetIp(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetHealthyThreshold(v int32) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.HealthyThreshold = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetVbrInstanceId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) SetVbrInstanceRegionId(v string) *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck {
	s.VbrInstanceRegionId = &v
	return s
}

func (s *DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck) Validate() error {
	return dara.Validate(s)
}
