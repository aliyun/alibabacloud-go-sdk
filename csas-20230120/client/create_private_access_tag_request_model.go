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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
