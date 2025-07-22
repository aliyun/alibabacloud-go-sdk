// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceInspectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetInstanceInspectionsRequest
	GetEndTime() *string
	SetEngine(v string) *GetInstanceInspectionsRequest
	GetEngine() *string
	SetInstanceArea(v string) *GetInstanceInspectionsRequest
	GetInstanceArea() *string
	SetPageNo(v string) *GetInstanceInspectionsRequest
	GetPageNo() *string
	SetPageSize(v string) *GetInstanceInspectionsRequest
	GetPageSize() *string
	SetResourceGroupId(v string) *GetInstanceInspectionsRequest
	GetResourceGroupId() *string
	SetSearchMap(v string) *GetInstanceInspectionsRequest
	GetSearchMap() *string
	SetStartTime(v string) *GetInstanceInspectionsRequest
	GetStartTime() *string
}

type GetInstanceInspectionsRequest struct {
	// The end of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// >  The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1655427625000
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The database engine. Valid values:
	//
	// 	- **MySQL**
	//
	// 	- **Redis**
	//
	// 	- **PolarDBMySQL**
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The type of the instance on which the database is deployed. Valid values:
	//
	// 	- **RDS**: an Alibaba Cloud database instance.
	//
	// 	- **ECS**: an ECS instance on which a self-managed database is deployed.
	//
	// 	- **IDC**: a self-managed database instance that is not deployed on Alibaba Cloud.
	//
	// >  The value IDC specifies that the instance is deployed in a data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// RDS
	InstanceArea *string `json:"InstanceArea,omitempty" xml:"InstanceArea,omitempty"`
	// The page number. The value must be a positive integer. Default value: 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2eil6npi****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The filter condition, which can be specified in one of the following formats:
	//
	// 	- Specify the ID of a single instance in the {"InstanceId":"Instance ID"} format.
	//
	// 	- Specify the IDs of multiple instances in the {"InstanceIds":["Instance ID1","Instance ID2"]} format. Separate the instance IDs with commas (,).
	//
	// 	- Specify the region in which the instance resides in the {"region":"Region of the instance"} format.
	//
	// example:
	//
	// {"InstanceId":"rm-bp10usoc1erj7****"}
	SearchMap *string `json:"SearchMap,omitempty" xml:"SearchMap,omitempty"`
	// The beginning of the time range to query. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1655416825000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetInstanceInspectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceInspectionsRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceInspectionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetInstanceInspectionsRequest) GetEngine() *string {
	return s.Engine
}

func (s *GetInstanceInspectionsRequest) GetInstanceArea() *string {
	return s.InstanceArea
}

func (s *GetInstanceInspectionsRequest) GetPageNo() *string {
	return s.PageNo
}

func (s *GetInstanceInspectionsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *GetInstanceInspectionsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceInspectionsRequest) GetSearchMap() *string {
	return s.SearchMap
}

func (s *GetInstanceInspectionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetInstanceInspectionsRequest) SetEndTime(v string) *GetInstanceInspectionsRequest {
	s.EndTime = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetEngine(v string) *GetInstanceInspectionsRequest {
	s.Engine = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetInstanceArea(v string) *GetInstanceInspectionsRequest {
	s.InstanceArea = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetPageNo(v string) *GetInstanceInspectionsRequest {
	s.PageNo = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetPageSize(v string) *GetInstanceInspectionsRequest {
	s.PageSize = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetResourceGroupId(v string) *GetInstanceInspectionsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetSearchMap(v string) *GetInstanceInspectionsRequest {
	s.SearchMap = &v
	return s
}

func (s *GetInstanceInspectionsRequest) SetStartTime(v string) *GetInstanceInspectionsRequest {
	s.StartTime = &v
	return s
}

func (s *GetInstanceInspectionsRequest) Validate() error {
	return dara.Validate(s)
}
