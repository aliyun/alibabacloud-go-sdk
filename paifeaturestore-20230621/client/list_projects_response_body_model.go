// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody
	GetProjects() []*ListProjectsResponseBodyProjects
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListProjectsResponseBody
	GetTotalCount() *int64
}

type ListProjectsResponseBody struct {
	Projects []*ListProjectsResponseBodyProjects `json:"Projects,omitempty" xml:"Projects,omitempty" type:"Repeated"`
	// example:
	//
	// 44933189-493B-5C43-A5C6-11EEC2A43520
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetProjects() []*ListProjectsResponseBodyProjects {
	return s.Projects
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListProjectsResponseBody) SetProjects(v []*ListProjectsResponseBodyProjects) *ListProjectsResponseBody {
	s.Projects = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetTotalCount(v int64) *ListProjectsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.Projects != nil {
		for _, item := range s.Projects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyProjects struct {
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 10
	FeatureEntityCount *int32 `json:"FeatureEntityCount,omitempty" xml:"FeatureEntityCount,omitempty"`
	// example:
	//
	// 10
	FeatureViewCount *int32 `json:"FeatureViewCount,omitempty" xml:"FeatureViewCount,omitempty"`
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
	// 5
	ModelCount *int32 `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	// example:
	//
	// project1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 4
	OfflineDatasourceId *string `json:"OfflineDatasourceId,omitempty" xml:"OfflineDatasourceId,omitempty"`
	// example:
	//
	// datasource1
	OfflineDatasourceName *string `json:"OfflineDatasourceName,omitempty" xml:"OfflineDatasourceName,omitempty"`
	// example:
	//
	// MaxCompute
	OfflineDatasourceType *string `json:"OfflineDatasourceType,omitempty" xml:"OfflineDatasourceType,omitempty"`
	// example:
	//
	// 10
	OfflineLifecycle *int32 `json:"OfflineLifecycle,omitempty" xml:"OfflineLifecycle,omitempty"`
	// example:
	//
	// 5
	OnlineDatasourceId *string `json:"OnlineDatasourceId,omitempty" xml:"OnlineDatasourceId,omitempty"`
	// example:
	//
	// datasource2
	OnlineDatasourceName *string `json:"OnlineDatasourceName,omitempty" xml:"OnlineDatasourceName,omitempty"`
	// example:
	//
	// Hologres
	OnlineDatasourceType *string `json:"OnlineDatasourceType,omitempty" xml:"OnlineDatasourceType,omitempty"`
	// example:
	//
	// 1232132543543****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 3
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListProjectsResponseBodyProjects) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyProjects) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyProjects) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyProjects) GetFeatureEntityCount() *int32 {
	return s.FeatureEntityCount
}

func (s *ListProjectsResponseBodyProjects) GetFeatureViewCount() *int32 {
	return s.FeatureViewCount
}

func (s *ListProjectsResponseBodyProjects) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListProjectsResponseBodyProjects) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListProjectsResponseBodyProjects) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *ListProjectsResponseBodyProjects) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyProjects) GetOfflineDatasourceId() *string {
	return s.OfflineDatasourceId
}

func (s *ListProjectsResponseBodyProjects) GetOfflineDatasourceName() *string {
	return s.OfflineDatasourceName
}

func (s *ListProjectsResponseBodyProjects) GetOfflineDatasourceType() *string {
	return s.OfflineDatasourceType
}

func (s *ListProjectsResponseBodyProjects) GetOfflineLifecycle() *int32 {
	return s.OfflineLifecycle
}

func (s *ListProjectsResponseBodyProjects) GetOnlineDatasourceId() *string {
	return s.OnlineDatasourceId
}

func (s *ListProjectsResponseBodyProjects) GetOnlineDatasourceName() *string {
	return s.OnlineDatasourceName
}

func (s *ListProjectsResponseBodyProjects) GetOnlineDatasourceType() *string {
	return s.OnlineDatasourceType
}

func (s *ListProjectsResponseBodyProjects) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsResponseBodyProjects) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListProjectsResponseBodyProjects) SetDescription(v string) *ListProjectsResponseBodyProjects {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetFeatureEntityCount(v int32) *ListProjectsResponseBodyProjects {
	s.FeatureEntityCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetFeatureViewCount(v int32) *ListProjectsResponseBodyProjects {
	s.FeatureViewCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtCreateTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtCreateTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetGmtModifiedTime(v string) *ListProjectsResponseBodyProjects {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetModelCount(v int32) *ListProjectsResponseBodyProjects {
	s.ModelCount = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetName(v string) *ListProjectsResponseBodyProjects {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceId(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceName(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineDatasourceType(v string) *ListProjectsResponseBodyProjects {
	s.OfflineDatasourceType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOfflineLifecycle(v int32) *ListProjectsResponseBodyProjects {
	s.OfflineLifecycle = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceId(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceName(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceName = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOnlineDatasourceType(v string) *ListProjectsResponseBodyProjects {
	s.OnlineDatasourceType = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetOwner(v string) *ListProjectsResponseBodyProjects {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) SetProjectId(v string) *ListProjectsResponseBodyProjects {
	s.ProjectId = &v
	return s
}

func (s *ListProjectsResponseBodyProjects) Validate() error {
	return dara.Validate(s)
}
