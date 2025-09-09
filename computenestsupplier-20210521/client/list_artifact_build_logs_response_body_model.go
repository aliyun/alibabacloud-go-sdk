// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactBuildLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuildLogs(v []*ListArtifactBuildLogsResponseBodyBuildLogs) *ListArtifactBuildLogsResponseBody
	GetBuildLogs() []*ListArtifactBuildLogsResponseBodyBuildLogs
	SetMaxResults(v int32) *ListArtifactBuildLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactBuildLogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListArtifactBuildLogsResponseBody
	GetRequestId() *string
}

type ListArtifactBuildLogsResponseBody struct {
	BuildLogs []*ListArtifactBuildLogsResponseBodyBuildLogs `json:"BuildLogs,omitempty" xml:"BuildLogs,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAVEKMJSB4yFi/EJc7fOJCkw=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 708AB976-F69C-5727-BED9-8C39D878B93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListArtifactBuildLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildLogsResponseBody) GetBuildLogs() []*ListArtifactBuildLogsResponseBodyBuildLogs {
	return s.BuildLogs
}

func (s *ListArtifactBuildLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactBuildLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactBuildLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactBuildLogsResponseBody) SetBuildLogs(v []*ListArtifactBuildLogsResponseBodyBuildLogs) *ListArtifactBuildLogsResponseBody {
	s.BuildLogs = v
	return s
}

func (s *ListArtifactBuildLogsResponseBody) SetMaxResults(v int32) *ListArtifactBuildLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactBuildLogsResponseBody) SetNextToken(v string) *ListArtifactBuildLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactBuildLogsResponseBody) SetRequestId(v string) *ListArtifactBuildLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactBuildLogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListArtifactBuildLogsResponseBodyBuildLogs struct {
	// example:
	//
	// NDAx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1738894304
	Timestamp *string `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s ListArtifactBuildLogsResponseBodyBuildLogs) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactBuildLogsResponseBodyBuildLogs) GoString() string {
	return s.String()
}

func (s *ListArtifactBuildLogsResponseBodyBuildLogs) GetContent() *string {
	return s.Content
}

func (s *ListArtifactBuildLogsResponseBodyBuildLogs) GetTimestamp() *string {
	return s.Timestamp
}

func (s *ListArtifactBuildLogsResponseBodyBuildLogs) SetContent(v string) *ListArtifactBuildLogsResponseBodyBuildLogs {
	s.Content = &v
	return s
}

func (s *ListArtifactBuildLogsResponseBodyBuildLogs) SetTimestamp(v string) *ListArtifactBuildLogsResponseBodyBuildLogs {
	s.Timestamp = &v
	return s
}

func (s *ListArtifactBuildLogsResponseBodyBuildLogs) Validate() error {
	return dara.Validate(s)
}
