// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProhibitedTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateProhibitedTagRequest
	GetDescription() *string
	SetName(v string) *UpdateProhibitedTagRequest
	GetName() *string
	SetTagId(v string) *UpdateProhibitedTagRequest
	GetTagId() *string
}

type UpdateProhibitedTagRequest struct {
	// The description of the prohibited software tag. The description can contain letters, digits, Chinese characters, spaces, periods (.), underscores (_), and hyphens (-), and cannot exceed 128 characters in length.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the prohibited software tag. The name must be 1 to 128 characters in length and can contain letters, digits, Chinese characters, periods (.), underscores (_), and hyphens (-). Spaces are not supported.
	//
	// example:
	//
	// Edge
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the custom prohibited software tag. Only custom tags under the current Alibaba Cloud account can be modified. Built-in system tags cannot be modified. You can obtain the value from the following operations:
	//
	// - [ListProhibitedTags](~~ListProhibitedTags~~): Lists prohibited software tags.
	//
	// - [CreateProhibitedTag](~~CreateProhibitedTag~~): Creates a custom prohibited software tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag-996078937c00****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s UpdateProhibitedTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProhibitedTagRequest) GoString() string {
	return s.String()
}

func (s *UpdateProhibitedTagRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateProhibitedTagRequest) GetName() *string {
	return s.Name
}

func (s *UpdateProhibitedTagRequest) GetTagId() *string {
	return s.TagId
}

func (s *UpdateProhibitedTagRequest) SetDescription(v string) *UpdateProhibitedTagRequest {
	s.Description = &v
	return s
}

func (s *UpdateProhibitedTagRequest) SetName(v string) *UpdateProhibitedTagRequest {
	s.Name = &v
	return s
}

func (s *UpdateProhibitedTagRequest) SetTagId(v string) *UpdateProhibitedTagRequest {
	s.TagId = &v
	return s
}

func (s *UpdateProhibitedTagRequest) Validate() error {
	return dara.Validate(s)
}
