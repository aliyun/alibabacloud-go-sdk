// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRayClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopRayClusterRequest
	GetInstanceId() *string
}

type StopRayClusterRequest struct {
	// example:
	//
	// ray-k7nm8ahl5te4tg91-ey7blpbg
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s StopRayClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s StopRayClusterRequest) GoString() string {
	return s.String()
}

func (s *StopRayClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopRayClusterRequest) SetInstanceId(v string) *StopRayClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *StopRayClusterRequest) Validate() error {
	return dara.Validate(s)
}
