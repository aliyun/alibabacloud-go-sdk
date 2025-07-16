// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DescribeDomainRealtimeLogDeliveryResponseBody
	GetLogstore() *string
	SetProject(v string) *DescribeDomainRealtimeLogDeliveryResponseBody
	GetProject() *string
	SetRegion(v string) *DescribeDomainRealtimeLogDeliveryResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeDomainRealtimeLogDeliveryResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeDomainRealtimeLogDeliveryResponseBody
	GetStatus() *string
}

type DescribeDomainRealtimeLogDeliveryResponseBody struct {
	// The name of the Logstore where log entries are stored.
	//
	// example:
	//
	// LogstoreName
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// ProjectName
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F8F3852-912F-42AC-80EB-F1CF4284DE93
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of real-time log delivery. Valid values:
	//
	// 	- **online**
	//
	// 	- **offline**
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDomainRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) GetProject() *string {
	return s.Project
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetLogstore(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetProject(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetRegion(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) SetStatus(v string) *DescribeDomainRealtimeLogDeliveryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDomainRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
