// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveRealTimeLogLogstoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogstore(v string) *DeleteLiveRealTimeLogLogstoreRequest
	GetLogstore() *string
	SetOwnerId(v int64) *DeleteLiveRealTimeLogLogstoreRequest
	GetOwnerId() *int64
	SetProject(v string) *DeleteLiveRealTimeLogLogstoreRequest
	GetProject() *string
	SetRegion(v string) *DeleteLiveRealTimeLogLogstoreRequest
	GetRegion() *string
	SetRegionId(v string) *DeleteLiveRealTimeLogLogstoreRequest
	GetRegionId() *string
}

type DeleteLiveRealTimeLogLogstoreRequest struct {
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

func (s DeleteLiveRealTimeLogLogstoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveRealTimeLogLogstoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) GetLogstore() *string {
	return s.Logstore
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) GetProject() *string {
	return s.Project
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetLogstore(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Logstore = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetOwnerId(v int64) *DeleteLiveRealTimeLogLogstoreRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetProject(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Project = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetRegion(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.Region = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) SetRegionId(v string) *DeleteLiveRealTimeLogLogstoreRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveRealTimeLogLogstoreRequest) Validate() error {
	return dara.Validate(s)
}
