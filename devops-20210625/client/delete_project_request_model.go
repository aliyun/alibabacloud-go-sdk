// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdentifier(v string) *DeleteProjectRequest
	GetIdentifier() *string
}

type DeleteProjectRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 7fb72727f32143574f7a...
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DeleteProjectRequest) SetIdentifier(v string) *DeleteProjectRequest {
	s.Identifier = &v
	return s
}

func (s *DeleteProjectRequest) Validate() error {
	return dara.Validate(s)
}
