// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABMetricGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricIds(v string) *UpdateABMetricGroupRequest
	GetABMetricIds() *string
	SetDescription(v string) *UpdateABMetricGroupRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateABMetricGroupRequest
	GetInstanceId() *string
	SetName(v string) *UpdateABMetricGroupRequest
	GetName() *string
	SetRealtime(v bool) *UpdateABMetricGroupRequest
	GetRealtime() *bool
	SetSceneId(v string) *UpdateABMetricGroupRequest
	GetSceneId() *string
}

type UpdateABMetricGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// This parameter is required.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s UpdateABMetricGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateABMetricGroupRequest) GetABMetricIds() *string {
	return s.ABMetricIds
}

func (s *UpdateABMetricGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateABMetricGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateABMetricGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateABMetricGroupRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *UpdateABMetricGroupRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *UpdateABMetricGroupRequest) SetABMetricIds(v string) *UpdateABMetricGroupRequest {
	s.ABMetricIds = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetDescription(v string) *UpdateABMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetInstanceId(v string) *UpdateABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetName(v string) *UpdateABMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetRealtime(v bool) *UpdateABMetricGroupRequest {
	s.Realtime = &v
	return s
}

func (s *UpdateABMetricGroupRequest) SetSceneId(v string) *UpdateABMetricGroupRequest {
	s.SceneId = &v
	return s
}

func (s *UpdateABMetricGroupRequest) Validate() error {
	return dara.Validate(s)
}
