// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanResources(v bool) *DeleteStackRequest
	GetCleanResources() *bool
}

type DeleteStackRequest struct {
	// example:
	//
	// true
	CleanResources *bool `json:"cleanResources,omitempty" xml:"cleanResources,omitempty"`
}

func (s DeleteStackRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackRequest) GoString() string {
	return s.String()
}

func (s *DeleteStackRequest) GetCleanResources() *bool {
	return s.CleanResources
}

func (s *DeleteStackRequest) SetCleanResources(v bool) *DeleteStackRequest {
	s.CleanResources = &v
	return s
}

func (s *DeleteStackRequest) Validate() error {
	return dara.Validate(s)
}
