// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetLayerResponseBody
	GetDescription() *string
	SetGmtCreateTime(v string) *GetLayerResponseBody
	GetGmtCreateTime() *string
	SetLaboratoryId(v string) *GetLayerResponseBody
	GetLaboratoryId() *string
	SetName(v string) *GetLayerResponseBody
	GetName() *string
	SetRequestId(v string) *GetLayerResponseBody
	GetRequestId() *string
	SetResidualFlow(v int64) *GetLayerResponseBody
	GetResidualFlow() *int64
	SetSceneId(v string) *GetLayerResponseBody
	GetSceneId() *string
}

type GetLayerResponseBody struct {
	// The description of the layer.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the experiment layer was created.
	//
	// example:
	//
	// 2022-10-13 17:34:52.0
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The ID of the laboratory.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// The name of the layer.
	//
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EE97D06A-2AA0-5AD9-B6CF-8A267924D691
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The residual traffic in the experiment layer.
	//
	// example:
	//
	// 10
	ResidualFlow *int64 `json:"ResidualFlow,omitempty" xml:"ResidualFlow,omitempty"`
	// The ID of the scene.
	//
	// example:
	//
	// 4
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetLayerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLayerResponseBody) GoString() string {
	return s.String()
}

func (s *GetLayerResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetLayerResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetLayerResponseBody) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *GetLayerResponseBody) GetName() *string {
	return s.Name
}

func (s *GetLayerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLayerResponseBody) GetResidualFlow() *int64 {
	return s.ResidualFlow
}

func (s *GetLayerResponseBody) GetSceneId() *string {
	return s.SceneId
}

func (s *GetLayerResponseBody) SetDescription(v string) *GetLayerResponseBody {
	s.Description = &v
	return s
}

func (s *GetLayerResponseBody) SetGmtCreateTime(v string) *GetLayerResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetLayerResponseBody) SetLaboratoryId(v string) *GetLayerResponseBody {
	s.LaboratoryId = &v
	return s
}

func (s *GetLayerResponseBody) SetName(v string) *GetLayerResponseBody {
	s.Name = &v
	return s
}

func (s *GetLayerResponseBody) SetRequestId(v string) *GetLayerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLayerResponseBody) SetResidualFlow(v int64) *GetLayerResponseBody {
	s.ResidualFlow = &v
	return s
}

func (s *GetLayerResponseBody) SetSceneId(v string) *GetLayerResponseBody {
	s.SceneId = &v
	return s
}

func (s *GetLayerResponseBody) Validate() error {
	return dara.Validate(s)
}
