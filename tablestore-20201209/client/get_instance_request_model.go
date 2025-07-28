// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *GetInstanceRequest
	GetInstanceName() *string
}

type GetInstanceRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *GetInstanceRequest) SetInstanceName(v string) *GetInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *GetInstanceRequest) Validate() error {
	return dara.Validate(s)
}
