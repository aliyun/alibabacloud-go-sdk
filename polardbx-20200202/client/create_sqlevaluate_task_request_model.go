// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSQLEvaluateTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *CreateSQLEvaluateTaskRequest
	GetDBInstanceName() *string
	SetDstDb(v string) *CreateSQLEvaluateTaskRequest
	GetDstDb() *string
	SetDstPassword(v string) *CreateSQLEvaluateTaskRequest
	GetDstPassword() *string
	SetDstResId(v string) *CreateSQLEvaluateTaskRequest
	GetDstResId() *string
	SetDstUserName(v string) *CreateSQLEvaluateTaskRequest
	GetDstUserName() *string
	SetRegionId(v string) *CreateSQLEvaluateTaskRequest
	GetRegionId() *string
	SetSlinkTaskDesc(v string) *CreateSQLEvaluateTaskRequest
	GetSlinkTaskDesc() *string
	SetSlinkTaskId(v string) *CreateSQLEvaluateTaskRequest
	GetSlinkTaskId() *string
	SetSrcDb(v string) *CreateSQLEvaluateTaskRequest
	GetSrcDb() *string
	SetSrcPassword(v string) *CreateSQLEvaluateTaskRequest
	GetSrcPassword() *string
	SetSrcResId(v string) *CreateSQLEvaluateTaskRequest
	GetSrcResId() *string
	SetSrcResType(v string) *CreateSQLEvaluateTaskRequest
	GetSrcResType() *string
	SetSrcUserName(v string) *CreateSQLEvaluateTaskRequest
	GetSrcUserName() *string
}

type CreateSQLEvaluateTaskRequest struct {
	// The instance ID. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of all instances in the target region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// pxc-shr4idrgogti89
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The ID of the instance on which the target SQL statement is executed. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of the target SQL statement, including the instance ID.
	//
	// example:
	//
	// transfer_for_st
	DstDb *string `json:"DstDb,omitempty" xml:"DstDb,omitempty"`
	// The password of the target SQL statement. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of the target SQL statement, including the password.
	//
	// example:
	//
	// ******
	DstPassword *string `json:"DstPassword,omitempty" xml:"DstPassword,omitempty"`
	// The ID of the target SQL statement. > You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the details of the target SQL statement, including the ID.
	//
	// example:
	//
	// pxc-zkrc1****l54rc
	DstResId *string `json:"DstResId,omitempty" xml:"DstResId,omitempty"`
	// The username of the destination instance.
	//
	// example:
	//
	// drds_test
	DstUserName *string `json:"DstUserName,omitempty" xml:"DstUserName,omitempty"`
	// The region ID of the instance. > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196841.html) operation to query the regions supported by PolarDB-X, including region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the target SQL statement.
	//
	// example:
	//
	// test-drds->auto
	SlinkTaskDesc *string `json:"SlinkTaskDesc,omitempty" xml:"SlinkTaskDesc,omitempty"`
	// The task ID for executing the target SQL statement. > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/196830.html) operation to query the IDs of target SQL statements supported by PolarDB-X.
	//
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
	// Specifies whether to enable the source ApsaraDB RDS instance. Valid values: ***enable**: enabled. ***disabled**: disabled.
	//
	// example:
	//
	// transfer_test3
	SrcDb *string `json:"SrcDb,omitempty" xml:"SrcDb,omitempty"`
	// The port number of the source instance. Valid values: 3200 to 3999. > This parameter is available and required only when **DBEndpointInstanceType*	- is set to **polardb-o**.
	//
	// example:
	//
	// ******
	SrcPassword *string `json:"SrcPassword,omitempty" xml:"SrcPassword,omitempty"`
	// The ID of the source ApsaraDB RDS instance. > You can call the [DescribeDropSystemEventMetaList](https://help.aliyun.com/document_detail/196836.html) operation to query the details of all source ApsaraDB RDS instances in the target region, including instance IDs.
	//
	// example:
	//
	// pxc-shr****rgkh87z
	SrcResId *string `json:"SrcResId,omitempty" xml:"SrcResId,omitempty"`
	// The type of the source instance. Valid values: RDS_MYSQL, POLARX1, POLARX2_STANDARD, POLARX2_ENTERPRISE, and POLARDB_M.
	//
	// example:
	//
	// POLARX2
	SrcResType *string `json:"SrcResType,omitempty" xml:"SrcResType,omitempty"`
	// The username of the source instance.
	//
	// example:
	//
	// drds_test
	SrcUserName *string `json:"SrcUserName,omitempty" xml:"SrcUserName,omitempty"`
}

func (s CreateSQLEvaluateTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSQLEvaluateTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSQLEvaluateTaskRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *CreateSQLEvaluateTaskRequest) GetDstDb() *string {
	return s.DstDb
}

func (s *CreateSQLEvaluateTaskRequest) GetDstPassword() *string {
	return s.DstPassword
}

func (s *CreateSQLEvaluateTaskRequest) GetDstResId() *string {
	return s.DstResId
}

func (s *CreateSQLEvaluateTaskRequest) GetDstUserName() *string {
	return s.DstUserName
}

func (s *CreateSQLEvaluateTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSQLEvaluateTaskRequest) GetSlinkTaskDesc() *string {
	return s.SlinkTaskDesc
}

func (s *CreateSQLEvaluateTaskRequest) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateSQLEvaluateTaskRequest) GetSrcDb() *string {
	return s.SrcDb
}

func (s *CreateSQLEvaluateTaskRequest) GetSrcPassword() *string {
	return s.SrcPassword
}

func (s *CreateSQLEvaluateTaskRequest) GetSrcResId() *string {
	return s.SrcResId
}

func (s *CreateSQLEvaluateTaskRequest) GetSrcResType() *string {
	return s.SrcResType
}

func (s *CreateSQLEvaluateTaskRequest) GetSrcUserName() *string {
	return s.SrcUserName
}

func (s *CreateSQLEvaluateTaskRequest) SetDBInstanceName(v string) *CreateSQLEvaluateTaskRequest {
	s.DBInstanceName = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetDstDb(v string) *CreateSQLEvaluateTaskRequest {
	s.DstDb = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetDstPassword(v string) *CreateSQLEvaluateTaskRequest {
	s.DstPassword = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetDstResId(v string) *CreateSQLEvaluateTaskRequest {
	s.DstResId = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetDstUserName(v string) *CreateSQLEvaluateTaskRequest {
	s.DstUserName = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetRegionId(v string) *CreateSQLEvaluateTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSlinkTaskDesc(v string) *CreateSQLEvaluateTaskRequest {
	s.SlinkTaskDesc = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSlinkTaskId(v string) *CreateSQLEvaluateTaskRequest {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSrcDb(v string) *CreateSQLEvaluateTaskRequest {
	s.SrcDb = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSrcPassword(v string) *CreateSQLEvaluateTaskRequest {
	s.SrcPassword = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSrcResId(v string) *CreateSQLEvaluateTaskRequest {
	s.SrcResId = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSrcResType(v string) *CreateSQLEvaluateTaskRequest {
	s.SrcResType = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) SetSrcUserName(v string) *CreateSQLEvaluateTaskRequest {
	s.SrcUserName = &v
	return s
}

func (s *CreateSQLEvaluateTaskRequest) Validate() error {
	return dara.Validate(s)
}
