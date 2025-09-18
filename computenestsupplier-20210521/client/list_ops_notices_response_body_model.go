// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsNoticesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListOpsNoticesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpsNoticesResponseBody
	GetNextToken() *string
	SetOpsNotices(v []*ListOpsNoticesResponseBodyOpsNotices) *ListOpsNoticesResponseBody
	GetOpsNotices() []*ListOpsNoticesResponseBodyOpsNotices
	SetRequestId(v string) *ListOpsNoticesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListOpsNoticesResponseBody
	GetTotalCount() *int32
}

type ListOpsNoticesResponseBody struct {
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OpsNotices []*ListOpsNoticesResponseBodyOpsNotices `json:"OpsNotices,omitempty" xml:"OpsNotices,omitempty" type:"Repeated"`
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListOpsNoticesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOpsNoticesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOpsNoticesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpsNoticesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpsNoticesResponseBody) GetOpsNotices() []*ListOpsNoticesResponseBodyOpsNotices {
	return s.OpsNotices
}

func (s *ListOpsNoticesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOpsNoticesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListOpsNoticesResponseBody) SetMaxResults(v int32) *ListOpsNoticesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListOpsNoticesResponseBody) SetNextToken(v string) *ListOpsNoticesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListOpsNoticesResponseBody) SetOpsNotices(v []*ListOpsNoticesResponseBodyOpsNotices) *ListOpsNoticesResponseBody {
	s.OpsNotices = v
	return s
}

func (s *ListOpsNoticesResponseBody) SetRequestId(v string) *ListOpsNoticesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOpsNoticesResponseBody) SetTotalCount(v int32) *ListOpsNoticesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListOpsNoticesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListOpsNoticesResponseBodyOpsNotices struct {
	// example:
	//
	// {"cveId":"CVE-2021-4034"}
	Attributes *string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// Security
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// example:
	//
	// content
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// notice-1749693290
	NoticeId *string `json:"NoticeId,omitempty" xml:"NoticeId,omitempty"`
	// example:
	//
	// service-e10349089de34exxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// test
	ServiceName     *string   `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	ServiceVersions []*string `json:"ServiceVersions,omitempty" xml:"ServiceVersions,omitempty" type:"Repeated"`
	// example:
	//
	// High
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// You need to upgrade service instance
	Solutions *string `json:"Solutions,omitempty" xml:"Solutions,omitempty"`
	// example:
	//
	// 2022-10-12T02:03:37Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// Vulnerability
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListOpsNoticesResponseBodyOpsNotices) String() string {
	return dara.Prettify(s)
}

func (s ListOpsNoticesResponseBodyOpsNotices) GoString() string {
	return s.String()
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetAttributes() *string {
	return s.Attributes
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetCategory() *string {
	return s.Category
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetContent() *string {
	return s.Content
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetNoticeId() *string {
	return s.NoticeId
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetServiceVersions() []*string {
	return s.ServiceVersions
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetSeverity() *string {
	return s.Severity
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetSolutions() *string {
	return s.Solutions
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetSuccess() *string {
	return s.Success
}

func (s *ListOpsNoticesResponseBodyOpsNotices) GetType() *string {
	return s.Type
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetAttributes(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Attributes = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetCategory(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Category = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetContent(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Content = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetNoticeId(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.NoticeId = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetServiceId(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.ServiceId = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetServiceName(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.ServiceName = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetServiceVersions(v []*string) *ListOpsNoticesResponseBodyOpsNotices {
	s.ServiceVersions = v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetSeverity(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Severity = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetSolutions(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Solutions = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetStartTime(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.StartTime = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetSuccess(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Success = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) SetType(v string) *ListOpsNoticesResponseBodyOpsNotices {
	s.Type = &v
	return s
}

func (s *ListOpsNoticesResponseBodyOpsNotices) Validate() error {
	return dara.Validate(s)
}
