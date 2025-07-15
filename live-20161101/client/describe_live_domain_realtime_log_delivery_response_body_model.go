// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainRealtimeLogDeliveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody
	GetLogstore() *string
	SetProject(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody
	GetProject() *string
	SetRegion(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody
	GetRegion() *string
	SetRequestId(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody
	GetStatus() *string
}

type DescribeLiveDomainRealtimeLogDeliveryResponseBody struct {
	// The name of the Logstore to which log entries are delivered.
	//
	// example:
	//
	// logstore_example
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// example:
	//
	// project_example
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F8F3852-912F-42AC-80EB-F1CF4284DE93
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of real-time log delivery. Valid values:
	//
	// 	- online: Real-time log delivery is enabled.
	//
	// 	- offline: Real-time log delivery is disabled.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainRealtimeLogDeliveryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) GetLogstore() *string {
	return s.Logstore
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) GetProject() *string {
	return s.Project
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) GetRegion() *string {
	return s.Region
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) SetLogstore(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	s.Logstore = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) SetProject(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	s.Project = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) SetRegion(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	s.Region = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) SetRequestId(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) SetStatus(v string) *DescribeLiveDomainRealtimeLogDeliveryResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainRealtimeLogDeliveryResponseBody) Validate() error {
	return dara.Validate(s)
}
