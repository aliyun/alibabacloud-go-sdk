// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyRelatedApprovalsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMyRelatedApprovalsResponseBodyData) *ListMyRelatedApprovalsResponseBody
	GetData() *ListMyRelatedApprovalsResponseBodyData
	SetRequestId(v string) *ListMyRelatedApprovalsResponseBody
	GetRequestId() *string
}

type ListMyRelatedApprovalsResponseBody struct {
	// The paginated results.
	Data *ListMyRelatedApprovalsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the API request. The ID is a UUID.
	//
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBody) GetData() *ListMyRelatedApprovalsResponseBodyData {
	return s.Data
}

func (s *ListMyRelatedApprovalsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMyRelatedApprovalsResponseBody) SetData(v *ListMyRelatedApprovalsResponseBodyData) *ListMyRelatedApprovalsResponseBody {
	s.Data = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBody) SetRequestId(v string) *ListMyRelatedApprovalsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyRelatedApprovalsResponseBodyData struct {
	// The permission request details.
	Data []*ListMyRelatedApprovalsResponseBodyDataData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Indicates whether more data exists.
	//
	// example:
	//
	// false
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The pagination cursor.
	//
	// example:
	//
	// eyJpZCI6NDU2fQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 200.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBodyData) GetData() []*ListMyRelatedApprovalsResponseBodyDataData {
	return s.Data
}

