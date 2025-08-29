// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushAllExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PushAllExperimentRequest
	GetInstanceId() *string
}

type PushAllExperimentRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PushAllExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s PushAllExperimentRequest) GoString() string {
	return s.String()
}

func (s *PushAllExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PushAllExperimentRequest) SetInstanceId(v string) *PushAllExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *PushAllExperimentRequest) Validate() error {
	return dara.Validate(s)
}
