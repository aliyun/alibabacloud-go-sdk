// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostModelDetailRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*MetricDefRespDTO) *CostModelDetailRespDTO
	GetColumns() []*MetricDefRespDTO
	SetGranularity(v string) *CostModelDetailRespDTO
	GetGranularity() *string
	SetModelId(v int64) *CostModelDetailRespDTO
	GetModelId() *int64
	SetModelName(v string) *CostModelDetailRespDTO
	GetModelName() *string
	SetPage(v int32) *CostModelDetailRespDTO
	GetPage() *int32
	SetPageSize(v int32) *CostModelDetailRespDTO
	GetPageSize() *int32
	SetRows(v []*CostModelDetailRowDTO) *CostModelDetailRespDTO
	GetRows() []*CostModelDetailRowDTO
	SetTotal(v int32) *CostModelDetailRespDTO
	GetTotal() *int32
}

type CostModelDetailRespDTO struct {
	Columns []*MetricDefRespDTO `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// hourly
	Granularity *string `json:"granularity,omitempty" xml:"granularity,omitempty"`
	// example:
	//
	// 1
	ModelId *int64 `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// 通义千问-Plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 20
	PageSize *int32                   `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	Rows     []*CostModelDetailRowDTO `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s CostModelDetailRespDTO) String() string {
	return dara.Prettify(s)
}

func (s CostModelDetailRespDTO) GoString() string {
	return s.String()
}

func (s *CostModelDetailRespDTO) GetColumns() []*MetricDefRespDTO {
	return s.Columns
}

func (s *CostModelDetailRespDTO) GetGranularity() *string {
	return s.Granularity
}

func (s *CostModelDetailRespDTO) GetModelId() *int64 {
	return s.ModelId
}

func (s *CostModelDetailRespDTO) GetModelName() *string {
	return s.ModelName
}

func (s *CostModelDetailRespDTO) GetPage() *int32 {
	return s.Page
}

func (s *CostModelDetailRespDTO) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CostModelDetailRespDTO) GetRows() []*CostModelDetailRowDTO {
	return s.Rows
}

func (s *CostModelDetailRespDTO) GetTotal() *int32 {
	return s.Total
}

func (s *CostModelDetailRespDTO) SetColumns(v []*MetricDefRespDTO) *CostModelDetailRespDTO {
	s.Columns = v
	return s
}

func (s *CostModelDetailRespDTO) SetGranularity(v string) *CostModelDetailRespDTO {
	s.Granularity = &v
	return s
}

func (s *CostModelDetailRespDTO) SetModelId(v int64) *CostModelDetailRespDTO {
	s.ModelId = &v
	return s
}

func (s *CostModelDetailRespDTO) SetModelName(v string) *CostModelDetailRespDTO {
	s.ModelName = &v
	return s
}

func (s *CostModelDetailRespDTO) SetPage(v int32) *CostModelDetailRespDTO {
	s.Page = &v
	return s
}

func (s *CostModelDetailRespDTO) SetPageSize(v int32) *CostModelDetailRespDTO {
	s.PageSize = &v
	return s
}

func (s *CostModelDetailRespDTO) SetRows(v []*CostModelDetailRowDTO) *CostModelDetailRespDTO {
	s.Rows = v
	return s
}

func (s *CostModelDetailRespDTO) SetTotal(v int32) *CostModelDetailRespDTO {
	s.Total = &v
	return s
}

func (s *CostModelDetailRespDTO) Validate() error {
	if s.Columns != nil {
		for _, item := range s.Columns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Rows != nil {
		for _, item := range s.Rows {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
