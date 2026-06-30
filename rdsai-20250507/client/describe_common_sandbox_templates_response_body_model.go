// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonSandboxTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeCommonSandboxTemplatesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCommonSandboxTemplatesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCommonSandboxTemplatesResponseBody
	GetRequestId() *string
	SetTemplates(v []*DescribeCommonSandboxTemplatesResponseBodyTemplates) *DescribeCommonSandboxTemplatesResponseBody
	GetTemplates() []*DescribeCommonSandboxTemplatesResponseBodyTemplates
}

type DescribeCommonSandboxTemplatesResponseBody struct {
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
	// FE9C65D7-930F-57A5-A207-8C396329241C
	RequestId *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Templates []*DescribeCommonSandboxTemplatesResponseBodyTemplates `json:"Templates,omitempty" xml:"Templates,omitempty" type:"Repeated"`
}

func (s DescribeCommonSandboxTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonSandboxTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommonSandboxTemplatesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCommonSandboxTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCommonSandboxTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommonSandboxTemplatesResponseBody) GetTemplates() []*DescribeCommonSandboxTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *DescribeCommonSandboxTemplatesResponseBody) SetMaxResults(v int32) *DescribeCommonSandboxTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBody) SetNextToken(v string) *DescribeCommonSandboxTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBody) SetRequestId(v string) *DescribeCommonSandboxTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBody) SetTemplates(v []*DescribeCommonSandboxTemplatesResponseBodyTemplates) *DescribeCommonSandboxTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommonSandboxTemplatesResponseBodyTemplates struct {
	// example:
	//
	// 1
	DefaultCpu *string `json:"DefaultCpu,omitempty" xml:"DefaultCpu,omitempty"`
	// example:
	//
	// 1Gi
	DefaultMemory *string `json:"DefaultMemory,omitempty" xml:"DefaultMemory,omitempty"`
	// example:
	//
	// 1
	DefaultReplicas *int64 `json:"DefaultReplicas,omitempty" xml:"DefaultReplicas,omitempty"`
	// example:
	//
	// Execute user-provided Python code in the sandbox environment. Runs any Python script the user provides and returns the output.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// desktop
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeCommonSandboxTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonSandboxTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) GetDefaultCpu() *string {
	return s.DefaultCpu
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) GetDefaultMemory() *string {
	return s.DefaultMemory
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) GetDefaultReplicas() *int64 {
	return s.DefaultReplicas
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) SetDefaultCpu(v string) *DescribeCommonSandboxTemplatesResponseBodyTemplates {
	s.DefaultCpu = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) SetDefaultMemory(v string) *DescribeCommonSandboxTemplatesResponseBodyTemplates {
	s.DefaultMemory = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) SetDefaultReplicas(v int64) *DescribeCommonSandboxTemplatesResponseBodyTemplates {
	s.DefaultReplicas = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) SetDescription(v string) *DescribeCommonSandboxTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) SetName(v string) *DescribeCommonSandboxTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *DescribeCommonSandboxTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
