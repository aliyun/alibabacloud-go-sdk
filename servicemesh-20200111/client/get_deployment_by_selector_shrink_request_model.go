// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentBySelectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGuestCluster(v string) *GetDeploymentBySelectorShrinkRequest
	GetGuestCluster() *string
	SetLabelSelectorShrink(v string) *GetDeploymentBySelectorShrinkRequest
	GetLabelSelectorShrink() *string
	SetLimit(v int64) *GetDeploymentBySelectorShrinkRequest
	GetLimit() *int64
	SetMark(v string) *GetDeploymentBySelectorShrinkRequest
	GetMark() *string
	SetNameSpace(v string) *GetDeploymentBySelectorShrinkRequest
	GetNameSpace() *string
	SetServiceMeshId(v string) *GetDeploymentBySelectorShrinkRequest
	GetServiceMeshId() *string
}

type GetDeploymentBySelectorShrinkRequest struct {
	// The name of the cluster.
	//
	// example:
	//
	// cbe80a56d07ed45818b4d39273e23****
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The label selector information.
	LabelSelectorShrink *string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
	// The maximum number of returned data entries.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The marker of data queried last time.
	//
	// example:
	//
	// eyJ2IjoibWV0YS5rOHMuaW8vdjEiLCJydiI6NzgxODk3MCwic3RhcnQiOiJuZ2lueDQ1N1x1MDAw****
	Mark *string `json:"Mark,omitempty" xml:"Mark,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// The ID of the Alibaba Cloud Service Mesh (ASM) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ce9fc65def2aa4c918747b9360fbd****
	ServiceMeshId *string `json:"ServiceMeshId,omitempty" xml:"ServiceMeshId,omitempty"`
}

func (s GetDeploymentBySelectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentBySelectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorShrinkRequest) GetGuestCluster() *string {
	return s.GuestCluster
}

func (s *GetDeploymentBySelectorShrinkRequest) GetLabelSelectorShrink() *string {
	return s.LabelSelectorShrink
}

func (s *GetDeploymentBySelectorShrinkRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetDeploymentBySelectorShrinkRequest) GetMark() *string {
	return s.Mark
}

func (s *GetDeploymentBySelectorShrinkRequest) GetNameSpace() *string {
	return s.NameSpace
}

func (s *GetDeploymentBySelectorShrinkRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetDeploymentBySelectorShrinkRequest) SetGuestCluster(v string) *GetDeploymentBySelectorShrinkRequest {
	s.GuestCluster = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetLabelSelectorShrink(v string) *GetDeploymentBySelectorShrinkRequest {
	s.LabelSelectorShrink = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetLimit(v int64) *GetDeploymentBySelectorShrinkRequest {
	s.Limit = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetMark(v string) *GetDeploymentBySelectorShrinkRequest {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetNameSpace(v string) *GetDeploymentBySelectorShrinkRequest {
	s.NameSpace = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) SetServiceMeshId(v string) *GetDeploymentBySelectorShrinkRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetDeploymentBySelectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
