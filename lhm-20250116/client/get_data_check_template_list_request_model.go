// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCheckTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckType(v int32) *GetDataCheckTemplateListRequest
	GetCheckType() *int32
	SetGroupBy(v string) *GetDataCheckTemplateListRequest
	GetGroupBy() *string
	SetIdList(v []*string) *GetDataCheckTemplateListRequest
	GetIdList() []*string
	SetIsAdmin(v bool) *GetDataCheckTemplateListRequest
	GetIsAdmin() *bool
	SetIsBuiltin(v int32) *GetDataCheckTemplateListRequest
	GetIsBuiltin() *int32
	SetNeedTotalCount(v bool) *GetDataCheckTemplateListRequest
	GetNeedTotalCount() *bool
	SetOrderBy(v string) *GetDataCheckTemplateListRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *GetDataCheckTemplateListRequest
	GetOrderDirection() *string
	SetPageIndex(v int32) *GetDataCheckTemplateListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetDataCheckTemplateListRequest
	GetPageSize() *int32
	SetRequestId(v string) *GetDataCheckTemplateListRequest
	GetRequestId() *string
	SetTemplateName(v string) *GetDataCheckTemplateListRequest
	GetTemplateName() *string
	SetTenantId(v string) *GetDataCheckTemplateListRequest
	GetTenantId() *string
}

type GetDataCheckTemplateListRequest struct {
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
	IdList []*string `json:"idList,omitempty" xml:"idList,omitempty" type:"Repeated"`
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

func (s GetDataCheckTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCheckTemplateListRequest) GoString() string {
	return s.String()
}

func (s *GetDataCheckTemplateListRequest) GetCheckType() *int32 {
	return s.CheckType
}

func (s *GetDataCheckTemplateListRequest) GetGroupBy() *string {
	return s.GroupBy
}

func (s *GetDataCheckTemplateListRequest) GetIdList() []*string {
	return s.IdList
}

func (s *GetDataCheckTemplateListRequest) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *GetDataCheckTemplateListRequest) GetIsBuiltin() *int32 {
	return s.IsBuiltin
}

func (s *GetDataCheckTemplateListRequest) GetNeedTotalCount() *bool {
	return s.NeedTotalCount
}

func (s *GetDataCheckTemplateListRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *GetDataCheckTemplateListRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *GetDataCheckTemplateListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetDataCheckTemplateListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetDataCheckTemplateListRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataCheckTemplateListRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetDataCheckTemplateListRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDataCheckTemplateListRequest) SetCheckType(v int32) *GetDataCheckTemplateListRequest {
	s.CheckType = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetGroupBy(v string) *GetDataCheckTemplateListRequest {
	s.GroupBy = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetIdList(v []*string) *GetDataCheckTemplateListRequest {
	s.IdList = v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetIsAdmin(v bool) *GetDataCheckTemplateListRequest {
	s.IsAdmin = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetIsBuiltin(v int32) *GetDataCheckTemplateListRequest {
	s.IsBuiltin = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetNeedTotalCount(v bool) *GetDataCheckTemplateListRequest {
	s.NeedTotalCount = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetOrderBy(v string) *GetDataCheckTemplateListRequest {
	s.OrderBy = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetOrderDirection(v string) *GetDataCheckTemplateListRequest {
	s.OrderDirection = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetPageIndex(v int32) *GetDataCheckTemplateListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetPageSize(v int32) *GetDataCheckTemplateListRequest {
	s.PageSize = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetRequestId(v string) *GetDataCheckTemplateListRequest {
	s.RequestId = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetTemplateName(v string) *GetDataCheckTemplateListRequest {
	s.TemplateName = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) SetTenantId(v string) *GetDataCheckTemplateListRequest {
	s.TenantId = &v
	return s
}

func (s *GetDataCheckTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
