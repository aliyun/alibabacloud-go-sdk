// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEdgeContainerAppResourceReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetDurationTime() *string
	SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetEnable() *bool
	SetForever(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetForever() *bool
	SetRequestId(v string) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetRequestId() *string
	SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) *UpdateEdgeContainerAppResourceReserveResponseBody
	GetReserveSet() []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet
}

type UpdateEdgeContainerAppResourceReserveResponseBody struct {
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
	Forever *bool `json:"Forever,omitempty" xml:"Forever,omitempty"`
	// example:
	//
	// 1AB799CF-562A-5CAF-A99E-4354053D814F
	RequestId  *string                                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserveSet []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s UpdateEdgeContainerAppResourceReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetDurationTime() *string {
	return s.DurationTime
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetForever() *bool {
	return s.Forever
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) GetReserveSet() []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	return s.ReserveSet
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetDurationTime(v string) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.DurationTime = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetEnable(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.Enable = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetForever(v bool) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.Forever = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetRequestId(v string) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) SetReserveSet(v []*UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) *UpdateEdgeContainerAppResourceReserveResponseBody {
	s.ReserveSet = v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet struct {
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

func (s UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) String() string {
	return dara.Prettify(s)
}

func (s UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GoString() string {
	return s.String()
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetRegion() *string {
	return s.Region
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetIsp(v string) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Isp = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetRegion(v string) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Region = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) SetReplicas(v int32) *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Replicas = &v
	return s
}

func (s *UpdateEdgeContainerAppResourceReserveResponseBodyReserveSet) Validate() error {
	return dara.Validate(s)
}
