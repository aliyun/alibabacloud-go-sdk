// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneRunningDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlanId(v string) *GetPtsSceneRunningDataRequest
	GetPlanId() *string
	SetSceneId(v string) *GetPtsSceneRunningDataRequest
	GetSceneId() *string
}

type GetPtsSceneRunningDataRequest struct {
	// The ID of the stress testing task. You can obtain the task ID by calling the StartPtsScene operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// NHBGVF8
	PlanId *string `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// NKKI6GB
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsSceneRunningDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneRunningDataRequest) GoString() string {
	return s.String()
}

func (s *GetPtsSceneRunningDataRequest) GetPlanId() *string {
	return s.PlanId
}

func (s *GetPtsSceneRunningDataRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsSceneRunningDataRequest) SetPlanId(v string) *GetPtsSceneRunningDataRequest {
	s.PlanId = &v
	return s
}

func (s *GetPtsSceneRunningDataRequest) SetSceneId(v string) *GetPtsSceneRunningDataRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsSceneRunningDataRequest) Validate() error {
	return dara.Validate(s)
}
