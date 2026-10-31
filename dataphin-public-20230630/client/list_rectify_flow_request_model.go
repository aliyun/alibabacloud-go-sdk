// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRectifyFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListRectifyFlowRequestListQuery) *ListRectifyFlowRequest
	GetListQuery() *ListRectifyFlowRequestListQuery
	SetOpTenantId(v int64) *ListRectifyFlowRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListRectifyFlowRequest
	GetOpUserId() *string
}

type ListRectifyFlowRequest struct {
	// The business request parameters.
	//
	// This parameter is required.
	ListQuery *ListRectifyFlowRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s ListRectifyFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowRequest) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowRequest) GetListQuery() *ListRectifyFlowRequestListQuery {
	return s.ListQuery
}

func (s *ListRectifyFlowRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListRectifyFlowRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListRectifyFlowRequest) SetListQuery(v *ListRectifyFlowRequestListQuery) *ListRectifyFlowRequest {
	s.ListQuery = v
	return s
}

func (s *ListRectifyFlowRequest) SetOpTenantId(v int64) *ListRectifyFlowRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListRectifyFlowRequest) SetOpUserId(v string) *ListRectifyFlowRequest {
	s.OpUserId = &v
	return s
}

func (s *ListRectifyFlowRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRectifyFlowRequestListQuery struct {
	// The governance object type. Valid values: TABLE, DATASOURCE_TABLE, DATASOURCE, INDEX, REALTIME_LOGICAL_TABLE, and QD_FEATURE.
	//
	// This parameter is required.
	//
	// example:
	//
	// TABLE
	GovernItemType *string `json:"GovernItemType,omitempty" xml:"GovernItemType,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 7381905
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 8322096
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The priority filters. Valid values: HIGH, MIDDLE, and LOW.
	PriorityFilters []*string `json:"PriorityFilters,omitempty" xml:"PriorityFilters,omitempty" type:"Repeated"`
	// The list of rectifier IDs.
	RectifyUserFilters []*string `json:"RectifyUserFilters,omitempty" xml:"RectifyUserFilters,omitempty" type:"Repeated"`
	// SUBMIT_RECTIFY, SUBMIT_VERIFY, FINISH_VERIFY, and REJECT_VERIFY.
	StatusFilters []*string `json:"StatusFilters,omitempty" xml:"StatusFilters,omitempty" type:"Repeated"`
	// The list of verifier IDs.
	VerifyUserFilters []*string `json:"VerifyUserFilters,omitempty" xml:"VerifyUserFilters,omitempty" type:"Repeated"`
}

func (s ListRectifyFlowRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListRectifyFlowRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListRectifyFlowRequestListQuery) GetGovernItemType() *string {
	return s.GovernItemType
}

func (s *ListRectifyFlowRequestListQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRectifyFlowRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRectifyFlowRequestListQuery) GetPriorityFilters() []*string {
	return s.PriorityFilters
}

func (s *ListRectifyFlowRequestListQuery) GetRectifyUserFilters() []*string {
	return s.RectifyUserFilters
}

func (s *ListRectifyFlowRequestListQuery) GetStatusFilters() []*string {
	return s.StatusFilters
}

func (s *ListRectifyFlowRequestListQuery) GetVerifyUserFilters() []*string {
	return s.VerifyUserFilters
}

func (s *ListRectifyFlowRequestListQuery) SetGovernItemType(v string) *ListRectifyFlowRequestListQuery {
	s.GovernItemType = &v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetPageNumber(v int32) *ListRectifyFlowRequestListQuery {
	s.PageNumber = &v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetPageSize(v int32) *ListRectifyFlowRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetPriorityFilters(v []*string) *ListRectifyFlowRequestListQuery {
	s.PriorityFilters = v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetRectifyUserFilters(v []*string) *ListRectifyFlowRequestListQuery {
	s.RectifyUserFilters = v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetStatusFilters(v []*string) *ListRectifyFlowRequestListQuery {
	s.StatusFilters = v
	return s
}

func (s *ListRectifyFlowRequestListQuery) SetVerifyUserFilters(v []*string) *ListRectifyFlowRequestListQuery {
	s.VerifyUserFilters = v
	return s
}

func (s *ListRectifyFlowRequestListQuery) Validate() error {
	return dara.Validate(s)
}
