// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccounts(v []*ListCloudAccountsResponseBodyCloudAccounts) *ListCloudAccountsResponseBody
	GetCloudAccounts() []*ListCloudAccountsResponseBodyCloudAccounts
	SetMaxResults(v int32) *ListCloudAccountsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCloudAccountsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCloudAccountsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCloudAccountsResponseBody
	GetTotalCount() *int32
}

type ListCloudAccountsResponseBody struct {
	// The list of cloud accounts.
	CloudAccounts []*ListCloudAccountsResponseBodyCloudAccounts `json:"CloudAccounts,omitempty" xml:"CloudAccounts,omitempty" type:"Repeated"`
	// The number of rows per page for paging.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned for the current call.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsResponseBody) GetCloudAccounts() []*ListCloudAccountsResponseBodyCloudAccounts {
	return s.CloudAccounts
}

func (s *ListCloudAccountsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCloudAccountsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCloudAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAccountsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloudAccountsResponseBody) SetCloudAccounts(v []*ListCloudAccountsResponseBodyCloudAccounts) *ListCloudAccountsResponseBody {
	s.CloudAccounts = v
	return s
}

func (s *ListCloudAccountsResponseBody) SetMaxResults(v int32) *ListCloudAccountsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCloudAccountsResponseBody) SetNextToken(v string) *ListCloudAccountsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCloudAccountsResponseBody) SetRequestId(v string) *ListCloudAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAccountsResponseBody) SetTotalCount(v int32) *ListCloudAccountsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAccountsResponseBody) Validate() error {
	if s.CloudAccounts != nil {
		for _, item := range s.CloudAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAccountsResponseBodyCloudAccounts struct {
	// The external unique identifier of the cloud account.
	//
	// example:
	//
	// 1234567
	CloudAccountExternalId *string `json:"CloudAccountExternalId,omitempty" xml:"CloudAccountExternalId,omitempty"`
	// The health status of the cloud account. Valid values:
	//
	// - healthy: Healthy.
	//
	// - unhealthy: Unhealthy.
	//
	// - unknown: Unknown.
	//
	// example:
	//
	// healthy
	CloudAccountHealth *string `json:"CloudAccountHealth,omitempty" xml:"CloudAccountHealth,omitempty"`
	// The health check result of the cloud account.
	CloudAccountHealthCheckResult *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult `json:"CloudAccountHealthCheckResult,omitempty" xml:"CloudAccountHealthCheckResult,omitempty" type:"Struct"`
	// The cloud account ID.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The cloud account name.
	//
	// example:
	//
	// cloud_accout_xxxx
	CloudAccountName *string `json:"CloudAccountName,omitempty" xml:"CloudAccountName,omitempty"`
	// The identity provider name.
	//
	// example:
	//
	// idaas-eiam-oidc-provider
	CloudAccountProviderName *string `json:"CloudAccountProviderName,omitempty" xml:"CloudAccountProviderName,omitempty"`
	CloudAccountSite         *string `json:"CloudAccountSite,omitempty" xml:"CloudAccountSite,omitempty"`
	// The cloud account type. Valid values:
	//
	// - alibaba_cloud: Alibaba Cloud.
	//
	// example:
	//
	// alibaba_cloud
	CloudAccountVendorType *string `json:"CloudAccountVendorType,omitempty" xml:"CloudAccountVendorType,omitempty"`
	// The time when the cloud account was created. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830225000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the cloud account.
	//
	// example:
	//
	// cloud_accout_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The time when the cloud account was last updated. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830227000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListCloudAccountsResponseBodyCloudAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsResponseBodyCloudAccounts) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountExternalId() *string {
	return s.CloudAccountExternalId
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountHealth() *string {
	return s.CloudAccountHealth
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountHealthCheckResult() *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult {
	return s.CloudAccountHealthCheckResult
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountName() *string {
	return s.CloudAccountName
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountProviderName() *string {
	return s.CloudAccountProviderName
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountSite() *string {
	return s.CloudAccountSite
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCloudAccountVendorType() *string {
	return s.CloudAccountVendorType
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetDescription() *string {
	return s.Description
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountExternalId(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountExternalId = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountHealth(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountHealth = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountHealthCheckResult(v *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountHealthCheckResult = v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountId(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountId = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountName(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountName = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountProviderName(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountProviderName = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountSite(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountSite = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCloudAccountVendorType(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CloudAccountVendorType = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetCreateTime(v int64) *ListCloudAccountsResponseBodyCloudAccounts {
	s.CreateTime = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetDescription(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.Description = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetInstanceId(v string) *ListCloudAccountsResponseBodyCloudAccounts {
	s.InstanceId = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) SetUpdateTime(v int64) *ListCloudAccountsResponseBodyCloudAccounts {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccounts) Validate() error {
	if s.CloudAccountHealthCheckResult != nil {
		if err := s.CloudAccountHealthCheckResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult struct {
	// The error reason. This field is returned when the health check status is unhealthy.
	ErrorReason *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason `json:"ErrorReason,omitempty" xml:"ErrorReason,omitempty" type:"Struct"`
	// The time of the last health check. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1649830226000
	LastCheckTime *int64 `json:"LastCheckTime,omitempty" xml:"LastCheckTime,omitempty"`
	// The health check result of the cloud account. Valid values:
	//
	// - success: Succeeded.
	//
	// - failed: Failed.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) GetErrorReason() *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason {
	return s.ErrorReason
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) GetLastCheckTime() *int64 {
	return s.LastCheckTime
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) GetResult() *string {
	return s.Result
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) SetErrorReason(v *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult {
	s.ErrorReason = v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) SetLastCheckTime(v int64) *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult {
	s.LastCheckTime = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) SetResult(v string) *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult {
	s.Result = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResult) Validate() error {
	if s.ErrorReason != nil {
		if err := s.ErrorReason.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason struct {
	// The error code.
	//
	// example:
	//
	// AuthenticationFail.NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// There is no permission.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) GoString() string {
	return s.String()
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) SetErrorCode(v string) *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason {
	s.ErrorCode = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) SetErrorMessage(v string) *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason {
	s.ErrorMessage = &v
	return s
}

func (s *ListCloudAccountsResponseBodyCloudAccountsCloudAccountHealthCheckResultErrorReason) Validate() error {
	return dara.Validate(s)
}
