// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABMetricGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricIds(v string) *CreateABMetricGroupRequest
	GetABMetricIds() *string
	SetDescription(v string) *CreateABMetricGroupRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateABMetricGroupRequest
	GetInstanceId() *string
	SetName(v string) *CreateABMetricGroupRequest
	GetName() *string
	SetRealtime(v bool) *CreateABMetricGroupRequest
	GetRealtime() *bool
	SetSceneId(v string) *CreateABMetricGroupRequest
	GetSceneId() *string
}

type CreateABMetricGroupRequest struct {
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

func (s CreateABMetricGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateABMetricGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateABMetricGroupRequest) GetABMetricIds() *string {
	return s.ABMetricIds
}

func (s *CreateABMetricGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateABMetricGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateABMetricGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateABMetricGroupRequest) GetRealtime() *bool {
	return s.Realtime
}

func (s *CreateABMetricGroupRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *CreateABMetricGroupRequest) SetABMetricIds(v string) *CreateABMetricGroupRequest {
	s.ABMetricIds = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetDescription(v string) *CreateABMetricGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetInstanceId(v string) *CreateABMetricGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetName(v string) *CreateABMetricGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetRealtime(v bool) *CreateABMetricGroupRequest {
	s.Realtime = &v
	return s
}

func (s *CreateABMetricGroupRequest) SetSceneId(v string) *CreateABMetricGroupRequest {
	s.SceneId = &v
	return s
}

func (s *CreateABMetricGroupRequest) Validate() error {
	return dara.Validate(s)
}
