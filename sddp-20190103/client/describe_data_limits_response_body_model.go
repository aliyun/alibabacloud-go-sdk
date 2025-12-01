// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataLimitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeDataLimitsResponseBody
	GetCurrentPage() *int32
	SetItems(v []*DescribeDataLimitsResponseBodyItems) *DescribeDataLimitsResponseBody
	GetItems() []*DescribeDataLimitsResponseBodyItems
	SetPageSize(v int32) *DescribeDataLimitsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDataLimitsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDataLimitsResponseBody
	GetTotalCount() *int32
}

type DescribeDataLimitsResponseBody struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The data assets.
	Items []*DescribeDataLimitsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 769FB3C1-F4C9-42DF-9B72-7077A8989***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDataLimitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeDataLimitsResponseBody) GetItems() []*DescribeDataLimitsResponseBodyItems {
	return s.Items
}

func (s *DescribeDataLimitsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDataLimitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataLimitsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataLimitsResponseBody) SetCurrentPage(v int32) *DescribeDataLimitsResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetItems(v []*DescribeDataLimitsResponseBodyItems) *DescribeDataLimitsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetPageSize(v int32) *DescribeDataLimitsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetRequestId(v string) *DescribeDataLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) SetTotalCount(v int32) *DescribeDataLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDataLimitsResponseBodyItems struct {
	// Indicates whether the security audit feature is enabled. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	AuditStatus *int32 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// Indicates whether the data asset can be automatically scanned. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	AutoScan *int32 `json:"AutoScan,omitempty" xml:"AutoScan,omitempty"`
	// The data detection status. Valid values:
	//
	// 	- **0**: The data detection is ready.
	//
	// 	- **1**: The data detection is running.
	//
	// 	- **2**: The connectivity test is in progress.
	//
	// 	- **3**: The connectivity test is passed.
	//
	// 	- **4**: The connectivity test failed.
	//
	// example:
	//
	// 3
	CheckStatus *int32 `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// The name of the data detection status.
	//
	// example:
	//
	// Connectivity test status
	CheckStatusName *string `json:"CheckStatusName,omitempty" xml:"CheckStatusName,omitempty"`
	// Indicates whether DSC has the data de-identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	DatamaskStatus *int32 `json:"DatamaskStatus,omitempty" xml:"DatamaskStatus,omitempty"`
	// The database engine version.
	//
	// example:
	//
	// 2.0
	DbVersion *string `json:"DbVersion,omitempty" xml:"DbVersion,omitempty"`
	// Indicates whether DSC has the data identification permissions on the data asset. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The type of the database engine. Valid values include **MySQL**, **SQLServer**, **Oracle**, **PostgreSQL**, and **MongoDB**.
	//
	// example:
	//
	// MySQL
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// The error code.
	//
	// example:
	//
	// connect_network_error
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The reason for the failure.
	//
	// example:
	//
	// The password is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether the data leak prevention feature is enabled. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes (default)
	//
	// example:
	//
	// 1
	EventStatus *int32 `json:"EventStatus,omitempty" xml:"EventStatus,omitempty"`
	// The time when the data asset was created. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The unique ID of the data asset.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// 123
	InstanceDescription *string `json:"InstanceDescription,omitempty" xml:"InstanceDescription,omitempty"`
	// The ID of the data asset to which the table belongs.
	//
	// example:
	//
	// 12332
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the last scan is performed.
	//
	// 	- The value is a UNIX timestamp.
	//
	// 	- Unit: milliseconds.
	//
	// example:
	//
	// 145600000
	LastFinishedTime *int64 `json:"LastFinishedTime,omitempty" xml:"LastFinishedTime,omitempty"`
	// The last scan start time of data assets, in milliseconds.
	//
	// example:
	//
	// 145600000
	LastStartTime *int64 `json:"LastStartTime,omitempty" xml:"LastStartTime,omitempty"`
	// The region in which the data asset resides.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// The retention period of raw logs. Unit: days.
	//
	// example:
	//
	// 30
	LogStoreDay *int32 `json:"LogStoreDay,omitempty" xml:"LogStoreDay,omitempty"`
	// The ID of the member.
	//
	// example:
	//
	// **********8103
	MemberAccount *int64 `json:"MemberAccount,omitempty" xml:"MemberAccount,omitempty"`
	// The next time when the data asset is scanned. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1676620236000
	NextStartTime *int64 `json:"NextStartTime,omitempty" xml:"NextStartTime,omitempty"`
	// Indicates whether the optical character recognition (OCR) feature is enabled. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 1
	OcrStatus *int32 `json:"OcrStatus,omitempty" xml:"OcrStatus,omitempty"`
	// The parent ID of the data asset that you want to query. Valid values include **bucket, db, and project**.
	//
	// example:
	//
	// project
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The port number of the self-managed database.
	//
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The status of the data asset scan. Valid values:
	//
	// 	- **-1**: invalid
	//
	// 	- **0**: waiting
	//
	// 	- **1**: being scanned
	//
	// 	- **2**: suspended
	//
	// 	- **3**: completed
	//
	// example:
	//
	// 3
	ProcessStatus *int32 `json:"ProcessStatus,omitempty" xml:"ProcessStatus,omitempty"`
	// The total number of data tables or files.
	//
	// example:
	//
	// 100
	ProcessTotalCount *int32 `json:"ProcessTotalCount,omitempty" xml:"ProcessTotalCount,omitempty"`
	// The region in which the asset resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the service to which the data asset belongs. Valid values:
	//
	// 	- **1**: MaxCompute
	//
	// 	- **2**: OSS
	//
	// 	- **3**: AnalyticDB for MySQL
	//
	// 	- **4**: Tablestore
	//
	// 	- **5**: ApsaraDB RDS
	//
	// 	- **6**: self-managed database
	//
	// example:
	//
	// 5
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The code of the service to which the data asset belongs. Valid values: **MaxCompute, OSS, ADS, OTS, and RDS**.
	//
	// example:
	//
	// RDS
	ResourceTypeCode *string `json:"ResourceTypeCode,omitempty" xml:"ResourceTypeCode,omitempty"`
	// The number of sensitive data samples. Valid values: **0**, **5**, and **10**. Unit: data entries.
	//
	// example:
	//
	// 5
	SamplingSize *int32 `json:"SamplingSize,omitempty" xml:"SamplingSize,omitempty"`
	// A list of the IDs of the security groups that are used by PrivateLink when you install the DSC agent.
	SecurityGroupIdList []*string `json:"SecurityGroupIdList,omitempty" xml:"SecurityGroupIdList,omitempty" type:"Repeated"`
	// Indicates whether the security audit feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportAudit *bool `json:"SupportAudit,omitempty" xml:"SupportAudit,omitempty"`
	// Indicates whether the data de-identification feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportDatamask *bool `json:"SupportDatamask,omitempty" xml:"SupportDatamask,omitempty"`
	// Indicates whether anomalous event detection is supported. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SupportEvent *bool `json:"SupportEvent,omitempty" xml:"SupportEvent,omitempty"`
	// Indicates whether OCR is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportOcr *bool `json:"SupportOcr,omitempty" xml:"SupportOcr,omitempty"`
	// Indicates whether the data asset scan feature is supported. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	SupportScan *bool `json:"SupportScan,omitempty" xml:"SupportScan,omitempty"`
	// The alias of the tenant.
	//
	// example:
	//
	// insta_gram
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	// The total number of fields in the table.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The username that is used to access the data asset.
	//
	// example:
	//
	// tsts
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// A list of the IDs of the vSwitches that are used by PrivateLink when you install the DSC agent.
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the data asset belongs.
	//
	// example:
	//
	// vpc-2zevcqke6hh09c41****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeDataLimitsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataLimitsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDataLimitsResponseBodyItems) GetAuditStatus() *int32 {
	return s.AuditStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetAutoScan() *int32 {
	return s.AutoScan
}

func (s *DescribeDataLimitsResponseBodyItems) GetCheckStatus() *int32 {
	return s.CheckStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetCheckStatusName() *string {
	return s.CheckStatusName
}

func (s *DescribeDataLimitsResponseBodyItems) GetDatamaskStatus() *int32 {
	return s.DatamaskStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetDbVersion() *string {
	return s.DbVersion
}

func (s *DescribeDataLimitsResponseBodyItems) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeDataLimitsResponseBodyItems) GetEngineType() *string {
	return s.EngineType
}

func (s *DescribeDataLimitsResponseBodyItems) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeDataLimitsResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeDataLimitsResponseBodyItems) GetEventStatus() *int32 {
	return s.EventStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeDataLimitsResponseBodyItems) GetId() *int64 {
	return s.Id
}

func (s *DescribeDataLimitsResponseBodyItems) GetInstanceDescription() *string {
	return s.InstanceDescription
}

func (s *DescribeDataLimitsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDataLimitsResponseBodyItems) GetLastFinishedTime() *int64 {
	return s.LastFinishedTime
}

func (s *DescribeDataLimitsResponseBodyItems) GetLastStartTime() *int64 {
	return s.LastStartTime
}

func (s *DescribeDataLimitsResponseBodyItems) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeDataLimitsResponseBodyItems) GetLogStoreDay() *int32 {
	return s.LogStoreDay
}

func (s *DescribeDataLimitsResponseBodyItems) GetMemberAccount() *int64 {
	return s.MemberAccount
}

func (s *DescribeDataLimitsResponseBodyItems) GetNextStartTime() *int64 {
	return s.NextStartTime
}

func (s *DescribeDataLimitsResponseBodyItems) GetOcrStatus() *int32 {
	return s.OcrStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeDataLimitsResponseBodyItems) GetPort() *int32 {
	return s.Port
}

func (s *DescribeDataLimitsResponseBodyItems) GetProcessStatus() *int32 {
	return s.ProcessStatus
}

func (s *DescribeDataLimitsResponseBodyItems) GetProcessTotalCount() *int32 {
	return s.ProcessTotalCount
}

func (s *DescribeDataLimitsResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataLimitsResponseBodyItems) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *DescribeDataLimitsResponseBodyItems) GetResourceTypeCode() *string {
	return s.ResourceTypeCode
}

func (s *DescribeDataLimitsResponseBodyItems) GetSamplingSize() *int32 {
	return s.SamplingSize
}

func (s *DescribeDataLimitsResponseBodyItems) GetSecurityGroupIdList() []*string {
	return s.SecurityGroupIdList
}

func (s *DescribeDataLimitsResponseBodyItems) GetSupportAudit() *bool {
	return s.SupportAudit
}

func (s *DescribeDataLimitsResponseBodyItems) GetSupportDatamask() *bool {
	return s.SupportDatamask
}

func (s *DescribeDataLimitsResponseBodyItems) GetSupportEvent() *bool {
	return s.SupportEvent
}

func (s *DescribeDataLimitsResponseBodyItems) GetSupportOcr() *bool {
	return s.SupportOcr
}

func (s *DescribeDataLimitsResponseBodyItems) GetSupportScan() *bool {
	return s.SupportScan
}

func (s *DescribeDataLimitsResponseBodyItems) GetTenantName() *string {
	return s.TenantName
}

func (s *DescribeDataLimitsResponseBodyItems) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDataLimitsResponseBodyItems) GetUserName() *string {
	return s.UserName
}

func (s *DescribeDataLimitsResponseBodyItems) GetVSwitchIdList() []*string {
	return s.VSwitchIdList
}

func (s *DescribeDataLimitsResponseBodyItems) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeDataLimitsResponseBodyItems) SetAuditStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AuditStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetAutoScan(v int32) *DescribeDataLimitsResponseBodyItems {
	s.AutoScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetCheckStatusName(v string) *DescribeDataLimitsResponseBodyItems {
	s.CheckStatusName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDatamaskStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.DatamaskStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetDbVersion(v string) *DescribeDataLimitsResponseBodyItems {
	s.DbVersion = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEnable(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Enable = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEngineType(v string) *DescribeDataLimitsResponseBodyItems {
	s.EngineType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetErrorMessage(v string) *DescribeDataLimitsResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetEventStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.EventStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetGmtCreate(v int64) *DescribeDataLimitsResponseBodyItems {
	s.GmtCreate = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetId(v int64) *DescribeDataLimitsResponseBodyItems {
	s.Id = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceDescription(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceDescription = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetInstanceId(v string) *DescribeDataLimitsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLastFinishedTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.LastFinishedTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLastStartTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.LastStartTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLocalName(v string) *DescribeDataLimitsResponseBodyItems {
	s.LocalName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetLogStoreDay(v int32) *DescribeDataLimitsResponseBodyItems {
	s.LogStoreDay = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetMemberAccount(v int64) *DescribeDataLimitsResponseBodyItems {
	s.MemberAccount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetNextStartTime(v int64) *DescribeDataLimitsResponseBodyItems {
	s.NextStartTime = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetOcrStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.OcrStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetParentId(v string) *DescribeDataLimitsResponseBodyItems {
	s.ParentId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetPort(v int32) *DescribeDataLimitsResponseBodyItems {
	s.Port = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessStatus(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessStatus = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetProcessTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.ProcessTotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetRegionId(v string) *DescribeDataLimitsResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceType(v int64) *DescribeDataLimitsResponseBodyItems {
	s.ResourceType = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetResourceTypeCode(v string) *DescribeDataLimitsResponseBodyItems {
	s.ResourceTypeCode = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSamplingSize(v int32) *DescribeDataLimitsResponseBodyItems {
	s.SamplingSize = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSecurityGroupIdList(v []*string) *DescribeDataLimitsResponseBodyItems {
	s.SecurityGroupIdList = v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportAudit(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportAudit = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportDatamask(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportDatamask = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportEvent(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportEvent = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportOcr(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportOcr = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetSupportScan(v bool) *DescribeDataLimitsResponseBodyItems {
	s.SupportScan = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTenantName(v string) *DescribeDataLimitsResponseBodyItems {
	s.TenantName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetTotalCount(v int32) *DescribeDataLimitsResponseBodyItems {
	s.TotalCount = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetUserName(v string) *DescribeDataLimitsResponseBodyItems {
	s.UserName = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetVSwitchIdList(v []*string) *DescribeDataLimitsResponseBodyItems {
	s.VSwitchIdList = v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) SetVpcId(v string) *DescribeDataLimitsResponseBodyItems {
	s.VpcId = &v
	return s
}

func (s *DescribeDataLimitsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
