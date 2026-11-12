// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDatasetVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDatasetVersionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDatasetVersionsResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListDatasetVersionsResponseBody
	GetTotal() *int32
	SetVersions(v []*ListDatasetVersionsResponseBodyVersions) *ListDatasetVersionsResponseBody
	GetVersions() []*ListDatasetVersionsResponseBodyVersions
}

type ListDatasetVersionsResponseBody struct {
	// The maximum number of versions per page used for this query, which is not the actual number of versions returned on the current page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page. This parameter is returned when a next page exists. To query the next page, pass this value as is to nextToken. If this parameter is not returned or is empty, no next page exists. The example is for demonstration purposes only.
	//
	// example:
	//
	// umaQfI7x758Ns4TgWrj8yA3fYlnk7dJgsfhMrSViRY8=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique identifier of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// D17DE39E-6C62-50E3-9EB7-FDE41BB0D43D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of versions of the dataset, excluding deleted versions. This value is not limited by the current page size.
	//
	// example:
	//
	// 35
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
	// The list of historical versions of the dataset on the current page. An empty array is returned if no versions meet the conditions.
	Versions []*ListDatasetVersionsResponseBodyVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListDatasetVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDatasetVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDatasetVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetVersionsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListDatasetVersionsResponseBody) GetVersions() []*ListDatasetVersionsResponseBodyVersions {
	return s.Versions
}

func (s *ListDatasetVersionsResponseBody) SetMaxResults(v int32) *ListDatasetVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetNextToken(v string) *ListDatasetVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetRequestId(v string) *ListDatasetVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetTotal(v int32) *ListDatasetVersionsResponseBody {
	s.Total = &v
	return s
}

func (s *ListDatasetVersionsResponseBody) SetVersions(v []*ListDatasetVersionsResponseBodyVersions) *ListDatasetVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListDatasetVersionsResponseBody) Validate() error {
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

type ListDatasetVersionsResponseBodyVersions struct {
	// The creation time of the version, in RFC 3339 format, including time zone information. For example, the Z in 2026-11-11T06:00:00Z indicates UTC.
	//
	// example:
	//
	// 2026-11-11T06:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description provided when the version was created. If no description was provided, an empty string is returned.
	//
	// example:
	//
	// The dataset version used for the first evaluation
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the historical version, which corresponds to the name specified when the version was created. This value can be used as the version parameter for DeleteDatasetVersion.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListDatasetVersionsResponseBodyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetVersionsResponseBodyVersions) GoString() string {
	return s.String()
}

func (s *ListDatasetVersionsResponseBodyVersions) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDatasetVersionsResponseBodyVersions) GetDescription() *string {
	return s.Description
}

func (s *ListDatasetVersionsResponseBodyVersions) GetVersion() *string {
	return s.Version
}

func (s *ListDatasetVersionsResponseBodyVersions) SetCreateTime(v string) *ListDatasetVersionsResponseBodyVersions {
	s.CreateTime = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyVersions) SetDescription(v string) *ListDatasetVersionsResponseBodyVersions {
	s.Description = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyVersions) SetVersion(v string) *ListDatasetVersionsResponseBodyVersions {
	s.Version = &v
	return s
}

func (s *ListDatasetVersionsResponseBodyVersions) Validate() error {
	return dara.Validate(s)
}
