// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseRegionServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHBaseRegionServersRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHBaseRegionServersRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorHBaseRegionServersRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHBaseRegionServersRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHBaseRegionServersRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHBaseRegionServersRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHBaseRegionServersRequest
	GetRegionId() *string
	SetRegionServerHosts(v []*string) *ListDoctorHBaseRegionServersRequest
	GetRegionServerHosts() []*string
}

type ListDoctorHBaseRegionServersRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The query date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The maximum number of entries to return on each page.
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
	// The field that you use to sort the query results. Valid value:
	//
	// 	- regionCount: the number of regions.
	//
	// example:
	//
	// regionCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid value:
	//
	// 	- ASC: in ascending order
	//
	// 	- DESC: in descending order
	//
	// example:
	//
	// ASC
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The RegionServer hosts.
	//
	// example:
	//
	// null
	RegionServerHosts []*string `json:"RegionServerHosts,omitempty" xml:"RegionServerHosts,omitempty" type:"Repeated"`
}

func (s ListDoctorHBaseRegionServersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHBaseRegionServersRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHBaseRegionServersRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHBaseRegionServersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHBaseRegionServersRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHBaseRegionServersRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHBaseRegionServersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHBaseRegionServersRequest) GetRegionServerHosts() []*string {
	return s.RegionServerHosts
}

func (s *ListDoctorHBaseRegionServersRequest) SetClusterId(v string) *ListDoctorHBaseRegionServersRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetDateTime(v string) *ListDoctorHBaseRegionServersRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetMaxResults(v int32) *ListDoctorHBaseRegionServersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetNextToken(v string) *ListDoctorHBaseRegionServersRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetOrderBy(v string) *ListDoctorHBaseRegionServersRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetOrderType(v string) *ListDoctorHBaseRegionServersRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetRegionId(v string) *ListDoctorHBaseRegionServersRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) SetRegionServerHosts(v []*string) *ListDoctorHBaseRegionServersRequest {
	s.RegionServerHosts = v
	return s
}

func (s *ListDoctorHBaseRegionServersRequest) Validate() error {
	return dara.Validate(s)
}
