// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWaypointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteWaypointRequest
	GetClusterId() *string
	SetName(v string) *DeleteWaypointRequest
	GetName() *string
	SetNamespace(v string) *DeleteWaypointRequest
	GetNamespace() *string
	SetServiceMeshId(v string) *DeleteWaypointRequest
	GetServiceMeshId() *string
}

type DeleteWaypointRequest struct {
	// The ID of the cluster on the data plane.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce3c25e247da24f3aab9b7edfae83****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Waypoint name.
	//
	// This parameter is required.
	//
	// example:
	//
	// bookinfo-reviews
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// This parameter is required.
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
	// cb8963379255149cb98c8686f274x****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s DeleteWaypointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWaypointRequest) GoString() string {
	return s.String()
}

func (s *DeleteWaypointRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteWaypointRequest) GetName() *string {
	return s.Name
}

func (s *DeleteWaypointRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteWaypointRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *DeleteWaypointRequest) SetClusterId(v string) *DeleteWaypointRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteWaypointRequest) SetName(v string) *DeleteWaypointRequest {
	s.Name = &v
	return s
}

func (s *DeleteWaypointRequest) SetNamespace(v string) *DeleteWaypointRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteWaypointRequest) SetServiceMeshId(v string) *DeleteWaypointRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *DeleteWaypointRequest) Validate() error {
	return dara.Validate(s)
}
