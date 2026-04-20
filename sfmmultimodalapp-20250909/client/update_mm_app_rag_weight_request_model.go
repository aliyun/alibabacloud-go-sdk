// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmAppRagWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMmAppRagWeightRequest
	GetAppId() *string
	SetRankWeights(v map[string]*float64) *UpdateMmAppRagWeightRequest
	GetRankWeights() map[string]*float64
	SetWorkspaceId(v string) *UpdateMmAppRagWeightRequest
	GetWorkspaceId() *string
}

type UpdateMmAppRagWeightRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId       *string             `json:"AppId,omitempty" xml:"AppId,omitempty"`
	RankWeights map[string]*float64 `json:"RankWeights,omitempty" xml:"RankWeights,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-6uhm7nfev4k8pwcz
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMmAppRagWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmAppRagWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmAppRagWeightRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMmAppRagWeightRequest) GetRankWeights() map[string]*float64 {
	return s.RankWeights
}

func (s *UpdateMmAppRagWeightRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMmAppRagWeightRequest) SetAppId(v string) *UpdateMmAppRagWeightRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMmAppRagWeightRequest) SetRankWeights(v map[string]*float64) *UpdateMmAppRagWeightRequest {
	s.RankWeights = v
	return s
}

func (s *UpdateMmAppRagWeightRequest) SetWorkspaceId(v string) *UpdateMmAppRagWeightRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMmAppRagWeightRequest) Validate() error {
	return dara.Validate(s)
}
