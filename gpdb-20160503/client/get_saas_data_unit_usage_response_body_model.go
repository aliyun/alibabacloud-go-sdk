// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasDataUnitUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistribution(v []*GetSaasDataUnitUsageResponseBodyDistribution) *GetSaasDataUnitUsageResponseBody
	GetDistribution() []*GetSaasDataUnitUsageResponseBodyDistribution
	SetEndTime(v string) *GetSaasDataUnitUsageResponseBody
	GetEndTime() *string
	SetRequestId(v string) *GetSaasDataUnitUsageResponseBody
	GetRequestId() *string
	SetServiceId(v string) *GetSaasDataUnitUsageResponseBody
	GetServiceId() *string
	SetStartTime(v string) *GetSaasDataUnitUsageResponseBody
	GetStartTime() *string
	SetSummary(v *GetSaasDataUnitUsageResponseBodySummary) *GetSaasDataUnitUsageResponseBody
	GetSummary() *GetSaasDataUnitUsageResponseBodySummary
	SetTimeSeries(v []*GetSaasDataUnitUsageResponseBodyTimeSeries) *GetSaasDataUnitUsageResponseBody
	GetTimeSeries() []*GetSaasDataUnitUsageResponseBodyTimeSeries
	SetWarnings(v []*GetSaasDataUnitUsageResponseBodyWarnings) *GetSaasDataUnitUsageResponseBody
	GetWarnings() []*GetSaasDataUnitUsageResponseBodyWarnings
}

type GetSaasDataUnitUsageResponseBody struct {
	// 按DataUnit降序，未知金额排末尾，同值按场景码、项目ID和名称升序。按场景和项目ID分组；无ID但有名称按名称分组；两者均缺失的记录为未关联项目。完整结果各项之和等于汇总；缺失或空白的日志场景归other（其他）；有值但未配置映射的场景保留原值作为编码和展示名，正常计入总量；无记录的other不创建空行。
	Distribution []*GetSaasDataUnitUsageResponseBodyDistribution `json:"Distribution,omitempty" xml:"Distribution,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-10-20T02:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// mock_req_success
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// agdb-example
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// 2026-10-20T00:00:00Z
	StartTime *string                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Summary   *GetSaasDataUnitUsageResponseBodySummary `json:"Summary,omitempty" xml:"Summary,omitempty" type:"Struct"`
	// 按UTC桶起点升序。完整查询补齐零桶；未知金额为null，不以零替代。
	TimeSeries []*GetSaasDataUnitUsageResponseBodyTimeSeries `json:"TimeSeries,omitempty" xml:"TimeSeries,omitempty" type:"Repeated"`
	Warnings   []*GetSaasDataUnitUsageResponseBodyWarnings   `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
}

func (s GetSaasDataUnitUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBody) GetDistribution() []*GetSaasDataUnitUsageResponseBodyDistribution {
	return s.Distribution
}

func (s *GetSaasDataUnitUsageResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *GetSaasDataUnitUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSaasDataUnitUsageResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasDataUnitUsageResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *GetSaasDataUnitUsageResponseBody) GetSummary() *GetSaasDataUnitUsageResponseBodySummary {
	return s.Summary
}

func (s *GetSaasDataUnitUsageResponseBody) GetTimeSeries() []*GetSaasDataUnitUsageResponseBodyTimeSeries {
	return s.TimeSeries
}

func (s *GetSaasDataUnitUsageResponseBody) GetWarnings() []*GetSaasDataUnitUsageResponseBodyWarnings {
	return s.Warnings
}

