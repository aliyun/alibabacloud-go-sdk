// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiByAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ListApiByAppRequest
	GetOpTenantId() *int64
	SetPageQuery(v *ListApiByAppRequestPageQuery) *ListApiByAppRequest
	GetPageQuery() *ListApiByAppRequestPageQuery
}

type ListApiByAppRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	PageQuery *ListApiByAppRequestPageQuery `json:"PageQuery,omitempty" xml:"PageQuery,omitempty" type:"Struct"`
}

func (s ListApiByAppRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppRequest) GoString() string {
	return s.String()
}

func (s *ListApiByAppRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ListApiByAppRequest) GetPageQuery() *ListApiByAppRequestPageQuery {
	return s.PageQuery
}

func (s *ListApiByAppRequest) SetOpTenantId(v int64) *ListApiByAppRequest {
	s.OpTenantId = &v
	return s
}

func (s *ListApiByAppRequest) SetPageQuery(v *ListApiByAppRequestPageQuery) *ListApiByAppRequest {
	s.PageQuery = v
	return s
}

func (s *ListApiByAppRequest) Validate() error {
	if s.PageQuery != nil {
		if err := s.PageQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListApiByAppRequestPageQuery struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 10121101
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// apiName
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApiByAppRequestPageQuery) String() string {
	return dara.Prettify(s)
}

func (s ListApiByAppRequestPageQuery) GoString() string {
	return s.String()
}

func (s *ListApiByAppRequestPageQuery) GetAppKey() *int64 {
	return s.AppKey
}

func (s *ListApiByAppRequestPageQuery) GetKeyword() *string {
	return s.Keyword
}

func (s *ListApiByAppRequestPageQuery) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListApiByAppRequestPageQuery) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListApiByAppRequestPageQuery) SetAppKey(v int64) *ListApiByAppRequestPageQuery {
	s.AppKey = &v
	return s
}

func (s *ListApiByAppRequestPageQuery) SetKeyword(v string) *ListApiByAppRequestPageQuery {
	s.Keyword = &v
	return s
}

func (s *ListApiByAppRequestPageQuery) SetPageNum(v int32) *ListApiByAppRequestPageQuery {
	s.PageNum = &v
	return s
}

func (s *ListApiByAppRequestPageQuery) SetPageSize(v int32) *ListApiByAppRequestPageQuery {
	s.PageSize = &v
	return s
}

func (s *ListApiByAppRequestPageQuery) Validate() error {
	return dara.Validate(s)
}
