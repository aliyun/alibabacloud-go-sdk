// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountPrivilegesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAccountPrivilegesResponseBodyData) *DescribeAccountPrivilegesResponseBody
	GetData() []*DescribeAccountPrivilegesResponseBodyData
	SetPageNumber(v int64) *DescribeAccountPrivilegesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeAccountPrivilegesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeAccountPrivilegesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeAccountPrivilegesResponseBody
	GetTotalCount() *int64
}

type DescribeAccountPrivilegesResponseBody struct {
	// Details of the permissions.
	Data []*DescribeAccountPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DA32480A-E3E5-1BE7-BA98-724551DC04C8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccountPrivilegesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBody) GetData() []*DescribeAccountPrivilegesResponseBodyData {
	return s.Data
}

func (s *DescribeAccountPrivilegesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeAccountPrivilegesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAccountPrivilegesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountPrivilegesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeAccountPrivilegesResponseBody) SetData(v []*DescribeAccountPrivilegesResponseBodyData) *DescribeAccountPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetPageNumber(v int64) *DescribeAccountPrivilegesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetPageSize(v int64) *DescribeAccountPrivilegesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetRequestId(v string) *DescribeAccountPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) SetTotalCount(v int64) *DescribeAccountPrivilegesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountPrivilegesResponseBodyData struct {
	// The objects on which the permission takes effect, including databases, tables, columns, and additional descriptions.
	PrivilegeObject *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level of the permission. Valid values: `Global`, `Database`, `Table`, and `Column`. You can call the `DescribeEnabledPrivileges` parameter to query the permission level of a specific permission.
	//
	// example:
	//
	// Column
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The name of the permission. You can call the `DescribeEnabledPrivileges` operation to query the name of the permission.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountPrivilegesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBodyData) GetPrivilegeObject() *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	return s.PrivilegeObject
}

func (s *DescribeAccountPrivilegesResponseBodyData) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *DescribeAccountPrivilegesResponseBodyData) GetPrivileges() []*string {
	return s.Privileges
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivilegeObject(v *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) *DescribeAccountPrivilegesResponseBodyData {
	s.PrivilegeObject = v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivilegeType(v string) *DescribeAccountPrivilegesResponseBodyData {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyData) SetPrivileges(v []*string) *DescribeAccountPrivilegesResponseBodyData {
	s.Privileges = v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyData) Validate() error {
	if s.PrivilegeObject != nil {
		if err := s.PrivilegeObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountPrivilegesResponseBodyDataPrivilegeObject struct {
	// The name of the column.
	//
	// example:
	//
	// column1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// db1
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description of the permission object.
	//
	// example:
	//
	// a test column
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tabl1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GoString() string {
	return s.String()
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GetColumn() *string {
	return s.Column
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GetDatabase() *string {
	return s.Database
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) GetTable() *string {
	return s.Table
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetColumn(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Column = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetDatabase(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Database = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetDescription(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Description = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) SetTable(v string) *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject {
	s.Table = &v
	return s
}

func (s *DescribeAccountPrivilegesResponseBodyDataPrivilegeObject) Validate() error {
	return dara.Validate(s)
}
