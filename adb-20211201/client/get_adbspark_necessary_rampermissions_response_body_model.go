// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetADBSparkNecessaryRAMPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetADBSparkNecessaryRAMPermissionsResponseBodyData) *GetADBSparkNecessaryRAMPermissionsResponseBody
	GetData() *GetADBSparkNecessaryRAMPermissionsResponseBodyData
	SetRequestId(v string) *GetADBSparkNecessaryRAMPermissionsResponseBody
	GetRequestId() *string
}

type GetADBSparkNecessaryRAMPermissionsResponseBody struct {
	// The returned result.
	Data *GetADBSparkNecessaryRAMPermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 642F3512-C628-5D0C-8815-F6670C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBody) GetData() *GetADBSparkNecessaryRAMPermissionsResponseBodyData {
	return s.Data
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBody) SetData(v *GetADBSparkNecessaryRAMPermissionsResponseBodyData) *GetADBSparkNecessaryRAMPermissionsResponseBody {
	s.Data = v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBody) SetRequestId(v string) *GetADBSparkNecessaryRAMPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetADBSparkNecessaryRAMPermissionsResponseBodyData struct {
	// When permission check fails, returns diagnostic information related to permission check failure.
	DeniedDetail *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail `json:"DeniedDetail,omitempty" xml:"DeniedDetail,omitempty" type:"Struct"`
	// Check whether the use has the basic permissions to use Analytic DB for Spark.
	//
	// 	- true: The check is passed and the basic permissions are granted.
	//
	// 	- false: The check fails and some permissions are missing.
	//
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
	// Based on diagnostic information, recommends configurations for customers to perform in the RAM system.
	//
	// example:
	//
	// Grant the system RAM policy \\"AliyunADBDeveloperAccess\\" to current RAM user can quickly solve this issue.
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) GetDeniedDetail() *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	return s.DeniedDetail
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) GetPassed() *bool {
	return s.Passed
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) GetSuggestion() *string {
	return s.Suggestion
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) SetDeniedDetail(v *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) *GetADBSparkNecessaryRAMPermissionsResponseBodyData {
	s.DeniedDetail = v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) SetPassed(v bool) *GetADBSparkNecessaryRAMPermissionsResponseBodyData {
	s.Passed = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) SetSuggestion(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyData {
	s.Suggestion = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyData) Validate() error {
	if s.DeniedDetail != nil {
		if err := s.DeniedDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail struct {
	// The name of the RAM action that failed the authentication.
	//
	// example:
	//
	// ListSparkApps
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// The type of the policy denial. Valid values:
	//
	// 	- ImplicitDeny: Resource owner has not configured relevant permission policies for the current user, default denial of unauthorized operations.
	//
	// 	- ExplicitDeny: RAM policies configured by the resource owner explicitly deny the current user access to corresponding resources
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The type of the policy that causes the access denied error.
	//
	// 	- ControlPolicy: control policy
	//
	// 	- SessionPolicy: Temporary Token additional permission policy
	//
	// 	- AssumeRolePolicy: RAM role trust policy
	//
	// 	- AccountLevelIdentityBasedPolicy: Principal policy within account authorization scope, including custom policies and system policies
	//
	// 	- ResourceGroupLevelIdentityBasedPolicy: Principal policy within resource group authorization scope, including custom policies and system policies.
	//
	// example:
	//
	// ControlPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The identity type of the current user. Valid values:
	//
	// 	- SubUser: a RAM user
	//
	// 	- AssumedRoleUser: a RAM role
	//
	// example:
	//
	// SubUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// Authentication object information, can be the current user\\"s RAM account ID, or the role information corresponding to the current visitor.
	//
	// example:
	//
	// 223345695632****
	ResourceAuthTargetInfo *string `json:"ResourceAuthTargetInfo,omitempty" xml:"ResourceAuthTargetInfo,omitempty"`
	// The ID of the resource owner.
	//
	// example:
	//
	// 11685695632****
	ResourceOwnerId *string `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GoString() string {
	return s.String()
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetAction() *string {
	return s.Action
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetResourceAuthTargetInfo() *string {
	return s.ResourceAuthTargetInfo
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetAction(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.Action = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetNoPermissionType(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetPolicyType(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetPrincipalType(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.PrincipalType = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetResourceAuthTargetInfo(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.ResourceAuthTargetInfo = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) SetResourceOwnerId(v string) *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail) Validate() error {
	return dara.Validate(s)
}
