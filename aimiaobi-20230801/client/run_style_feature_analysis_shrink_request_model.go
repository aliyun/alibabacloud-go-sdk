// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleFeatureAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest
	GetContentsShrink() *string
	SetMaterialIdsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest
	GetMaterialIdsShrink() *string
	SetWorkspaceId(v string) *RunStyleFeatureAnalysisShrinkRequest
	GetWorkspaceId() *string
}

type RunStyleFeatureAnalysisShrinkRequest struct {
	ContentsShrink    *string `json:"Contents,omitempty" xml:"Contents,omitempty"`
	MaterialIdsShrink *string `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunStyleFeatureAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisShrinkRequest) GetContentsShrink() *string {
	return s.ContentsShrink
}

func (s *RunStyleFeatureAnalysisShrinkRequest) GetMaterialIdsShrink() *string {
	return s.MaterialIdsShrink
}

func (s *RunStyleFeatureAnalysisShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetContentsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.ContentsShrink = &v
	return s
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetMaterialIdsShrink(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.MaterialIdsShrink = &v
	return s
}

func (s *RunStyleFeatureAnalysisShrinkRequest) SetWorkspaceId(v string) *RunStyleFeatureAnalysisShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStyleFeatureAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
