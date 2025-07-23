// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *PauseInstanceRequest
	GetInstanceId() *string
}

type PauseInstanceRequest struct {
	// The ID of the HSM.
	//
	// This parameter is required.
	//
	// example:
	//
	// hsm-cn-vj30bil8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s PauseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseInstanceRequest) GoString() string {
	return s.String()
}

func (s *PauseInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *PauseInstanceRequest) SetInstanceId(v string) *PauseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *PauseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
