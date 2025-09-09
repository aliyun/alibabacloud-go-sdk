// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDrdsDBsResponseBodyData) *DescribeDrdsDBsResponseBody
	GetData() *DescribeDrdsDBsResponseBodyData
	SetPageNumber(v string) *DescribeDrdsDBsResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeDrdsDBsResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeDrdsDBsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDBsResponseBody
	GetSuccess() *bool
	SetTotal(v string) *DescribeDrdsDBsResponseBody
	GetTotal() *string
}

type DescribeDrdsDBsResponseBody struct {
	// The list of returned databases.
	Data *DescribeDrdsDBsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of databases returned on each page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 006B7D19-8CDB-4AA6-AAE7-23C107GS3W2T
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The number of returned databases.
	//
	// example:
	//
	// 1
	Total *string `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDrdsDBsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBody) GetData() *DescribeDrdsDBsResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsDBsResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeDrdsDBsResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeDrdsDBsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDBsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDBsResponseBody) GetTotal() *string {
	return s.Total
}

func (s *DescribeDrdsDBsResponseBody) SetData(v *DescribeDrdsDBsResponseBodyData) *DescribeDrdsDBsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageNumber(v string) *DescribeDrdsDBsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetPageSize(v string) *DescribeDrdsDBsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetRequestId(v string) *DescribeDrdsDBsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetSuccess(v bool) *DescribeDrdsDBsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) SetTotal(v string) *DescribeDrdsDBsResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDrdsDBsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBsResponseBodyData struct {
	Db []*DescribeDrdsDBsResponseBodyDataDb `json:"Db,omitempty" xml:"Db,omitempty" type:"Repeated"`
}

func (s DescribeDrdsDBsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyData) GetDb() []*DescribeDrdsDBsResponseBodyDataDb {
	return s.Db
}

func (s *DescribeDrdsDBsResponseBodyData) SetDb(v []*DescribeDrdsDBsResponseBodyDataDb) *DescribeDrdsDBsResponseBodyData {
	s.Db = v
	return s
}

func (s *DescribeDrdsDBsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBsResponseBodyDataDb struct {
	// The time when the database is created. The value of this parameter is a UNIX timestamp. Unit: ms.
	//
	// example:
	//
	// 1563773824000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The type of the database. Valid values: **RDS*	- and **POLARDB**.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// drds_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The partitioning mode of the database. Valid values:
	//
	// 	- **HORIZONTAL**: The database is horizontally partitioned.
	//
	// 	- **VERTICAL**: The database is vertically partitioned.
	//
	// example:
	//
	// HORIZONTAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The schema ID that is assigned to the partitioned database.
	//
	// example:
	//
	// drds_test_1563773871118kxqd
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// The state of the database.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBsResponseBodyDataDb) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBsResponseBodyDataDb) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetMode() *string {
	return s.Mode
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetSchema() *string {
	return s.Schema
}

func (s *DescribeDrdsDBsResponseBodyDataDb) GetStatus() *string {
	return s.Status
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetCreateTime(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbInstType(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetDbName(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetMode(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetSchema(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) SetStatus(v string) *DescribeDrdsDBsResponseBodyDataDb {
	s.Status = &v
	return s
}

func (s *DescribeDrdsDBsResponseBodyDataDb) Validate() error {
	return dara.Validate(s)
}
