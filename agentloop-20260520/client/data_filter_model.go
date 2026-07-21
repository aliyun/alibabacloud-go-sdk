// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataFilter interface {
	dara.Model
	String() string
	GoString() string
	SetMaxRecords(v int32) *DataFilter
	GetMaxRecords() *int32
	SetProvided(v map[string]interface{}) *DataFilter
	GetProvided() map[string]interface{}
	SetQuery(v string) *DataFilter
	GetQuery() *string
	SetSamplingRate(v int32) *DataFilter
	GetSamplingRate() *int32
	SetServiceNames(v []*string) *DataFilter
	GetServiceNames() []*string
}

type DataFilter struct {
	// The maximum number of evaluation records. This takes effect for both backfill and continuous runs. If not specified, the backend does not write a default value.
	//
	// example:
	//
	// 10
	MaxRecords *int32 `json:"maxRecords,omitempty" xml:"maxRecords,omitempty"`
	// The one-time temporary evaluation input content, primarily used for oneshot tasks. The value is stored as a string. Object or array values are serialized to a JSON string.
	//
	// example:
	//
	// {"input":"用户查询订单状态","output":"已查询到订单状态"}
	Provided map[string]interface{} `json:"provided,omitempty" xml:"provided,omitempty"`
	// The data query filter condition. This takes effect together with the evaluator-level filters.query. In Trace scenarios, you can specify filter expressions such as service name, environment, or labels.
	//
	// example:
	//
	// serviceName=\\"checkout-service\\"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The sampling rate percentage. Valid values: 0 to 100. A value of 0 or not specified indicates no sampling. A value of 100 indicates full data. If the value is less than 100, random sampling is applied first, and then the maxRecords limit is applied.
	//
	// example:
	//
	// 100
	SamplingRate *int32    `json:"samplingRate,omitempty" xml:"samplingRate,omitempty"`
	ServiceNames []*string `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
}

func (s DataFilter) String() string {
	return dara.Prettify(s)
}

func (s DataFilter) GoString() string {
	return s.String()
}

func (s *DataFilter) GetMaxRecords() *int32 {
	return s.MaxRecords
}

func (s *DataFilter) GetProvided() map[string]interface{} {
	return s.Provided
}

func (s *DataFilter) GetQuery() *string {
	return s.Query
}

func (s *DataFilter) GetSamplingRate() *int32 {
	return s.SamplingRate
}

func (s *DataFilter) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *DataFilter) SetMaxRecords(v int32) *DataFilter {
	s.MaxRecords = &v
	return s
}

func (s *DataFilter) SetProvided(v map[string]interface{}) *DataFilter {
	s.Provided = v
	return s
}

func (s *DataFilter) SetQuery(v string) *DataFilter {
	s.Query = &v
	return s
}

func (s *DataFilter) SetSamplingRate(v int32) *DataFilter {
	s.SamplingRate = &v
	return s
}

func (s *DataFilter) SetServiceNames(v []*string) *DataFilter {
	s.ServiceNames = v
	return s
}

func (s *DataFilter) Validate() error {
	return dara.Validate(s)
}
