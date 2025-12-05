// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *GetPtsReportDetailsRequest
	GetPlanId() *string
	SetSceneId(v string) *GetPtsReportDetailsRequest
	GetSceneId() *string
}

type GetPtsReportDetailsRequest struct {
	// The ID of the performance testing task. A task ID is generated each time a PTS scenario is started.
	//
	// This parameter is required.
	//
	// example:
	//
	// OH5HA3VB
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The ID of the scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// G5HCVS
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsReportDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportDetailsRequest) GoString() string {
	return s.String()
}

func (s *GetPtsReportDetailsRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetPtsReportDetailsRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsReportDetailsRequest) SetPlanId(v string) *GetPtsReportDetailsRequest {
	s.PlanId = &v
	return s
}

func (s *GetPtsReportDetailsRequest) SetSceneId(v string) *GetPtsReportDetailsRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsReportDetailsRequest) Validate() error {
	return dara.Validate(s)
}
