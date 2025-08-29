// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTrafficControlTaskRequest
	GetInstanceId() *string
}

type DeleteTrafficControlTaskRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTrafficControlTaskRequest) SetInstanceId(v string) *DeleteTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
