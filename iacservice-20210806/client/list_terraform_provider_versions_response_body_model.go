// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTerraformProviderVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListTerraformProviderVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListTerraformProviderVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTerraformProviderVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListTerraformProviderVersionsResponseBody
	GetTotalCount() *int32
	SetVersions(v []*ListTerraformProviderVersionsResponseBodyVersions) *ListTerraformProviderVersionsResponseBody
	GetVersions() []*ListTerraformProviderVersionsResponseBodyVersions
}

type ListTerraformProviderVersionsResponseBody struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 73588ebb-9d40-4660-a59f-764636ae6034
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32                                               `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	Versions   []*ListTerraformProviderVersionsResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListTerraformProviderVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTerraformProviderVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTerraformProviderVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTerraformProviderVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTerraformProviderVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListTerraformProviderVersionsResponseBody) GetVersions() []*ListTerraformProviderVersionsResponseBodyVersions {
	return s.Versions
}

func (s *ListTerraformProviderVersionsResponseBody) SetMaxResults(v int32) *ListTerraformProviderVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetNextToken(v string) *ListTerraformProviderVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetRequestId(v string) *ListTerraformProviderVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetTotalCount(v int32) *ListTerraformProviderVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) SetVersions(v []*ListTerraformProviderVersionsResponseBodyVersions) *ListTerraformProviderVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListTerraformProviderVersionsResponseBody) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTerraformProviderVersionsResponseBodyVersions struct {
	PublishedTime *string `json:"publishedTime,omitempty" xml:"publishedTime,omitempty"`
	Status        *string `json:"status,omitempty" xml:"status,omitempty"`
	Version       *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListTerraformProviderVersionsResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListTerraformProviderVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) GetPublishedTime() *string {
	return s.PublishedTime
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) GetStatus() *string {
	return s.Status
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) GetVersion() *string {
	return s.Version
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetPublishedTime(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.PublishedTime = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetStatus(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.Status = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) SetVersion(v string) *ListTerraformProviderVersionsResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *ListTerraformProviderVersionsResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
