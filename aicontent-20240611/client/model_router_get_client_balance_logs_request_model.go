// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeType(v string) *ModelRouterGetClientBalanceLogsRequest
	GetChangeType() *string
	SetMaxResults(v int32) *ModelRouterGetClientBalanceLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterGetClientBalanceLogsRequest
	GetNextToken() *string
	SetPage(v int32) *ModelRouterGetClientBalanceLogsRequest
	GetPage() *int32
	SetSize(v int32) *ModelRouterGetClientBalanceLogsRequest
	GetSize() *int32
}

type ModelRouterGetClientBalanceLogsRequest struct {
	// example:
	//
	// recharge
	ChangeType *string `json:"changeType,omitempty" xml:"changeType,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ModelRouterGetClientBalanceLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceLogsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceLogsRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *ModelRouterGetClientBalanceLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterGetClientBalanceLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterGetClientBalanceLogsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterGetClientBalanceLogsRequest) GetSize() *int32 {
	return s.Size
}

func (s *ModelRouterGetClientBalanceLogsRequest) SetChangeType(v string) *ModelRouterGetClientBalanceLogsRequest {
	s.ChangeType = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsRequest) SetMaxResults(v int32) *ModelRouterGetClientBalanceLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsRequest) SetNextToken(v string) *ModelRouterGetClientBalanceLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsRequest) SetPage(v int32) *ModelRouterGetClientBalanceLogsRequest {
	s.Page = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsRequest) SetSize(v int32) *ModelRouterGetClientBalanceLogsRequest {
	s.Size = &v
	return s
}

func (s *ModelRouterGetClientBalanceLogsRequest) Validate() error {
	return dara.Validate(s)
}
