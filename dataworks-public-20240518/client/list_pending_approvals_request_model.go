// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPendingApprovalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTypes(v []*string) *ListPendingApprovalsRequest
	GetAccessTypes() []*string
	SetDefSchema(v string) *ListPendingApprovalsRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListPendingApprovalsRequest
	GetEndTime() *int64
	SetGrantee(v *ListPendingApprovalsRequestGrantee) *ListPendingApprovalsRequest
	GetGrantee() *ListPendingApprovalsRequestGrantee
	SetNextToken(v string) *ListPendingApprovalsRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListPendingApprovalsRequest
	GetPageSize() *int32
	SetResource(v *ListPendingApprovalsRequestResource) *ListPendingApprovalsRequest
	GetResource() *ListPendingApprovalsRequestResource
	SetResourceType(v []*string) *ListPendingApprovalsRequest
	GetResourceType() []*string
	SetStartTime(v int64) *ListPendingApprovalsRequest
	GetStartTime() *int64
}

type ListPendingApprovalsRequest struct {
	// Filters by requested permissions.
	//
	// Note: The permission types supported for application vary across resource levels and are uniformly constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).isValidLeaf, accessTypeRestrictions, and authMethodAccessTypes.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	AccessTypes []*string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty" type:"Repeated"`
	// Filters by resource type.
	//
	// Note: The resource types for which permissions can be requested are constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).name.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The end of the time range for request submission, as a timestamp in milliseconds.
	//
	// example:
	//
	// 1779724799999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Filters by authorization principal.
	//
	// Note: The authorization principal types supported by the system are constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).authPrincipal.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	Grantee *ListPendingApprovalsRequestGrantee `json:"Grantee,omitempty" xml:"Grantee,omitempty" type:"Struct"`
	// The cursor.
	//
	// example:
	//
	// eyJpZCI6MTIzfQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 200.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by exact or generalized resource. The resource description is constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	Resource *ListPendingApprovalsRequestResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Filters by least privilege resource type.
	//
	// Note: The least privilege resource types are constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).resources[*].isValidLeaf being true.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// This parameter is required.
	//
	// example:
	//
	// ["table", "column"]
	ResourceType []*string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" type:"Repeated"`
	// The start of the time range for request submission, as a timestamp in milliseconds.
	//
	// example:
	//
	// 1771948800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListPendingApprovalsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsRequest) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsRequest) GetAccessTypes() []*string {
	return s.AccessTypes
}

func (s *ListPendingApprovalsRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListPendingApprovalsRequest) GetGrantee() *ListPendingApprovalsRequestGrantee {
	return s.Grantee
}

func (s *ListPendingApprovalsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPendingApprovalsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPendingApprovalsRequest) GetResource() *ListPendingApprovalsRequestResource {
	return s.Resource
}

func (s *ListPendingApprovalsRequest) GetResourceType() []*string {
	return s.ResourceType
}

func (s *ListPendingApprovalsRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListPendingApprovalsRequest) SetAccessTypes(v []*string) *ListPendingApprovalsRequest {
	s.AccessTypes = v
	return s
}

func (s *ListPendingApprovalsRequest) SetDefSchema(v string) *ListPendingApprovalsRequest {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsRequest) SetEndTime(v int64) *ListPendingApprovalsRequest {
	s.EndTime = &v
	return s
}

func (s *ListPendingApprovalsRequest) SetGrantee(v *ListPendingApprovalsRequestGrantee) *ListPendingApprovalsRequest {
	s.Grantee = v
	return s
}

func (s *ListPendingApprovalsRequest) SetNextToken(v string) *ListPendingApprovalsRequest {
	s.NextToken = &v
	return s
}

func (s *ListPendingApprovalsRequest) SetPageSize(v int32) *ListPendingApprovalsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPendingApprovalsRequest) SetResource(v *ListPendingApprovalsRequestResource) *ListPendingApprovalsRequest {
	s.Resource = v
	return s
}

func (s *ListPendingApprovalsRequest) SetResourceType(v []*string) *ListPendingApprovalsRequest {
	s.ResourceType = v
	return s
}

func (s *ListPendingApprovalsRequest) SetStartTime(v int64) *ListPendingApprovalsRequest {
	s.StartTime = &v
	return s
}

func (s *ListPendingApprovalsRequest) Validate() error {
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

type ListPendingApprovalsRequestGrantee struct {
	// The ID of the authorization principal.
	//
	// example:
	//
	// ROLE_3133343434
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// The type of the authorization principal.
	//
	// example:
	//
	// RamRole
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s ListPendingApprovalsRequestGrantee) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsRequestGrantee) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsRequestGrantee) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *ListPendingApprovalsRequestGrantee) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *ListPendingApprovalsRequestGrantee) SetPrincipalId(v string) *ListPendingApprovalsRequestGrantee {
	s.PrincipalId = &v
	return s
}

func (s *ListPendingApprovalsRequestGrantee) SetPrincipalType(v string) *ListPendingApprovalsRequestGrantee {
	s.PrincipalType = &v
	return s
}

func (s *ListPendingApprovalsRequestGrantee) Validate() error {
	return dara.Validate(s)
}

type ListPendingApprovalsRequestResource struct {
	// The resource type.
	//
	// Note: The resource types for which permissions can be requested are constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).name.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// MaxCompute
	DefSchema *string `json:"DefSchema,omitempty" xml:"DefSchema,omitempty"`
	// The resource parsing version is constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).version.
	//
	// [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// v1.0.0
	DefVersion *string `json:"DefVersion,omitempty" xml:"DefVersion,omitempty"`
	// The resource metadata.
	//
	// Note: The metadata is constrained by [ResourceSchema](https://www.alibabacloud.com/help/en/dataworks/developer-reference/resourceschema-template-instructions).resources. A valid resource declaration must include the full-path metadata declaration from level 0 to the validLeaf level.
	//
	// Appendix: [ResourceSchema documentation for the international site](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	//
	// example:
	//
	// {"workspace": "449056", "project": "sync_destination"}
	MetaData map[string]interface{} `json:"MetaData,omitempty" xml:"MetaData,omitempty"`
}

func (s ListPendingApprovalsRequestResource) String() string {
	return dara.Prettify(s)
}

func (s ListPendingApprovalsRequestResource) GoString() string {
	return s.String()
}

func (s *ListPendingApprovalsRequestResource) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListPendingApprovalsRequestResource) GetDefVersion() *string {
	return s.DefVersion
}

func (s *ListPendingApprovalsRequestResource) GetMetaData() map[string]interface{} {
	return s.MetaData
}

func (s *ListPendingApprovalsRequestResource) SetDefSchema(v string) *ListPendingApprovalsRequestResource {
	s.DefSchema = &v
	return s
}

func (s *ListPendingApprovalsRequestResource) SetDefVersion(v string) *ListPendingApprovalsRequestResource {
	s.DefVersion = &v
	return s
}

func (s *ListPendingApprovalsRequestResource) SetMetaData(v map[string]interface{}) *ListPendingApprovalsRequestResource {
	s.MetaData = v
	return s
}

func (s *ListPendingApprovalsRequestResource) Validate() error {
	return dara.Validate(s)
}
