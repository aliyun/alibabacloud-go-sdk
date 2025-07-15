// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveRealTimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CreateLiveRealTimeLogDeliveryRequest
	GetDomainName() *string
	SetLogstore(v string) *CreateLiveRealTimeLogDeliveryRequest
	GetLogstore() *string
	SetOwnerId(v int64) *CreateLiveRealTimeLogDeliveryRequest
	GetOwnerId() *int64
	SetProject(v string) *CreateLiveRealTimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *CreateLiveRealTimeLogDeliveryRequest
	GetRegion() *string
	SetRegionId(v string) *CreateLiveRealTimeLogDeliveryRequest
	GetRegionId() *string
}

type CreateLiveRealTimeLogDeliveryRequest struct {
	// The streaming domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the Logstore to which log entries are delivered.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_logstore
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The ID of the region where the Log Service project is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateLiveRealTimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveRealTimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *CreateLiveRealTimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetDomainName(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetLogstore(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetOwnerId(v int64) *CreateLiveRealTimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetProject(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetRegion(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) SetRegionId(v string) *CreateLiveRealTimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLiveRealTimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
