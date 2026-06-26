// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyRelatedApprovalsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessTypesShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetAccessTypesShrink() *string
	SetDefSchema(v string) *ListMyRelatedApprovalsShrinkRequest
	GetDefSchema() *string
	SetEndTime(v int64) *ListMyRelatedApprovalsShrinkRequest
	GetEndTime() *int64
	SetGranteeShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetGranteeShrink() *string
	SetNextToken(v string) *ListMyRelatedApprovalsShrinkRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListMyRelatedApprovalsShrinkRequest
	GetPageSize() *int32
	SetResourceShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetResourceShrink() *string
	SetResourceTypeShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetResourceTypeShrink() *string
	SetStartTime(v int64) *ListMyRelatedApprovalsShrinkRequest
	GetStartTime() *int64
	SetStatusesShrink(v string) *ListMyRelatedApprovalsShrinkRequest
	GetStatusesShrink() *string
}

type ListMyRelatedApprovalsShrinkRequest struct {
	// Filter by requested permissions.
	//
	// Note: Different resource levels support different application permission types, all constrained by [ResourceSchema](https://help.aliyun.com/zh/dataworks/developer-reference/resourceschema-template-instructions).isValidLeaf, accessTypeRestrictions, and authMethodAccessTypes.
	//
	// Reference: [ResourceSchema International Site Documentation](https://www.alibabacloud.com/help/zh/dataworks/developer-reference/resourceschema-template-instructions)
	AccessTypesShrink *string `json:"AccessTypes,omitempty" xml:"AccessTypes,omitempty"`
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
	GranteeShrink *string `json:"Grantee,omitempty" xml:"Grantee,omitempty"`
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
	ResourceShrink *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
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
	ResourceTypeShrink *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
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
	StatusesShrink *string `json:"Statuses,omitempty" xml:"Statuses,omitempty"`
}

func (s ListMyRelatedApprovalsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyRelatedApprovalsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetAccessTypesShrink() *string {
	return s.AccessTypesShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetDefSchema() *string {
	return s.DefSchema
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetGranteeShrink() *string {
	return s.GranteeShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetResourceShrink() *string {
	return s.ResourceShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetResourceTypeShrink() *string {
	return s.ResourceTypeShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListMyRelatedApprovalsShrinkRequest) GetStatusesShrink() *string {
	return s.StatusesShrink
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetAccessTypesShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.AccessTypesShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetDefSchema(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.DefSchema = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetEndTime(v int64) *ListMyRelatedApprovalsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetGranteeShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.GranteeShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetNextToken(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetPageSize(v int32) *ListMyRelatedApprovalsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetResourceShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.ResourceShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetResourceTypeShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.ResourceTypeShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetStartTime(v int64) *ListMyRelatedApprovalsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) SetStatusesShrink(v string) *ListMyRelatedApprovalsShrinkRequest {
	s.StatusesShrink = &v
	return s
}

func (s *ListMyRelatedApprovalsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
