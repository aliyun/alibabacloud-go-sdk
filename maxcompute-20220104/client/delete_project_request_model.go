// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsLogical(v bool) *DeleteProjectRequest
	GetIsLogical() *bool
}

type DeleteProjectRequest struct {
	// Specifies whether to logically delete the project. Valid values: true and false. Default value: true. The value false indicates that the project is physically deleted.
	//
	// example:
	//
	// true
	IsLogical *bool `json:"isLogical,omitempty" xml:"isLogical,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetIsLogical() *bool {
	return s.IsLogical
}

func (s *DeleteProjectRequest) SetIsLogical(v bool) *DeleteProjectRequest {
	s.IsLogical = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
