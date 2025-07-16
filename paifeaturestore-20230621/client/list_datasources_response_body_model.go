// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasources(v []*ListDatasourcesResponseBodyDatasources) *ListDatasourcesResponseBody
	GetDatasources() []*ListDatasourcesResponseBodyDatasources
	SetRequestId(v string) *ListDatasourcesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListDatasourcesResponseBody
	GetTotalCount() *int64
}

type ListDatasourcesResponseBody struct {
	Datasources []*ListDatasourcesResponseBodyDatasources `json:"Datasources,omitempty" xml:"Datasources,omitempty" type:"Repeated"`
	// example:
	//
	// 44933189-493B-5C43-A5C6-11EEC2A43520
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBody) GetDatasources() []*ListDatasourcesResponseBodyDatasources {
	return s.Datasources
}

func (s *ListDatasourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDatasourcesResponseBody) SetDatasources(v []*ListDatasourcesResponseBodyDatasources) *ListDatasourcesResponseBody {
	s.Datasources = v
	return s
}

func (s *ListDatasourcesResponseBody) SetRequestId(v string) *ListDatasourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasourcesResponseBody) SetTotalCount(v int64) *ListDatasourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDatasourcesResponseBodyDatasources struct {
	// example:
	//
	// {"address": ""}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// 3
	DatasourceId *string `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-12-15T23:24:33.132+08:00
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// datasource1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// igraph_instance1
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// example:
	//
	// 32324
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDatasourcesResponseBodyDatasources) String() string {
	return dara.Prettify(s)
}

func (s ListDatasourcesResponseBodyDatasources) GoString() string {
	return s.String()
}

func (s *ListDatasourcesResponseBodyDatasources) GetConfig() *string {
	return s.Config
}

func (s *ListDatasourcesResponseBodyDatasources) GetDatasourceId() *string {
	return s.DatasourceId
}

func (s *ListDatasourcesResponseBodyDatasources) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListDatasourcesResponseBodyDatasources) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListDatasourcesResponseBodyDatasources) GetName() *string {
	return s.Name
}

func (s *ListDatasourcesResponseBodyDatasources) GetType() *string {
	return s.Type
}

func (s *ListDatasourcesResponseBodyDatasources) GetUri() *string {
	return s.Uri
}

func (s *ListDatasourcesResponseBodyDatasources) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDatasourcesResponseBodyDatasources) SetConfig(v string) *ListDatasourcesResponseBodyDatasources {
	s.Config = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetDatasourceId(v string) *ListDatasourcesResponseBodyDatasources {
	s.DatasourceId = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetGmtCreateTime(v string) *ListDatasourcesResponseBodyDatasources {
	s.GmtCreateTime = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetGmtModifiedTime(v string) *ListDatasourcesResponseBodyDatasources {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetName(v string) *ListDatasourcesResponseBodyDatasources {
	s.Name = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetType(v string) *ListDatasourcesResponseBodyDatasources {
	s.Type = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetUri(v string) *ListDatasourcesResponseBodyDatasources {
	s.Uri = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) SetWorkspaceId(v string) *ListDatasourcesResponseBodyDatasources {
	s.WorkspaceId = &v
	return s
}

func (s *ListDatasourcesResponseBodyDatasources) Validate() error {
	return dara.Validate(s)
}
