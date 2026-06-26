// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyRelatedApprovalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTypes(v []*string) *ListMyRelatedApprovalsRequest
	GetAccessTypes() []*string
	SetDefSchema(v string) *ListMyRelatedApprovalsRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListMyRelatedApprovalsRequest
	GetEndTime() *int64
	SetGrantee(v *ListMyRelatedApprovalsRequestGrantee) *ListMyRelatedApprovalsRequest
	GetGrantee() *ListMyRelatedApprovalsRequestGrantee
	SetNextToken(v string) *ListMyRelatedApprovalsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMyRelatedApprovalsRequest
	GetPageSize() *int32
	SetResource(v *ListMyRelatedApprovalsRequestResource) *ListMyRelatedApprovalsRequest
	GetResource() *ListMyRelatedApprovalsRequestResource
	SetResourceType(v []*string) *ListMyRelatedApprovalsRequest
	GetResourceType() []*string
	SetStartTime(v int64) *ListMyRelatedApprovalsRequest
	GetStartTime() *int64
	SetStatuses(v []*string) *ListMyRelatedApprovalsRequest
	GetStatuses() []*string
}

type ListMyRelatedApprovalsRequest struct {
	// Filter by requested permissions.
	//
	// Note: Different resource levels support different application permission types, all constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).isValidLeaf, accessTypeRestrictions, and authMethodAccessTypes.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// Filter by resource type.
	//
	// Note: The resource types supported by the system for applications are constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).name.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// Application time end (millisecond timestamp)
	//
	// example:
	//
	// 1779724799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Filter by authorization principal.
	//
	// Note: The authorization principal types supported by the system are constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).authPrincipal.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	Grantee *ListMyRelatedApprovalsRequestGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// Pagination cursor
	//
	// example:
	//
	// eyJpZCI6MTIzfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Page size (default 10, maximum 200)
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filter by resource with exact/generalized matching. The resource description is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	Resource *ListMyRelatedApprovalsRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Filter by minimum permission resource type.
	//
	// Note: The minimum permission resource type is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).resources[*].isValidLeaf being true.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	//
	// example:
	//
	// ["table", "column"]
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// Application time start (millisecond timestamp)
	//
	// example:
	//
	// 1771948800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Filter by approval status. Enum values:
	//
	// - WaitApproval: Pending approval
	//
	// - Confirmed: Pending authorization
	//
	// - RejectApproval: Approval rejected
	//
	// - AuthorizeSucceed: Authorization succeeded
	//
	// - AuthorizeFailed: Authorization failed
	//
	// - Deleted: Deleted
	//
	// - Canceled: Withdrawn
	//
	// example:
	//
	// WAIT_APPROVAL
	Statuses []*string `json:"Statuses,omitempty" xml:"Statuses,omitempty" type:"Repeated"`
}

func (s ListMyRelatedApprovalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsRequest) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsRequest) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ListMyRelatedApprovalsRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMyRelatedApprovalsRequest) GetGrantee() *ListMyRelatedApprovalsRequestGrantee {
	return s.Grantee
}

func (s *ListMyRelatedApprovalsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyRelatedApprovalsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyRelatedApprovalsRequest) GetResource() *ListMyRelatedApprovalsRequestResource {
	return s.Resource
}

func (s *ListMyRelatedApprovalsRequest) GetResourceType() []*string {
	return s.ResourceType
}

func (s *ListMyRelatedApprovalsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMyRelatedApprovalsRequest) GetStatuses() []*string {
	return s.Statuses
}

func (s *ListMyRelatedApprovalsRequest) SetAccessTypes(v []*string) *ListMyRelatedApprovalsRequest {
	s.AccessTypes = v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetDefSchema(v string) *ListMyRelatedApprovalsRequest {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetEndTime(v int64) *ListMyRelatedApprovalsRequest {
	s.EndTime = &v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetGrantee(v *ListMyRelatedApprovalsRequestGrantee) *ListMyRelatedApprovalsRequest {
	s.Grantee = v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetNextToken(v string) *ListMyRelatedApprovalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetPageSize(v int32) *ListMyRelatedApprovalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetResource(v *ListMyRelatedApprovalsRequestResource) *ListMyRelatedApprovalsRequest {
	s.Resource = v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetResourceType(v []*string) *ListMyRelatedApprovalsRequest {
	s.ResourceType = v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetStartTime(v int64) *ListMyRelatedApprovalsRequest {
	s.StartTime = &v
	return s
}

func (s *ListMyRelatedApprovalsRequest) SetStatuses(v []*string) *ListMyRelatedApprovalsRequest {
	s.Statuses = v
	return s
}

func (s *ListMyRelatedApprovalsRequest) Validate() error {
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

type ListMyRelatedApprovalsRequestGrantee struct {
	// Authorization principal ID:
	//
	// - `RamUser`: Dataworks UserId
	//
	// - `RamRole`: Dataworks UserId prefixed with "ROLE_"
	//
	// - `DataworksTenantMember`: Dataworks UserId
	//
	// - `DataworksTenantRole`: Dataworks tenant roleCode
	//
	// - `DataworksProjectRole`: Dataworks workspace roleCode
	//
	// - `DataworksProjectMember`: Dataworks UserId
	//
	// - `DlfRole`: DlfNext role name
	//
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// Authorization principal type:
	//
	// - `RamRole`
	//
	// - `RamUser`
	//
	// - `DataworksTenantMember`
	//
	// - `DataworksTenantRole`
	//
	// - `DataworksProjectMember`
	//
	// - `DataworksProjectRole`
	//
	// - `DlfRole`
	//
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListMyRelatedApprovalsRequestGrantee) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsRequestGrantee) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsRequestGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListMyRelatedApprovalsRequestGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListMyRelatedApprovalsRequestGrantee) SetPrincipalId(v string) *ListMyRelatedApprovalsRequestGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ListMyRelatedApprovalsRequestGrantee) SetPrincipalType(v string) *ListMyRelatedApprovalsRequestGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ListMyRelatedApprovalsRequestGrantee) Validate() error {
	return dara.Validate(s)
}

type ListMyRelatedApprovalsRequestResource struct {
	// Resource type.
	//
	// Note: The resource types supported by the system for applications are constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).name.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The resource parsing version is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).version.
	//
	// [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// Resource metadata.
	//
	// Note: The metadata is constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).resources. A valid resource declaration must include the full-path metadata declaration from level 0 to validLeaf layer.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	MetaData map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ListMyRelatedApprovalsRequestResource) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsRequestResource) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsRequestResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsRequestResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListMyRelatedApprovalsRequestResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListMyRelatedApprovalsRequestResource) SetDefSchema(v string) *ListMyRelatedApprovalsRequestResource {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsRequestResource) SetDefVersion(v string) *ListMyRelatedApprovalsRequestResource {
	s.DefVersion = &v
	return s
}

func (s *ListMyRelatedApprovalsRequestResource) SetMetaData(v map[string]interface{}) *ListMyRelatedApprovalsRequestResource {
	s.MetaData = v
	return s
}

func (s *ListMyRelatedApprovalsRequestResource) Validate() error {
	return dara.Validate(s)
}
