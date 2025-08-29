// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneTrafficControlTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CloneTrafficControlTaskRequest
	GetInstanceId() *string
}

type CloneTrafficControlTaskRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CloneTrafficControlTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneTrafficControlTaskRequest) GoString() string {
	return s.String()
}

func (s *CloneTrafficControlTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CloneTrafficControlTaskRequest) SetInstanceId(v string) *CloneTrafficControlTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CloneTrafficControlTaskRequest) Validate() error {
	return dara.Validate(s)
}
