// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostQueryModelsDTO interface {
	dara.Model
	String() string
	GoString() string
	SetColumns(v []*MetricDefRespDTO) *CostQueryModelsDTO
	GetColumns() []*MetricDefRespDTO
	SetIdField(v string) *CostQueryModelsDTO
	GetIdField() *string
	SetNameField(v string) *CostQueryModelsDTO
	GetNameField() *string
	SetRows(v []*ModelRowDTO) *CostQueryModelsDTO
	GetRows() []*ModelRowDTO
}

type CostQueryModelsDTO struct {
	Columns []*MetricDefRespDTO `json:"columns,omitempty" xml:"columns,omitempty" type:"Repeated"`
	// example:
	//
	// model_id
	IdField *string `json:"idField,omitempty" xml:"idField,omitempty"`
	// example:
	//
	// model_name
	NameField *string        `json:"nameField,omitempty" xml:"nameField,omitempty"`
	Rows      []*ModelRowDTO `json:"rows,omitempty" xml:"rows,omitempty" type:"Repeated"`
}

func (s CostQueryModelsDTO) String() string {
	return dara.Prettify(s)
}

func (s CostQueryModelsDTO) GoString() string {
	return s.String()
}

func (s *CostQueryModelsDTO) GetColumns() []*MetricDefRespDTO {
	return s.Columns
}

func (s *CostQueryModelsDTO) GetIdField() *string {
	return s.IdField
}

func (s *CostQueryModelsDTO) GetNameField() *string {
	return s.NameField
}

func (s *CostQueryModelsDTO) GetRows() []*ModelRowDTO {
	return s.Rows
}

func (s *CostQueryModelsDTO) SetColumns(v []*MetricDefRespDTO) *CostQueryModelsDTO {
	s.Columns = v
	return s
}

func (s *CostQueryModelsDTO) SetIdField(v string) *CostQueryModelsDTO {
	s.IdField = &v
	return s
}

func (s *CostQueryModelsDTO) SetNameField(v string) *CostQueryModelsDTO {
	s.NameField = &v
	return s
}

func (s *CostQueryModelsDTO) SetRows(v []*ModelRowDTO) *CostQueryModelsDTO {
	s.Rows = v
	return s
}

func (s *CostQueryModelsDTO) Validate() error {
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
