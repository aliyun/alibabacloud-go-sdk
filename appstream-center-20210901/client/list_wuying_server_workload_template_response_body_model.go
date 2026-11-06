// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWuyingServerWorkloadTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListWuyingServerWorkloadTemplateResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWuyingServerWorkloadTemplateResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWuyingServerWorkloadTemplateResponseBody
	GetRequestId() *string
	SetWorkloadTemplates(v []*ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) *ListWuyingServerWorkloadTemplateResponseBody
	GetWorkloadTemplates() []*ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates
}

type ListWuyingServerWorkloadTemplateResponseBody struct {
	// A compatibility field. The current version does not use this field to indicate the number of returned entries.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A compatibility field. The current version does not generate a next page token.
	//
	// example:
	//
	// unused
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot request issues.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of workload templates. Up to 100 templates are returned, sorted by template configuration order.
	WorkloadTemplates []*ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates `json:"WorkloadTemplates,omitempty" xml:"WorkloadTemplates,omitempty" type:"Repeated"`
}

func (s ListWuyingServerWorkloadTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerWorkloadTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) GetWorkloadTemplates() []*ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates {
	return s.WorkloadTemplates
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) SetMaxResults(v int32) *ListWuyingServerWorkloadTemplateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) SetNextToken(v string) *ListWuyingServerWorkloadTemplateResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) SetRequestId(v string) *ListWuyingServerWorkloadTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) SetWorkloadTemplates(v []*ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) *ListWuyingServerWorkloadTemplateResponseBody {
	s.WorkloadTemplates = v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBody) Validate() error {
	if s.WorkloadTemplates != nil {
		for _, item := range s.WorkloadTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates struct {
	// The rendered template content. The original text is Base64-encoded after being encoded in UTF-8. You must decode the Base64 string before use.
	//
	// example:
	//
	// IyB3b3JrbG9hZCB0ZW1wbGF0ZQo=
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the workload template.
	//
	// example:
	//
	// Used to submit general-purpose computing workloads
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the workload template.
	//
	// example:
	//
	// General-purpose computing task template
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) GoString() string {
	return s.String()
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) GetContent() *string {
	return s.Content
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) GetName() *string {
	return s.Name
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) SetContent(v string) *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates {
	s.Content = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) SetDescription(v string) *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates {
	s.Description = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) SetName(v string) *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates {
	s.Name = &v
	return s
}

func (s *ListWuyingServerWorkloadTemplateResponseBodyWorkloadTemplates) Validate() error {
	return dara.Validate(s)
}
