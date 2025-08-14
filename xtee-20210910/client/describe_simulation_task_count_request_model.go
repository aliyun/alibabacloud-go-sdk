// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimulationTaskCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSimulationTaskCountRequest
	GetLang() *string
	SetDataSourceConfig(v string) *DescribeSimulationTaskCountRequest
	GetDataSourceConfig() *string
	SetDataSourceType(v string) *DescribeSimulationTaskCountRequest
	GetDataSourceType() *string
	SetEndTime(v int64) *DescribeSimulationTaskCountRequest
	GetEndTime() *int64
	SetEventCode(v string) *DescribeSimulationTaskCountRequest
	GetEventCode() *string
	SetFiltersStr(v string) *DescribeSimulationTaskCountRequest
	GetFiltersStr() *string
	SetRegId(v string) *DescribeSimulationTaskCountRequest
	GetRegId() *string
	SetStartTime(v int64) *DescribeSimulationTaskCountRequest
	GetStartTime() *int64
}

type DescribeSimulationTaskCountRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Data source configuration
	//
	// example:
	//
	// {}
	DataSourceConfig *string `json:"dataSourceConfig,omitempty" xml:"dataSourceConfig,omitempty"`
	// Data source type
	//
	// example:
	//
	// SLS
	DataSourceType *string `json:"dataSourceType,omitempty" xml:"dataSourceType,omitempty"`
	// Task end time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1740016411000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// Event code
	//
	// This parameter is required.
	//
	// example:
	//
	// de_ayfofy4941
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Filter
	//
	// example:
	//
	// {"left":"score","operate":"bw","right":"222,333"}
	FiltersStr *string `json:"filtersStr,omitempty" xml:"filtersStr,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Task start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1739496651000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s DescribeSimulationTaskCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimulationTaskCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSimulationTaskCountRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSimulationTaskCountRequest) GetDataSourceConfig() *string {
	return s.DataSourceConfig
}

func (s *DescribeSimulationTaskCountRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *DescribeSimulationTaskCountRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSimulationTaskCountRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeSimulationTaskCountRequest) GetFiltersStr() *string {
	return s.FiltersStr
}

func (s *DescribeSimulationTaskCountRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSimulationTaskCountRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSimulationTaskCountRequest) SetLang(v string) *DescribeSimulationTaskCountRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetDataSourceConfig(v string) *DescribeSimulationTaskCountRequest {
	s.DataSourceConfig = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetDataSourceType(v string) *DescribeSimulationTaskCountRequest {
	s.DataSourceType = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetEndTime(v int64) *DescribeSimulationTaskCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetEventCode(v string) *DescribeSimulationTaskCountRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetFiltersStr(v string) *DescribeSimulationTaskCountRequest {
	s.FiltersStr = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetRegId(v string) *DescribeSimulationTaskCountRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) SetStartTime(v int64) *DescribeSimulationTaskCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSimulationTaskCountRequest) Validate() error {
	return dara.Validate(s)
}
