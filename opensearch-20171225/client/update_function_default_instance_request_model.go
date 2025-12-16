// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFunctionDefaultInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *UpdateFunctionDefaultInstanceRequest
	GetInstanceName() *string
}

type UpdateFunctionDefaultInstanceRequest struct {
	// The name of the instance.
	//
	// example:
	//
	// "pop_test"
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateFunctionDefaultInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateFunctionDefaultInstanceRequest) SetInstanceName(v string) *UpdateFunctionDefaultInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceRequest) Validate() error {
	return dara.Validate(s)
}
