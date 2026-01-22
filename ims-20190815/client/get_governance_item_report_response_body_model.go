// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceItemReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColumnsSchema(v *GetGovernanceItemReportResponseBodyColumnsSchema) *GetGovernanceItemReportResponseBody
	GetColumnsSchema() *GetGovernanceItemReportResponseBodyColumnsSchema
	SetColumnsValue(v *GetGovernanceItemReportResponseBodyColumnsValue) *GetGovernanceItemReportResponseBody
	GetColumnsValue() *GetGovernanceItemReportResponseBodyColumnsValue
	SetGenerateTime(v string) *GetGovernanceItemReportResponseBody
	GetGenerateTime() *string
	SetIsTruncated(v bool) *GetGovernanceItemReportResponseBody
	GetIsTruncated() *bool
	SetMarker(v string) *GetGovernanceItemReportResponseBody
	GetMarker() *string
	SetMetricType(v string) *GetGovernanceItemReportResponseBody
	GetMetricType() *string
	SetMetricValue(v interface{}) *GetGovernanceItemReportResponseBody
	GetMetricValue() interface{}
	SetRequestId(v string) *GetGovernanceItemReportResponseBody
	GetRequestId() *string
}

type GetGovernanceItemReportResponseBody struct {
	ColumnsSchema *GetGovernanceItemReportResponseBodyColumnsSchema `json:"ColumnsSchema,omitempty" xml:"ColumnsSchema,omitempty" type:"Struct"`
	ColumnsValue  *GetGovernanceItemReportResponseBodyColumnsValue  `json:"ColumnsValue,omitempty" xml:"ColumnsValue,omitempty" type:"Struct"`
	// example:
	//
	// 2020-10-19T15:06:52Z
	GenerateTime *string `json:"GenerateTime,omitempty" xml:"GenerateTime,omitempty"`
	// example:
	//
	// true
	IsTruncated *bool `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	// example:
	//
	// EXAMPLE
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// example:
	//
	// Number
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// example:
	//
	// 100
	MetricValue interface{} `json:"MetricValue,omitempty" xml:"MetricValue,omitempty"`
	// example:
	//
	// F2CE9688-AA85-5F23-8C22-0EC23473405F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGovernanceItemReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBody) GetColumnsSchema() *GetGovernanceItemReportResponseBodyColumnsSchema {
	return s.ColumnsSchema
}

func (s *GetGovernanceItemReportResponseBody) GetColumnsValue() *GetGovernanceItemReportResponseBodyColumnsValue {
	return s.ColumnsValue
}

func (s *GetGovernanceItemReportResponseBody) GetGenerateTime() *string {
	return s.GenerateTime
}

func (s *GetGovernanceItemReportResponseBody) GetIsTruncated() *bool {
	return s.IsTruncated
}

func (s *GetGovernanceItemReportResponseBody) GetMarker() *string {
	return s.Marker
}

func (s *GetGovernanceItemReportResponseBody) GetMetricType() *string {
	return s.MetricType
}

func (s *GetGovernanceItemReportResponseBody) GetMetricValue() interface{} {
	return s.MetricValue
}

func (s *GetGovernanceItemReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGovernanceItemReportResponseBody) SetColumnsSchema(v *GetGovernanceItemReportResponseBodyColumnsSchema) *GetGovernanceItemReportResponseBody {
	s.ColumnsSchema = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetColumnsValue(v *GetGovernanceItemReportResponseBodyColumnsValue) *GetGovernanceItemReportResponseBody {
	s.ColumnsValue = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetGenerateTime(v string) *GetGovernanceItemReportResponseBody {
	s.GenerateTime = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetIsTruncated(v bool) *GetGovernanceItemReportResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMarker(v string) *GetGovernanceItemReportResponseBody {
	s.Marker = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMetricType(v string) *GetGovernanceItemReportResponseBody {
	s.MetricType = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetMetricValue(v interface{}) *GetGovernanceItemReportResponseBody {
	s.MetricValue = v
	return s
}

func (s *GetGovernanceItemReportResponseBody) SetRequestId(v string) *GetGovernanceItemReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGovernanceItemReportResponseBody) Validate() error {
	if s.ColumnsSchema != nil {
		if err := s.ColumnsSchema.Validate(); err != nil {
			return err
		}
	}
	if s.ColumnsValue != nil {
		if err := s.ColumnsValue.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGovernanceItemReportResponseBodyColumnsSchema struct {
	ColumnSchema []*GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema `json:"ColumnSchema,omitempty" xml:"ColumnSchema,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsSchema) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsSchema) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchema) GetColumnSchema() []*GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema {
	return s.ColumnSchema
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchema) SetColumnSchema(v []*GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) *GetGovernanceItemReportResponseBodyColumnsSchema {
	s.ColumnSchema = v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchema) Validate() error {
	if s.ColumnSchema != nil {
		for _, item := range s.ColumnSchema {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema struct {
	// example:
	//
	// UserPrincipalName
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// example:
	//
	// String
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
}

func (s GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) GetColumnName() *string {
	return s.ColumnName
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) GetColumnType() *string {
	return s.ColumnType
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) SetColumnName(v string) *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema {
	s.ColumnName = &v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) SetColumnType(v string) *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema {
	s.ColumnType = &v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsSchemaColumnSchema) Validate() error {
	return dara.Validate(s)
}

type GetGovernanceItemReportResponseBodyColumnsValue struct {
	ColumnRow []*GetGovernanceItemReportResponseBodyColumnsValueColumnRow `json:"ColumnRow,omitempty" xml:"ColumnRow,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsValue) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsValue) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsValue) GetColumnRow() []*GetGovernanceItemReportResponseBodyColumnsValueColumnRow {
	return s.ColumnRow
}

func (s *GetGovernanceItemReportResponseBodyColumnsValue) SetColumnRow(v []*GetGovernanceItemReportResponseBodyColumnsValueColumnRow) *GetGovernanceItemReportResponseBodyColumnsValue {
	s.ColumnRow = v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsValue) Validate() error {
	if s.ColumnRow != nil {
		for _, item := range s.ColumnRow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetGovernanceItemReportResponseBodyColumnsValueColumnRow struct {
	ColumnValue []interface{} `json:"ColumnValue,omitempty" xml:"ColumnValue,omitempty" type:"Repeated"`
}

func (s GetGovernanceItemReportResponseBodyColumnsValueColumnRow) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceItemReportResponseBodyColumnsValueColumnRow) GoString() string {
	return s.String()
}

func (s *GetGovernanceItemReportResponseBodyColumnsValueColumnRow) GetColumnValue() []interface{} {
	return s.ColumnValue
}

func (s *GetGovernanceItemReportResponseBodyColumnsValueColumnRow) SetColumnValue(v []interface{}) *GetGovernanceItemReportResponseBodyColumnsValueColumnRow {
	s.ColumnValue = v
	return s
}

func (s *GetGovernanceItemReportResponseBodyColumnsValueColumnRow) Validate() error {
	return dara.Validate(s)
}
