// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorReportsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListDoctorReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorReportsRequest
	GetNextToken() *string
	SetRegionId(v string) *ListDoctorReportsRequest
	GetRegionId() *string
}

type ListDoctorReportsRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
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

func (s ListDoctorReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorReportsRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorReportsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorReportsRequest) SetClusterId(v string) *ListDoctorReportsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorReportsRequest) SetMaxResults(v int32) *ListDoctorReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorReportsRequest) SetNextToken(v string) *ListDoctorReportsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorReportsRequest) SetRegionId(v string) *ListDoctorReportsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorReportsRequest) Validate() error {
	return dara.Validate(s)
}
