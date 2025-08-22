// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *CreateClusterShrinkRequest
	GetClusterName() *string
	SetClusterSpec(v string) *CreateClusterShrinkRequest
	GetClusterSpec() *string
	SetEngineType(v string) *CreateClusterShrinkRequest
	GetEngineType() *string
	SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest
	GetTag() []*CreateClusterShrinkRequestTag
	SetVSwitchesShrink(v string) *CreateClusterShrinkRequest
	GetVSwitchesShrink() *string
	SetVpcId(v string) *CreateClusterShrinkRequest
	GetVpcId() *string
}

type CreateClusterShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// qianxi-test-0812
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// scx.dev.x1
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob
	EngineType *string                          `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	Tag        []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	VSwitchesShrink *string `json:"VSwitches,omitempty" xml:"VSwitches,omitempty"`
	// VPC id
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-aa1a18236n90rqhuhhnhh
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterShrinkRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateClusterShrinkRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateClusterShrinkRequest) GetTag() []*CreateClusterShrinkRequestTag {
	return s.Tag
}

func (s *CreateClusterShrinkRequest) GetVSwitchesShrink() *string {
	return s.VSwitchesShrink
}

func (s *CreateClusterShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterSpec(v string) *CreateClusterShrinkRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetEngineType(v string) *CreateClusterShrinkRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterShrinkRequest) SetVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.VSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetVpcId(v string) *CreateClusterShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type CreateClusterShrinkRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
