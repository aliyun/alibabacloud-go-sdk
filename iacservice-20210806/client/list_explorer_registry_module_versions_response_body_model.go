// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModuleVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExplorerRegistryModuleVersions(v []*ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) *ListExplorerRegistryModuleVersionsResponseBody
	GetExplorerRegistryModuleVersions() []*ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions
	SetMaxResults(v int32) *ListExplorerRegistryModuleVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExplorerRegistryModuleVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExplorerRegistryModuleVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExplorerRegistryModuleVersionsResponseBody
	GetTotalCount() *int64
}

type ListExplorerRegistryModuleVersionsResponseBody struct {
	ExplorerRegistryModuleVersions []*ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions `json:"explorerRegistryModuleVersions,omitempty" xml:"explorerRegistryModuleVersions,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// oIM6ssGyh00noi5zoDR1hJ4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// F2D40488-3F74-568B-87EC-1C04D098DF8B
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 22
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExplorerRegistryModuleVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) GetExplorerRegistryModuleVersions() []*ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	return s.ExplorerRegistryModuleVersions
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) SetExplorerRegistryModuleVersions(v []*ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) *ListExplorerRegistryModuleVersionsResponseBody {
	s.ExplorerRegistryModuleVersions = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) SetMaxResults(v int32) *ListExplorerRegistryModuleVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) SetNextToken(v string) *ListExplorerRegistryModuleVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) SetRequestId(v string) *ListExplorerRegistryModuleVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) SetTotalCount(v int64) *ListExplorerRegistryModuleVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBody) Validate() error {
	if s.ExplorerRegistryModuleVersions != nil {
		for _, item := range s.ExplorerRegistryModuleVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions struct {
	// example:
	//
	// {}
	ModuleDetail map[string]interface{} `json:"moduleDetail,omitempty" xml:"moduleDetail,omitempty"`
	// example:
	//
	// {}
	ModuleFile map[string]interface{} `json:"moduleFile,omitempty" xml:"moduleFile,omitempty"`
	// example:
	//
	// eip-slb-ecs-polardb
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// namespace-test
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// {}
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties,omitempty"`
	// example:
	//
	// test_namespace/RegistryModule-test4
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// 1.5.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetModuleDetail() map[string]interface{} {
	return s.ModuleDetail
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetModuleFile() map[string]interface{} {
	return s.ModuleFile
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetSource() *string {
	return s.Source
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) GetVersion() *string {
	return s.Version
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetModuleDetail(v map[string]interface{}) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.ModuleDetail = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetModuleFile(v map[string]interface{}) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.ModuleFile = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetModuleName(v string) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetNamespaceName(v string) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.NamespaceName = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetProperties(v map[string]interface{}) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.Properties = v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetSource(v string) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.Source = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) SetVersion(v string) *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions {
	s.Version = &v
	return s
}

func (s *ListExplorerRegistryModuleVersionsResponseBodyExplorerRegistryModuleVersions) Validate() error {
	return dara.Validate(s)
}
