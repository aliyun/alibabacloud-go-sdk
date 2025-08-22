// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetClusterSpec(v string) *CreateClusterRequest
	GetClusterSpec() *string
	SetEngineType(v string) *CreateClusterRequest
	GetEngineType() *string
	SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest
	GetTag() []*CreateClusterRequestTag
	SetVSwitches(v []*CreateClusterRequestVSwitches) *CreateClusterRequest
	GetVSwitches() []*CreateClusterRequestVSwitches
	SetVpcId(v string) *CreateClusterRequest
	GetVpcId() *string
}

type CreateClusterRequest struct {
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
	EngineType *string                    `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	Tag        []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// This parameter is required.
	VSwitches []*CreateClusterRequestVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	// VPC id
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-aa1a18236n90rqhuhhnhh
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *CreateClusterRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateClusterRequest) GetTag() []*CreateClusterRequestTag {
	return s.Tag
}

func (s *CreateClusterRequest) GetVSwitches() []*CreateClusterRequestVSwitches {
	return s.VSwitches
}

func (s *CreateClusterRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterSpec(v string) *CreateClusterRequest {
	s.ClusterSpec = &v
	return s
}

func (s *CreateClusterRequest) SetEngineType(v string) *CreateClusterRequest {
	s.EngineType = &v
	return s
}

func (s *CreateClusterRequest) SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterRequest) SetVSwitches(v []*CreateClusterRequestVSwitches) *CreateClusterRequest {
	s.VSwitches = v
	return s
}

func (s *CreateClusterRequest) SetVpcId(v string) *CreateClusterRequest {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestVSwitches struct {
	// This parameter is required.
	//
	// example:
	//
	// vsw-2ze745n3r2sfqtahhubpl
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestVSwitches) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestVSwitches) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequestVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequestVSwitches) SetVSwitchId(v string) *CreateClusterRequestVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestVSwitches) SetZoneId(v string) *CreateClusterRequestVSwitches {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequestVSwitches) Validate() error {
	return dara.Validate(s)
}
