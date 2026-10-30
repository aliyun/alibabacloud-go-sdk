// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAifgJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrent(v int64) *ListAifgJobsRequest
	GetCurrent() *int64
	SetMaxResults(v int32) *ListAifgJobsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAifgJobsRequest
	GetNextToken() *string
	SetPageSize(v int64) *ListAifgJobsRequest
	GetPageSize() *int64
	SetRegion(v string) *ListAifgJobsRequest
	GetRegion() *string
	SetServiceName(v string) *ListAifgJobsRequest
	GetServiceName() *string
}

type ListAifgJobsRequest struct {
	// The current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Current *int64 `json:"current,omitempty" xml:"current,omitempty"`
	// The maximum number of records to retrieve at a time.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination cursor. A non-empty value indicates that more data is available.
	//
	// example:
	//
	// c2f78a783f49457caba6bace6f6f79e4
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of entries per page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The diagnostic service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// slowNode
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
}

func (s ListAifgJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAifgJobsRequest) GoString() string {
	return s.String()
}

func (s *ListAifgJobsRequest) GetCurrent() *int64 {
	return s.Current
}

func (s *ListAifgJobsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAifgJobsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAifgJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAifgJobsRequest) GetRegion() *string {
	return s.Region
}

func (s *ListAifgJobsRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAifgJobsRequest) SetCurrent(v int64) *ListAifgJobsRequest {
	s.Current = &v
	return s
}

func (s *ListAifgJobsRequest) SetMaxResults(v int32) *ListAifgJobsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAifgJobsRequest) SetNextToken(v string) *ListAifgJobsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAifgJobsRequest) SetPageSize(v int64) *ListAifgJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAifgJobsRequest) SetRegion(v string) *ListAifgJobsRequest {
	s.Region = &v
	return s
}

func (s *ListAifgJobsRequest) SetServiceName(v string) *ListAifgJobsRequest {
	s.ServiceName = &v
	return s
}

func (s *ListAifgJobsRequest) Validate() error {
	return dara.Validate(s)
}
