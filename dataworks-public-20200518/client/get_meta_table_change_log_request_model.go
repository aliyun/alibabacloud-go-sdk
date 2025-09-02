// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableChangeLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChangeType(v string) *GetMetaTableChangeLogRequest
	GetChangeType() *string
	SetEndDate(v string) *GetMetaTableChangeLogRequest
	GetEndDate() *string
	SetObjectType(v string) *GetMetaTableChangeLogRequest
	GetObjectType() *string
	SetPageNumber(v int32) *GetMetaTableChangeLogRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetMetaTableChangeLogRequest
	GetPageSize() *int32
	SetStartDate(v string) *GetMetaTableChangeLogRequest
	GetStartDate() *string
	SetTableGuid(v string) *GetMetaTableChangeLogRequest
	GetTableGuid() *string
}

type GetMetaTableChangeLogRequest struct {
	// The type of the change. Valid values: CREATE_TABLE, ALTER_TABLE, DROP_TABLE, ADD_PARTITION, and DROP_PARTITION.
	//
	// example:
	//
	// ALTER_TABLE
	ChangeType *string `json:"ChangeType,omitempty" xml:"ChangeType,omitempty"`
	// The end of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// 	- By default, the system uses the current time as the value of this parameter if the time that you specify is invalid.
	//
	// 	- If both the values of the StartDate and EndDate parameters are invalid, the system automatically queries the change logs that are generated within the last 30 days.
	//
	// example:
	//
	// 2020-06-02 00:00:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The entity on which the change is made. Valid values: TABLE and PARTITION.
	//
	// example:
	//
	// TABLE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The beginning of the time range to query. Specify the time in the yyyy-MM-dd HH:mm:ss format.
	//
	// 	- By default, the system uses the current time as the value of this parameter if the time that you specify is invalid.
	//
	// 	- If both the values of the StartDate and EndDate parameters are invalid, the system automatically queries the change logs that are generated within the last 30 days.
	//
	// example:
	//
	// 2020-06-01 00:00:00
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The GUID of the table. Specify the GUID in the odps.projectName.tableName format. You can call the [GetMetaDBTableList](https://help.aliyun.com/document_detail/2780086.html) operation to query the GUID.
	//
	// > To query the change logs of a MaxCompute table, you must call the [GetMetaTableChangeLog](https://help.aliyun.com/document_detail/2780094.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s GetMetaTableChangeLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableChangeLogRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableChangeLogRequest) GetChangeType() *string {
	return s.ChangeType
}

func (s *GetMetaTableChangeLogRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetMetaTableChangeLogRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *GetMetaTableChangeLogRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetMetaTableChangeLogRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMetaTableChangeLogRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetMetaTableChangeLogRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableChangeLogRequest) SetChangeType(v string) *GetMetaTableChangeLogRequest {
	s.ChangeType = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetEndDate(v string) *GetMetaTableChangeLogRequest {
	s.EndDate = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetObjectType(v string) *GetMetaTableChangeLogRequest {
	s.ObjectType = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetPageNumber(v int32) *GetMetaTableChangeLogRequest {
	s.PageNumber = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetPageSize(v int32) *GetMetaTableChangeLogRequest {
	s.PageSize = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetStartDate(v string) *GetMetaTableChangeLogRequest {
	s.StartDate = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) SetTableGuid(v string) *GetMetaTableChangeLogRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableChangeLogRequest) Validate() error {
	return dara.Validate(s)
}
