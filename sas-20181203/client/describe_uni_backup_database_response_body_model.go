// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseList(v []*DescribeUniBackupDatabaseResponseBodyDatabaseList) *DescribeUniBackupDatabaseResponseBody
	GetDatabaseList() []*DescribeUniBackupDatabaseResponseBodyDatabaseList
	SetPageInfo(v *DescribeUniBackupDatabaseResponseBodyPageInfo) *DescribeUniBackupDatabaseResponseBody
	GetPageInfo() *DescribeUniBackupDatabaseResponseBodyPageInfo
	SetRequestId(v string) *DescribeUniBackupDatabaseResponseBody
	GetRequestId() *string
}

type DescribeUniBackupDatabaseResponseBody struct {
	// An array that consists of the information about the databases.
	DatabaseList []*DescribeUniBackupDatabaseResponseBodyDatabaseList `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeUniBackupDatabaseResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUniBackupDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupDatabaseResponseBody) GetDatabaseList() []*DescribeUniBackupDatabaseResponseBodyDatabaseList {
	return s.DatabaseList
}

func (s *DescribeUniBackupDatabaseResponseBody) GetPageInfo() *DescribeUniBackupDatabaseResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeUniBackupDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniBackupDatabaseResponseBody) SetDatabaseList(v []*DescribeUniBackupDatabaseResponseBodyDatabaseList) *DescribeUniBackupDatabaseResponseBody {
	s.DatabaseList = v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBody) SetPageInfo(v *DescribeUniBackupDatabaseResponseBodyPageInfo) *DescribeUniBackupDatabaseResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBody) SetRequestId(v string) *DescribeUniBackupDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUniBackupDatabaseResponseBodyDatabaseList struct {
	// The status of the anti-ransomware agent. Valid values:
	//
	// 	- **UNKNOWN**: unknown
	//
	// 	- **INSTALLED**: installed
	//
	// 	- **INSTALL_FAILED**: installation failed
	//
	// 	- **UNINSTALL_FAILED**: uninstallation failed
	//
	// example:
	//
	// INSTALLED
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The service from which the database is created. Valid values:
	//
	// 	- **HBR**: HBR
	//
	// 	- **AEGIS**: Security Center
	//
	// example:
	//
	// AEGIS
	CreatedByProduct *string `json:"CreatedByProduct,omitempty" xml:"CreatedByProduct,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// apns_tt180
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **MYSQL**
	//
	// 	- **MSSQL**
	//
	// 	- **Oracle**
	//
	// example:
	//
	// MYSQL
	DatabaseType *string `json:"DatabaseType,omitempty" xml:"DatabaseType,omitempty"`
	// The version of the database engine.
	//
	// example:
	//
	// 12.0.4100.1
	DatabaseVersion *string `json:"DatabaseVersion,omitempty" xml:"DatabaseVersion,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-bp15aho9hhftvmhw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance to which the database belongs.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The UUID of the Hybrid Backup Recovery (HBR) agent that is used to back up the data of the database.
	//
	// example:
	//
	// 85878b284df911ec800000163e19****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The status of the ECS instance. Valid values:
	//
	// 	- **Stopped**
	//
	// 	- **Running**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeUniBackupDatabaseResponseBodyDatabaseList) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupDatabaseResponseBodyDatabaseList) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetCreatedByProduct() *string {
	return s.CreatedByProduct
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetDatabaseVersion() *string {
	return s.DatabaseVersion
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) GetStatus() *string {
	return s.Status
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetAgentStatus(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.AgentStatus = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetCreatedByProduct(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.CreatedByProduct = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetDatabaseName(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.DatabaseName = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetDatabaseType(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.DatabaseType = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetDatabaseVersion(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.DatabaseVersion = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetInstanceId(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.InstanceId = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetInstanceName(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.InstanceName = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetInstanceUuid(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.InstanceUuid = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetPolicyId(v int64) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.PolicyId = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) SetStatus(v string) *DescribeUniBackupDatabaseResponseBodyDatabaseList {
	s.Status = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyDatabaseList) Validate() error {
	return dara.Validate(s)
}

type DescribeUniBackupDatabaseResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUniBackupDatabaseResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupDatabaseResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) SetCount(v int32) *DescribeUniBackupDatabaseResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeUniBackupDatabaseResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) SetPageSize(v int32) *DescribeUniBackupDatabaseResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) SetTotalCount(v int32) *DescribeUniBackupDatabaseResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeUniBackupDatabaseResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
