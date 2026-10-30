// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgenticSecurityPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListAgenticSecurityPoliciesResponseBodyData) *ListAgenticSecurityPoliciesResponseBody
	GetData() []*ListAgenticSecurityPoliciesResponseBodyData
	SetErrorCode(v string) *ListAgenticSecurityPoliciesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListAgenticSecurityPoliciesResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListAgenticSecurityPoliciesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgenticSecurityPoliciesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListAgenticSecurityPoliciesResponseBody
	GetTotalCount() *int64
}

type ListAgenticSecurityPoliciesResponseBody struct {
	// The list of security policies returned by this query. Enabled system default policies are placed at the beginning, followed by the user policies on the current page. You can use Readonly to distinguish between the two types of policies.
	Data []*ListAgenticSecurityPoliciesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when the request fails. This parameter is empty when the request is successful.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request fails. This parameter is empty when the request is successful.
	//
	// example:
	//
	// invalid policyStatus: ARCHIVED
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate and trace this call.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of policies that meet the query conditions. The value is the sum of the number of all enabled system default policies and the number of user policies of the current account that meet the filter conditions.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAgenticSecurityPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSecurityPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetData() []*ListAgenticSecurityPoliciesResponseBodyData {
	return s.Data
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgenticSecurityPoliciesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetData(v []*ListAgenticSecurityPoliciesResponseBodyData) *ListAgenticSecurityPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetErrorCode(v string) *ListAgenticSecurityPoliciesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetErrorMessage(v string) *ListAgenticSecurityPoliciesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetRequestId(v string) *ListAgenticSecurityPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetSuccess(v bool) *ListAgenticSecurityPoliciesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) SetTotalCount(v int64) *ListAgenticSecurityPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgenticSecurityPoliciesResponseBodyData struct {
	// The UID of the Alibaba Cloud account to which the user policy belongs. This value is not returned for system default policies.
	//
	// example:
	//
	// 123456789012****
	AliyunAccountUid *string `json:"AliyunAccountUid,omitempty" xml:"AliyunAccountUid,omitempty"`
	// The ID of the system default policy inherited by the user policy. This value is empty if no inheritance relationship is configured. This value is not returned for system default policies.
	//
	// example:
	//
	// sp-sys-default
	BasePolicyId *string `json:"BasePolicyId,omitempty" xml:"BasePolicyId,omitempty"`
	// The description of the security policy. The description of a system default policy is returned in the language of the request context.
	//
	// example:
	//
	// Read-only access policy for analytics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the security policy was created. The time is in the yyyy-MM-dd HH:mm:ss format and does not contain a time zone offset.
	//
	// example:
	//
	// 2026-10-13 14:30:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the security policy was last modified. The time is in the yyyy-MM-dd HH:mm:ss format and does not contain a time zone offset.
	//
	// example:
	//
	// 2026-10-13 14:35:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The business ID of the security policy. A system default policy ID starts with sp-sys-, and a user policy ID starts with sp-usr-.
	//
	// example:
	//
	// sp-sys-default
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The Policy Name of the security policy. The Policy Name of a system default policy is returned in the language of the request context.
	//
	// example:
	//
	// team-readonly-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The current state of the security policy. Valid values:
	//
	// - ON: Enabled.
	//
	// - OFF: Disabled.
	//
	// - ARCHIVED: Archived.
	//
	// All system default policies returned by this operation are in the ON state.
	//
	// example:
	//
	// ON
	PolicyStatus *string `json:"PolicyStatus,omitempty" xml:"PolicyStatus,omitempty"`
	// Specifies whether the policy is read-only. Valid values:
	//
	// - true: The policy is a system default policy.
	//
	// - false: The policy is a user policy of the current account.
	Readonly *bool `json:"Readonly,omitempty" xml:"Readonly,omitempty"`
	// The tags of the system default policy. Multiple tags are separated by commas (,). This value is not returned for user policies.
	//
	// example:
	//
	// default,phase1
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The identifier of the operator that last modified the policy. For a user policy, the operator ID is recorded first. If no operator ID is available, the UID of the current Alibaba Cloud account is recorded.
	//
	// example:
	//
	// system
	UpdatedBy *string `json:"UpdatedBy,omitempty" xml:"UpdatedBy,omitempty"`
}

func (s ListAgenticSecurityPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgenticSecurityPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetAliyunAccountUid() *string {
	return s.AliyunAccountUid
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetBasePolicyId() *string {
	return s.BasePolicyId
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetPolicyId() *string {
	return s.PolicyId
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetPolicyStatus() *string {
	return s.PolicyStatus
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetReadonly() *bool {
	return s.Readonly
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetTags() *string {
	return s.Tags
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) GetUpdatedBy() *string {
	return s.UpdatedBy
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetAliyunAccountUid(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.AliyunAccountUid = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetBasePolicyId(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.BasePolicyId = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetDescription(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetGmtCreate(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.GmtCreate = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetGmtModified(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetPolicyId(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.PolicyId = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetPolicyName(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.PolicyName = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetPolicyStatus(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.PolicyStatus = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetReadonly(v bool) *ListAgenticSecurityPoliciesResponseBodyData {
	s.Readonly = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetTags(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.Tags = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) SetUpdatedBy(v string) *ListAgenticSecurityPoliciesResponseBodyData {
	s.UpdatedBy = &v
	return s
}

func (s *ListAgenticSecurityPoliciesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
