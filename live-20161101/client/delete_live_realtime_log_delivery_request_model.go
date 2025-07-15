// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DeleteLiveRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetLogstore(v string) *DeleteLiveRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetOwnerId(v int64) *DeleteLiveRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetProject(v string) *DeleteLiveRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *DeleteLiveRealtimeLogDeliveryRequest
	GetRegion() *string
	SetRegionId(v string) *DeleteLiveRealtimeLogDeliveryRequest
	GetRegionId() *string
}

type DeleteLiveRealtimeLogDeliveryRequest struct {
	// The streaming domain.
	//
	// Separate multiple streaming domains with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com,example.aliyundoc.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The name of the Logstore to which log entries are delivered.
	//
	// This parameter is required.
	//
	// example:
	//
	// logstore_example
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Log Service project that is used for real-time log delivery.
	//
	// This parameter is required.
	//
	// example:
	//
	// project_example
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

func (s DeleteLiveRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetLogstore(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetOwnerId(v int64) *DeleteLiveRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetProject(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetRegion(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) SetRegionId(v string) *DeleteLiveRealtimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
