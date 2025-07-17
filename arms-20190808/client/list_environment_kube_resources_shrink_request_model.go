// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentKubeResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvironmentKubeResourcesShrinkRequest
	GetEnvironmentId() *string
	SetKind(v string) *ListEnvironmentKubeResourcesShrinkRequest
	GetKind() *string
	SetLabelSelectorsShrink(v string) *ListEnvironmentKubeResourcesShrinkRequest
	GetLabelSelectorsShrink() *string
	SetNamespace(v string) *ListEnvironmentKubeResourcesShrinkRequest
	GetNamespace() *string
	SetRegionId(v string) *ListEnvironmentKubeResourcesShrinkRequest
	GetRegionId() *string
}

type ListEnvironmentKubeResourcesShrinkRequest struct {
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
	LabelSelectorsShrink *string `json:"LabelSelectors,omitempty" xml:"LabelSelectors,omitempty"`
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

func (s ListEnvironmentKubeResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentKubeResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) GetKind() *string {
	return s.Kind
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) GetLabelSelectorsShrink() *string {
	return s.LabelSelectorsShrink
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) SetEnvironmentId(v string) *ListEnvironmentKubeResourcesShrinkRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) SetKind(v string) *ListEnvironmentKubeResourcesShrinkRequest {
	s.Kind = &v
	return s
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) SetLabelSelectorsShrink(v string) *ListEnvironmentKubeResourcesShrinkRequest {
	s.LabelSelectorsShrink = &v
	return s
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) SetNamespace(v string) *ListEnvironmentKubeResourcesShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) SetRegionId(v string) *ListEnvironmentKubeResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentKubeResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
