// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPythonModulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListQuery(v *ListPythonModulesRequestListQuery) *ListPythonModulesRequest
	GetListQuery() *ListPythonModulesRequestListQuery
	SetOpTenantId(v int64) *ListPythonModulesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *ListPythonModulesRequest
	GetOpUserId() *string
}

type ListPythonModulesRequest struct {
	// The paging conditions for the paged query.
	//
	// This parameter is required.
	ListQuery *ListPythonModulesRequestListQuery `json:"ListQuery,omitempty" xml:"ListQuery,omitempty" type:"Struct"`
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

func (s ListPythonModulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesRequest) GoString() string {
	return s.String()
}

func (s *ListPythonModulesRequest) GetListQuery() *ListPythonModulesRequestListQuery {
	return s.ListQuery
}

func (s *ListPythonModulesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListPythonModulesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ListPythonModulesRequest) SetListQuery(v *ListPythonModulesRequestListQuery) *ListPythonModulesRequest {
	s.ListQuery = v
	return s
}

func (s *ListPythonModulesRequest) SetOpTenantId(v int64) *ListPythonModulesRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListPythonModulesRequest) SetOpUserId(v string) *ListPythonModulesRequest {
	s.OpUserId = &v
	return s
}

func (s *ListPythonModulesRequest) Validate() error {
	if s.ListQuery != nil {
		if err := s.ListQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPythonModulesRequestListQuery struct {
	// The keyword of the third-party package name.
	//
	// example:
	//
	// numpy
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the Python third-party package.
	//
	// example:
	//
	// 1001
	PythonModuleId *int64 `json:"PythonModuleId,omitempty" xml:"PythonModuleId,omitempty"`
	// The Python version. Valid values: PYTHON_2_7, PYTHON_3_7, and PYTHON_3_11.
	//
	// example:
	//
	// PYTHON_3_7
	PythonVersion *string `json:"PythonVersion,omitempty" xml:"PythonVersion,omitempty"`
}

func (s ListPythonModulesRequestListQuery) String() string {
	return dara.Prettify(s)
}

func (s ListPythonModulesRequestListQuery) GoString() string {
	return s.String()
}

func (s *ListPythonModulesRequestListQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListPythonModulesRequestListQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListPythonModulesRequestListQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPythonModulesRequestListQuery) GetPythonModuleId() *int64 {
	return s.PythonModuleId
}

func (s *ListPythonModulesRequestListQuery) GetPythonVersion() *string {
	return s.PythonVersion
}

func (s *ListPythonModulesRequestListQuery) SetKeyword(v string) *ListPythonModulesRequestListQuery {
	s.Keyword = &v
	return s
}

func (s *ListPythonModulesRequestListQuery) SetPageNum(v int32) *ListPythonModulesRequestListQuery {
	s.PageNum = &v
	return s
}

func (s *ListPythonModulesRequestListQuery) SetPageSize(v int32) *ListPythonModulesRequestListQuery {
	s.PageSize = &v
	return s
}

func (s *ListPythonModulesRequestListQuery) SetPythonModuleId(v int64) *ListPythonModulesRequestListQuery {
	s.PythonModuleId = &v
	return s
}

func (s *ListPythonModulesRequestListQuery) SetPythonVersion(v string) *ListPythonModulesRequestListQuery {
	s.PythonVersion = &v
	return s
}

func (s *ListPythonModulesRequestListQuery) Validate() error {
	return dara.Validate(s)
}