func (s *GetSaasDataUnitUsageResponseBody) SetDistribution(v []*GetSaasDataUnitUsageResponseBodyDistribution) *GetSaasDataUnitUsageResponseBody {
	s.Distribution = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetEndTime(v string) *GetSaasDataUnitUsageResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetRequestId(v string) *GetSaasDataUnitUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetServiceId(v string) *GetSaasDataUnitUsageResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetStartTime(v string) *GetSaasDataUnitUsageResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetSummary(v *GetSaasDataUnitUsageResponseBodySummary) *GetSaasDataUnitUsageResponseBody {
	s.Summary = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetTimeSeries(v []*GetSaasDataUnitUsageResponseBodyTimeSeries) *GetSaasDataUnitUsageResponseBody {
	s.TimeSeries = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) SetWarnings(v []*GetSaasDataUnitUsageResponseBodyWarnings) *GetSaasDataUnitUsageResponseBody {
	s.Warnings = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBody) Validate() error {
	if s.Distribution != nil {
		for _, item := range s.Distribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Summary != nil {
		if err := s.Summary.Validate(); err != nil {
			return err
		}
	}
	if s.TimeSeries != nil {
		for _, item := range s.TimeSeries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Warnings != nil {
		for _, item := range s.Warnings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSaasDataUnitUsageResponseBodyDistribution struct {
	// 精确非负十进制字符串，不使用浮点或指数格式；金额不完整时为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// 调用时的项目归因；缺失时为null，不影响总量。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// project-1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 同一场景和项目ID取查询范围内最新非空名称快照，同秒取字典序最大值；缺失为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 场景重制
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// drama_game_oneshot
	UsageScene *string `json:"UsageScene,omitempty" xml:"UsageScene,omitempty"`
	// 由Meta原值到中文名的字典补充，不是SLS字段。UsageScene始终保留日志原值；无中文名则名称也用原值；空白归其他。同中文名不同编码不合并。
	//
	// example:
	//
	// 游戏生成
	UsageSceneName *string `json:"UsageSceneName,omitempty" xml:"UsageSceneName,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyDistribution) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyDistribution) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetUsageScene() *string {
	return s.UsageScene
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) GetUsageSceneName() *string {
	return s.UsageSceneName
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetProjectId(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.ProjectId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetProjectName(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.ProjectName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetUsageScene(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.UsageScene = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) SetUsageSceneName(v string) *GetSaasDataUnitUsageResponseBodyDistribution {
	s.UsageSceneName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyDistribution) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodySummary struct {
	// 精确非负十进制字符串，不使用浮点或指数格式；金额不完整时为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodySummary) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodySummary) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodySummary) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodySummary) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodySummary {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodySummary) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodyTimeSeries struct {
	// 精确非负十进制字符串，不使用浮点或指数格式；金额不完整时为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string                                             `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	Scenes   []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes `json:"Scenes,omitempty" xml:"Scenes,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-10-20T00:00:00Z
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeries) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeries) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetScenes() []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	return s.Scenes
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) GetTime() *string {
	return s.Time
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetScenes(v []*GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.Scenes = v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) SetTime(v string) *GetSaasDataUnitUsageResponseBodyTimeSeries {
	s.Time = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeries) Validate() error {
	if s.Scenes != nil {
		for _, item := range s.Scenes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSaasDataUnitUsageResponseBodyTimeSeriesScenes struct {
	// 精确非负十进制字符串，不使用浮点或指数格式；金额不完整时为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2.82
	DataUnit *string `json:"DataUnit,omitempty" xml:"DataUnit,omitempty"`
	// 调用时的项目归因；缺失时为null，不影响总量。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// project-1
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// 同一场景和项目ID取查询范围内最新非空名称快照，同秒取字典序最大值；缺失为null。
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 场景重制
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// game_translation
	UsageScene *string `json:"UsageScene,omitempty" xml:"UsageScene,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetDataUnit() *string {
	return s.DataUnit
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) GetUsageScene() *string {
	return s.UsageScene
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetDataUnit(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.DataUnit = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetProjectId(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.ProjectId = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetProjectName(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.ProjectName = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) SetUsageScene(v string) *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes {
	s.UsageScene = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyTimeSeriesScenes) Validate() error {
	return dara.Validate(s)
}

type GetSaasDataUnitUsageResponseBodyWarnings struct {
	// 包括DATA_INCOMPLETE和DATA_UNIT_PENDING。
	//
	// example:
	//
	// DATA_INCOMPLETE
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 部分用量尚未就绪
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetSaasDataUnitUsageResponseBodyWarnings) String() string {
	return dara.Prettify(s)
}

func (s GetSaasDataUnitUsageResponseBodyWarnings) GoString() string {
	return s.String()
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) GetCode() *string {
	return s.Code
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) GetMessage() *string {
	return s.Message
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) SetCode(v string) *GetSaasDataUnitUsageResponseBodyWarnings {
	s.Code = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) SetMessage(v string) *GetSaasDataUnitUsageResponseBodyWarnings {
	s.Message = &v
	return s
}

func (s *GetSaasDataUnitUsageResponseBodyWarnings) Validate() error {
	return dara.Validate(s)
}
