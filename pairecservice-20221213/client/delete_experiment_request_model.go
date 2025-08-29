// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteExperimentRequest
	GetInstanceId() *string
}

type DeleteExperimentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteExperimentRequest) SetInstanceId(v string) *DeleteExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExperimentRequest) Validate() error {
	return dara.Validate(s)
}
