// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDrdsDBResponseBodyData) *DescribeDrdsDBResponseBody
	GetData() *DescribeDrdsDBResponseBodyData
	SetRequestId(v string) *DescribeDrdsDBResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsDBResponseBody
	GetSuccess() *bool
}

type DescribeDrdsDBResponseBody struct {
	// Indicates the details about the database.
	Data *DescribeDrdsDBResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the ID of the request.
	//
	// example:
	//
	// 58FB0EC7-CF71-4E48-92FB-CF070D******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBody) GetData() *DescribeDrdsDBResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsDBResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsDBResponseBody) SetData(v *DescribeDrdsDBResponseBodyData) *DescribeDrdsDBResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetRequestId(v string) *DescribeDrdsDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) SetSuccess(v bool) *DescribeDrdsDBResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsDBResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsDBResponseBodyData struct {
	// Indicates the time when the database was created. The value is in the UNIX timestamp format. Unit: ms.
	//
	// example:
	//
	// 1602050276000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates the storage type of the database.
	//
	// example:
	//
	// RDS
	DbInstType *string `json:"DbInstType,omitempty" xml:"DbInstType,omitempty"`
	// Indicates the name of the database.
	//
	// example:
	//
	// db_test
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// Indicates the type of the instance in which the database is deployed. Valid values:
	//
	// 	- **MASTER**: The instance is a primary instance.
	//
	// 	- **SLAVE**: The instance is a read-only instance.
	//
	// example:
	//
	// MASTER
	InstRole *string `json:"InstRole,omitempty" xml:"InstRole,omitempty"`
	// Indicates the database sharding method.
	//
	// 	- **HORIZONTAL**: The database is horizontally sharded.
	//
	// 	- **VERTICAL**: The database is vertically sharded.
	//
	// example:
	//
	// HORIZONTAL
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// Indicates the schema name of the database.
	//
	// example:
	//
	// db_test*******************
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates the state of the database. Valid values:
	//
	// 	- **TO_BE_INIT**: The database is being created.
	//
	// 	- **NORMAL**: The database is running.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDrdsDBResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDrdsDBResponseBodyData) GetDbInstType() *string {
	return s.DbInstType
}

func (s *DescribeDrdsDBResponseBodyData) GetDbName() *string {
	return s.DbName
}

func (s *DescribeDrdsDBResponseBodyData) GetInstRole() *string {
	return s.InstRole
}

func (s *DescribeDrdsDBResponseBodyData) GetMode() *string {
	return s.Mode
}

func (s *DescribeDrdsDBResponseBodyData) GetSchema() *string {
	return s.Schema
}

func (s *DescribeDrdsDBResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *DescribeDrdsDBResponseBodyData) SetCreateTime(v string) *DescribeDrdsDBResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbInstType(v string) *DescribeDrdsDBResponseBodyData {
	s.DbInstType = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetDbName(v string) *DescribeDrdsDBResponseBodyData {
	s.DbName = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetInstRole(v string) *DescribeDrdsDBResponseBodyData {
	s.InstRole = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetMode(v string) *DescribeDrdsDBResponseBodyData {
	s.Mode = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetSchema(v string) *DescribeDrdsDBResponseBodyData {
	s.Schema = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) SetStatus(v string) *DescribeDrdsDBResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeDrdsDBResponseBodyData) Validate() error {
	return dara.Validate(s)
}
