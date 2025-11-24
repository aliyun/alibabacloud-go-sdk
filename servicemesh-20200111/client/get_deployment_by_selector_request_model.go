// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentBySelectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGuestCluster(v string) *GetDeploymentBySelectorRequest
	GetGuestCluster() *string
	SetLabelSelector(v map[string]*string) *GetDeploymentBySelectorRequest
	GetLabelSelector() map[string]*string
	SetLimit(v int64) *GetDeploymentBySelectorRequest
	GetLimit() *int64
	SetMark(v string) *GetDeploymentBySelectorRequest
	GetMark() *string
	SetNameSpace(v string) *GetDeploymentBySelectorRequest
	GetNameSpace() *string
	SetServiceMeshId(v string) *GetDeploymentBySelectorRequest
	GetServiceMeshId() *string
}

type GetDeploymentBySelectorRequest struct {
	// The name of the cluster.
	//
	// example:
	//
	// cbe80a56d07ed45818b4d39273e23****
	GuestCluster *string `json:"GuestCluster,omitempty" xml:"GuestCluster,omitempty"`
	// The label selector information.
	LabelSelector map[string]*string `json:"LabelSelector,omitempty" xml:"LabelSelector,omitempty"`
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

func (s GetDeploymentBySelectorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentBySelectorRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentBySelectorRequest) GetGuestCluster() *string {
	return s.GuestCluster
}

func (s *GetDeploymentBySelectorRequest) GetLabelSelector() map[string]*string {
	return s.LabelSelector
}

func (s *GetDeploymentBySelectorRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *GetDeploymentBySelectorRequest) GetMark() *string {
	return s.Mark
}

func (s *GetDeploymentBySelectorRequest) GetNameSpace() *string {
	return s.NameSpace
}

func (s *GetDeploymentBySelectorRequest) GetServiceMeshId() *string {
	return s.ServiceMeshId
}

func (s *GetDeploymentBySelectorRequest) SetGuestCluster(v string) *GetDeploymentBySelectorRequest {
	s.GuestCluster = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetLabelSelector(v map[string]*string) *GetDeploymentBySelectorRequest {
	s.LabelSelector = v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetLimit(v int64) *GetDeploymentBySelectorRequest {
	s.Limit = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetMark(v string) *GetDeploymentBySelectorRequest {
	s.Mark = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetNameSpace(v string) *GetDeploymentBySelectorRequest {
	s.NameSpace = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) SetServiceMeshId(v string) *GetDeploymentBySelectorRequest {
	s.ServiceMeshId = &v
	return s
}

func (s *GetDeploymentBySelectorRequest) Validate() error {
	return dara.Validate(s)
}
