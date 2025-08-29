// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficControlTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteTrafficControlTargetRequest
	GetInstanceId() *string
}

type DeleteTrafficControlTargetRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteTrafficControlTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficControlTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteTrafficControlTargetRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteTrafficControlTargetRequest) SetInstanceId(v string) *DeleteTrafficControlTargetRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTrafficControlTargetRequest) Validate() error {
	return dara.Validate(s)
}
