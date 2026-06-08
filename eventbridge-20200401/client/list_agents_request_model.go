// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfter(v string) *ListAgentsRequest
	GetAfter() *string
	SetLimit(v string) *ListAgentsRequest
	GetLimit() *string
	SetOrder(v string) *ListAgentsRequest
	GetOrder() *string
}

type ListAgentsRequest struct {
	// example:
	//
	// my-agent
	After *string `json:"After,omitempty" xml:"After,omitempty"`
	// example:
	//
	// 50
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// acs
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetAfter() *string {
	return s.After
}

func (s *ListAgentsRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListAgentsRequest) GetOrder() *string {
	return s.Order
}

func (s *ListAgentsRequest) SetAfter(v string) *ListAgentsRequest {
	s.After = &v
	return s
}

func (s *ListAgentsRequest) SetLimit(v string) *ListAgentsRequest {
	s.Limit = &v
	return s
}

func (s *ListAgentsRequest) SetOrder(v string) *ListAgentsRequest {
	s.Order = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
