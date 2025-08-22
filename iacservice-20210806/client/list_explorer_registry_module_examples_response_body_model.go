// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleExamplesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExplorerRegistryModuleExamples(v []*ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) *ListExplorerRegistryModuleExamplesResponseBody
	GetExplorerRegistryModuleExamples() []*ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples
	SetMaxResults(v int32) *ListExplorerRegistryModuleExamplesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExplorerRegistryModuleExamplesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExplorerRegistryModuleExamplesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExplorerRegistryModuleExamplesResponseBody
	GetTotalCount() *int64
}

type ListExplorerRegistryModuleExamplesResponseBody struct {
	ExplorerRegistryModuleExamples []*ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples `json:"explorerRegistryModuleExamples,omitempty" xml:"explorerRegistryModuleExamples,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// VSjwD+sJ8OZJ8fNjV89AZs7o2AdSD25ZQLeWZ8REjXA=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// BA8F6459-EED6-556B-8130-D150A3866E56
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 132
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExplorerRegistryModuleExamplesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleExamplesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) GetExplorerRegistryModuleExamples() []*ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	return s.ExplorerRegistryModuleExamples
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) SetExplorerRegistryModuleExamples(v []*ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) *ListExplorerRegistryModuleExamplesResponseBody {
	s.ExplorerRegistryModuleExamples = v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) SetMaxResults(v int32) *ListExplorerRegistryModuleExamplesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) SetNextToken(v string) *ListExplorerRegistryModuleExamplesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) SetRequestId(v string) *ListExplorerRegistryModuleExamplesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) SetTotalCount(v int64) *ListExplorerRegistryModuleExamplesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples struct {
	// example:
	//
	// 201-use-case-create-actiontrail-trail
	ExampleName *string `json:"exampleName,omitempty" xml:"exampleName,omitempty"`
	// example:
	//
	// /
	ExamplePath *string `json:"examplePath,omitempty" xml:"examplePath,omitempty"`
	// example:
	//
	// {}
	ExampleSchema map[string]interface{} `json:"exampleSchema,omitempty" xml:"exampleSchema,omitempty"`
	// example:
	//
	// iactestname1
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// 1.11.0
	ModuleVersion *string `json:"moduleVersion,omitempty" xml:"moduleVersion,omitempty"`
	// example:
	//
	// alibaba
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetExampleName() *string {
	return s.ExampleName
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetExamplePath() *string {
	return s.ExamplePath
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetExampleSchema() map[string]interface{} {
	return s.ExampleSchema
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetModuleVersion() *string {
	return s.ModuleVersion
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) GetStatus() *string {
	return s.Status
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetExampleName(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.ExampleName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetExamplePath(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.ExamplePath = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetExampleSchema(v map[string]interface{}) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.ExampleSchema = v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetModuleName(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetModuleVersion(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.ModuleVersion = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetNamespaceName(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.NamespaceName = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) SetStatus(v string) *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples {
	s.Status = &v
	return s
}

func (s *ListExplorerRegistryModuleExamplesResponseBodyExplorerRegistryModuleExamples) Validate() error {
	return dara.Validate(s)
}
