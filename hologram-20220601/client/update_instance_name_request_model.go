// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *UpdateInstanceNameRequest
	GetInstanceName() *string
}

type UpdateInstanceNameRequest struct {
	// The new name of the instance. The name must be 2 to 64 characters in length.
	//
	// example:
	//
	// new_name
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