func (s *ListMyRelatedApprovalsResponseBodyData) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListMyRelatedApprovalsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyRelatedApprovalsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyRelatedApprovalsResponseBodyData) SetData(v []*ListMyRelatedApprovalsResponseBodyDataData) *ListMyRelatedApprovalsResponseBodyData {
	s.Data = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyData) SetHasMore(v bool) *ListMyRelatedApprovalsResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyData) SetNextToken(v string) *ListMyRelatedApprovalsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyData) SetPageSize(v int32) *ListMyRelatedApprovalsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyData) Validate() error {
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

type ListMyRelatedApprovalsResponseBodyDataData struct {
	// The time when the request was submitted.
	//
	// example:
	//
	// 1779695088000
	ApplicationTime *int64 `json:"ApplicationTime,omitempty" xml:"ApplicationTime,omitempty"`
	// The request details.
	Contents []*ListMyRelatedApprovalsResponseBodyDataDataContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The process instance ID.
	//
	// example:
	//
	// 176906667488145
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// The reason for the request.
	//
	// example:
	//
	// Business requirement
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Filters by approval status. Valid values:
	//
	// - WaitApproval: Pending approval.
	//
	// - Confirmed: Pending authorization.
	//
	// - RejectApproval: Rejected.
	//
	// - AuthorizeSucceed: Authorized.
	//
	// - AuthorizeFailed: Authorization failed.
	//
	// - Deleted: Deleted.
	//
	// - Canceled: Withdrawn.
	//
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetApplicationTime() *int64 {
	return s.ApplicationTime
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetContents() []*ListMyRelatedApprovalsResponseBodyDataDataContents {
	return s.Contents
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetReason() *string {
	return s.Reason
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) GetStatus() *string {
	return s.Status
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetApplicationTime(v int64) *ListMyRelatedApprovalsResponseBodyDataData {
	s.ApplicationTime = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetContents(v []*ListMyRelatedApprovalsResponseBodyDataDataContents) *ListMyRelatedApprovalsResponseBodyDataData {
	s.Contents = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetDefSchema(v string) *ListMyRelatedApprovalsResponseBodyDataData {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetProcessInstanceId(v string) *ListMyRelatedApprovalsResponseBodyDataData {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetReason(v string) *ListMyRelatedApprovalsResponseBodyDataData {
	s.Reason = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) SetStatus(v string) *ListMyRelatedApprovalsResponseBodyDataData {
	s.Status = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataData) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMyRelatedApprovalsResponseBodyDataDataContents struct {
	// The permissions requested for the resource operations.
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// The authorization method.
	//
	// example:
	//
	// default
	AuthMethod *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2025-09-11 10:13:21
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource type.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The expiration time of the permissions. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1779695088000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The resource operation permissions that are finally approved.
	FinalAccessTypes []*string `json:"FinalAccessTypes,omitempty" xml:"FinalAccessTypes,omitempty" type:"Repeated"`
	// The grantee.
	Grantee *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// The unique identifier of the application content.
	//
	// example:
	//
	// 1009516415
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the approval process instance for the request.
	//
	// example:
	//
	// 777799223
	ProcessInstanceId *string `json:"ProcessInstanceId,omitempty" xml:"ProcessInstanceId,omitempty"`
	// The resource declaration.
	Resource *ListMyRelatedApprovalsResponseBodyDataDataContentsResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// The least privilege resource type.
	//
	// example:
	//
	// table
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// Filters by approval status. Valid values:
	//
	// - WaitApproval: Pending approval.
	//
	// - Confirmed: Pending authorization.
	//
	// - RejectApproval: Rejected.
	//
	// - AuthorizeSucceed: Authorized.
	//
	// - AuthorizeFailed: Authorization failed.
	//
	// - Deleted: Deleted.
	//
	// - Canceled: Withdrawn.
	//
	// example:
	//
	// Deleted
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 69973837489
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-03-05 18:24:13
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContents) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContents) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetFinalAccessTypes() []*string {
	return s.FinalAccessTypes
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetGrantee() *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee {
	return s.Grantee
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetId() *string {
	return s.Id
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetProcessInstanceId() *string {
	return s.ProcessInstanceId
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetResource() *ListMyRelatedApprovalsResponseBodyDataDataContentsResource {
	return s.Resource
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetStatus() *string {
	return s.Status
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetAccessTypes(v []*string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.AccessTypes = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetAuthMethod(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.AuthMethod = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetCreateTime(v int64) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.CreateTime = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetDefSchema(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetExpirationTime(v int64) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.ExpirationTime = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetFinalAccessTypes(v []*string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.FinalAccessTypes = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetGrantee(v *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.Grantee = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetId(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.Id = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetProcessInstanceId(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.ProcessInstanceId = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetResource(v *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.Resource = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetResourceName(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.ResourceName = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetStatus(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.Status = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetTenantId(v string) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.TenantId = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) SetUpdateTime(v int64) *ListMyRelatedApprovalsResponseBodyDataDataContents {
	s.UpdateTime = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContents) Validate() error {
	if s.Grantee != nil {
		if err := s.Grantee.Validate(); err != nil {
			return err
		}
	}
	if s.Resource != nil {
		if err := s.Resource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee struct {
	// The ID of the principal.
	//
	// Note: The meaning of the ID varies based on the principal type:
	//
	// - RamUser: DataWorks user ID.
	//
	// - RamRole: DataWorks user ID that starts with "ROLE_".
	//
	// - DataworksTenantMember: DataWorks user ID.
	//
	// - DataworksTenantRole: DataWorks tenant role code.
	//
	// - DataworksProjectRole: DataWorks workspace role code.
	//
	// - DataworksProjectMember: DataWorks user ID.
	//
	// - DlfRole: DLF Next role name.
	//
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the principal. Valid values:
	//
	// - RamRole
	//
	// - RamUser
	//
	// - DataworksTenantMember
	//
	// - DataworksTenantRole
	//
	// - DataworksProjectMember
	//
	// - DataworksProjectRole
	//
	// - DlfRole
	//
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) SetPrincipalId(v string) *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) SetPrincipalType(v string) *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsGrantee) Validate() error {
	return dara.Validate(s)
}

type ListMyRelatedApprovalsResponseBodyDataDataContentsResource struct {
	// The ResourceSchema.name that the resource parsing depends on.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The ResourceSchema.version that the resource parsing depends on.
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// The resource metadata. The data content is constrained by ResourceSchema.
	//
	// example:
	//
	// {"schema": "default", "threeTierModel": false, "project": "sync_dtestn", "workspace": "4477656", "table": "order", "tenant": "524257424864736"}
	MetaData map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContentsResource) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsResponseBodyDataDataContentsResource) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) SetDefSchema(v string) *ListMyRelatedApprovalsResponseBodyDataDataContentsResource {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) SetDefVersion(v string) *ListMyRelatedApprovalsResponseBodyDataDataContentsResource {
	s.DefVersion = &v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) SetMetaData(v map[string]interface{}) *ListMyRelatedApprovalsResponseBodyDataDataContentsResource {
	s.MetaData = v
	return s
}

func (s *ListMyRelatedApprovalsResponseBodyDataDataContentsResource) Validate() error {
	return dara.Validate(s)
}
