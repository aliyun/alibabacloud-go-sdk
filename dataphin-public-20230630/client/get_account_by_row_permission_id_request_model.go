// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountByRowPermissionIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGetAccountByRowPermissionIdQuery(v *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) *GetAccountByRowPermissionIdRequest
	GetGetAccountByRowPermissionIdQuery() *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery
	SetOpTenantId(v int64) *GetAccountByRowPermissionIdRequest
	GetOpTenantId() *int64
}

type GetAccountByRowPermissionIdRequest struct {
	// This parameter is required.
	GetAccountByRowPermissionIdQuery *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery `json:"GetAccountByRowPermissionIdQuery,omitempty" xml:"GetAccountByRowPermissionIdQuery,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s GetAccountByRowPermissionIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdRequest) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdRequest) GetGetAccountByRowPermissionIdQuery() *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery {
	return s.GetAccountByRowPermissionIdQuery
}

func (s *GetAccountByRowPermissionIdRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAccountByRowPermissionIdRequest) SetGetAccountByRowPermissionIdQuery(v *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) *GetAccountByRowPermissionIdRequest {
	s.GetAccountByRowPermissionIdQuery = v
	return s
}

func (s *GetAccountByRowPermissionIdRequest) SetOpTenantId(v int64) *GetAccountByRowPermissionIdRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAccountByRowPermissionIdRequest) Validate() error {
	if s.GetAccountByRowPermissionIdQuery != nil {
		if err := s.GetAccountByRowPermissionIdQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery struct {
	// This parameter is required.
	//
	// example:
	//
	// 300001234
	RowPermissionId *int64 `json:"RowPermissionId,omitempty" xml:"RowPermissionId,omitempty"`
	// This parameter is required.
	RuleIds []*int64 `json:"RuleIds,omitempty" xml:"RuleIds,omitempty" type:"Repeated"`
}

func (s GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) String() string {
	return dara.Prettify(s)
}

func (s GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) GoString() string {
	return s.String()
}

func (s *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) GetRowPermissionId() *int64 {
	return s.RowPermissionId
}

func (s *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) GetRuleIds() []*int64 {
	return s.RuleIds
}

func (s *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) SetRowPermissionId(v int64) *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery {
	s.RowPermissionId = &v
	return s
}

func (s *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) SetRuleIds(v []*int64) *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery {
	s.RuleIds = v
	return s
}

func (s *GetAccountByRowPermissionIdRequestGetAccountByRowPermissionIdQuery) Validate() error {
	return dara.Validate(s)
}
