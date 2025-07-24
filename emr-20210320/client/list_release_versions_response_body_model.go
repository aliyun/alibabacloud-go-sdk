// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReleaseVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListReleaseVersionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListReleaseVersionsResponseBody
	GetNextToken() *string
	SetReleaseVersions(v []*ListReleaseVersionsResponseBodyReleaseVersions) *ListReleaseVersionsResponseBody
	GetReleaseVersions() []*ListReleaseVersionsResponseBodyReleaseVersions
	SetRequestId(v string) *ListReleaseVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListReleaseVersionsResponseBody
	GetTotalCount() *int32
}

type ListReleaseVersionsResponseBody struct {
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Returns the location of the data that was read.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The major EMR versions.
	ReleaseVersions []*ListReleaseVersionsResponseBodyReleaseVersions `json:"ReleaseVersions,omitempty" xml:"ReleaseVersions,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListReleaseVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListReleaseVersionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListReleaseVersionsResponseBody) GetReleaseVersions() []*ListReleaseVersionsResponseBodyReleaseVersions {
	return s.ReleaseVersions
}

func (s *ListReleaseVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListReleaseVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListReleaseVersionsResponseBody) SetMaxResults(v int32) *ListReleaseVersionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetNextToken(v string) *ListReleaseVersionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetReleaseVersions(v []*ListReleaseVersionsResponseBodyReleaseVersions) *ListReleaseVersionsResponseBody {
	s.ReleaseVersions = v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetRequestId(v string) *ListReleaseVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) SetTotalCount(v int32) *ListReleaseVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListReleaseVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListReleaseVersionsResponseBodyReleaseVersions struct {
	// The IaaS type.
	//
	// example:
	//
	// ECS
	IaasType *string `json:"IaasType,omitempty" xml:"IaasType,omitempty"`
	// The EMR version.
	//
	// example:
	//
	// EMR-5.3.0
	ReleaseVersion *string `json:"ReleaseVersion,omitempty" xml:"ReleaseVersion,omitempty"`
	// The version series.
	//
	// example:
	//
	// EMR-6.X
	Series *string `json:"Series,omitempty" xml:"Series,omitempty"`
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetIaasType() *string {
	return s.IaasType
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetSeries() *string {
	return s.Series
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetIaasType(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.IaasType = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetReleaseVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.ReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetSeries(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.Series = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) Validate() error {
	return dara.Validate(s)
}
