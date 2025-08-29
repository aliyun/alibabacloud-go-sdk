// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneLaboratoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloneExperimentGroup(v bool) *CloneLaboratoryRequest
	GetCloneExperimentGroup() *bool
	SetEnvironment(v string) *CloneLaboratoryRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CloneLaboratoryRequest
	GetInstanceId() *string
}

type CloneLaboratoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// true
	CloneExperimentGroup *bool `json:"CloneExperimentGroup,omitempty" xml:"CloneExperimentGroup,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneLaboratoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneLaboratoryRequest) GoString() string {
	return s.String()
}

func (s *CloneLaboratoryRequest) GetCloneExperimentGroup() *bool {
	return s.CloneExperimentGroup
}

func (s *CloneLaboratoryRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CloneLaboratoryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneLaboratoryRequest) SetCloneExperimentGroup(v bool) *CloneLaboratoryRequest {
	s.CloneExperimentGroup = &v
	return s
}

func (s *CloneLaboratoryRequest) SetEnvironment(v string) *CloneLaboratoryRequest {
	s.Environment = &v
	return s
}

func (s *CloneLaboratoryRequest) SetInstanceId(v string) *CloneLaboratoryRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneLaboratoryRequest) Validate() error {
	return dara.Validate(s)
}
