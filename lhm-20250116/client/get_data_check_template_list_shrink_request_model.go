// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateListShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *GetDataCheckTemplateListShrinkRequest
	GetCheckType() *int32
	SetGroupBy(v string) *GetDataCheckTemplateListShrinkRequest
	GetGroupBy() *string
	SetIdListShrink(v string) *GetDataCheckTemplateListShrinkRequest
	GetIdListShrink() *string
	SetIsAdmin(v bool) *GetDataCheckTemplateListShrinkRequest
	GetIsAdmin() *bool
	SetIsBuiltin(v int32) *GetDataCheckTemplateListShrinkRequest
	GetIsBuiltin() *int32
	SetNeedTotalCount(v bool) *GetDataCheckTemplateListShrinkRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *GetDataCheckTemplateListShrinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *GetDataCheckTemplateListShrinkRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *GetDataCheckTemplateListShrinkRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetDataCheckTemplateListShrinkRequest
	GetPageSize() *int32
	SetRequestId(v string) *GetDataCheckTemplateListShrinkRequest
	GetRequestId() *string
	SetTemplateName(v string) *GetDataCheckTemplateListShrinkRequest
	GetTemplateName() *string
	SetTenantId(v string) *GetDataCheckTemplateListShrinkRequest
	GetTenantId() *string
}

type GetDataCheckTemplateListShrinkRequest struct {
	// The validation rule type. Valid values:
	//
	// - 0: data volume comparison.
	//
	// - 1: metric comparison.
	//
	// - 2: weak content comparison.
	//
	// - 3: custom comparison.
	//
	// - 4: full-text comparison.
	//
	// - 5: null rate comparison.
	//
	// example:
	//
	// 1
	CheckType *int32 `json:"checkType,omitempty" xml:"checkType,omitempty"`
	// The field used for grouping (GROUP BY condition). Configure this parameter as needed.
	//
	// example:
	//
	// order_date
	GroupBy *string `json:"groupBy,omitempty" xml:"groupBy,omitempty"`
	// The list of validation template UUIDs. The source code of CheckTemplatePagedQry indicates that this parameter has no actual effect and does not need to be exposed externally. It is retained only for backward compatibility with legacy calls. Passing this parameter does not affect query results.
	IdListShrink *string `json:"idList,omitempty" xml:"idList,omitempty"`
	// **[Deprecated]*	- This parameter is deprecated and does not need to be passed. The source code of CheckTemplatePagedQry marks this parameter with @Deprecated.
	IsAdmin *bool `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	// Specifies whether the template is built-in. Valid values:
	//
	// - 0: No. The template is a custom template.
	//
	// - 1: Yes. The template is a built-in template.
	//
	// example:
	//
	// 0
	IsBuiltin *int32 `json:"isBuiltin,omitempty" xml:"isBuiltin,omitempty"`
	// Specifies whether to return the total record count in the paginated results.
	NeedTotalCount *bool `json:"needTotalCount,omitempty" xml:"needTotalCount,omitempty"`
	// The field used for sorting. Configure this parameter as needed.
	//
	// example:
	//
	// gmtCreate
	OrderBy *string `json:"orderBy,omitempty" xml:"orderBy,omitempty"`
	// The sort direction. Valid values:
	//
	// - ASC: ascending order.
	//
	// - DESC: descending order.
	//
	// example:
	//
	// DESC
	OrderDirection *string `json:"orderDirection,omitempty" xml:"orderDirection,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The page size, which specifies the number of records returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4B7A-9C3D-2A8F7E6B5C4D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The name of the validation template.
	//
	// example:
	//
	// DataVolumeValidationTemplate
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10001
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDataCheckTemplateListShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateListShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateListShrinkRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTemplateListShrinkRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetDataCheckTemplateListShrinkRequest) GetIdListShrink() *string {
	return s.IdListShrink
}

func (s *GetDataCheckTemplateListShrinkRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *GetDataCheckTemplateListShrinkRequest) GetIsBuiltin() *int32 {
	return s.IsBuiltin
}

func (s *GetDataCheckTemplateListShrinkRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetDataCheckTemplateListShrinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetDataCheckTemplateListShrinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *GetDataCheckTemplateListShrinkRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetDataCheckTemplateListShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDataCheckTemplateListShrinkRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTemplateListShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTemplateListShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDataCheckTemplateListShrinkRequest) SetCheckType(v int32) *GetDataCheckTemplateListShrinkRequest {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetGroupBy(v string) *GetDataCheckTemplateListShrinkRequest {
	s.GroupBy = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetIdListShrink(v string) *GetDataCheckTemplateListShrinkRequest {
	s.IdListShrink = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetIsAdmin(v bool) *GetDataCheckTemplateListShrinkRequest {
	s.IsAdmin = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetIsBuiltin(v int32) *GetDataCheckTemplateListShrinkRequest {
	s.IsBuiltin = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetNeedTotalCount(v bool) *GetDataCheckTemplateListShrinkRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetOrderBy(v string) *GetDataCheckTemplateListShrinkRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetOrderDirection(v string) *GetDataCheckTemplateListShrinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetPageIndex(v int32) *GetDataCheckTemplateListShrinkRequest {
	s.PageIndex = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetPageSize(v int32) *GetDataCheckTemplateListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetRequestId(v string) *GetDataCheckTemplateListShrinkRequest {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetTemplateName(v string) *GetDataCheckTemplateListShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) SetTenantId(v string) *GetDataCheckTemplateListShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataCheckTemplateListShrinkRequest) Validate() error {
	return dara.Validate(s)
}
