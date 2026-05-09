// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeSandboxTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSandboxTemplatesResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeSandboxTemplatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSandboxTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSandboxTemplatesResponseBody
	GetRequestId() *string
	SetSandboxTemplates(v []*DescribeSandboxTemplatesResponseBodySandboxTemplates) *DescribeSandboxTemplatesResponseBody
	GetSandboxTemplates() []*DescribeSandboxTemplatesResponseBodySandboxTemplates
	SetTotalCount(v int64) *DescribeSandboxTemplatesResponseBody
	GetTotalCount() *int64
}

type DescribeSandboxTemplatesResponseBody struct {
	// example:
	//
	// None
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAc3HCuYhJi/wvpk4xOr0VLYz/NvD85HpgBeRBCusEIeVQ0dHZH9jr+NP3X9Jx0iSoql55b9nd4PIDm252/a0f+U=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SandboxTemplates []*DescribeSandboxTemplatesResponseBodySandboxTemplates `json:"SandboxTemplates,omitempty" xml:"SandboxTemplates,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSandboxTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSandboxTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSandboxTemplatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSandboxTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSandboxTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSandboxTemplatesResponseBody) GetSandboxTemplates() []*DescribeSandboxTemplatesResponseBodySandboxTemplates {
	return s.SandboxTemplates
}

func (s *DescribeSandboxTemplatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeSandboxTemplatesResponseBody) SetMaxResults(v int32) *DescribeSandboxTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetNextToken(v string) *DescribeSandboxTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetPageNumber(v int64) *DescribeSandboxTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetPageSize(v int64) *DescribeSandboxTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetRequestId(v string) *DescribeSandboxTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetSandboxTemplates(v []*DescribeSandboxTemplatesResponseBodySandboxTemplates) *DescribeSandboxTemplatesResponseBody {
	s.SandboxTemplates = v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) SetTotalCount(v int64) *DescribeSandboxTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBody) Validate() error {
	if s.SandboxTemplates != nil {
		for _, item := range s.SandboxTemplates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSandboxTemplatesResponseBodySandboxTemplates struct {
	// example:
	//
	// 2
	DefaultCpu *string `json:"DefaultCpu,omitempty" xml:"DefaultCpu,omitempty"`
	// example:
	//
	// 4Gi
	DefaultMemory *string `json:"DefaultMemory,omitempty" xml:"DefaultMemory,omitempty"`
	// example:
	//
	// code-interpreter-vpc
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// true
	EnableVpcAccess *string `json:"EnableVpcAccess,omitempty" xml:"EnableVpcAccess,omitempty"`
	// example:
	//
	// code-interpreter
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Replicas *int64  `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
	// example:
	//
	// code-interpreter-asdxxxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DescribeSandboxTemplatesResponseBodySandboxTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxTemplatesResponseBodySandboxTemplates) GoString() string {
	return s.String()
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetEnableVpcAccess() *string {
	return s.EnableVpcAccess
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetReplicas() *int64 {
	return s.Replicas
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetDefaultCpu(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.DefaultCpu = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetDefaultMemory(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.DefaultMemory = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetDescription(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.Description = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetEnableVpcAccess(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.EnableVpcAccess = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetName(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.Name = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetReplicas(v int64) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.Replicas = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) SetTemplateId(v string) *DescribeSandboxTemplatesResponseBodySandboxTemplates {
	s.TemplateId = &v
	return s
}

func (s *DescribeSandboxTemplatesResponseBodySandboxTemplates) Validate() error {
	return dara.Validate(s)
}
