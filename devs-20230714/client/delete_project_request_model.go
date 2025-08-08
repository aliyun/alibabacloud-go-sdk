// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteProjectRequest
	GetForce() *bool
}

type DeleteProjectRequest struct {
	// example:
	//
	// true
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteProjectRequest) SetForce(v bool) *DeleteProjectRequest {
	s.Force = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
