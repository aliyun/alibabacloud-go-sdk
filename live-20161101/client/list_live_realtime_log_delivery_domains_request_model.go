// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLiveRealtimeLogDeliveryDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *ListLiveRealtimeLogDeliveryDomainsRequest
	GetLogstore() *string
	SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryDomainsRequest
	GetOwnerId() *int64
	SetProject(v string) *ListLiveRealtimeLogDeliveryDomainsRequest
	GetProject() *string
	SetRegion(v string) *ListLiveRealtimeLogDeliveryDomainsRequest
	GetRegion() *string
	SetRegionId(v string) *ListLiveRealtimeLogDeliveryDomainsRequest
	GetRegionId() *string
}

type ListLiveRealtimeLogDeliveryDomainsRequest struct {
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

func (s ListLiveRealtimeLogDeliveryDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLiveRealtimeLogDeliveryDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) GetProject() *string {
	return s.Project
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetLogstore(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Logstore = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetOwnerId(v int64) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetProject(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Project = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetRegion(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.Region = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) SetRegionId(v string) *ListLiveRealtimeLogDeliveryDomainsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLiveRealtimeLogDeliveryDomainsRequest) Validate() error {
	return dara.Validate(s)
}
