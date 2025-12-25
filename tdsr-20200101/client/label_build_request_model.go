// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLabelBuildRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *LabelBuildRequest
	GetMode() *string
	SetModelStyle(v string) *LabelBuildRequest
	GetModelStyle() *string
	SetOptimizeWallWidth(v string) *LabelBuildRequest
	GetOptimizeWallWidth() *string
	SetPlanStyle(v string) *LabelBuildRequest
	GetPlanStyle() *string
	SetSceneId(v string) *LabelBuildRequest
	GetSceneId() *string
	SetWallHeight(v int64) *LabelBuildRequest
	GetWallHeight() *int64
}

type LabelBuildRequest struct {
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// PATCH
	ModelStyle *string `json:"ModelStyle,omitempty" xml:"ModelStyle,omitempty"`
	// example:
	//
	// OFF
	OptimizeWallWidth *string `json:"OptimizeWallWidth,omitempty" xml:"OptimizeWallWidth,omitempty"`
	// example:
	//
	// DEFAULT
	PlanStyle *string `json:"PlanStyle,omitempty" xml:"PlanStyle,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234****
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// example:
	//
	// 0
	WallHeight *int64 `json:"WallHeight,omitempty" xml:"WallHeight,omitempty"`
}

func (s LabelBuildRequest) String() string {
	return dara.Prettify(s)
}

func (s LabelBuildRequest) GoString() string {
	return s.String()
}

func (s *LabelBuildRequest) GetMode() *string {
	return s.Mode
}

func (s *LabelBuildRequest) GetModelStyle() *string {
	return s.ModelStyle
}

func (s *LabelBuildRequest) GetOptimizeWallWidth() *string {
	return s.OptimizeWallWidth
}

func (s *LabelBuildRequest) GetPlanStyle() *string {
	return s.PlanStyle
}

func (s *LabelBuildRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *LabelBuildRequest) GetWallHeight() *int64 {
	return s.WallHeight
}

func (s *LabelBuildRequest) SetMode(v string) *LabelBuildRequest {
	s.Mode = &v
	return s
}

func (s *LabelBuildRequest) SetModelStyle(v string) *LabelBuildRequest {
	s.ModelStyle = &v
	return s
}

func (s *LabelBuildRequest) SetOptimizeWallWidth(v string) *LabelBuildRequest {
	s.OptimizeWallWidth = &v
	return s
}

func (s *LabelBuildRequest) SetPlanStyle(v string) *LabelBuildRequest {
	s.PlanStyle = &v
	return s
}

func (s *LabelBuildRequest) SetSceneId(v string) *LabelBuildRequest {
	s.SceneId = &v
	return s
}

func (s *LabelBuildRequest) SetWallHeight(v int64) *LabelBuildRequest {
	s.WallHeight = &v
	return s
}

func (s *LabelBuildRequest) Validate() error {
	return dara.Validate(s)
}
