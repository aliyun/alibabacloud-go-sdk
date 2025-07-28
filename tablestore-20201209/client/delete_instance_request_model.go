// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteInstanceRequest
	GetInstanceName() *string
}

type DeleteInstanceRequest struct {
	// The name of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteInstanceRequest) SetInstanceName(v string) *DeleteInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
