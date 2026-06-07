// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetSkillRequest
	GetName() *string
}

type GetSkillRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetSkillRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRequest) GoString() string {
	return s.String()
}

func (s *GetSkillRequest) GetName() *string {
	return s.Name
}

func (s *GetSkillRequest) SetName(v string) *GetSkillRequest {
	s.Name = &v
	return s
}

func (s *GetSkillRequest) Validate() error {
	return dara.Validate(s)
}
