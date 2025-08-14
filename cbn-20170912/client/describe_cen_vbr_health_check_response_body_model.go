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
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The health check configuration of the VBR.
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
	return dara.Validate(s)
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
	return dara.Validate(s)
}

type DescribeCenVbrHealthCheckResponseBodyVbrHealthChecksVbrHealthCheck struct {
	// The ID of the CEN instance.
	//
	// example:
	//
	// cen-6hpdgj7ni6pz1k****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The description.
	//
	// example:
	//
	// healthcheck_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time interval at which probe packets are sent during the health check. Unit: seconds.
	//
	// example:
	//
	// 2
	HealthCheckInterval *int32 `json:"HealthCheckInterval,omitempty" xml:"HealthCheckInterval,omitempty"`
	// Indicates whether probing is enabled. Valid values:
	//
	// 	- **true**: Probing is enabled.
	//
	//         If you enable probing, the system does not switch to another route if the detected route is not reachable.
	//
	// 	- **false**: Probing is disabled.
	//
	//           If probing is disabled and a redundant route is specified, the system switches to the redundant route when the detected route is not reachable.
	//
	// example:
	//
	// false
	HealthCheckOnly *bool `json:"HealthCheckOnly,omitempty" xml:"HealthCheckOnly,omitempty"`
	// The source IP address of the health check.
	//
	// example:
	//
	// 172.XX.XX.1
	HealthCheckSourceIp *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	// The destination IP address of the health check.
	//
	// example:
	//
	// 192.XX.XX.1
	HealthCheckTargetIp *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	// The number of probe packets that are sent during the health check.
	//
	// example:
	//
	// 8
	HealthyThreshold *int32 `json:"HealthyThreshold,omitempty" xml:"HealthyThreshold,omitempty"`
	// The VBR ID.
	//
	// example:
	//
	// vbr-bp1kznorjeembsuhl****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
	// The ID of the region where the VBR is deployed.
	//
	// example:
	//
	// cn-hangzhou
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
