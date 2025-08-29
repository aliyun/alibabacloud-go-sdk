// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetABMetricGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetABMetricIds(v string) *GetABMetricGroupResponseBody
	GetABMetricIds() *string
	SetABMetricNames(v string) *GetABMetricGroupResponseBody
	GetABMetricNames() *string
	SetDescription(v string) *GetABMetricGroupResponseBody
	GetDescription() *string
	SetName(v string) *GetABMetricGroupResponseBody
	GetName() *string
	SetOwner(v string) *GetABMetricGroupResponseBody
	GetOwner() *string
	SetRealtime(v bool) *GetABMetricGroupResponseBody
	GetRealtime() *bool
	SetRequestId(v string) *GetABMetricGroupResponseBody
	GetRequestId() *string
	SetSceneId(v string) *GetABMetricGroupResponseBody
	GetSceneId() *string
}

type GetABMetricGroupResponseBody struct {
	// example:
	//
	// 1,2
	ABMetricIds *string `json:"ABMetricIds,omitempty" xml:"ABMetricIds,omitempty"`
	// example:
	//
	// pv,uv
	ABMetricNames *string `json:"ABMetricNames,omitempty" xml:"ABMetricNames,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// visits
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2799614***
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// false
	Realtime *bool `json:"Realtime,omitempty" xml:"Realtime,omitempty"`
	// example:
	//
	// 01D22D08-BA20-5F35-8302-99115F288220
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetABMetricGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetABMetricGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetABMetricGroupResponseBody) GetABMetricIds() *string {
	return s.ABMetricIds
}

func (s *GetABMetricGroupResponseBody) GetABMetricNames() *string {
	return s.ABMetricNames
}

func (s *GetABMetricGroupResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetABMetricGroupResponseBody) GetName() *string {
	return s.Name
}

func (s *GetABMetricGroupResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetABMetricGroupResponseBody) GetRealtime() *bool {
	return s.Realtime
}

func (s *GetABMetricGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetABMetricGroupResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetABMetricGroupResponseBody) SetABMetricIds(v string) *GetABMetricGroupResponseBody {
	s.ABMetricIds = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetABMetricNames(v string) *GetABMetricGroupResponseBody {
	s.ABMetricNames = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetDescription(v string) *GetABMetricGroupResponseBody {
	s.Description = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetName(v string) *GetABMetricGroupResponseBody {
	s.Name = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetOwner(v string) *GetABMetricGroupResponseBody {
	s.Owner = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetRealtime(v bool) *GetABMetricGroupResponseBody {
	s.Realtime = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetRequestId(v string) *GetABMetricGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetABMetricGroupResponseBody) SetSceneId(v string) *GetABMetricGroupResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetABMetricGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
