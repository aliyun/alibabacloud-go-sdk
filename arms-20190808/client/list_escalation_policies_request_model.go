// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEscalationPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListEscalationPoliciesRequest
	GetName() *string
	SetPage(v int64) *ListEscalationPoliciesRequest
	GetPage() *int64
	SetSize(v int64) *ListEscalationPoliciesRequest
	GetSize() *int64
}

type ListEscalationPoliciesRequest struct {
	// The name of the escalation policy.
	//
	// example:
	//
	// prod escalation policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListEscalationPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEscalationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListEscalationPoliciesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListEscalationPoliciesRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListEscalationPoliciesRequest) SetName(v string) *ListEscalationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListEscalationPoliciesRequest) SetPage(v int64) *ListEscalationPoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListEscalationPoliciesRequest) SetSize(v int64) *ListEscalationPoliciesRequest {
	s.Size = &v
	return s
}

func (s *ListEscalationPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
