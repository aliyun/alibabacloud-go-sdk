// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *CloneExperimentGroupRequest
	GetEnvironment() *string
	SetInstanceId(v string) *CloneExperimentGroupRequest
	GetInstanceId() *string
	SetLayerId(v string) *CloneExperimentGroupRequest
	GetLayerId() *string
}

type CloneExperimentGroupRequest struct {
	// The target environment for the experiment group. ● Daily: Daily environment ● Pre: pre-release environment ● Prod: production environment
	//
	// This parameter is required.
	//
	// example:
	//
	// Daily
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// The instance ID. You can get this ID by calling the ListInstances operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The layer ID. You can get this ID by calling the ListLayers operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	LayerId *string `json:"LayerId,omitempty" xml:"LayerId,omitempty"`
}

func (s CloneExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *CloneExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneExperimentGroupRequest) GetLayerId() *string {
	return s.LayerId
}

func (s *CloneExperimentGroupRequest) SetEnvironment(v string) *CloneExperimentGroupRequest {
	s.Environment = &v
	return s
}

func (s *CloneExperimentGroupRequest) SetInstanceId(v string) *CloneExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneExperimentGroupRequest) SetLayerId(v string) *CloneExperimentGroupRequest {
	s.LayerId = &v
	return s
}

func (s *CloneExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
