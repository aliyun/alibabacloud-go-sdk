// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *GetProjectResponseBody
	GetDescription() *string
	SetFeatureEntityCount(v int32) *GetProjectResponseBody
	GetFeatureEntityCount() *int32
	SetFeatureViewCount(v int32) *GetProjectResponseBody
	GetFeatureViewCount() *int32
	SetGmtCreateTime(v string) *GetProjectResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetProjectResponseBody
	GetGmtModifiedTime() *string
	SetModelCount(v int32) *GetProjectResponseBody
	GetModelCount() *int32
	SetName(v string) *GetProjectResponseBody
	GetName() *string
	SetOfflineDatasourceId(v string) *GetProjectResponseBody
	GetOfflineDatasourceId() *string
	SetOfflineDatasourceName(v string) *GetProjectResponseBody
	GetOfflineDatasourceName() *string
	SetOfflineDatasourceType(v string) *GetProjectResponseBody
	GetOfflineDatasourceType() *string
	SetOfflineLifecycle(v int32) *GetProjectResponseBody
	GetOfflineLifecycle() *int32
	SetOnlineDatasourceId(v string) *GetProjectResponseBody
	GetOnlineDatasourceId() *string
	SetOnlineDatasourceName(v string) *GetProjectResponseBody
	GetOnlineDatasourceName() *string
	SetOnlineDatasourceType(v string) *GetProjectResponseBody
	GetOnlineDatasourceType() *string
	SetOwner(v string) *GetProjectResponseBody
	GetOwner() *string
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
}

type GetProjectResponseBody struct {
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
	// 90
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
	// C33E160C-BFCA-5719-B958-942850E949F6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetProjectResponseBody) GetFeatureEntityCount() *int32 {
	return s.FeatureEntityCount
}

func (s *GetProjectResponseBody) GetFeatureViewCount() *int32 {
	return s.FeatureViewCount
}

func (s *GetProjectResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetProjectResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetProjectResponseBody) GetModelCount() *int32 {
	return s.ModelCount
}

func (s *GetProjectResponseBody) GetName() *string {
	return s.Name
}

func (s *GetProjectResponseBody) GetOfflineDatasourceId() *string {
	return s.OfflineDatasourceId
}

func (s *GetProjectResponseBody) GetOfflineDatasourceName() *string {
	return s.OfflineDatasourceName
}

func (s *GetProjectResponseBody) GetOfflineDatasourceType() *string {
	return s.OfflineDatasourceType
}

func (s *GetProjectResponseBody) GetOfflineLifecycle() *int32 {
	return s.OfflineLifecycle
}

func (s *GetProjectResponseBody) GetOnlineDatasourceId() *string {
	return s.OnlineDatasourceId
}

func (s *GetProjectResponseBody) GetOnlineDatasourceName() *string {
	return s.OnlineDatasourceName
}

func (s *GetProjectResponseBody) GetOnlineDatasourceType() *string {
	return s.OnlineDatasourceType
}

func (s *GetProjectResponseBody) GetOwner() *string {
	return s.Owner
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) SetDescription(v string) *GetProjectResponseBody {
	s.Description = &v
	return s
}

func (s *GetProjectResponseBody) SetFeatureEntityCount(v int32) *GetProjectResponseBody {
	s.FeatureEntityCount = &v
	return s
}

func (s *GetProjectResponseBody) SetFeatureViewCount(v int32) *GetProjectResponseBody {
	s.FeatureViewCount = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtCreateTime(v string) *GetProjectResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetGmtModifiedTime(v string) *GetProjectResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetProjectResponseBody) SetModelCount(v int32) *GetProjectResponseBody {
	s.ModelCount = &v
	return s
}

func (s *GetProjectResponseBody) SetName(v string) *GetProjectResponseBody {
	s.Name = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceId(v string) *GetProjectResponseBody {
	s.OfflineDatasourceId = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceName(v string) *GetProjectResponseBody {
	s.OfflineDatasourceName = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineDatasourceType(v string) *GetProjectResponseBody {
	s.OfflineDatasourceType = &v
	return s
}

func (s *GetProjectResponseBody) SetOfflineLifecycle(v int32) *GetProjectResponseBody {
	s.OfflineLifecycle = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceId(v string) *GetProjectResponseBody {
	s.OnlineDatasourceId = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceName(v string) *GetProjectResponseBody {
	s.OnlineDatasourceName = &v
	return s
}

func (s *GetProjectResponseBody) SetOnlineDatasourceType(v string) *GetProjectResponseBody {
	s.OnlineDatasourceType = &v
	return s
}

func (s *GetProjectResponseBody) SetOwner(v string) *GetProjectResponseBody {
	s.Owner = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
