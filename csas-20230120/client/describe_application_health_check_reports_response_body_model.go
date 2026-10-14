// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationHealthCheckReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReports(v []*DescribeApplicationHealthCheckReportsResponseBodyReports) *DescribeApplicationHealthCheckReportsResponseBody
	GetReports() []*DescribeApplicationHealthCheckReportsResponseBodyReports
	SetRequestId(v string) *DescribeApplicationHealthCheckReportsResponseBody
	GetRequestId() *string
}

type DescribeApplicationHealthCheckReportsResponseBody struct {
	// The list of health check reports.
	Reports []*DescribeApplicationHealthCheckReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 748CFDC7-1EB6-5B8B-9405-DA76ED5BB60D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeApplicationHealthCheckReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationHealthCheckReportsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationHealthCheckReportsResponseBody) GetReports() []*DescribeApplicationHealthCheckReportsResponseBodyReports {
	return s.Reports
}

func (s *DescribeApplicationHealthCheckReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationHealthCheckReportsResponseBody) SetReports(v []*DescribeApplicationHealthCheckReportsResponseBodyReports) *DescribeApplicationHealthCheckReportsResponseBody {
	s.Reports = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBody) SetRequestId(v string) *DescribeApplicationHealthCheckReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBody) Validate() error {
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationHealthCheckReportsResponseBodyReports struct {
	// The ID of the internal-facing application.
	//
	// example:
	//
	// pa-application-e12860ef6c48****
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The list of addresses where health checks failed.
	FailedAddresses []*string `json:"FailedAddresses,omitempty" xml:"FailedAddresses,omitempty" type:"Repeated"`
	// The list of regions where health checks failed.
	FailedRegions []*string `json:"FailedRegions,omitempty" xml:"FailedRegions,omitempty" type:"Repeated"`
	// The number of failed health checks within the query time range.
	//
	// example:
	//
	// 2
	FailureCount *int64 `json:"FailureCount,omitempty" xml:"FailureCount,omitempty"`
	// The list of reasons for health check failures.
	FailureReasons []*string `json:"FailureReasons,omitempty" xml:"FailureReasons,omitempty" type:"Repeated"`
	// The status calculated based on the latest check result. Valid values:
	//
	// 	- normal: The success rate of the latest check is greater than or equal to 95%.
	//
	// 	- partialError: The success rate of the latest check is greater than 0 and less than 95%.
	//
	// 	- error: The success rate of the latest check is 0.
	//
	// 	- uncompleted: No check results are available.
	//
	// example:
	//
	// normal
	HealthCheckStatus *string `json:"HealthCheckStatus,omitempty" xml:"HealthCheckStatus,omitempty"`
	// The health check mode. Valid values:
	//
	// 	- OnAccess: on-access check
	//
	// 	- Continuous: continuous check
	//
	// example:
	//
	// OnAccess
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The number of successful health checks within the query time range.
	//
	// example:
	//
	// 100
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s DescribeApplicationHealthCheckReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationHealthCheckReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetFailedAddresses() []*string {
	return s.FailedAddresses
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetFailedRegions() []*string {
	return s.FailedRegions
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetFailureCount() *int64 {
	return s.FailureCount
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetFailureReasons() []*string {
	return s.FailureReasons
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetHealthCheckStatus() *string {
	return s.HealthCheckStatus
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetMode() *string {
	return s.Mode
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetApplicationId(v string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetFailedAddresses(v []*string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.FailedAddresses = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetFailedRegions(v []*string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.FailedRegions = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetFailureCount(v int64) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.FailureCount = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetFailureReasons(v []*string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.FailureReasons = v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetHealthCheckStatus(v string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.HealthCheckStatus = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetMode(v string) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.Mode = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) SetSuccessCount(v int64) *DescribeApplicationHealthCheckReportsResponseBodyReports {
	s.SuccessCount = &v
	return s
}

func (s *DescribeApplicationHealthCheckReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}
