// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunStyleFeatureAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*string) *RunStyleFeatureAnalysisRequest
	GetContents() []*string
	SetMaterialIds(v []*int64) *RunStyleFeatureAnalysisRequest
	GetMaterialIds() []*int64
	SetWorkspaceId(v string) *RunStyleFeatureAnalysisRequest
	GetWorkspaceId() *string
}

type RunStyleFeatureAnalysisRequest struct {
	Contents    []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	MaterialIds []*int64  `json:"MaterialIds,omitempty" xml:"MaterialIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunStyleFeatureAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunStyleFeatureAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunStyleFeatureAnalysisRequest) GetContents() []*string {
	return s.Contents
}

func (s *RunStyleFeatureAnalysisRequest) GetMaterialIds() []*int64 {
	return s.MaterialIds
}

func (s *RunStyleFeatureAnalysisRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunStyleFeatureAnalysisRequest) SetContents(v []*string) *RunStyleFeatureAnalysisRequest {
	s.Contents = v
	return s
}

func (s *RunStyleFeatureAnalysisRequest) SetMaterialIds(v []*int64) *RunStyleFeatureAnalysisRequest {
	s.MaterialIds = v
	return s
}

func (s *RunStyleFeatureAnalysisRequest) SetWorkspaceId(v string) *RunStyleFeatureAnalysisRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunStyleFeatureAnalysisRequest) Validate() error {
	return dara.Validate(s)
}
