// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogs(v []*ListCatalogsResponseBodyCatalogs) *ListCatalogsResponseBody
	GetCatalogs() []*ListCatalogsResponseBodyCatalogs
	SetMaxResults(v int32) *ListCatalogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListCatalogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListCatalogsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCatalogsResponseBody
	GetTotalCount() *int32
}

type ListCatalogsResponseBody struct {
	Catalogs []*ListCatalogsResponseBodyCatalogs `json:"catalogs,omitempty" xml:"catalogs,omitempty" type:"Repeated"`
	// 一次获取的最大记录数。
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// 下一页TOKEN。
	//
	// example:
	//
	// 1
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 记录总数。
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) GetCatalogs() []*ListCatalogsResponseBodyCatalogs {
	return s.Catalogs
}

func (s *ListCatalogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCatalogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCatalogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCatalogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCatalogsResponseBody) SetCatalogs(v []*ListCatalogsResponseBodyCatalogs) *ListCatalogsResponseBody {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBody) SetMaxResults(v int32) *ListCatalogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListCatalogsResponseBody) SetNextToken(v string) *ListCatalogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListCatalogsResponseBody) SetRequestId(v string) *ListCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCatalogsResponseBody) SetTotalCount(v int32) *ListCatalogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCatalogsResponseBody) Validate() error {
	if s.Catalogs != nil {
		for _, item := range s.Catalogs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCatalogsResponseBodyCatalogs struct {
	// example:
	//
	// alias
	Alias *string `json:"alias,omitempty" xml:"alias,omitempty"`
	// regionId。
	//
	// example:
	//
	// 15097**********
	CatalogId *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	// example:
	//
	// HMS
	CatalogProvider *string `json:"catalogProvider,omitempty" xml:"catalogProvider,omitempty"`
	// example:
	//
	// default_catalog
	CatalogType  *string            `json:"catalogType,omitempty" xml:"catalogType,omitempty"`
	Environments []*string          `json:"environments,omitempty" xml:"environments,omitempty" type:"Repeated"`
	Extras       map[string]*string `json:"extras,omitempty" xml:"extras,omitempty"`
	// example:
	//
	// 1760604889
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1760604889
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// 15097**********
	ResourceOwnerId *string `json:"resourceOwnerId,omitempty" xml:"resourceOwnerId,omitempty"`
	// 工作空间id。
	//
	// example:
	//
	// w-d2d82aa09155****
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s ListCatalogsResponseBodyCatalogs) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsResponseBodyCatalogs) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBodyCatalogs) GetAlias() *string {
	return s.Alias
}

func (s *ListCatalogsResponseBodyCatalogs) GetCatalogId() *string {
	return s.CatalogId
}

func (s *ListCatalogsResponseBodyCatalogs) GetCatalogProvider() *string {
	return s.CatalogProvider
}

func (s *ListCatalogsResponseBodyCatalogs) GetCatalogType() *string {
	return s.CatalogType
}

func (s *ListCatalogsResponseBodyCatalogs) GetEnvironments() []*string {
	return s.Environments
}

func (s *ListCatalogsResponseBodyCatalogs) GetExtras() map[string]*string {
	return s.Extras
}

func (s *ListCatalogsResponseBodyCatalogs) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListCatalogsResponseBodyCatalogs) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListCatalogsResponseBodyCatalogs) GetResourceOwnerId() *string {
	return s.ResourceOwnerId
}

func (s *ListCatalogsResponseBodyCatalogs) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListCatalogsResponseBodyCatalogs) SetAlias(v string) *ListCatalogsResponseBodyCatalogs {
	s.Alias = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetCatalogId(v string) *ListCatalogsResponseBodyCatalogs {
	s.CatalogId = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetCatalogProvider(v string) *ListCatalogsResponseBodyCatalogs {
	s.CatalogProvider = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetCatalogType(v string) *ListCatalogsResponseBodyCatalogs {
	s.CatalogType = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetEnvironments(v []*string) *ListCatalogsResponseBodyCatalogs {
	s.Environments = v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetExtras(v map[string]*string) *ListCatalogsResponseBodyCatalogs {
	s.Extras = v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetGmtCreate(v int64) *ListCatalogsResponseBodyCatalogs {
	s.GmtCreate = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetGmtModified(v int64) *ListCatalogsResponseBodyCatalogs {
	s.GmtModified = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetResourceOwnerId(v string) *ListCatalogsResponseBodyCatalogs {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) SetWorkspaceId(v string) *ListCatalogsResponseBodyCatalogs {
	s.WorkspaceId = &v
	return s
}

func (s *ListCatalogsResponseBodyCatalogs) Validate() error {
	return dara.Validate(s)
}
