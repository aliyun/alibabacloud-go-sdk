// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSaseUserTagRequest
	GetDescription() *string
	SetName(v string) *CreateSaseUserTagRequest
	GetName() *string
}

type CreateSaseUserTagRequest struct {
	// The description of the employee tag.
	//
	// example:
	//
	// These are the company\\"s employees
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the user tag.
	//
	// example:
	//
	// boss
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *CreateSaseUserTagRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSaseUserTagRequest) GetName() *string {
	return s.Name
}

func (s *CreateSaseUserTagRequest) SetDescription(v string) *CreateSaseUserTagRequest {
	s.Description = &v
	return s
}

func (s *CreateSaseUserTagRequest) SetName(v string) *CreateSaseUserTagRequest {
	s.Name = &v
	return s
}

func (s *CreateSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
