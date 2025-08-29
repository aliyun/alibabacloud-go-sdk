// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnlineExperimentGroupRequest
	GetInstanceId() *string
}

type OnlineExperimentGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineExperimentGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentGroupRequest) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnlineExperimentGroupRequest) SetInstanceId(v string) *OnlineExperimentGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *OnlineExperimentGroupRequest) Validate() error {
	return dara.Validate(s)
}
