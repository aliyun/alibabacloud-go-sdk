// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentKubeResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvironmentKubeResourcesRequest
	GetEnvironmentId() *string
	SetKind(v string) *ListEnvironmentKubeResourcesRequest
	GetKind() *string
	SetLabelSelectors(v map[string]*string) *ListEnvironmentKubeResourcesRequest
	GetLabelSelectors() map[string]*string
	SetNamespace(v string) *ListEnvironmentKubeResourcesRequest
	GetNamespace() *string
	SetRegionId(v string) *ListEnvironmentKubeResourcesRequest
	GetRegionId() *string
}

type ListEnvironmentKubeResourcesRequest struct {
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The resource type. Valid values: Pod, Deployment, and Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// Pod
	Kind *string `json:"Kind,omitempty" xml:"Kind,omitempty"`
	// The tags.
	LabelSelectors map[string]*string `json:"LabelSelectors,omitempty" xml:"LabelSelectors,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnvironmentKubeResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentKubeResourcesRequest) GetKind() *string {
	return s.Kind
}

func (s *ListEnvironmentKubeResourcesRequest) GetLabelSelectors() map[string]*string {
	return s.LabelSelectors
}

func (s *ListEnvironmentKubeResourcesRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEnvironmentKubeResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentKubeResourcesRequest) SetEnvironmentId(v string) *ListEnvironmentKubeResourcesRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentKubeResourcesRequest) SetKind(v string) *ListEnvironmentKubeResourcesRequest {
	s.Kind = &v
	return s
}

func (s *ListEnvironmentKubeResourcesRequest) SetLabelSelectors(v map[string]*string) *ListEnvironmentKubeResourcesRequest {
	s.LabelSelectors = v
	return s
}

func (s *ListEnvironmentKubeResourcesRequest) SetNamespace(v string) *ListEnvironmentKubeResourcesRequest {
	s.Namespace = &v
	return s
}

func (s *ListEnvironmentKubeResourcesRequest) SetRegionId(v string) *ListEnvironmentKubeResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentKubeResourcesRequest) Validate() error {
	return dara.Validate(s)
}
