// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateDataImportTaskRequest
	GetDBInstanceName() *string
	SetDstDb(v string) *CreateDataImportTaskRequest
	GetDstDb() *string
	SetDstPassword(v string) *CreateDataImportTaskRequest
	GetDstPassword() *string
	SetDstResId(v string) *CreateDataImportTaskRequest
	GetDstResId() *string
	SetDstUserName(v string) *CreateDataImportTaskRequest
	GetDstUserName() *string
	SetRegionId(v string) *CreateDataImportTaskRequest
	GetRegionId() *string
	SetSlinkTaskId(v string) *CreateDataImportTaskRequest
	GetSlinkTaskId() *string
	SetSrcDb(v string) *CreateDataImportTaskRequest
	GetSrcDb() *string
	SetSrcPassword(v string) *CreateDataImportTaskRequest
	GetSrcPassword() *string
	SetSrcResId(v string) *CreateDataImportTaskRequest
	GetSrcResId() *string
	SetSrcUserName(v string) *CreateDataImportTaskRequest
	GetSrcUserName() *string
}

type CreateDataImportTaskRequest struct {
	// The instance ID. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the specified region, including instance IDs.
	//
	// example:
	//
	// pxc-********
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The execution status of the target SQL import. Valid values: 	- **importing**: importing. 	- **success**: import succeeded. 	- **fail**: import failed.
	//
	// example:
	//
	// transfer_test3
	DstDb *string `json:"DstDb,omitempty" xml:"DstDb,omitempty"`
	// The password of the privileged account for the target ApsaraDB RDS instance. > 	- You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the privileged account information of the target instance, including the password. 	- This parameter takes effect only when DstPassword is set to true.
	//
	// example:
	//
	// ******
	DstPassword *string `json:"DstPassword,omitempty" xml:"DstPassword,omitempty"`
	// The migration task ID.
	//
	// example:
	//
	// pxc-shr8****k36vrn
	DstResId *string `json:"DstResId,omitempty" xml:"DstResId,omitempty"`
	// The username of the target.
	//
	// example:
	//
	// bbt_cms_prod
	DstUserName *string `json:"DstUserName,omitempty" xml:"DstUserName,omitempty"`
	// The region in which the instance resides. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The import task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// The database information of the source when the source database is ApsaraDB RDS for MySQL. > The source database must be consistent with the target database.
	//
	// example:
	//
	// transfer_for_st
	SrcDb *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	// The read/write mode for executing the import task on the source. Valid values: 	- **rw**: read and write. 	- **ro**: read-only.
	//
	// example:
	//
	// ******
	SrcPassword *string `json:"SrcPassword,omitempty" xml:"SrcPassword,omitempty"`
	// The ID of the source ApsaraDB RDS instance. > You can call the [DescribeDrivingAccess](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all source ApsaraDB RDS instances in the specified region, including instance IDs.
	//
	// example:
	//
	// pxc-shrnv****kh87z
	SrcResId *string `json:"SrcResId,omitempty" xml:"SrcResId,omitempty"`
	// The username of the source.
	//
	// example:
	//
	// drds_test
	SrcUserName *string `json:"SrcUserName,omitempty" xml:"SrcUserName,omitempty"`
}

func (s CreateDataImportTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDataImportTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateDataImportTaskRequest) GetDstDb() *string {
	return s.DstDb
}

func (s *CreateDataImportTaskRequest) GetDstPassword() *string {
	return s.DstPassword
}

func (s *CreateDataImportTaskRequest) GetDstResId() *string {
	return s.DstResId
}

func (s *CreateDataImportTaskRequest) GetDstUserName() *string {
	return s.DstUserName
}

func (s *CreateDataImportTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataImportTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateDataImportTaskRequest) GetSrcDb() *string {
	return s.SrcDb
}

func (s *CreateDataImportTaskRequest) GetSrcPassword() *string {
	return s.SrcPassword
}

func (s *CreateDataImportTaskRequest) GetSrcResId() *string {
	return s.SrcResId
}

func (s *CreateDataImportTaskRequest) GetSrcUserName() *string {
	return s.SrcUserName
}

func (s *CreateDataImportTaskRequest) SetDBInstanceName(v string) *CreateDataImportTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstDb(v string) *CreateDataImportTaskRequest {
	s.DstDb = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstPassword(v string) *CreateDataImportTaskRequest {
	s.DstPassword = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstResId(v string) *CreateDataImportTaskRequest {
	s.DstResId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetDstUserName(v string) *CreateDataImportTaskRequest {
	s.DstUserName = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetRegionId(v string) *CreateDataImportTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSlinkTaskId(v string) *CreateDataImportTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcDb(v string) *CreateDataImportTaskRequest {
	s.SrcDb = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcPassword(v string) *CreateDataImportTaskRequest {
	s.SrcPassword = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcResId(v string) *CreateDataImportTaskRequest {
	s.SrcResId = &v
	return s
}

func (s *CreateDataImportTaskRequest) SetSrcUserName(v string) *CreateDataImportTaskRequest {
	s.SrcUserName = &v
	return s
}

func (s *CreateDataImportTaskRequest) Validate() error {
	return dara.Validate(s)
}
