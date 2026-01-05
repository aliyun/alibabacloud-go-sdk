// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSlotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteSlotRequest
	GetForce() *bool
}

type DeleteSlotRequest struct {
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteSlotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSlotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSlotRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteSlotRequest) SetForce(v bool) *DeleteSlotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSlotRequest) Validate() error {
	return dara.Validate(s)
}
