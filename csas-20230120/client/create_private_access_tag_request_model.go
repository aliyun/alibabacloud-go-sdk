// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrivateAccessTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePrivateAccessTagRequest
	GetDescription() *string
	SetName(v string) *CreatePrivateAccessTagRequest
	GetName() *string
}

type CreatePrivateAccessTagRequest struct {
	// The description of the internal-facing access tag. The description must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and spaces. Chinese characters are supported.
	//
	// example:
	//
	// This is an internal-facing access tag
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the internal-facing access tag. The name must be 1 to 128 characters in length and can contain letters, digits, periods (.), underscores (_), and hyphens (-). Chinese characters are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// tag_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePrivateAccessTagRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrivateAccessTagRequest) GoString() string {
	return s.String()
}

func (s *CreatePrivateAccessTagRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePrivateAccessTagRequest) GetName() *string {
	return s.Name
}

func (s *CreatePrivateAccessTagRequest) SetDescription(v string) *CreatePrivateAccessTagRequest {
	s.Description = &v
	return s
}

func (s *CreatePrivateAccessTagRequest) SetName(v string) *CreatePrivateAccessTagRequest {
	s.Name = &v
	return s
}

func (s *CreatePrivateAccessTagRequest) Validate() error {
	return dara.Validate(s)
}
