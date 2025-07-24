// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody
	GetApplications() []*ListApplicationsResponseBodyApplications
	SetMaxResults(v int32) *ListApplicationsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListApplicationsResponseBody
	GetTotalCount() *int32
}

type ListApplicationsResponseBody struct {
	// The applications.
	Applications []*ListApplicationsResponseBodyApplications `json:"Applications,omitempty" xml:"Applications,omitempty" type:"Repeated"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// 2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBody) GetApplications() []*ListApplicationsResponseBodyApplications {
	return s.Applications
}

func (s *ListApplicationsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListApplicationsResponseBody) SetApplications(v []*ListApplicationsResponseBodyApplications) *ListApplicationsResponseBody {
	s.Applications = v
	return s
}

func (s *ListApplicationsResponseBody) SetMaxResults(v int32) *ListApplicationsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsResponseBody) SetNextToken(v string) *ListApplicationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsResponseBody) SetRequestId(v string) *ListApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsResponseBody) SetTotalCount(v int32) *ListApplicationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListApplicationsResponseBodyApplications struct {
	// The application name.
	//
	// example:
	//
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The status of the applications. Valid values:
	//
	// 	- STOPPED: At least one application is in the Stopped state.
	//
	// 	- RUNNING: All applications are in the Running state.
	//
	// This parameter is returned only for DataLake, OLAP, Dataflow, DataServing, and custom clusters. For other types of clusters, no value is returned for this parameter.
	//
	// example:
	//
	// RUNNING
	ApplicationState *string `json:"ApplicationState,omitempty" xml:"ApplicationState,omitempty"`
	// Deprecated
	//
	// The version of the application.
	//
	// example:
	//
	// 2.8.1
	ApplicationVersion *string `json:"ApplicationVersion,omitempty" xml:"ApplicationVersion,omitempty"`
	// The community edition.
	//
	// example:
	//
	// 2.8.5
	CommunityVersion *string `json:"CommunityVersion,omitempty" xml:"CommunityVersion,omitempty"`
}

func (s ListApplicationsResponseBodyApplications) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsResponseBodyApplications) GoString() string {
	return s.String()
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationState() *string {
	return s.ApplicationState
}

func (s *ListApplicationsResponseBodyApplications) GetApplicationVersion() *string {
	return s.ApplicationVersion
}

func (s *ListApplicationsResponseBodyApplications) GetCommunityVersion() *string {
	return s.CommunityVersion
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationName(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationState(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationState = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetApplicationVersion(v string) *ListApplicationsResponseBodyApplications {
	s.ApplicationVersion = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) SetCommunityVersion(v string) *ListApplicationsResponseBodyApplications {
	s.CommunityVersion = &v
	return s
}

func (s *ListApplicationsResponseBodyApplications) Validate() error {
	return dara.Validate(s)
}
