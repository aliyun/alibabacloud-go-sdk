// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeUniBackupPoliciesResponseBodyPageInfo) *DescribeUniBackupPoliciesResponseBody
	GetPageInfo() *DescribeUniBackupPoliciesResponseBodyPageInfo
	SetRequestId(v string) *DescribeUniBackupPoliciesResponseBody
	GetRequestId() *string
	SetUniBackupPolicies(v []*DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) *DescribeUniBackupPoliciesResponseBody
	GetUniBackupPolicies() []*DescribeUniBackupPoliciesResponseBodyUniBackupPolicies
}

type DescribeUniBackupPoliciesResponseBody struct {
	// The pagination information.
	PageInfo *DescribeUniBackupPoliciesResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// ACF97412-FD09-4D1F-994F-34DF12BR****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the anti-ransomware policies.
	UniBackupPolicies []*DescribeUniBackupPoliciesResponseBodyUniBackupPolicies `json:"UniBackupPolicies,omitempty" xml:"UniBackupPolicies,omitempty" type:"Repeated"`
}

func (s DescribeUniBackupPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPoliciesResponseBody) GetPageInfo() *DescribeUniBackupPoliciesResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeUniBackupPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUniBackupPoliciesResponseBody) GetUniBackupPolicies() []*DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	return s.UniBackupPolicies
}

func (s *DescribeUniBackupPoliciesResponseBody) SetPageInfo(v *DescribeUniBackupPoliciesResponseBodyPageInfo) *DescribeUniBackupPoliciesResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBody) SetRequestId(v string) *DescribeUniBackupPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBody) SetUniBackupPolicies(v []*DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) *DescribeUniBackupPoliciesResponseBody {
	s.UniBackupPolicies = v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UniBackupPolicies != nil {
		for _, item := range s.UniBackupPolicies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUniBackupPoliciesResponseBodyPageInfo struct {
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
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeUniBackupPoliciesResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPoliciesResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) SetCount(v int32) *DescribeUniBackupPoliciesResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeUniBackupPoliciesResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) SetPageSize(v int32) *DescribeUniBackupPoliciesResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) SetTotalCount(v int32) *DescribeUniBackupPoliciesResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeUniBackupPoliciesResponseBodyUniBackupPolicies struct {
	// The error message for the anti-ransomware agent.
	//
	// example:
	//
	// INSTALL_TIMEOUT
	AgentErrorMessage *string `json:"AgentErrorMessage,omitempty" xml:"AgentErrorMessage,omitempty"`
	// The status of the agent. Valid values:
	//
	// 	- **UNKNOWN**
	//
	// 	- **INSTALLED**
	//
	// 	- **INSTALL_FAILED**
	//
	// 	- **UNINSTALL_FAILED**
	//
	// example:
	//
	// INSTALLED
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// abc123
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
	// The error code returned when the backup task fails.
	//
	// example:
	//
	// EXPIRED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message for the anti-ransomware policy.
	//
	// example:
	//
	// AttachRamRoleError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the server.
	//
	// example:
	//
	// i-9dp7mubt5wit6g0h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The status of the Elastic Compute Service (ECS) instance. Valid values:
	//
	// 	- **Stopped**
	//
	// 	- **Running**
	//
	// example:
	//
	// Running
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The UUID of the agent that is used to back up the data of the database.
	//
	// example:
	//
	// cf1bcad4063f11ed800000163e0e****
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The execution result of the last backup task.
	//
	// example:
	//
	// completed
	LatestBackResult *string `json:"LatestBackResult,omitempty" xml:"LatestBackResult,omitempty"`
	// The time when the last backup task was executed.
	//
	// example:
	//
	// 2022-01-01 00:00:11
	LatestBackupTime *string `json:"LatestBackupTime,omitempty" xml:"LatestBackupTime,omitempty"`
	// The status of the backup task. Valid values:
	//
	// 	- **init**
	//
	// 	- **running**
	//
	// 	- **completed**
	//
	// 	- **restoring**
	//
	// 	- **creating**
	//
	// 	- **created**
	//
	// example:
	//
	// creating
	PlanStatus *string `json:"PlanStatus,omitempty" xml:"PlanStatus,omitempty"`
	// The ID of the anti-ransomware policy.
	//
	// example:
	//
	// 123
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The name of the anti-ransomware policy.
	//
	// example:
	//
	// auto_oracle_37f
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The status of the anti-ransomware policy. Valid values:
	//
	// 	- **initiating**
	//
	// 	- **opening**
	//
	// 	- **closing**
	//
	// 	- **deleting**
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// example:
	//
	// opening
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// The region ID of the server that hosts the database.
	//
	// example:
	//
	// cn-hangzhou
	UniRegionId *string `json:"UniRegionId,omitempty" xml:"UniRegionId,omitempty"`
}

func (s DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetAgentErrorMessage() *string {
	return s.AgentErrorMessage
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetDatabaseType() *string {
	return s.DatabaseType
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetLatestBackResult() *string {
	return s.LatestBackResult
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetLatestBackupTime() *string {
	return s.LatestBackupTime
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetPlanStatus() *string {
	return s.PlanStatus
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) GetUniRegionId() *string {
	return s.UniRegionId
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetAgentErrorMessage(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.AgentErrorMessage = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetAgentStatus(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.AgentStatus = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetDatabaseName(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.DatabaseName = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetDatabaseType(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.DatabaseType = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetErrorCode(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.ErrorCode = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetErrorMessage(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetInstanceId(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.InstanceId = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetInstanceName(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.InstanceName = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetInstanceStatus(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetInstanceUuid(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.InstanceUuid = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetLatestBackResult(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.LatestBackResult = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetLatestBackupTime(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.LatestBackupTime = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetPlanStatus(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.PlanStatus = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetPolicyId(v int64) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.PolicyId = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetPolicyName(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.PolicyName = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetPolicyStatus(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.PolicyStatus = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) SetUniRegionId(v string) *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies {
	s.UniRegionId = &v
	return s
}

func (s *DescribeUniBackupPoliciesResponseBodyUniBackupPolicies) Validate() error {
	return dara.Validate(s)
}
