// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalcEnginesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalcEngineType(v string) *ListCalcEnginesRequest
	GetCalcEngineType() *string
	SetEnvType(v string) *ListCalcEnginesRequest
	GetEnvType() *string
	SetName(v string) *ListCalcEnginesRequest
	GetName() *string
	SetPageNumber(v int32) *ListCalcEnginesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCalcEnginesRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListCalcEnginesRequest
	GetProjectId() *int64
}

type ListCalcEnginesRequest struct {
	// The type of the compute engine. The value of this parameter is not case-sensitive. Valid values:
	//
	// 	- **ODPS**
	//
	// 	- **EMR**
	//
	// 	- **BLINK**
	//
	// 	- **HOLO**
	//
	// 	- **MaxGraph**
	//
	// 	- **HYBRIDDB_FOR_POSTGRESQL**
	//
	// 	- **ADB_MYSQL**
	//
	// 	- **HADOOP_CDH**
	//
	// 	- **CLICKHOUSE**
	//
	// This parameter is required.
	//
	// example:
	//
	// ODPS
	CalcEngineType *string `json:"CalcEngineType,omitempty" xml:"CalcEngineType,omitempty"`
	// The environment in which the compute engine is used. Valid values:
	//
	// 	- **DEV**
	//
	// 	- **PRD**
	//
	// example:
	//
	// PRD
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The name of the compute engine, which must be exactly matched.
	//
	// example:
	//
	// abc
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace with which the compute engine is associated.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListCalcEnginesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCalcEnginesRequest) GoString() string {
	return s.String()
}

func (s *ListCalcEnginesRequest) GetCalcEngineType() *string {
	return s.CalcEngineType
}

func (s *ListCalcEnginesRequest) GetEnvType() *string {
	return s.EnvType
}

func (s *ListCalcEnginesRequest) GetName() *string {
	return s.Name
}

func (s *ListCalcEnginesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCalcEnginesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCalcEnginesRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListCalcEnginesRequest) SetCalcEngineType(v string) *ListCalcEnginesRequest {
	s.CalcEngineType = &v
	return s
}

func (s *ListCalcEnginesRequest) SetEnvType(v string) *ListCalcEnginesRequest {
	s.EnvType = &v
	return s
}

func (s *ListCalcEnginesRequest) SetName(v string) *ListCalcEnginesRequest {
	s.Name = &v
	return s
}

func (s *ListCalcEnginesRequest) SetPageNumber(v int32) *ListCalcEnginesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCalcEnginesRequest) SetPageSize(v int32) *ListCalcEnginesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCalcEnginesRequest) SetProjectId(v int64) *ListCalcEnginesRequest {
	s.ProjectId = &v
	return s
}

func (s *ListCalcEnginesRequest) Validate() error {
	return dara.Validate(s)
}
