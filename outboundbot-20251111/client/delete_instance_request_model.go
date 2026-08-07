// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
}

type DeleteInstanceRequest struct {
	// The instance ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b04
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
