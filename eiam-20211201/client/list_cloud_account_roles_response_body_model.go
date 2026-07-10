// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountRolesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountRoles(v []*ListCloudAccountRolesResponseBodyCloudAccountRoles) *ListCloudAccountRolesResponseBody
	GetCloudAccountRoles() []*ListCloudAccountRolesResponseBodyCloudAccountRoles
	SetMaxResults(v int32) *ListCloudAccountRolesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCloudAccountRolesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCloudAccountRolesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCloudAccountRolesResponseBody
	GetTotalCount() *int32
}

type ListCloudAccountRolesResponseBody struct {
	// The list of cloud roles.
	CloudAccountRoles []*ListCloudAccountRolesResponseBodyCloudAccountRoles `json:"CloudAccountRoles,omitempty" xml:"CloudAccountRoles,omitempty" type:"Repeated"`
	// The number of rows per page in the paging query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAccountRolesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesResponseBody) GetCloudAccountRoles() []*ListCloudAccountRolesResponseBodyCloudAccountRoles {
	return s.CloudAccountRoles
}

func (s *ListCloudAccountRolesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCloudAccountRolesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountRolesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAccountRolesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloudAccountRolesResponseBody) SetCloudAccountRoles(v []*ListCloudAccountRolesResponseBodyCloudAccountRoles) *ListCloudAccountRolesResponseBody {
	s.CloudAccountRoles = v
	return s
}

func (s *ListCloudAccountRolesResponseBody) SetMaxResults(v int32) *ListCloudAccountRolesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountRolesResponseBody) SetNextToken(v string) *ListCloudAccountRolesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountRolesResponseBody) SetRequestId(v string) *ListCloudAccountRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccountRolesResponseBody) SetTotalCount(v int32) *ListCloudAccountRolesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAccountRolesResponseBody) Validate() error {
	if s.CloudAccountRoles != nil {
		for _, item := range s.CloudAccountRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAccountRolesResponseBodyCloudAccountRoles struct {
	// The cloud account ID.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The cloud role identifier.
	//
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"CloudAccountRoleExternalId,omitempty" xml:"CloudAccountRoleExternalId,omitempty"`
	// The health status of the cloud role. Valid values:
	//
	// - healthy: healthy.
	//
	// - unhealthy: unhealthy.
	//
	// - unknown: unknown.
	//
	// example:
	//
	// healthy
	CloudAccountRoleHealth *string `json:"CloudAccountRoleHealth,omitempty" xml:"CloudAccountRoleHealth,omitempty"`
	// The health check result of the cloud role.
	CloudAccountRoleHealthCheckResult *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult `json:"CloudAccountRoleHealthCheckResult,omitempty" xml:"CloudAccountRoleHealthCheckResult,omitempty" type:"Struct"`
	// The cloud role ID.
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
	// The cloud role name.
	//
	// example:
	//
	// role-test
	CloudAccountRoleName *string `json:"CloudAccountRoleName,omitempty" xml:"CloudAccountRoleName,omitempty"`
	// The cloud role type. The specific format depends on the cloud account type. Valid values:
	//
	// - role: applicable to Alibaba Cloud accounts.
	//
	// example:
	//
	// role
	CloudAccountRoleType *string `json:"CloudAccountRoleType,omitempty" xml:"CloudAccountRoleType,omitempty"`
	// The usage type of the cloud role. Valid values:
	//
	// - system: system.
	//
	// - user: user.
	//
	// example:
	//
	// system
	CloudAccountRoleUsageType *string `json:"CloudAccountRoleUsageType,omitempty" xml:"CloudAccountRoleUsageType,omitempty"`
	// The creation time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1719320115000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cloud role description.
	//
	// example:
	//
	// cloud_account_role_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The cloud role status. Valid values:
	//
	// - enabled: enabled.
	//
	// - disable: disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The last update time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1719320117000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRoles) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRoles) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleExternalId() *string {
	return s.CloudAccountRoleExternalId
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleHealth() *string {
	return s.CloudAccountRoleHealth
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleHealthCheckResult() *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult {
	return s.CloudAccountRoleHealthCheckResult
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleName() *string {
	return s.CloudAccountRoleName
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleType() *string {
	return s.CloudAccountRoleType
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCloudAccountRoleUsageType() *string {
	return s.CloudAccountRoleUsageType
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetDescription() *string {
	return s.Description
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetStatus() *string {
	return s.Status
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountId(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleExternalId(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleExternalId = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleHealth(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleHealth = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleHealthCheckResult(v *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleHealthCheckResult = v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleId(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleId = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleName(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleName = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleType(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleType = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCloudAccountRoleUsageType(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CloudAccountRoleUsageType = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetCreateTime(v int64) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.CreateTime = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetDescription(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.Description = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetInstanceId(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetStatus(v string) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.Status = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) SetUpdateTime(v int64) *ListCloudAccountRolesResponseBodyCloudAccountRoles {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRoles) Validate() error {
	if s.CloudAccountRoleHealthCheckResult != nil {
		if err := s.CloudAccountRoleHealthCheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult struct {
	// The error reason. This field is returned when the health check status is unhealthy.
	ErrorReason *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// The last check time, in UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The health check result of the cloud role. Valid values:
	//
	// - success: succeeded.
	//
	// - failed: failed.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) GetErrorReason() *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason {
	return s.ErrorReason
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) GetResult() *string {
	return s.Result
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) SetErrorReason(v *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult {
	s.ErrorReason = v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) SetLastCheckTime(v int64) *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult {
	s.LastCheckTime = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) SetResult(v string) *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult {
	s.Result = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult) Validate() error {
	if s.ErrorReason != nil {
		if err := s.ErrorReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason struct {
	// The error code.
	//
	// example:
	//
	// AuthenticationFail.NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error description.
	//
	// example:
	//
	// There is no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) GoString() string {
	return s.String()
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) SetErrorCode(v string) *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason {
	s.ErrorCode = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) SetErrorMessage(v string) *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason {
	s.ErrorMessage = &v
	return s
}

func (s *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason) Validate() error {
	return dara.Validate(s)
}
