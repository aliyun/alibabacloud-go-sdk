// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetExperimentGroupRequest
	GetInstanceId() *string
}

type GetExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetExperimentGroupRequest) SetInstanceId(v string) *GetExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *GetExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
