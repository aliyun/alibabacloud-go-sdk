// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEscalationPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListEscalationPoliciesResponseBodyPageBean) *ListEscalationPoliciesResponseBody
	GetPageBean() *ListEscalationPoliciesResponseBodyPageBean
	SetRequestId(v string) *ListEscalationPoliciesResponseBody
	GetRequestId() *string
}

type ListEscalationPoliciesResponseBody struct {
	// The returned objects.
	PageBean *ListEscalationPoliciesResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 78901766-3806-4E96-8E47-CFEF59E4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEscalationPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEscalationPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBody) GetPageBean() *ListEscalationPoliciesResponseBodyPageBean {
	return s.PageBean
}

func (s *ListEscalationPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEscalationPoliciesResponseBody) SetPageBean(v *ListEscalationPoliciesResponseBodyPageBean) *ListEscalationPoliciesResponseBody {
	s.PageBean = v
	return s
}

func (s *ListEscalationPoliciesResponseBody) SetRequestId(v string) *ListEscalationPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEscalationPoliciesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEscalationPoliciesResponseBodyPageBean struct {
	// The list of escalation policies.
	EscalationPolicies []*ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies `json:"EscalationPolicies,omitempty" xml:"EscalationPolicies,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEscalationPoliciesResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListEscalationPoliciesResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBodyPageBean) GetEscalationPolicies() []*ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies {
	return s.EscalationPolicies
}

func (s *ListEscalationPoliciesResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListEscalationPoliciesResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListEscalationPoliciesResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetEscalationPolicies(v []*ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) *ListEscalationPoliciesResponseBodyPageBean {
	s.EscalationPolicies = v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetPage(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetSize(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) SetTotal(v int64) *ListEscalationPoliciesResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies struct {
	// The ID of the escalation policy.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the escalation policy.
	//
	// example:
	//
	// prod escalation policy
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) String() string {
	return dara.Prettify(s)
}

func (s ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) GoString() string {
	return s.String()
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) GetId() *int64 {
	return s.Id
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) GetName() *string {
	return s.Name
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) SetId(v int64) *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies {
	s.Id = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) SetName(v string) *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies {
	s.Name = &v
	return s
}

func (s *ListEscalationPoliciesResponseBodyPageBeanEscalationPolicies) Validate() error {
	return dara.Validate(s)
}
