// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHiveDatabasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHiveDatabasesRequest
	GetClusterId() *string
	SetDatabaseNames(v []*string) *ListDoctorHiveDatabasesRequest
	GetDatabaseNames() []*string
	SetDateTime(v string) *ListDoctorHiveDatabasesRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorHiveDatabasesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHiveDatabasesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHiveDatabasesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHiveDatabasesRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHiveDatabasesRequest
	GetRegionId() *string
}

type ListDoctorHiveDatabasesRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database names.
	//
	// example:
	//
	// null
	DatabaseNames []*string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty" type:"Repeated"`
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
	// The basis on which you want to sort the query results. Valid values:
	//
	// 	- tableCount: the number of tables
	//
	// example:
	//
	// tableCount
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// The order in which you want to sort the query results. Valid values:
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
}

func (s ListDoctorHiveDatabasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHiveDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHiveDatabasesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHiveDatabasesRequest) GetDatabaseNames() []*string {
	return s.DatabaseNames
}

func (s *ListDoctorHiveDatabasesRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHiveDatabasesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHiveDatabasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHiveDatabasesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHiveDatabasesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHiveDatabasesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHiveDatabasesRequest) SetClusterId(v string) *ListDoctorHiveDatabasesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetDatabaseNames(v []*string) *ListDoctorHiveDatabasesRequest {
	s.DatabaseNames = v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetDateTime(v string) *ListDoctorHiveDatabasesRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetMaxResults(v int32) *ListDoctorHiveDatabasesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetNextToken(v string) *ListDoctorHiveDatabasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetOrderBy(v string) *ListDoctorHiveDatabasesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetOrderType(v string) *ListDoctorHiveDatabasesRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) SetRegionId(v string) *ListDoctorHiveDatabasesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHiveDatabasesRequest) Validate() error {
	return dara.Validate(s)
}
