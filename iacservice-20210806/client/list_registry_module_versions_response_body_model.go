// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegistryModuleVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int64) *ListRegistryModuleVersionsResponseBody
	GetCount() *int64
	SetMaxResults(v int32) *ListRegistryModuleVersionsResponseBody
	GetMaxResults() *int32
	SetModuleVersions(v []*ListRegistryModuleVersionsResponseBodyModuleVersions) *ListRegistryModuleVersionsResponseBody
	GetModuleVersions() []*ListRegistryModuleVersionsResponseBodyModuleVersions
	SetNextToken(v string) *ListRegistryModuleVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListRegistryModuleVersionsResponseBody
	GetRequestId() *string
}

type ListRegistryModuleVersionsResponseBody struct {
	// example:
	//
	// 21
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// 20
	MaxResults     *int32                                                  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	ModuleVersions []*ListRegistryModuleVersionsResponseBodyModuleVersions `json:"moduleVersions,omitempty" xml:"moduleVersions,omitempty" type:"Repeated"`
	// example:
	//
	// IbuvZbAXFOiB4nKg8iOH447bhHWDavGTOMijI2Jep7c=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 8606B880-3485-54E2-89E1-43361C468C85
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRegistryModuleVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModuleVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegistryModuleVersionsResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *ListRegistryModuleVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListRegistryModuleVersionsResponseBody) GetModuleVersions() []*ListRegistryModuleVersionsResponseBodyModuleVersions {
	return s.ModuleVersions
}

func (s *ListRegistryModuleVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListRegistryModuleVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegistryModuleVersionsResponseBody) SetCount(v int64) *ListRegistryModuleVersionsResponseBody {
	s.Count = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBody) SetMaxResults(v int32) *ListRegistryModuleVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBody) SetModuleVersions(v []*ListRegistryModuleVersionsResponseBodyModuleVersions) *ListRegistryModuleVersionsResponseBody {
	s.ModuleVersions = v
	return s
}

func (s *ListRegistryModuleVersionsResponseBody) SetNextToken(v string) *ListRegistryModuleVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBody) SetRequestId(v string) *ListRegistryModuleVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBody) Validate() error {
	if s.ModuleVersions != nil {
		for _, item := range s.ModuleVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRegistryModuleVersionsResponseBodyModuleVersions struct {
	CreateTime    *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	ModuleName    *string `json:"moduleName,omitempty" xml:"moduleName,omitempty"`
	NamespaceName *string `json:"namespaceName,omitempty" xml:"namespaceName,omitempty"`
	Provider      *string `json:"provider,omitempty" xml:"provider,omitempty"`
	Source        *string `json:"source,omitempty" xml:"source,omitempty"`
	SourceUrl     *string `json:"sourceUrl,omitempty" xml:"sourceUrl,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListRegistryModuleVersionsResponseBodyModuleVersions) String() string {
	return dara.Prettify(s)
}

func (s ListRegistryModuleVersionsResponseBodyModuleVersions) GoString() string {
	return s.String()
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetModuleName() *string {
	return s.ModuleName
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetProvider() *string {
	return s.Provider
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetSource() *string {
	return s.Source
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetSourceUrl() *string {
	return s.SourceUrl
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) GetVersion() *string {
	return s.Version
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetCreateTime(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.CreateTime = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetModuleName(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.ModuleName = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetNamespaceName(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.NamespaceName = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetProvider(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.Provider = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetSource(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.Source = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetSourceUrl(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.SourceUrl = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) SetVersion(v string) *ListRegistryModuleVersionsResponseBodyModuleVersions {
	s.Version = &v
	return s
}

func (s *ListRegistryModuleVersionsResponseBodyModuleVersions) Validate() error {
	return dara.Validate(s)
}
