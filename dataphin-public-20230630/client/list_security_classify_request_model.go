// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecurityClassifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListSecurityClassifyRequestListQuery) *ListSecurityClassifyRequest
	GetListQuery() *ListSecurityClassifyRequestListQuery
	SetOpTenantId(v int64) *ListSecurityClassifyRequest
	GetOpTenantId() *int64
}

type ListSecurityClassifyRequest struct {
	// The query conditions.
	ListQuery *ListSecurityClassifyRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s ListSecurityClassifyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyRequest) GetListQuery() *ListSecurityClassifyRequestListQuery {
	return s.ListQuery
}

func (s *ListSecurityClassifyRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListSecurityClassifyRequest) SetListQuery(v *ListSecurityClassifyRequestListQuery) *ListSecurityClassifyRequest {
	s.ListQuery = v
	return s
}

func (s *ListSecurityClassifyRequest) SetOpTenantId(v int64) *ListSecurityClassifyRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListSecurityClassifyRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSecurityClassifyRequestListQuery struct {
	// The data level ID.
	//
	// example:
	//
	// 1
	LevelIndex *int64 `json:"LevelIndex,omitempty" xml:"LevelIndex,omitempty"`
	// The classification name. Fuzzy match is supported.
	//
	// example:
	//
	// Personal Information
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The list of effective statuses. Valid values: ENABLE, DISABLE.
	//
	// example:
	//
	// ["ENABLE"]
	StatusList []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
}

func (s ListSecurityClassifyRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListSecurityClassifyRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListSecurityClassifyRequestListQuery) GetLevelIndex() *int64 {
	return s.LevelIndex
}

func (s *ListSecurityClassifyRequestListQuery) GetName() *string {
	return s.Name
}

func (s *ListSecurityClassifyRequestListQuery) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSecurityClassifyRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecurityClassifyRequestListQuery) GetStatusList() []*string {
	return s.StatusList
}

func (s *ListSecurityClassifyRequestListQuery) SetLevelIndex(v int64) *ListSecurityClassifyRequestListQuery {
	s.LevelIndex = &v
	return s
}

func (s *ListSecurityClassifyRequestListQuery) SetName(v string) *ListSecurityClassifyRequestListQuery {
	s.Name = &v
	return s
}

func (s *ListSecurityClassifyRequestListQuery) SetPageNo(v int32) *ListSecurityClassifyRequestListQuery {
	s.PageNo = &v
	return s
}

func (s *ListSecurityClassifyRequestListQuery) SetPageSize(v int32) *ListSecurityClassifyRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListSecurityClassifyRequestListQuery) SetStatusList(v []*string) *ListSecurityClassifyRequestListQuery {
	s.StatusList = v
	return s
}

func (s *ListSecurityClassifyRequestListQuery) Validate() error {
	return dara.Validate(s)
}
