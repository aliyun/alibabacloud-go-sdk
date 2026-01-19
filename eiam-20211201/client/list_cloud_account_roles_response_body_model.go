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
	CloudAccountRoles []*ListCloudAccountRolesResponseBodyCloudAccountRoles `json:"CloudAccountRoles,omitempty" xml:"CloudAccountRoles,omitempty" type:"Repeated"`
	// 分页查询时每页行数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 云账号ID
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// 云账号角色外部唯一ID
	//
	// example:
	//
	// acs:ram::xxx:role/role-test
	CloudAccountRoleExternalId *string `json:"CloudAccountRoleExternalId,omitempty" xml:"CloudAccountRoleExternalId,omitempty"`
	// 云账号角色可用性
	//
	// example:
	//
	// healthy
	CloudAccountRoleHealth            *string                                                                              `json:"CloudAccountRoleHealth,omitempty" xml:"CloudAccountRoleHealth,omitempty"`
	CloudAccountRoleHealthCheckResult *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResult `json:"CloudAccountRoleHealthCheckResult,omitempty" xml:"CloudAccountRoleHealthCheckResult,omitempty" type:"Struct"`
	// 云账号角色ID
	//
	// example:
	//
	// carole_01kmek49aqxxxx
	CloudAccountRoleId *string `json:"CloudAccountRoleId,omitempty" xml:"CloudAccountRoleId,omitempty"`
	// 云账号名称
	//
	// example:
	//
	// role-test
	CloudAccountRoleName *string `json:"CloudAccountRoleName,omitempty" xml:"CloudAccountRoleName,omitempty"`
	// 云账号角色用途
	//
	// example:
	//
	// role
	CloudAccountRoleType *string `json:"CloudAccountRoleType,omitempty" xml:"CloudAccountRoleType,omitempty"`
	// 云账号角色类别
	//
	// example:
	//
	// system
	CloudAccountRoleUsageType *string `json:"CloudAccountRoleUsageType,omitempty" xml:"CloudAccountRoleUsageType,omitempty"`
	// example:
	//
	// 1719320115000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 云账号描述
	//
	// example:
	//
	// cloud_account_role_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// IDaaS EIAM 实例Id
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 云账号角色状态
	//
	// example:
	//
	// enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	ErrorReason *ListCloudAccountRolesResponseBodyCloudAccountRolesCloudAccountRoleHealthCheckResultErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// example:
	//
	// 1649830226000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
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
	// example:
	//
	// AuthenticationFail.NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
