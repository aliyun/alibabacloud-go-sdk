// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDynamicChartJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAxisParams(v string) *SubmitDynamicChartJobRequest
	GetAxisParams() *string
	SetBackground(v string) *SubmitDynamicChartJobRequest
	GetBackground() *string
	SetChartConfig(v string) *SubmitDynamicChartJobRequest
	GetChartConfig() *string
	SetChartTitle(v string) *SubmitDynamicChartJobRequest
	GetChartTitle() *string
	SetChartType(v string) *SubmitDynamicChartJobRequest
	GetChartType() *string
	SetDataSource(v string) *SubmitDynamicChartJobRequest
	GetDataSource() *string
	SetDescription(v string) *SubmitDynamicChartJobRequest
	GetDescription() *string
	SetInput(v string) *SubmitDynamicChartJobRequest
	GetInput() *string
	SetOutputConfig(v string) *SubmitDynamicChartJobRequest
	GetOutputConfig() *string
	SetSubtitle(v string) *SubmitDynamicChartJobRequest
	GetSubtitle() *string
	SetTitle(v string) *SubmitDynamicChartJobRequest
	GetTitle() *string
	SetUnit(v string) *SubmitDynamicChartJobRequest
	GetUnit() *string
	SetUserData(v string) *SubmitDynamicChartJobRequest
	GetUserData() *string
}

type SubmitDynamicChartJobRequest struct {
	// The axis style settings. If `XAxisFontInterval` is not specified or is set to 0, the system automatically calculates the interval.
	//
	// example:
	//
	// {"FontFile":"Microsoft YaHei","XAxisFontSize":"30","YAxisFontSize":"30","XAxisFontInterval":"30","AxisColor":"30"}
	AxisParams *string `json:"AxisParams,omitempty" xml:"AxisParams,omitempty"`
	// The settings for the chart background. You can specify a Background Color or a Background Image.
	//
	// example:
	//
	// {"Color":"#000000","ImageUrl":"http://your-bucket.oss-cn-shanghai.aliyuncs.com/obj.jpg"}
	Background *string `json:"Background,omitempty" xml:"Background,omitempty"`
	// The overall configuration for the chart\\"s appearance and timing.
	//
	// example:
	//
	// {"Style":"Normal","TitleStartTime":"3000","ChartStartTime":"3000","VideoDuration":"15000"}
	ChartConfig *string `json:"ChartConfig,omitempty" xml:"ChartConfig,omitempty"`
	// The chart title settings.
	//
	// example:
	//
	// {"Text":" primary title","FontSize":"30","FontColor":"#000000","FontFile":"Microsoft YaHei"}
	ChartTitle *string `json:"ChartTitle,omitempty" xml:"ChartTitle,omitempty"`
	// The chart type to generate.
	//
	// This parameter is required.
	//
	// example:
	//
	// Line
	ChartType *string `json:"ChartType,omitempty" xml:"ChartType,omitempty"`
	// The data source label settings.
	//
	// example:
	//
	// {"Text":"Data source","FontSize":"30","FontColor":"#000000","FontFile":"Microsoft YaHei"}
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The job description.
	//
	// example:
	//
	// Task description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The input data configuration, which specifies the Excel file location.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"XlsFile":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/obj.xls"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The output video configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"MediaURL":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/obj.mp4","Bitrate":2000,"Width":800,"Height":680}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The subtitle settings.
	//
	// example:
	//
	// {"Text":"Subtitle","FontSize":"30","FontColor":"#000000","FontFile":"Microsoft YaHei"}
	Subtitle *string `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The job title.
	//
	// example:
	//
	// Task title
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The data unit label settings.
	//
	// example:
	//
	// {"Text":"Unit","FontSize":"30","FontColor":"#000000","FontFile":"Microsoft YaHei"}
	Unit *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
	// The custom user data in JSON format.
	//
	// example:
	//
	// {"user":"data"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitDynamicChartJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDynamicChartJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitDynamicChartJobRequest) GetAxisParams() *string {
	return s.AxisParams
}

func (s *SubmitDynamicChartJobRequest) GetBackground() *string {
	return s.Background
}

func (s *SubmitDynamicChartJobRequest) GetChartConfig() *string {
	return s.ChartConfig
}

func (s *SubmitDynamicChartJobRequest) GetChartTitle() *string {
	return s.ChartTitle
}

func (s *SubmitDynamicChartJobRequest) GetChartType() *string {
	return s.ChartType
}

func (s *SubmitDynamicChartJobRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *SubmitDynamicChartJobRequest) GetDescription() *string {
	return s.Description
}

func (s *SubmitDynamicChartJobRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitDynamicChartJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitDynamicChartJobRequest) GetSubtitle() *string {
	return s.Subtitle
}

func (s *SubmitDynamicChartJobRequest) GetTitle() *string {
	return s.Title
}

func (s *SubmitDynamicChartJobRequest) GetUnit() *string {
	return s.Unit
}

func (s *SubmitDynamicChartJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitDynamicChartJobRequest) SetAxisParams(v string) *SubmitDynamicChartJobRequest {
	s.AxisParams = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetBackground(v string) *SubmitDynamicChartJobRequest {
	s.Background = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetChartConfig(v string) *SubmitDynamicChartJobRequest {
	s.ChartConfig = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetChartTitle(v string) *SubmitDynamicChartJobRequest {
	s.ChartTitle = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetChartType(v string) *SubmitDynamicChartJobRequest {
	s.ChartType = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetDataSource(v string) *SubmitDynamicChartJobRequest {
	s.DataSource = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetDescription(v string) *SubmitDynamicChartJobRequest {
	s.Description = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetInput(v string) *SubmitDynamicChartJobRequest {
	s.Input = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetOutputConfig(v string) *SubmitDynamicChartJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetSubtitle(v string) *SubmitDynamicChartJobRequest {
	s.Subtitle = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetTitle(v string) *SubmitDynamicChartJobRequest {
	s.Title = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetUnit(v string) *SubmitDynamicChartJobRequest {
	s.Unit = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) SetUserData(v string) *SubmitDynamicChartJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitDynamicChartJobRequest) Validate() error {
	return dara.Validate(s)
}
