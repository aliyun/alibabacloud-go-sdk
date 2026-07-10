// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountRole(v *GetCloudAccountRoleResponseBodyCloudAccountRole) *GetCloudAccountRoleResponseBody
	GetCloudAccountRole() *GetCloudAccountRoleResponseBodyCloudAccountRole
	SetRequestId(v string) *GetCloudAccountRoleResponseBody
	GetRequestId() *string
}

type GetCloudAccountRoleResponseBody struct {
	// The cloud role details.
	CloudAccountRole *GetCloudAccountRoleResponseBodyCloudAccountRole `json:"CloudAccountRole,omitempty" xml:"CloudAccountRole,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAccountRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleResponseBody) GetCloudAccountRole() *GetCloudAccountRoleResponseBodyCloudAccountRole {
	return s.CloudAccountRole
}

func (s *GetCloudAccountRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAccountRoleResponseBody) SetCloudAccountRole(v *GetCloudAccountRoleResponseBodyCloudAccountRole) *GetCloudAccountRoleResponseBody {
	s.CloudAccountRole = v
	return s
}

func (s *GetCloudAccountRoleResponseBody) SetRequestId(v string) *GetCloudAccountRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAccountRoleResponseBody) Validate() error {
	if s.CloudAccountRole != nil {
		if err := s.CloudAccountRole.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountRoleResponseBodyCloudAccountRole struct {
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
	// The cloud role health status. Valid values:
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
	// The cloud role health check result.
	CloudAccountRoleHealthCheckResult *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult `json:"CloudAccountRoleHealthCheckResult,omitempty" xml:"CloudAccountRoleHealthCheckResult,omitempty" type:"Struct"`
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
	// The cloud role usage type. Valid values:
	//
	// - system: system.
	//
	// - user: user.
	//
	// example:
	//
	// system
	CloudAccountRoleUsageType *string `json:"CloudAccountRoleUsageType,omitempty" xml:"CloudAccountRoleUsageType,omitempty"`
	// The creation time. The value is a UNIX timestamp in milliseconds.
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
	// The last update time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetCloudAccountRoleResponseBodyCloudAccountRole) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleResponseBodyCloudAccountRole) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleExternalId() *string {
	return s.CloudAccountRoleExternalId
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleHealth() *string {
	return s.CloudAccountRoleHealth
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleHealthCheckResult() *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult {
	return s.CloudAccountRoleHealthCheckResult
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleId() *string {
	return s.CloudAccountRoleId
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleName() *string {
	return s.CloudAccountRoleName
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleType() *string {
	return s.CloudAccountRoleType
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCloudAccountRoleUsageType() *string {
	return s.CloudAccountRoleUsageType
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetDescription() *string {
	return s.Description
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetStatus() *string {
	return s.Status
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountId(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleExternalId(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleExternalId = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleHealth(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleHealth = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleHealthCheckResult(v *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleHealthCheckResult = v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleId(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleId = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleName(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleName = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleType(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleType = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCloudAccountRoleUsageType(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CloudAccountRoleUsageType = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetCreateTime(v int64) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.CreateTime = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetDescription(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.Description = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetInstanceId(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetStatus(v string) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.Status = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) SetUpdateTime(v int64) *GetCloudAccountRoleResponseBodyCloudAccountRole {
	s.UpdateTime = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRole) Validate() error {
	if s.CloudAccountRoleHealthCheckResult != nil {
		if err := s.CloudAccountRoleHealthCheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult struct {
	// The error reason. This field is returned when the health check status is unhealthy.
	ErrorReason *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// The time of the last health check. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The cloud role health check result. Valid values:
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

func (s GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) GetErrorReason() *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason {
	return s.ErrorReason
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) GetResult() *string {
	return s.Result
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) SetErrorReason(v *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult {
	s.ErrorReason = v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) SetLastCheckTime(v int64) *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult {
	s.LastCheckTime = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) SetResult(v string) *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult {
	s.Result = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResult) Validate() error {
	if s.ErrorReason != nil {
		if err := s.ErrorReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason struct {
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

func (s GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) SetErrorCode(v string) *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason {
	s.ErrorCode = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) SetErrorMessage(v string) *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason {
	s.ErrorMessage = &v
	return s
}

func (s *GetCloudAccountRoleResponseBodyCloudAccountRoleCloudAccountRoleHealthCheckResultErrorReason) Validate() error {
	return dara.Validate(s)
}
