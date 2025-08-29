// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteExperimentGroupRequest
	GetInstanceId() *string
}

type DeleteExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteExperimentGroupRequest) SetInstanceId(v string) *DeleteExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
