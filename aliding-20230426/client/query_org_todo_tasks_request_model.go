// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgTodoTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *QueryOrgTodoTasksRequestTenantContext) *QueryOrgTodoTasksRequest
	GetTenantContext() *QueryOrgTodoTasksRequestTenantContext
	SetIsDone(v bool) *QueryOrgTodoTasksRequest
	GetIsDone() *bool
	SetNextToken(v string) *QueryOrgTodoTasksRequest
	GetNextToken() *string
}

type QueryOrgTodoTasksRequest struct {
	TenantContext *QueryOrgTodoTasksRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// example:
	//
	// true
	IsDone *bool `json:"isDone,omitempty" xml:"isDone,omitempty"`
	// example:
	//
	// 0
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s QueryOrgTodoTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksRequest) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksRequest) GetTenantContext() *QueryOrgTodoTasksRequestTenantContext {
	return s.TenantContext
}

func (s *QueryOrgTodoTasksRequest) GetIsDone() *bool {
	return s.IsDone
}

func (s *QueryOrgTodoTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryOrgTodoTasksRequest) SetTenantContext(v *QueryOrgTodoTasksRequestTenantContext) *QueryOrgTodoTasksRequest {
	s.TenantContext = v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetIsDone(v bool) *QueryOrgTodoTasksRequest {
	s.IsDone = &v
	return s
}

func (s *QueryOrgTodoTasksRequest) SetNextToken(v string) *QueryOrgTodoTasksRequest {
	s.NextToken = &v
	return s
}

func (s *QueryOrgTodoTasksRequest) Validate() error {
	return dara.Validate(s)
}

type QueryOrgTodoTasksRequestTenantContext struct {
	// example:
	//
	// 1
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s QueryOrgTodoTasksRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgTodoTasksRequestTenantContext) GoString() string {
	return s.String()
}

func (s *QueryOrgTodoTasksRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryOrgTodoTasksRequestTenantContext) SetTenantId(v string) *QueryOrgTodoTasksRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *QueryOrgTodoTasksRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
