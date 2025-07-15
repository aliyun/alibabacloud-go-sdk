// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInventorySchemaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregator(v bool) *GetInventorySchemaRequest
	GetAggregator() *bool
	SetMaxResults(v int32) *GetInventorySchemaRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetInventorySchemaRequest
	GetNextToken() *string
	SetRegionId(v string) *GetInventorySchemaRequest
	GetRegionId() *string
	SetTypeName(v string) *GetInventorySchemaRequest
	GetTypeName() *string
}

type GetInventorySchemaRequest struct {
	// Specifies whether to return only properties that support the aggregate feature in the configuration list. Valid values:
	//
	// 	- true: only returns properties that support the aggregate feature in the configuration list.
	//
	// 	- false: returns all properties in the configuration list.
	//
	// example:
	//
	// false
	Aggregator *bool `json:"Aggregator,omitempty" xml:"Aggregator,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// gAAAAABfh8MVLQI9AuKGACLgjbsXbWs-Mna47IDM6tr6wK7TZ1
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The configuration list type name. Valid values:
	//
	// 	- ACS:InstanceInformation
	//
	// 	- ACS:Application
	//
	// 	- ACS:File
	//
	// 	- ACS:Network
	//
	// 	- ACS:WindowsRole
	//
	// 	- ACS:Service
	//
	// 	- ACS:WindowsUpdate
	//
	// 	- ACS:WindowsRegistry
	//
	// example:
	//
	// ACS:Application
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s GetInventorySchemaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInventorySchemaRequest) GoString() string {
	return s.String()
}

func (s *GetInventorySchemaRequest) GetAggregator() *bool {
	return s.Aggregator
}

func (s *GetInventorySchemaRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetInventorySchemaRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetInventorySchemaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInventorySchemaRequest) GetTypeName() *string {
	return s.TypeName
}

func (s *GetInventorySchemaRequest) SetAggregator(v bool) *GetInventorySchemaRequest {
	s.Aggregator = &v
	return s
}

func (s *GetInventorySchemaRequest) SetMaxResults(v int32) *GetInventorySchemaRequest {
	s.MaxResults = &v
	return s
}

func (s *GetInventorySchemaRequest) SetNextToken(v string) *GetInventorySchemaRequest {
	s.NextToken = &v
	return s
}

func (s *GetInventorySchemaRequest) SetRegionId(v string) *GetInventorySchemaRequest {
	s.RegionId = &v
	return s
}

func (s *GetInventorySchemaRequest) SetTypeName(v string) *GetInventorySchemaRequest {
	s.TypeName = &v
	return s
}

func (s *GetInventorySchemaRequest) Validate() error {
	return dara.Validate(s)
}
