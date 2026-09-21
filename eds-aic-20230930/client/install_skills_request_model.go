// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *InstallSkillsRequest
	GetInstanceIds() []*string
	SetSkillIds(v []*string) *InstallSkillsRequest
	GetSkillIds() []*string
}

type InstallSkillsRequest struct {
	// The list of cloud phone instance IDs. You can specify 1 to 200 instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The list of skill IDs. You can specify 1 to 10 skill IDs.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s InstallSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallSkillsRequest) GoString() string {
	return s.String()
}

func (s *InstallSkillsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *InstallSkillsRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *InstallSkillsRequest) SetInstanceIds(v []*string) *InstallSkillsRequest {
	s.InstanceIds = v
	return s
}

func (s *InstallSkillsRequest) SetSkillIds(v []*string) *InstallSkillsRequest {
	s.SkillIds = v
	return s
}

func (s *InstallSkillsRequest) Validate() error {
	return dara.Validate(s)
}
