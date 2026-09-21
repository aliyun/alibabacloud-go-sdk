// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *UninstallSkillsRequest
	GetInstanceIds() []*string
	SetSkillIds(v []*string) *UninstallSkillsRequest
	GetSkillIds() []*string
}

type UninstallSkillsRequest struct {
	// The list of cloud phone instance IDs. You can specify 1 to 200 instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The list of skill IDs. You can specify 1 to 10 skill IDs.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
}

func (s UninstallSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallSkillsRequest) GoString() string {
	return s.String()
}

func (s *UninstallSkillsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *UninstallSkillsRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *UninstallSkillsRequest) SetInstanceIds(v []*string) *UninstallSkillsRequest {
	s.InstanceIds = v
	return s
}

func (s *UninstallSkillsRequest) SetSkillIds(v []*string) *UninstallSkillsRequest {
	s.SkillIds = v
	return s
}

func (s *UninstallSkillsRequest) Validate() error {
	return dara.Validate(s)
}
