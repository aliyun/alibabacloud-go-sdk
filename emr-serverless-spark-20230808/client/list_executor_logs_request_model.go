// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogType(v string) *ListExecutorLogsRequest
	GetLogType() *string
	SetMaxResults(v int32) *ListExecutorLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListExecutorLogsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListExecutorLogsRequest
	GetRegionId() *string
}

type ListExecutorLogsRequest struct {
	// The type of log. Only logs of this type are returned.
	//
	// example:
	//
	// log4j
	LogType *string `json:"logType,omitempty" xml:"logType,omitempty"`
	// The maximum number of results to return per page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. This is the `nextToken` from a previously truncated response.
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListExecutorLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorLogsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorLogsRequest) GetLogType() *string {
	return s.LogType
}

func (s *ListExecutorLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListExecutorLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutorLogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListExecutorLogsRequest) SetLogType(v string) *ListExecutorLogsRequest {
	s.LogType = &v
	return s
}

func (s *ListExecutorLogsRequest) SetMaxResults(v int32) *ListExecutorLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListExecutorLogsRequest) SetNextToken(v string) *ListExecutorLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListExecutorLogsRequest) SetRegionId(v string) *ListExecutorLogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListExecutorLogsRequest) Validate() error {
	return dara.Validate(s)
}
