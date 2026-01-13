// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVCUInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteVCUInstanceRequest
	GetInstanceName() *string
}

type DeleteVCUInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// instance_name
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s DeleteVCUInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVCUInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteVCUInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteVCUInstanceRequest) SetInstanceName(v string) *DeleteVCUInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteVCUInstanceRequest) Validate() error {
	return dara.Validate(s)
}
