// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagWeightShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppRagWeightShrinkRequest
	GetAppId() *string
	SetRankWeightsShrink(v string) *UpdateMmAppRagWeightShrinkRequest
	GetRankWeightsShrink() *string
	SetWorkspaceId(v string) *UpdateMmAppRagWeightShrinkRequest
	GetWorkspaceId() *string
}

type UpdateMmAppRagWeightShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId             *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RankWeightsShrink *string `json:"RankWeights,omitempty" xml:"RankWeights,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-6uhm7nfev4k8pwcz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppRagWeightShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagWeightShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagWeightShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppRagWeightShrinkRequest) GetRankWeightsShrink() *string {
	return s.RankWeightsShrink
}

func (s *UpdateMmAppRagWeightShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppRagWeightShrinkRequest) SetAppId(v string) *UpdateMmAppRagWeightShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppRagWeightShrinkRequest) SetRankWeightsShrink(v string) *UpdateMmAppRagWeightShrinkRequest {
	s.RankWeightsShrink = &v
	return s
}

func (s *UpdateMmAppRagWeightShrinkRequest) SetWorkspaceId(v string) *UpdateMmAppRagWeightShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppRagWeightShrinkRequest) Validate() error {
	return dara.Validate(s)
}
