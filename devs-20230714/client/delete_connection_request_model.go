// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteConnectionRequest
	GetForce() *bool
}

type DeleteConnectionRequest struct {
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteConnectionRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteConnectionRequest) SetForce(v bool) *DeleteConnectionRequest {
	s.Force = &v
	return s
}

func (s *DeleteConnectionRequest) Validate() error {
	return dara.Validate(s)
}
