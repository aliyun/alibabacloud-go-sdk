// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSavedQueryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateSavedQueryRequest
	GetDescription() *string
	SetExpression(v string) *UpdateSavedQueryRequest
	GetExpression() *string
	SetName(v string) *UpdateSavedQueryRequest
	GetName() *string
	SetQueryId(v string) *UpdateSavedQueryRequest
	GetQueryId() *string
}

type UpdateSavedQueryRequest struct {
	// The description of the template.
	//
	// The description must be 1 to 256 characters in length.
	//
	// example:
	//
	// 查询我的所有资源，返回所有的列。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expression of the template.
	//
	// example:
	//
	// SELECT 	- FROM resources;
	Expression *string `json:"Expression,omitempty" xml:"Expression,omitempty"`
	// The name of the template.
	//
	// - The name must be 1 to 64 characters in length.
	//
	// - It can contain letters, digits, underscores (_), and hyphens (-).
	//
	// - It must be globally unique.
	//
	// example:
	//
	// 查询我的所有资源。
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// cq-GeAck****
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s UpdateSavedQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSavedQueryRequest) GoString() string {
	return s.String()
}

func (s *UpdateSavedQueryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateSavedQueryRequest) GetExpression() *string {
	return s.Expression
}

func (s *UpdateSavedQueryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSavedQueryRequest) GetQueryId() *string {
	return s.QueryId
}

func (s *UpdateSavedQueryRequest) SetDescription(v string) *UpdateSavedQueryRequest {
	s.Description = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetExpression(v string) *UpdateSavedQueryRequest {
	s.Expression = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetName(v string) *UpdateSavedQueryRequest {
	s.Name = &v
	return s
}

func (s *UpdateSavedQueryRequest) SetQueryId(v string) *UpdateSavedQueryRequest {
	s.QueryId = &v
	return s
}

func (s *UpdateSavedQueryRequest) Validate() error {
	return dara.Validate(s)
}
