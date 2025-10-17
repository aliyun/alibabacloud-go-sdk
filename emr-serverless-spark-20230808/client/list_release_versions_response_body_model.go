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
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The versions.
	ReleaseVersions []*ListReleaseVersionsResponseBodyReleaseVersions `json:"releaseVersions,omitempty" xml:"releaseVersions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
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
	if s.ReleaseVersions != nil {
		for _, item := range s.ReleaseVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListReleaseVersionsResponseBodyReleaseVersions struct {
	// The version number of open source Spark.
	//
	// example:
	//
	// Spark 3.3.1
	CommunityVersion *string `json:"communityVersion,omitempty" xml:"communityVersion,omitempty"`
	// The CPU architectures.
	CpuArchitectures []*string `json:"cpuArchitectures,omitempty" xml:"cpuArchitectures,omitempty" type:"Repeated"`
	// The version number.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12)
	DisplayReleaseVersion *string `json:"displayReleaseVersion,omitempty" xml:"displayReleaseVersion,omitempty"`
	// Indicates whether the Fusion engine is used for acceleration.
	//
	// example:
	//
	// true
	Fusion *bool `json:"fusion,omitempty" xml:"fusion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1716215854101
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The type of the Infrastructure as a Service (IaaS) layer.
	//
	// example:
	//
	// ASI
	IaasType *string `json:"iaasType,omitempty" xml:"iaasType,omitempty"`
	// The version number.
	//
	// example:
	//
	// esr-2.1 (Spark 3.3.1, Scala 2.12, Java Runtime)
	ReleaseVersion *string `json:"releaseVersion,omitempty" xml:"releaseVersion,omitempty"`
	// The version of Scala.
	//
	// example:
	//
	// 2.12
	ScalaVersion *string `json:"scalaVersion,omitempty" xml:"scalaVersion,omitempty"`
	// The status of the version.
	//
	// example:
	//
	// ONLINE
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// The type of the version.
	//
	// example:
	//
	// stable
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) String() string {
	return dara.Prettify(s)
}

func (s ListReleaseVersionsResponseBodyReleaseVersions) GoString() string {
	return s.String()
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetCommunityVersion() *string {
	return s.CommunityVersion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetCpuArchitectures() []*string {
	return s.CpuArchitectures
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetDisplayReleaseVersion() *string {
	return s.DisplayReleaseVersion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetFusion() *bool {
	return s.Fusion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetIaasType() *string {
	return s.IaasType
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetReleaseVersion() *string {
	return s.ReleaseVersion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetScalaVersion() *string {
	return s.ScalaVersion
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetState() *string {
	return s.State
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) GetType() *string {
	return s.Type
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetCommunityVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.CommunityVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetCpuArchitectures(v []*string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.CpuArchitectures = v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetDisplayReleaseVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.DisplayReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetFusion(v bool) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.Fusion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetGmtCreate(v int64) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.GmtCreate = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetIaasType(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.IaasType = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetReleaseVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.ReleaseVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetScalaVersion(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.ScalaVersion = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetState(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.State = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) SetType(v string) *ListReleaseVersionsResponseBodyReleaseVersions {
	s.Type = &v
	return s
}

func (s *ListReleaseVersionsResponseBodyReleaseVersions) Validate() error {
	return dara.Validate(s)
}
