// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLayersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLayers(v []*ListLayersResponseBodyLayers) *ListLayersResponseBody
	GetLayers() []*ListLayersResponseBodyLayers
	SetRequestId(v string) *ListLayersResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListLayersResponseBody
	GetTotalCount() *int64
}

type ListLayersResponseBody struct {
	Layers []*ListLayersResponseBodyLayers `json:"Layers,omitempty" xml:"Layers,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 518C64F6-DFF7-11ED-85B0-00163E14B3D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLayersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLayersResponseBody) GoString() string {
	return s.String()
}

func (s *ListLayersResponseBody) GetLayers() []*ListLayersResponseBodyLayers {
	return s.Layers
}

func (s *ListLayersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLayersResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLayersResponseBody) SetLayers(v []*ListLayersResponseBodyLayers) *ListLayersResponseBody {
	s.Layers = v
	return s
}

func (s *ListLayersResponseBody) SetRequestId(v string) *ListLayersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLayersResponseBody) SetTotalCount(v int64) *ListLayersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLayersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLayersResponseBodyLayers struct {
	// example:
	//
	// This is a test.
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
	// example:
	//
	// layer1
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ResidualFlow *int64  `json:"ResidualFlow,omitempty" xml:"ResidualFlow,omitempty"`
	// example:
	//
	// 3
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s ListLayersResponseBodyLayers) String() string {
	return dara.Prettify(s)
}

func (s ListLayersResponseBodyLayers) GoString() string {
	return s.String()
}

func (s *ListLayersResponseBodyLayers) GetDescription() *string {
	return s.Description
}

func (s *ListLayersResponseBodyLayers) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListLayersResponseBodyLayers) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *ListLayersResponseBodyLayers) GetLayerId() *string {
	return s.LayerId
}

func (s *ListLayersResponseBodyLayers) GetName() *string {
	return s.Name
}

func (s *ListLayersResponseBodyLayers) GetResidualFlow() *int64 {
	return s.ResidualFlow
}

func (s *ListLayersResponseBodyLayers) GetSceneId() *string {
	return s.SceneId
}

func (s *ListLayersResponseBodyLayers) SetDescription(v string) *ListLayersResponseBodyLayers {
	s.Description = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetGmtCreateTime(v string) *ListLayersResponseBodyLayers {
	s.GmtCreateTime = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetLaboratoryId(v string) *ListLayersResponseBodyLayers {
	s.LaboratoryId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetLayerId(v string) *ListLayersResponseBodyLayers {
	s.LayerId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetName(v string) *ListLayersResponseBodyLayers {
	s.Name = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetResidualFlow(v int64) *ListLayersResponseBodyLayers {
	s.ResidualFlow = &v
	return s
}

func (s *ListLayersResponseBodyLayers) SetSceneId(v string) *ListLayersResponseBodyLayers {
	s.SceneId = &v
	return s
}

func (s *ListLayersResponseBodyLayers) Validate() error {
	return dara.Validate(s)
}
