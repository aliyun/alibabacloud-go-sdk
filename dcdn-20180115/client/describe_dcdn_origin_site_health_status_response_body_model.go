// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnOriginSiteHealthStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOriginSiteStatus(v []*DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) *DescribeDcdnOriginSiteHealthStatusResponseBody
	GetOriginSiteStatus() []*DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus
	SetRequestId(v string) *DescribeDcdnOriginSiteHealthStatusResponseBody
	GetRequestId() *string
}

type DescribeDcdnOriginSiteHealthStatusResponseBody struct {
	// The information about the origin server of the accelerated domain name.
	OriginSiteStatus []*DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus `json:"OriginSiteStatus,omitempty" xml:"OriginSiteStatus,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnOriginSiteHealthStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnOriginSiteHealthStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBody) GetOriginSiteStatus() []*DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus {
	return s.OriginSiteStatus
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBody) SetOriginSiteStatus(v []*DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) *DescribeDcdnOriginSiteHealthStatusResponseBody {
	s.OriginSiteStatus = v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBody) SetRequestId(v string) *DescribeDcdnOriginSiteHealthStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus struct {
	// The health status of the origin server. Each point of presence (POP) periodically initiates a probe request to the configured origin domain name. If the POP receives a response from the origin server in 5 seconds, the probe is considered successful. After the probe data for each POP is collected, the health status of the origin server is calculated based on the proportion of successful probes. Valid values:
	//
	// 	- unknown: The probe data of the origin server is not obtained because the configurations of the origin server have been changed recently. Try again later.
	//
	// 	- healthy: The proportion of successful probes is higher than 80%.
	//
	// 	- degraded: The proportion of successful probes is higher than 0% and lower than or equal to 80%.
	//
	// 	- critical: All probing requests to the origin server failed.
	//
	// example:
	//
	// healthy
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The origin domain name that you configured in the DCDN console, which can be an IPv4 address, IPv6 address, common domain name, or Object Storage Service (OSS) domain name.
	//
	// example:
	//
	// host.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
}

func (s DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) GoString() string {
	return s.String()
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) GetHost() *string {
	return s.Host
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) SetHealthStatus(v string) *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus {
	s.HealthStatus = &v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) SetHost(v string) *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus {
	s.Host = &v
	return s
}

func (s *DescribeDcdnOriginSiteHealthStatusResponseBodyOriginSiteStatus) Validate() error {
	return dara.Validate(s)
}
