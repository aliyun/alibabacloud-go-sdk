// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWaypointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListWaypointsRequest
	GetClusterId() *string
	SetContinue(v string) *ListWaypointsRequest
	GetContinue() *string
	SetLimit(v int64) *ListWaypointsRequest
	GetLimit() *int64
	SetName(v string) *ListWaypointsRequest
	GetName() *string
	SetNamespace(v string) *ListWaypointsRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *ListWaypointsRequest
	GetServiceMeshId() *string
}

type ListWaypointsRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Used in conjunction with Limit, it is used to query waypoints starting from a specified offset. When calling for the first time, no need to care abount it. If the Continue in the returned result is not empty, it means that there is still data that has not been returned, and you need to continue the call with the returned Continue.
	//
	// example:
	//
	// eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6MjY4Njc5Miwic3RhcnQiOiJkZWZhdWx0L2Jvb2tpbmZvLXByb2R1Y3RwYWdlXHUwMDAwIn0
	Continue *string `json:"Continue,omitempty" xml:"Continue,omitempty"`
	// Limit the number of waypoints returned.
	//
	// example:
	//
	// 30
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// Waypoint which you want to get. If empty, return all waypoints.
	//
	// example:
	//
	// bookinfo-productpage
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The Service Mesh (ASM) instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c08ba3fd1e6484b0f8cc1ad8fe10d****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s ListWaypointsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWaypointsRequest) GoString() string {
	return s.String()
}

func (s *ListWaypointsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListWaypointsRequest) GetContinue() *string {
	return s.Continue
}

func (s *ListWaypointsRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *ListWaypointsRequest) GetName() *string {
	return s.Name
}

func (s *ListWaypointsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListWaypointsRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *ListWaypointsRequest) SetClusterId(v string) *ListWaypointsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListWaypointsRequest) SetContinue(v string) *ListWaypointsRequest {
	s.Continue = &v
	return s
}

func (s *ListWaypointsRequest) SetLimit(v int64) *ListWaypointsRequest {
	s.Limit = &v
	return s
}

func (s *ListWaypointsRequest) SetName(v string) *ListWaypointsRequest {
	s.Name = &v
	return s
}

func (s *ListWaypointsRequest) SetNamespace(v string) *ListWaypointsRequest {
	s.Namespace = &v
	return s
}

func (s *ListWaypointsRequest) SetServiceMeshId(v string) *ListWaypointsRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *ListWaypointsRequest) Validate() error {
	return dara.Validate(s)
}
