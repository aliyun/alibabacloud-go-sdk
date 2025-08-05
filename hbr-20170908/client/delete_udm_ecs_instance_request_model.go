// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUdmEcsInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteUdmEcsInstanceRequest
	GetInstanceId() *string
}

type DeleteUdmEcsInstanceRequest struct {
	// The ID of the ECS instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-2zed************tlrm
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteUdmEcsInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUdmEcsInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteUdmEcsInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteUdmEcsInstanceRequest) SetInstanceId(v string) *DeleteUdmEcsInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteUdmEcsInstanceRequest) Validate() error {
	return dara.Validate(s)
}
