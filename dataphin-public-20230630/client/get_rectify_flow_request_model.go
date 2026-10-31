// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRectifyFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetQuery(v *GetRectifyFlowRequestGetQuery) *GetRectifyFlowRequest
	GetGetQuery() *GetRectifyFlowRequestGetQuery
	SetOpTenantId(v int64) *GetRectifyFlowRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetRectifyFlowRequest
	GetOpUserId() *string
}

type GetRectifyFlowRequest struct {
	// The business request parameters.
	//
	// This parameter is required.
	GetQuery *GetRectifyFlowRequestGetQuery `json:"GetQuery,omitempty" xml:"GetQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The operator ID.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetRectifyFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowRequest) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowRequest) GetGetQuery() *GetRectifyFlowRequestGetQuery {
	return s.GetQuery
}

func (s *GetRectifyFlowRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetRectifyFlowRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetRectifyFlowRequest) SetGetQuery(v *GetRectifyFlowRequestGetQuery) *GetRectifyFlowRequest {
	s.GetQuery = v
	return s
}

func (s *GetRectifyFlowRequest) SetOpTenantId(v int64) *GetRectifyFlowRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetRectifyFlowRequest) SetOpUserId(v string) *GetRectifyFlowRequest {
	s.OpUserId = &v
	return s
}

func (s *GetRectifyFlowRequest) Validate() error {
	if s.GetQuery != nil {
		if err := s.GetQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetRectifyFlowRequestGetQuery struct {
	// The page number. Default value: 1.
	//
	// example:
	//
	// 6500387
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of items per page. Default value: 20.
	//
	// example:
	//
	// 9919509
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the rectification process.
	//
	// This parameter is required.
	//
	// example:
	//
	// 442777862946
	RectifyId *int64 `json:"RectifyId,omitempty" xml:"RectifyId,omitempty"`
}

func (s GetRectifyFlowRequestGetQuery) String() string {
	return dara.Prettify(s)
}

func (s GetRectifyFlowRequestGetQuery) GoString() string {
	return s.String()
}

func (s *GetRectifyFlowRequestGetQuery) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetRectifyFlowRequestGetQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRectifyFlowRequestGetQuery) GetRectifyId() *int64 {
	return s.RectifyId
}

func (s *GetRectifyFlowRequestGetQuery) SetPageNumber(v int32) *GetRectifyFlowRequestGetQuery {
	s.PageNumber = &v
	return s
}

func (s *GetRectifyFlowRequestGetQuery) SetPageSize(v int32) *GetRectifyFlowRequestGetQuery {
	s.PageSize = &v
	return s
}

func (s *GetRectifyFlowRequestGetQuery) SetRectifyId(v int64) *GetRectifyFlowRequestGetQuery {
	s.RectifyId = &v
	return s
}

func (s *GetRectifyFlowRequestGetQuery) Validate() error {
	return dara.Validate(s)
}
