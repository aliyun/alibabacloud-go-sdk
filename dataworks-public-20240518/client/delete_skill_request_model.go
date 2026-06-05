// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteSkillRequest
	GetName() *string
}

type DeleteSkillRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillRequest) GoString() string {
	return s.String()
}

func (s *DeleteSkillRequest) GetName() *string {
	return s.Name
}

func (s *DeleteSkillRequest) SetName(v string) *DeleteSkillRequest {
	s.Name = &v
	return s
}

func (s *DeleteSkillRequest) Validate() error {
	return dara.Validate(s)
}
