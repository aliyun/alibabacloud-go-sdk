// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExplorerRegistryModulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExplorerRegistryModules(v []*ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) *ListExplorerRegistryModulesResponseBody
	GetExplorerRegistryModules() []*ListExplorerRegistryModulesResponseBodyExplorerRegistryModules
	SetMaxResults(v int32) *ListExplorerRegistryModulesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListExplorerRegistryModulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExplorerRegistryModulesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListExplorerRegistryModulesResponseBody
	GetTotalCount() *int64
}

type ListExplorerRegistryModulesResponseBody struct {
	ExplorerRegistryModules []*ListExplorerRegistryModulesResponseBodyExplorerRegistryModules `json:"explorerRegistryModules,omitempty" xml:"explorerRegistryModules,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// DxEkv+3w0EDAQgcRFBp8Ep4dD+2BRJj42DLT6GrZysw=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1A662F56-CA76-55F6-869D-7F26293B8E67
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 170
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListExplorerRegistryModulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModulesResponseBody) GetExplorerRegistryModules() []*ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	return s.ExplorerRegistryModules
}

func (s *ListExplorerRegistryModulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExplorerRegistryModulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExplorerRegistryModulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExplorerRegistryModulesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListExplorerRegistryModulesResponseBody) SetExplorerRegistryModules(v []*ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) *ListExplorerRegistryModulesResponseBody {
	s.ExplorerRegistryModules = v
	return s
}

func (s *ListExplorerRegistryModulesResponseBody) SetMaxResults(v int32) *ListExplorerRegistryModulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBody) SetNextToken(v string) *ListExplorerRegistryModulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBody) SetRequestId(v string) *ListExplorerRegistryModulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBody) SetTotalCount(v int64) *ListExplorerRegistryModulesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBody) Validate() error {
	if s.ExplorerRegistryModules != nil {
		for _, item := range s.ExplorerRegistryModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExplorerRegistryModulesResponseBodyExplorerRegistryModules struct {
	// example:
	//
	// demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 34
	Downloads *int64 `json:"downloads,omitempty" xml:"downloads,omitempty"`
	// example:
	//
	// v1
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// example:
	//
	// terraform-alicloud-modules/mongodb
	ModuleName *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	// example:
	//
	// test_namespace
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	// example:
	//
	// terraform-alicloud-modules/mongodb/alicloud
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// Default
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) String() string {
	return dara.Prettify(s)
}

func (s ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GoString() string {
	return s.String()
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetDescription() *string {
	return s.Description
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetDownloads() *int64 {
	return s.Downloads
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetSource() *string {
	return s.Source
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) GetStatus() *string {
	return s.Status
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetDescription(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.Description = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetDownloads(v int64) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.Downloads = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetLatestVersion(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.LatestVersion = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetModuleName(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.ModuleName = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetNamespaceName(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.NamespaceName = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetSource(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.Source = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) SetStatus(v string) *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules {
	s.Status = &v
	return s
}

func (s *ListExplorerRegistryModulesResponseBodyExplorerRegistryModules) Validate() error {
	return dara.Validate(s)
}
