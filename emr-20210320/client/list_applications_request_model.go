// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationNames(v []*string) *ListApplicationsRequest
	GetApplicationNames() []*string
	SetClusterId(v string) *ListApplicationsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListApplicationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListApplicationsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListApplicationsRequest
	GetRegionId() *string
}

type ListApplicationsRequest struct {
	// The application names.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ApplicationNames []*string `json:"ApplicationNames,omitempty" xml:"ApplicationNames,omitempty" type:"Repeated"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e6a9d46e92675****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The page number of the next page returned.
	//
	// example:
	//
	// 0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApplicationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationsRequest) GetApplicationNames() []*string {
	return s.ApplicationNames
}

func (s *ListApplicationsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListApplicationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListApplicationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApplicationsRequest) SetApplicationNames(v []*string) *ListApplicationsRequest {
	s.ApplicationNames = v
	return s
}

func (s *ListApplicationsRequest) SetClusterId(v string) *ListApplicationsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListApplicationsRequest) SetMaxResults(v int32) *ListApplicationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListApplicationsRequest) SetNextToken(v string) *ListApplicationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsRequest) SetRegionId(v string) *ListApplicationsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApplicationsRequest) Validate() error {
	return dara.Validate(s)
}
