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
	Data *GetADBSparkNecessaryRAMPermissionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	return dara.Validate(s)
}

type GetADBSparkNecessaryRAMPermissionsResponseBodyData struct {
	DeniedDetail *GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail `json:"DeniedDetail,omitempty" xml:"DeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// true
	Passed *bool `json:"Passed,omitempty" xml:"Passed,omitempty"`
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
	return dara.Validate(s)
}

type GetADBSparkNecessaryRAMPermissionsResponseBodyDataDeniedDetail struct {
	// example:
	//
	// ListSparkApps
	Action *string `json:"Action,omitempty" xml:"Action,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// example:
	//
	// ControlPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// example:
	//
	// SubUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// example:
	//
	// 223345695632****
	ResourceAuthTargetInfo *string `json:"ResourceAuthTargetInfo,omitempty" xml:"ResourceAuthTargetInfo,omitempty"`
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
