// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureViewsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeatureViews(v []*ListFeatureViewsResponseBodyFeatureViews) *ListFeatureViewsResponseBody
	GetFeatureViews() []*ListFeatureViewsResponseBodyFeatureViews
	SetRequestId(v string) *ListFeatureViewsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListFeatureViewsResponseBody
	GetTotalCount() *int64
}

type ListFeatureViewsResponseBody struct {
	FeatureViews []*ListFeatureViewsResponseBodyFeatureViews `json:"FeatureViews,omitempty" xml:"FeatureViews,omitempty" type:"Repeated"`
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFeatureViewsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponseBody) GetFeatureViews() []*ListFeatureViewsResponseBodyFeatureViews {
	return s.FeatureViews
}

func (s *ListFeatureViewsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureViewsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListFeatureViewsResponseBody) SetFeatureViews(v []*ListFeatureViewsResponseBodyFeatureViews) *ListFeatureViewsResponseBody {
	s.FeatureViews = v
	return s
}

func (s *ListFeatureViewsResponseBody) SetRequestId(v string) *ListFeatureViewsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureViewsResponseBody) SetTotalCount(v int64) *ListFeatureViewsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFeatureViewsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFeatureViewsResponseBodyFeatureViews struct {
	// example:
	//
	// featureEntity1
	FeatureEntityName *string `json:"FeatureEntityName,omitempty" xml:"FeatureEntityName,omitempty"`
	// example:
	//
	// 3
	FeatureViewId *string `json:"FeatureViewId,omitempty" xml:"FeatureViewId,omitempty"`
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
	// featureView1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12321421412****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// project1
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 4
	RegisterDatasourceId *string `json:"RegisterDatasourceId,omitempty" xml:"RegisterDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	RegisterDatasourceName *string `json:"RegisterDatasourceName,omitempty" xml:"RegisterDatasourceName,omitempty"`
	// example:
	//
	// table1
	RegisterTable *string `json:"RegisterTable,omitempty" xml:"RegisterTable,omitempty"`
	// example:
	//
	// 90
	TTL *int32 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// example:
	//
	// Batch
	Type             *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WriteToFeatureDB *bool   `json:"WriteToFeatureDB,omitempty" xml:"WriteToFeatureDB,omitempty"`
}

func (s ListFeatureViewsResponseBodyFeatureViews) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureViewsResponseBodyFeatureViews) GoString() string {
	return s.String()
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetFeatureEntityName() *string {
	return s.FeatureEntityName
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetFeatureViewId() *string {
	return s.FeatureViewId
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetName() *string {
	return s.Name
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetOwner() *string {
	return s.Owner
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetRegisterDatasourceId() *string {
	return s.RegisterDatasourceId
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetRegisterDatasourceName() *string {
	return s.RegisterDatasourceName
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetRegisterTable() *string {
	return s.RegisterTable
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetTTL() *int32 {
	return s.TTL
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetType() *string {
	return s.Type
}

func (s *ListFeatureViewsResponseBodyFeatureViews) GetWriteToFeatureDB() *bool {
	return s.WriteToFeatureDB
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetFeatureEntityName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.FeatureEntityName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetFeatureViewId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.FeatureViewId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetGmtCreateTime(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.GmtCreateTime = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetGmtModifiedTime(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Name = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetOwner(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Owner = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetProjectId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.ProjectId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetProjectName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.ProjectName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterDatasourceId(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterDatasourceId = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterDatasourceName(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterDatasourceName = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetRegisterTable(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.RegisterTable = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetTTL(v int32) *ListFeatureViewsResponseBodyFeatureViews {
	s.TTL = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetType(v string) *ListFeatureViewsResponseBodyFeatureViews {
	s.Type = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) SetWriteToFeatureDB(v bool) *ListFeatureViewsResponseBodyFeatureViews {
	s.WriteToFeatureDB = &v
	return s
}

func (s *ListFeatureViewsResponseBodyFeatureViews) Validate() error {
	return dara.Validate(s)
}
