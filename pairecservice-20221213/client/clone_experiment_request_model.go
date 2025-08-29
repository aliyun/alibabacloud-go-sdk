// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CloneExperimentRequest
	GetInstanceId() *string
}

type CloneExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentRequest) GoString() string {
	return s.String()
}

func (s *CloneExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneExperimentRequest) SetInstanceId(v string) *CloneExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneExperimentRequest) Validate() error {
	return dara.Validate(s)
}
