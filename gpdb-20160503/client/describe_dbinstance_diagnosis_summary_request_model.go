// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDiagnosisSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDBInstanceDiagnosisSummaryRequest
	GetDBInstanceId() *string
	SetPageNumber(v int32) *DescribeDBInstanceDiagnosisSummaryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDBInstanceDiagnosisSummaryRequest
	GetPageSize() *int32
	SetRolePreferd(v string) *DescribeDBInstanceDiagnosisSummaryRequest
	GetRolePreferd() *string
	SetStartStatus(v string) *DescribeDBInstanceDiagnosisSummaryRequest
	GetStartStatus() *string
	SetSyncMode(v string) *DescribeDBInstanceDiagnosisSummaryRequest
	GetSyncMode() *string
}

type DescribeDBInstanceDiagnosisSummaryRequest struct {
	// The instance ID.
	//
	// > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query details about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The page number. Pages start from page 1. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **20**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The role state of the node. It specifies whether a primary/secondary switchover occurs. Valid values:
	//
	// 	- **normal**: The node role is normal. No primary/secondary switchover occurs.
	//
	// 	- **reverse**: The node role is reversed. A primary/secondary switchover occurs.
	//
	// example:
	//
	// normal
	RolePreferd *string `json:"RolePreferd,omitempty" xml:"RolePreferd,omitempty"`
	// The running state of the node. Valid values:
	//
	// 	- **UP**: The node is running.
	//
	// 	- **DOWN**: The node is faulty.
	//
	// If you do not specify this parameter, the information about nodes in all running states is returned.
	//
	// example:
	//
	// UP
	StartStatus *string `json:"StartStatus,omitempty" xml:"StartStatus,omitempty"`
	// The data synchronization state of the node. Valid values:
	//
	// 	- **synced**: The node data is synchronized.
	//
	// 	- **notSyncing**: The node data is not synchronized.
	//
	// If you do not specify this parameter, the information about nodes in all synchronization states is returned.
	//
	// example:
	//
	// synced
	SyncMode *string `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDiagnosisSummaryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetRolePreferd() *string {
	return s.RolePreferd
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetStartStatus() *string {
	return s.StartStatus
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) GetSyncMode() *string {
	return s.SyncMode
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetDBInstanceId(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageNumber(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetPageSize(v int32) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetRolePreferd(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.RolePreferd = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetStartStatus(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.StartStatus = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) SetSyncMode(v string) *DescribeDBInstanceDiagnosisSummaryRequest {
	s.SyncMode = &v
	return s
}

func (s *DescribeDBInstanceDiagnosisSummaryRequest) Validate() error {
	return dara.Validate(s)
}
