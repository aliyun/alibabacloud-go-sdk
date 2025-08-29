// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *OnlineExperimentRequest
	GetInstanceId() *string
}

type OnlineExperimentRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s OnlineExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentRequest) GoString() string {
	return s.String()
}

func (s *OnlineExperimentRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OnlineExperimentRequest) SetInstanceId(v string) *OnlineExperimentRequest {
	s.InstanceId = &v
	return s
}

func (s *OnlineExperimentRequest) Validate() error {
	return dara.Validate(s)
}
