// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSaseUserTagRequest
	GetDescription() *string
	SetName(v string) *UpdateSaseUserTagRequest
	GetName() *string
	SetTagId(v string) *UpdateSaseUserTagRequest
	GetTagId() *string
}

type UpdateSaseUserTagRequest struct {
	// The description of the user tag.
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
	// The ID of the user tag. You can obtain the tag ID from the following operations:
	//
	// - [ListSaseUserTags](~~ListSaseUserTags~~): Lists user tags.
	//
	// - [CreateSaseUserTag](~~CreateSaseUserTag~~): Creates a user tag.
	//
	// example:
	//
	// tag-d3f64e8bdd4a****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateSaseUserTagRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSaseUserTagRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSaseUserTagRequest) GetTagId() *string {
	return s.TagId
}

func (s *UpdateSaseUserTagRequest) SetDescription(v string) *UpdateSaseUserTagRequest {
	s.Description = &v
	return s
}

func (s *UpdateSaseUserTagRequest) SetName(v string) *UpdateSaseUserTagRequest {
	s.Name = &v
	return s
}

func (s *UpdateSaseUserTagRequest) SetTagId(v string) *UpdateSaseUserTagRequest {
	s.TagId = &v
	return s
}

func (s *UpdateSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
