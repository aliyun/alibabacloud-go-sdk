// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListDoctorHBaseTablesRequest
	GetClusterId() *string
	SetDateTime(v string) *ListDoctorHBaseTablesRequest
	GetDateTime() *string
	SetMaxResults(v int32) *ListDoctorHBaseTablesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDoctorHBaseTablesRequest
	GetNextToken() *string
	SetOrderBy(v string) *ListDoctorHBaseTablesRequest
	GetOrderBy() *string
	SetOrderType(v string) *ListDoctorHBaseTablesRequest
	GetOrderType() *string
	SetRegionId(v string) *ListDoctorHBaseTablesRequest
	GetRegionId() *string
	SetTableNames(v []*string) *ListDoctorHBaseTablesRequest
	GetTableNames() []*string
}

type ListDoctorHBaseTablesRequest struct {
	// The ID of the cluster.
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
	// The maximum number of entries that are returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Marks the current position to start reading. If this field is empty, the data is read from the beginning.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C89568980
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The field that you use to sort the query results.
	//
	// Valid values:
	//
	// 	- tableSize
	//
	// example:
	//
	// tableSize
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
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table names, which are used to filter the query results.
	//
	// example:
	//
	// null
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s ListDoctorHBaseTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseTablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDoctorHBaseTablesRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *ListDoctorHBaseTablesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDoctorHBaseTablesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDoctorHBaseTablesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListDoctorHBaseTablesRequest) GetOrderType() *string {
	return s.OrderType
}

func (s *ListDoctorHBaseTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDoctorHBaseTablesRequest) GetTableNames() []*string {
	return s.TableNames
}

func (s *ListDoctorHBaseTablesRequest) SetClusterId(v string) *ListDoctorHBaseTablesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetDateTime(v string) *ListDoctorHBaseTablesRequest {
	s.DateTime = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetMaxResults(v int32) *ListDoctorHBaseTablesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetNextToken(v string) *ListDoctorHBaseTablesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetOrderBy(v string) *ListDoctorHBaseTablesRequest {
	s.OrderBy = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetOrderType(v string) *ListDoctorHBaseTablesRequest {
	s.OrderType = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetRegionId(v string) *ListDoctorHBaseTablesRequest {
	s.RegionId = &v
	return s
}

func (s *ListDoctorHBaseTablesRequest) SetTableNames(v []*string) *ListDoctorHBaseTablesRequest {
	s.TableNames = v
	return s
}

func (s *ListDoctorHBaseTablesRequest) Validate() error {
	return dara.Validate(s)
}
