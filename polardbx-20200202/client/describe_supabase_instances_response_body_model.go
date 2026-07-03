// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupabaseInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) *DescribeSupabaseInstancesResponseBody
	GetAccessDeniedDetail() *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail
	SetData(v *DescribeSupabaseInstancesResponseBodyData) *DescribeSupabaseInstancesResponseBody
	GetData() *DescribeSupabaseInstancesResponseBodyData
	SetMaxResults(v int32) *DescribeSupabaseInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSupabaseInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeSupabaseInstancesResponseBody
	GetRequestId() *string
}

type DescribeSupabaseInstancesResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The paginated result of the instance list.
	Data *DescribeSupabaseInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries per page for a paged query. Maximum value: 100. Default value: If you do not specify this parameter or specify a value less than 10, the default value is 10. If you specify a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 1000
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next query.
	//
	// example:
	//
	// xxdds
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C457B28E-9CAB-4B77-B5C6-5D71B7870B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSupabaseInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesResponseBody) GetAccessDeniedDetail() *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeSupabaseInstancesResponseBody) GetData() *DescribeSupabaseInstancesResponseBodyData {
	return s.Data
}

func (s *DescribeSupabaseInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSupabaseInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSupabaseInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSupabaseInstancesResponseBody) SetAccessDeniedDetail(v *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) *DescribeSupabaseInstancesResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeSupabaseInstancesResponseBody) SetData(v *DescribeSupabaseInstancesResponseBodyData) *DescribeSupabaseInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSupabaseInstancesResponseBody) SetMaxResults(v int32) *DescribeSupabaseInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBody) SetNextToken(v string) *DescribeSupabaseInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBody) SetRequestId(v string) *DescribeSupabaseInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSupabaseInstancesResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The authentication principal type.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAacnceTlBNjlFODgyLTlBMDUtNUUyRC04MzE5LUQwMzZDRjJFQTM3NA==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the permission denial.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeSupabaseInstancesResponseBodyData struct {
	// The list of instances.
	Instances []*DescribeSupabaseInstancesResponseBodyDataInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalNumber *int32 `json:"TotalNumber,omitempty" xml:"TotalNumber,omitempty"`
}

func (s DescribeSupabaseInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesResponseBodyData) GetInstances() []*DescribeSupabaseInstancesResponseBodyDataInstances {
	return s.Instances
}

func (s *DescribeSupabaseInstancesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSupabaseInstancesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSupabaseInstancesResponseBodyData) GetTotalNumber() *int32 {
	return s.TotalNumber
}

func (s *DescribeSupabaseInstancesResponseBodyData) SetInstances(v []*DescribeSupabaseInstancesResponseBodyDataInstances) *DescribeSupabaseInstancesResponseBodyData {
	s.Instances = v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyData) SetPageNumber(v int32) *DescribeSupabaseInstancesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyData) SetPageSize(v int32) *DescribeSupabaseInstancesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyData) SetTotalNumber(v int32) *DescribeSupabaseInstancesResponseBodyData {
	s.TotalNumber = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyData) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSupabaseInstancesResponseBodyDataInstances struct {
	// The creation time.
	//
	// example:
	//
	// 2026-06-08T07:19:05.000+0000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance name.
	//
	// example:
	//
	// pxsp-jinwusp30
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The instance description.
	//
	// example:
	//
	// 我的 Supabase 项目
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance status.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Supabase specification code.
	//
	// example:
	//
	// polarx.supabase.x4.2xlarge
	SupabaseClassCode *string `json:"SupabaseClassCode,omitempty" xml:"SupabaseClassCode,omitempty"`
	// The multi-tenant mode.
	//
	// example:
	//
	// MySQL
	TenantMode *bool `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-beijing-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeSupabaseInstancesResponseBodyDataInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupabaseInstancesResponseBodyDataInstances) GoString() string {
	return s.String()
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetDescription() *string {
	return s.Description
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetStatus() *string {
	return s.Status
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetSupabaseClassCode() *string {
	return s.SupabaseClassCode
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetTenantMode() *bool {
	return s.TenantMode
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetCreateTime(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetDBInstanceName(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetDescription(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.Description = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetRegionId(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetStatus(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.Status = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetSupabaseClassCode(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.SupabaseClassCode = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetTenantMode(v bool) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.TenantMode = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) SetZoneId(v string) *DescribeSupabaseInstancesResponseBodyDataInstances {
	s.ZoneId = &v
	return s
}

func (s *DescribeSupabaseInstancesResponseBodyDataInstances) Validate() error {
	return dara.Validate(s)
}
