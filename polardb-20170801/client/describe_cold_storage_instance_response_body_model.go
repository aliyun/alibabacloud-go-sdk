// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColdStorageInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeColdStorageInstanceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeColdStorageInstanceResponseBody
	GetNextToken() *string
	SetObjectType(v string) *DescribeColdStorageInstanceResponseBody
	GetObjectType() *string
	SetOssClusterEnabled(v string) *DescribeColdStorageInstanceResponseBody
	GetOssClusterEnabled() *string
	SetOssClusterInfoList(v []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList) *DescribeColdStorageInstanceResponseBody
	GetOssClusterInfoList() []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList
	SetPageNumber(v int32) *DescribeColdStorageInstanceResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeColdStorageInstanceResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeColdStorageInstanceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeColdStorageInstanceResponseBody
	GetRequestId() *string
	SetSupportOssCluster(v string) *DescribeColdStorageInstanceResponseBody
	GetSupportOssCluster() *string
	SetTables(v []*DescribeColdStorageInstanceResponseBodyTables) *DescribeColdStorageInstanceResponseBody
	GetTables() []*DescribeColdStorageInstanceResponseBodyTables
	SetTotalRecord(v int32) *DescribeColdStorageInstanceResponseBody
	GetTotalRecord() *int32
}

type DescribeColdStorageInstanceResponseBody struct {
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// c2FpXzIwMjIwNjI5X2Jhay9zYWlfc3VtbWVyX3RyZWFzdXJlX3Bvb2xfbG9nLkNTVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// TABLE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// true
	OssClusterEnabled  *string                                                      `json:"OssClusterEnabled,omitempty" xml:"OssClusterEnabled,omitempty"`
	OssClusterInfoList []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList `json:"OssClusterInfoList,omitempty" xml:"OssClusterInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// C7A8EA8E-A140-5226-90D7-5BCB304D3DB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	SupportOssCluster *string                                          `json:"SupportOssCluster,omitempty" xml:"SupportOssCluster,omitempty"`
	Tables            []*DescribeColdStorageInstanceResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	TotalRecord *int32 `json:"TotalRecord,omitempty" xml:"TotalRecord,omitempty"`
}

func (s DescribeColdStorageInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeColdStorageInstanceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeColdStorageInstanceResponseBody) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeColdStorageInstanceResponseBody) GetOssClusterEnabled() *string {
	return s.OssClusterEnabled
}

func (s *DescribeColdStorageInstanceResponseBody) GetOssClusterInfoList() []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList {
	return s.OssClusterInfoList
}

func (s *DescribeColdStorageInstanceResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeColdStorageInstanceResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeColdStorageInstanceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeColdStorageInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeColdStorageInstanceResponseBody) GetSupportOssCluster() *string {
	return s.SupportOssCluster
}

func (s *DescribeColdStorageInstanceResponseBody) GetTables() []*DescribeColdStorageInstanceResponseBodyTables {
	return s.Tables
}

func (s *DescribeColdStorageInstanceResponseBody) GetTotalRecord() *int32 {
	return s.TotalRecord
}

func (s *DescribeColdStorageInstanceResponseBody) SetMaxResults(v int32) *DescribeColdStorageInstanceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetNextToken(v string) *DescribeColdStorageInstanceResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetObjectType(v string) *DescribeColdStorageInstanceResponseBody {
	s.ObjectType = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetOssClusterEnabled(v string) *DescribeColdStorageInstanceResponseBody {
	s.OssClusterEnabled = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetOssClusterInfoList(v []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList) *DescribeColdStorageInstanceResponseBody {
	s.OssClusterInfoList = v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetPageNumber(v int32) *DescribeColdStorageInstanceResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetPageRecordCount(v int32) *DescribeColdStorageInstanceResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetPageSize(v int32) *DescribeColdStorageInstanceResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetRequestId(v string) *DescribeColdStorageInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetSupportOssCluster(v string) *DescribeColdStorageInstanceResponseBody {
	s.SupportOssCluster = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetTables(v []*DescribeColdStorageInstanceResponseBodyTables) *DescribeColdStorageInstanceResponseBody {
	s.Tables = v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) SetTotalRecord(v int32) *DescribeColdStorageInstanceResponseBody {
	s.TotalRecord = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeColdStorageInstanceResponseBodyOssClusterInfoList struct {
	// example:
	//
	// 2023-05-10T17:01:16Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// pc-*****************
	OssClusterId *string `json:"OssClusterId,omitempty" xml:"OssClusterId,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 50
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeColdStorageInstanceResponseBodyOssClusterInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceResponseBodyOssClusterInfoList) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) GetOssClusterId() *string {
	return s.OssClusterId
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) GetRegion() *string {
	return s.Region
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) GetSize() *string {
	return s.Size
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) SetCreatedTime(v string) *DescribeColdStorageInstanceResponseBodyOssClusterInfoList {
	s.CreatedTime = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) SetOssClusterId(v string) *DescribeColdStorageInstanceResponseBodyOssClusterInfoList {
	s.OssClusterId = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) SetRegion(v string) *DescribeColdStorageInstanceResponseBodyOssClusterInfoList {
	s.Region = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) SetSize(v string) *DescribeColdStorageInstanceResponseBodyOssClusterInfoList {
	s.Size = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyOssClusterInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeColdStorageInstanceResponseBodyTables struct {
	ChildObjects []*DescribeColdStorageInstanceResponseBodyTablesChildObjects `json:"ChildObjects,omitempty" xml:"ChildObjects,omitempty" type:"Repeated"`
	// example:
	//
	// test_db
	DB *string `json:"DB,omitempty" xml:"DB,omitempty"`
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// example:
	//
	// user
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// example:
	//
	// pc-*****************
	OssClusterId *string `json:"OssClusterId,omitempty" xml:"OssClusterId,omitempty"`
	// example:
	//
	// 202509
	Partion *string `json:"Partion,omitempty" xml:"Partion,omitempty"`
	// example:
	//
	// 30
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// user
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// example:
	//
	// test_table
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeColdStorageInstanceResponseBodyTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceResponseBodyTables) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetChildObjects() []*DescribeColdStorageInstanceResponseBodyTablesChildObjects {
	return s.ChildObjects
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetDB() *string {
	return s.DB
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetDBName() *string {
	return s.DBName
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetOssClusterId() *string {
	return s.OssClusterId
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetPartion() *string {
	return s.Partion
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetSize() *string {
	return s.Size
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetStatus() *string {
	return s.Status
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetTable() *string {
	return s.Table
}

func (s *DescribeColdStorageInstanceResponseBodyTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetChildObjects(v []*DescribeColdStorageInstanceResponseBodyTablesChildObjects) *DescribeColdStorageInstanceResponseBodyTables {
	s.ChildObjects = v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetDB(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.DB = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetDBName(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.DBName = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetFieldName(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.FieldName = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetOssClusterId(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.OssClusterId = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetPartion(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.Partion = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetSize(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.Size = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetStatus(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.Status = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetTable(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.Table = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) SetTableName(v string) *DescribeColdStorageInstanceResponseBodyTables {
	s.TableName = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTables) Validate() error {
	return dara.Validate(s)
}

type DescribeColdStorageInstanceResponseBodyTablesChildObjects struct {
	// example:
	//
	// img/1728554006462.png
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// example:
	//
	// File
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// example:
	//
	// 10
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeColdStorageInstanceResponseBodyTablesChildObjects) String() string {
	return dara.Prettify(s)
}

func (s DescribeColdStorageInstanceResponseBodyTablesChildObjects) GoString() string {
	return s.String()
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) GetObjectName() *string {
	return s.ObjectName
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) GetObjectType() *string {
	return s.ObjectType
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) GetSize() *string {
	return s.Size
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) GetStatus() *string {
	return s.Status
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) SetObjectName(v string) *DescribeColdStorageInstanceResponseBodyTablesChildObjects {
	s.ObjectName = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) SetObjectType(v string) *DescribeColdStorageInstanceResponseBodyTablesChildObjects {
	s.ObjectType = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) SetSize(v string) *DescribeColdStorageInstanceResponseBodyTablesChildObjects {
	s.Size = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) SetStatus(v string) *DescribeColdStorageInstanceResponseBodyTablesChildObjects {
	s.Status = &v
	return s
}

func (s *DescribeColdStorageInstanceResponseBodyTablesChildObjects) Validate() error {
	return dara.Validate(s)
}
