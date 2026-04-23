// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricDefRespDTO interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *MetricDefRespDTO
	GetKey() *string
	SetLabel(v string) *MetricDefRespDTO
	GetLabel() *string
	SetSortable(v bool) *MetricDefRespDTO
	GetSortable() *bool
	SetUnit(v string) *MetricDefRespDTO
	GetUnit() *string
}

type MetricDefRespDTO struct {
	// example:
	//
	// total_calls
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// example:
	//
	// true
	Sortable *bool   `json:"sortable,omitempty" xml:"sortable,omitempty"`
	Unit     *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s MetricDefRespDTO) String() string {
	return dara.Prettify(s)
}

func (s MetricDefRespDTO) GoString() string {
	return s.String()
}

func (s *MetricDefRespDTO) GetKey() *string {
	return s.Key
}

func (s *MetricDefRespDTO) GetLabel() *string {
	return s.Label
}

func (s *MetricDefRespDTO) GetSortable() *bool {
	return s.Sortable
}

func (s *MetricDefRespDTO) GetUnit() *string {
	return s.Unit
}

func (s *MetricDefRespDTO) SetKey(v string) *MetricDefRespDTO {
	s.Key = &v
	return s
}

func (s *MetricDefRespDTO) SetLabel(v string) *MetricDefRespDTO {
	s.Label = &v
	return s
}

func (s *MetricDefRespDTO) SetSortable(v bool) *MetricDefRespDTO {
	s.Sortable = &v
	return s
}

func (s *MetricDefRespDTO) SetUnit(v string) *MetricDefRespDTO {
	s.Unit = &v
	return s
}

func (s *MetricDefRespDTO) Validate() error {
	return dara.Validate(s)
}
