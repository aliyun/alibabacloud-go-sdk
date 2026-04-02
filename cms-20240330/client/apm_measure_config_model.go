// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApmMeasureConfig interface {
	dara.Model
	String() string
	GoString() string
	SetGroupBy(v []*string) *ApmMeasureConfig
	GetGroupBy() []*string
	SetMeasureCode(v string) *ApmMeasureConfig
	GetMeasureCode() *string
	SetWindowSecs(v int32) *ApmMeasureConfig
	GetWindowSecs() *int32
}

type ApmMeasureConfig struct {
	// 分组维度
	GroupBy []*string `json:"groupBy,omitempty" xml:"groupBy,omitempty" type:"Repeated"`
	// 指标代码
	//
	// This parameter is required.
	MeasureCode *string `json:"measureCode,omitempty" xml:"measureCode,omitempty"`
	// 查询时间窗口（秒）
	//
	// This parameter is required.
	WindowSecs *int32 `json:"windowSecs,omitempty" xml:"windowSecs,omitempty"`
}

func (s ApmMeasureConfig) String() string {
	return dara.Prettify(s)
}

func (s ApmMeasureConfig) GoString() string {
	return s.String()
}

func (s *ApmMeasureConfig) GetGroupBy() []*string {
	return s.GroupBy
}

func (s *ApmMeasureConfig) GetMeasureCode() *string {
	return s.MeasureCode
}

func (s *ApmMeasureConfig) GetWindowSecs() *int32 {
	return s.WindowSecs
}

func (s *ApmMeasureConfig) SetGroupBy(v []*string) *ApmMeasureConfig {
	s.GroupBy = v
	return s
}

func (s *ApmMeasureConfig) SetMeasureCode(v string) *ApmMeasureConfig {
	s.MeasureCode = &v
	return s
}

func (s *ApmMeasureConfig) SetWindowSecs(v int32) *ApmMeasureConfig {
	s.WindowSecs = &v
	return s
}

func (s *ApmMeasureConfig) Validate() error {
	return dara.Validate(s)
}
