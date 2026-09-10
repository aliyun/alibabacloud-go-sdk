// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMetaDataComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryType(v string) *AddMetaDataComponentRequest
	GetCategoryType() *string
	SetComponentType(v int32) *AddMetaDataComponentRequest
	GetComponentType() *int32
	SetDsConfig(v string) *AddMetaDataComponentRequest
	GetDsConfig() *string
	SetDsDesc(v string) *AddMetaDataComponentRequest
	GetDsDesc() *string
	SetDsId(v string) *AddMetaDataComponentRequest
	GetDsId() *string
	SetDsName(v string) *AddMetaDataComponentRequest
	GetDsName() *string
	SetDsStatus(v int32) *AddMetaDataComponentRequest
	GetDsStatus() *int32
	SetDsType(v string) *AddMetaDataComponentRequest
	GetDsType() *string
	SetDsVersion(v string) *AddMetaDataComponentRequest
	GetDsVersion() *string
}

type AddMetaDataComponentRequest struct {
	// The data source category. Valid values: DATASET, WORKFLOW, ENGINE.
	//
	// example:
	//
	// WORKFLOW
	CategoryType *string `json:"categoryType,omitempty" xml:"categoryType,omitempty"`
	// The role of the data source in the migration pipeline. Valid values:
	//
	// - 0: source.
	//
	// - 1: destination.
	//
	// example:
	//
	// 0
	ComponentType *int32 `json:"componentType,omitempty" xml:"componentType,omitempty"`
	// The datasource config.
	//
	// example:
	//
	// {"endpoint":"...","token":"******"}
	DsConfig *string `json:"dsConfig,omitempty" xml:"dsConfig,omitempty"`
	// The description of the data source.
	//
	// example:
	//
	// Data source description
	DsDesc *string `json:"dsDesc,omitempty" xml:"dsDesc,omitempty"`
	// The external ID of the data source.
	//
	// example:
	//
	// 290
	DsId *string `json:"dsId,omitempty" xml:"dsId,omitempty"`
	// The name of the data source.
	//
	// example:
	//
	// test_ds318_hangzhou_0428
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
	// The connectivity status of the data source.
	//
	// example:
	//
	// 1
	DsStatus *int32 `json:"dsStatus,omitempty" xml:"dsStatus,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// Hive
	DsType *string `json:"dsType,omitempty" xml:"dsType,omitempty"`
	// The version of the data source.
	//
	// example:
	//
	// 3.2.0
	DsVersion *string `json:"dsVersion,omitempty" xml:"dsVersion,omitempty"`
}

func (s AddMetaDataComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMetaDataComponentRequest) GoString() string {
	return s.String()
}

func (s *AddMetaDataComponentRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddMetaDataComponentRequest) GetComponentType() *int32 {
	return s.ComponentType
}

func (s *AddMetaDataComponentRequest) GetDsConfig() *string {
	return s.DsConfig
}

func (s *AddMetaDataComponentRequest) GetDsDesc() *string {
	return s.DsDesc
}

func (s *AddMetaDataComponentRequest) GetDsId() *string {
	return s.DsId
}

func (s *AddMetaDataComponentRequest) GetDsName() *string {
	return s.DsName
}

func (s *AddMetaDataComponentRequest) GetDsStatus() *int32 {
	return s.DsStatus
}

func (s *AddMetaDataComponentRequest) GetDsType() *string {
	return s.DsType
}

func (s *AddMetaDataComponentRequest) GetDsVersion() *string {
	return s.DsVersion
}

func (s *AddMetaDataComponentRequest) SetCategoryType(v string) *AddMetaDataComponentRequest {
	s.CategoryType = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetComponentType(v int32) *AddMetaDataComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsConfig(v string) *AddMetaDataComponentRequest {
	s.DsConfig = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsDesc(v string) *AddMetaDataComponentRequest {
	s.DsDesc = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsId(v string) *AddMetaDataComponentRequest {
	s.DsId = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsName(v string) *AddMetaDataComponentRequest {
	s.DsName = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsStatus(v int32) *AddMetaDataComponentRequest {
	s.DsStatus = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsType(v string) *AddMetaDataComponentRequest {
	s.DsType = &v
	return s
}

func (s *AddMetaDataComponentRequest) SetDsVersion(v string) *AddMetaDataComponentRequest {
	s.DsVersion = &v
	return s
}

func (s *AddMetaDataComponentRequest) Validate() error {
	return dara.Validate(s)
}
