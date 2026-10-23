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
	// The maximum number of data entries to read in this request. Default value: 10.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous API call. If there are no more queries, do not specify this parameter.
	//
	// example:
	//
	// c2FpXzIwMjIwNjI5X2Jhay9zYWlfc3VtbWVyX3RyZWFzdXJlX3Bvb2xfbG9nLkNTVg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The object type.
	//
	// example:
	//
	// TABLE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// Specifies whether OSS storage space is enabled. Valid values:
	//
	// - **true**: enabled
	//
	// - **false**: disabled
	//
	// example:
	//
	// true
	OssClusterEnabled *string `json:"OssClusterEnabled,omitempty" xml:"OssClusterEnabled,omitempty"`
	// The list of OSS addresses for the cold storage instance.
	OssClusterInfoList []*DescribeColdStorageInstanceResponseBodyOssClusterInfoList `json:"OssClusterInfoList,omitempty" xml:"OssClusterInfoList,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7A8EA8E-A140-5226-90D7-5BCB304D3DB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the current cluster supports enabling cold storage. For clusters that do not support this feature, the switch is not displayed on the page.
	//
	// example:
	//
	// true
	SupportOssCluster *string `json:"SupportOssCluster,omitempty" xml:"SupportOssCluster,omitempty"`
	// The list of cold storage instances.
	Tables []*DescribeColdStorageInstanceResponseBodyTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	// The total number of records.
	//
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
	if s.OssClusterInfoList != nil {
		for _, item := range s.OssClusterInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tables != nil {
		for _, item := range s.Tables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeColdStorageInstanceResponseBodyOssClusterInfoList struct {
	// The creation time of the cluster. The time follows the ISO 8601 standard in the YYYY-MM-DDTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-05-10T17:01:16Z
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// The cold storage instance ID.
	//
	// example:
	//
	// pc-*****************
	OssClusterId *string `json:"OssClusterId,omitempty" xml:"OssClusterId,omitempty"`
	// The region ID of the region where the task resides.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The size of the cold storage table. Unit: GB.
	//
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
	// The list of child objects.
	ChildObjects []*DescribeColdStorageInstanceResponseBodyTablesChildObjects `json:"ChildObjects,omitempty" xml:"ChildObjects,omitempty" type:"Repeated"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	DB *string `json:"DB,omitempty" xml:"DB,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// test_db
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// The name of the LOB field.
	//
	// example:
	//
	// user
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The OSS cluster ID.
	//
	// example:
	//
	// pc-*****************
	OssClusterId *string `json:"OssClusterId,omitempty" xml:"OssClusterId,omitempty"`
	// The partition of the cold storage instance.
	//
	// example:
	//
	// 202509
	Partion *string `json:"Partion,omitempty" xml:"Partion,omitempty"`
	// The disk size of the cold storage instance. Unit: GiB.
	//
	// example:
	//
	// 30
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The task status.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// user
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
	// The name of the table.
	//
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
	if s.ChildObjects != nil {
		for _, item := range s.ChildObjects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeColdStorageInstanceResponseBodyTablesChildObjects struct {
	// The object name.
	//
	// example:
	//
	// img/1728554006462.png
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The object type.
	//
	// example:
	//
	// File
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The disk size. Unit: GiB.
	//
	// example:
	//
	// 10
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// The task status. Valid values:
	//
	// 	- **Scheduled**: waiting for execution
	//
	// 	- **Running**: running
	//
	// 	- **Succeed**: succeeded
	//
	// 	- **Cancelling**: cancelling
	//
	// 	- **Canceled**: canceled
	//
	// 	- **Waiting**: waiting for the preset time
	//
	// To query multiple statuses, separate them with commas (,). The default value is empty, which indicates that all statuses are selected.
	//
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
