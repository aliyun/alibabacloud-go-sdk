// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveRealtimeLogDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyLiveRealtimeLogDeliveryRequest
	GetDomainName() *string
	SetLogstore(v string) *ModifyLiveRealtimeLogDeliveryRequest
	GetLogstore() *string
	SetOwnerId(v int64) *ModifyLiveRealtimeLogDeliveryRequest
	GetOwnerId() *int64
	SetProject(v string) *ModifyLiveRealtimeLogDeliveryRequest
	GetProject() *string
	SetRegion(v string) *ModifyLiveRealtimeLogDeliveryRequest
	GetRegion() *string
	SetRegionId(v string) *ModifyLiveRealtimeLogDeliveryRequest
	GetRegionId() *string
}

type ModifyLiveRealtimeLogDeliveryRequest struct {
	// The main streaming domain for which you want to modify the configuration of real-time log delivery.
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

func (s ModifyLiveRealtimeLogDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveRealtimeLogDeliveryRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetProject() *string {
	return s.Project
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetRegion() *string {
	return s.Region
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetDomainName(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetLogstore(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Logstore = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetOwnerId(v int64) *ModifyLiveRealtimeLogDeliveryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetProject(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Project = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetRegion(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.Region = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) SetRegionId(v string) *ModifyLiveRealtimeLogDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLiveRealtimeLogDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
