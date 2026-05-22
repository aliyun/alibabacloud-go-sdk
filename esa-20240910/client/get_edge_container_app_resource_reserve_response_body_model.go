// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceReserveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDurationTime(v string) *GetEdgeContainerAppResourceReserveResponseBody
	GetDurationTime() *string
	SetEnable(v bool) *GetEdgeContainerAppResourceReserveResponseBody
	GetEnable() *bool
	SetForever(v bool) *GetEdgeContainerAppResourceReserveResponseBody
	GetForever() *bool
	SetRequestId(v string) *GetEdgeContainerAppResourceReserveResponseBody
	GetRequestId() *string
	SetReserveSet(v []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet) *GetEdgeContainerAppResourceReserveResponseBody
	GetReserveSet() []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet
}

type GetEdgeContainerAppResourceReserveResponseBody struct {
	DurationTime *string                                                     `json:"DurationTime,omitempty" xml:"DurationTime,omitempty"`
	Enable       *bool                                                       `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Forever      *bool                                                       `json:"Forever,omitempty" xml:"Forever,omitempty"`
	RequestId    *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ReserveSet   []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet `json:"ReserveSet,omitempty" xml:"ReserveSet,omitempty" type:"Repeated"`
}

func (s GetEdgeContainerAppResourceReserveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetDurationTime() *string {
	return s.DurationTime
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetForever() *bool {
	return s.Forever
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) GetReserveSet() []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	return s.ReserveSet
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetDurationTime(v string) *GetEdgeContainerAppResourceReserveResponseBody {
	s.DurationTime = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetEnable(v bool) *GetEdgeContainerAppResourceReserveResponseBody {
	s.Enable = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetForever(v bool) *GetEdgeContainerAppResourceReserveResponseBody {
	s.Forever = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetRequestId(v string) *GetEdgeContainerAppResourceReserveResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) SetReserveSet(v []*GetEdgeContainerAppResourceReserveResponseBodyReserveSet) *GetEdgeContainerAppResourceReserveResponseBody {
	s.ReserveSet = v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBody) Validate() error {
	if s.ReserveSet != nil {
		for _, item := range s.ReserveSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEdgeContainerAppResourceReserveResponseBodyReserveSet struct {
	Isp      *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	Region   *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Replicas *int32  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s GetEdgeContainerAppResourceReserveResponseBodyReserveSet) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetIsp() *string {
	return s.Isp
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetRegion() *string {
	return s.Region
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) GetReplicas() *int32 {
	return s.Replicas
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetIsp(v string) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Isp = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetRegion(v string) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Region = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) SetReplicas(v int32) *GetEdgeContainerAppResourceReserveResponseBodyReserveSet {
	s.Replicas = &v
	return s
}

func (s *GetEdgeContainerAppResourceReserveResponseBodyReserveSet) Validate() error {
	return dara.Validate(s)
}
