// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateEdgeContainerAppResourceReserveRequest
	GetAppId() *string
	SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveRequest
	GetDurationTime() *string
	SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveRequest
	GetEnable() *bool
	SetForever(v bool) *UpdateEdgeContainerAppResourceReserveRequest
	GetForever() *bool
	SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveRequestReserveSet) *UpdateEdgeContainerAppResourceReserveRequest
	GetReserveSet() []*UpdateEdgeContainerAppResourceReserveRequestReserveSet
}

type UpdateEdgeContainerAppResourceReserveRequest struct {
	// example:
	//
	// app-88068867578379****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05Z
	DurationTime *string `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// example:
	//
	// true
	Forever    *bool                                                     `json:"Forever,omitempty" xml:"Forever,omitempty"`
	ReserveSet []*UpdateEdgeContainerAppResourceReserveRequestReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s UpdateEdgeContainerAppResourceReserveRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveRequest) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetDurationTime() *string {
	return s.DurationTime
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetForever() *bool {
	return s.Forever
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) GetReserveSet() []*UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	return s.ReserveSet
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetAppId(v string) *UpdateEdgeContainerAppResourceReserveRequest {
	s.AppId = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveRequest {
	s.DurationTime = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveRequest {
	s.Enable = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetForever(v bool) *UpdateEdgeContainerAppResourceReserveRequest {
	s.Forever = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveRequestReserveSet) *UpdateEdgeContainerAppResourceReserveRequest {
	s.ReserveSet = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateEdgeContainerAppResourceReserveRequestReserveSet struct {
	// example:
	//
	// cmcc
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// huazhong
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s UpdateEdgeContainerAppResourceReserveRequestReserveSet) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveRequestReserveSet) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetRegion() *string {
	return s.Region
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetIsp(v string) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Isp = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetRegion(v string) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Region = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) SetReplicas(v int32) *UpdateEdgeContainerAppResourceReserveRequestReserveSet {
	s.Replicas = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveRequestReserveSet) Validate() error {
	return dara.Validate(s)
}
