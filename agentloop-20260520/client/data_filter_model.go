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
	// The maximum number of evaluation records. This parameter takes effect for both backfill and continuous runs. If this parameter is not specified, the backend does not apply an additional default value.
	//
	// example:
	//
	// 10
	MaxRecords *int32 `json:"maxRecords,omitempty" xml:"maxRecords,omitempty"`
	// The single-shot input content for temporary evaluation, which is primarily used for one-shot tasks. Values are saved as strings. Object or array values are serialized into JSON strings.
	//
	// example:
	//
	// {"input":"User queries order status","output":"Order status queried"}
	Provided map[string]interface{} `json:"provided,omitempty" xml:"provided,omitempty"`
	// The data query filter conditions that take effect together with the evaluator-level filters.query. For trace scenarios, you can specify filter expressions such as service name, environment, and tags.
	//
	// example:
	//
	// serviceName=\\"checkout-service\\"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The sampling rate percentage. Valid values: 0 to 100. A value of 0 or an unspecified value indicates no sampling. A value of 100 indicates full sampling. If the value is less than 100, random sampling is performed first, and then the maxRecords limit is applied.
	//
	// example:
	//
	// 100
	SamplingRate *int32 `json:"samplingRate,omitempty" xml:"samplingRate,omitempty"`
	// The list of service names used to restrict the services involved in data filtering.
	//
	// example:
	//
	// ["checkout-service"]
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
