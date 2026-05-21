// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSavedQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSavedQueryRequest
	GetDescription() *string
	SetExpression(v string) *CreateSavedQueryRequest
	GetExpression() *string
	SetName(v string) *CreateSavedQueryRequest
	GetName() *string
}

type CreateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// Queries all resources on which you have permissions and sorts the resources by resource type and resource ID.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The query statement in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// 	- The name must be 1 to 64 characters in length.
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// Query of All Alibaba Cloud Resources
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateSavedQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *CreateSavedQueryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSavedQueryRequest) GetExpression() *string {
	return s.Expression
}

func (s *CreateSavedQueryRequest) GetName() *string {
	return s.Name
}

func (s *CreateSavedQueryRequest) SetDescription(v string) *CreateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *CreateSavedQueryRequest) SetExpression(v string) *CreateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *CreateSavedQueryRequest) SetName(v string) *CreateSavedQueryRequest {
	s.Name = &v
	return s
}

func (s *CreateSavedQueryRequest) Validate() error {
	return dara.Validate(s)
}
