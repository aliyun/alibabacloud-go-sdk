// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTensorboard interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Tensorboard
	GetAccessibility() *string
	SetCpu(v int64) *Tensorboard
	GetCpu() *int64
	SetDataSourceId(v string) *Tensorboard
	GetDataSourceId() *string
	SetDataSourceType(v string) *Tensorboard
	GetDataSourceType() *string
	SetDisplayName(v string) *Tensorboard
	GetDisplayName() *string
	SetDuration(v string) *Tensorboard
	GetDuration() *string
	SetGmtCreateTime(v string) *Tensorboard
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *Tensorboard
	GetGmtFinishTime() *string
	SetGmtModifyTime(v string) *Tensorboard
	GetGmtModifyTime() *string
	SetJobId(v string) *Tensorboard
	GetJobId() *string
	SetMaxRunningTimeMinutes(v int64) *Tensorboard
	GetMaxRunningTimeMinutes() *int64
	SetMemory(v int64) *Tensorboard
	GetMemory() *int64
	SetOptions(v string) *Tensorboard
	GetOptions() *string
	SetPriority(v string) *Tensorboard
	GetPriority() *string
	SetQuotaId(v string) *Tensorboard
	GetQuotaId() *string
	SetQuotaName(v string) *Tensorboard
	GetQuotaName() *string
	SetReasonCode(v string) *Tensorboard
	GetReasonCode() *string
	SetReasonMessage(v string) *Tensorboard
	GetReasonMessage() *string
	SetRequestId(v string) *Tensorboard
	GetRequestId() *string
	SetStatus(v string) *Tensorboard
	GetStatus() *string
	SetSummaryPath(v string) *Tensorboard
	GetSummaryPath() *string
	SetSummaryRelativePath(v string) *Tensorboard
	GetSummaryRelativePath() *string
	SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *Tensorboard
	GetTensorboardDataSources() []*TensorboardDataSourceSpec
	SetTensorboardId(v string) *Tensorboard
	GetTensorboardId() *string
	SetTensorboardSpec(v *TensorboardSpec) *Tensorboard
	GetTensorboardSpec() *TensorboardSpec
	SetTensorboardUrl(v string) *Tensorboard
	GetTensorboardUrl() *string
	SetToken(v string) *Tensorboard
	GetToken() *string
	SetUserId(v string) *Tensorboard
	GetUserId() *string
	SetUsername(v string) *Tensorboard
	GetUsername() *string
	SetWorkspaceId(v string) *Tensorboard
	GetWorkspaceId() *string
}

type Tensorboard struct {
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Cpu           *int64  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// example:
	//
	// datasource-test
	DataSourceId   *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// example:
	//
	// test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 1234567
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:00Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:36:00Z
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// example:
	//
	// dlc-20210114104214-vf9lowjt3pso
	JobId                 *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Memory                *int64  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	Priority              *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	QuotaId               *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	QuotaName             *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// example:
	//
	// Delete by user
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// Tensorboard is deleted
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// /root/data
	SummaryPath            *string                      `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	SummaryRelativePath    *string                      `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	TensorboardDataSources []*TensorboardDataSourceSpec `json:"TensorboardDataSources,omitempty" xml:"TensorboardDataSources,omitempty" type:"Repeated"`
	// example:
	//
	// tensorboard-xxx
	TensorboardId   *string          `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	TensorboardSpec *TensorboardSpec `json:"TensorboardSpec,omitempty" xml:"TensorboardSpec,omitempty"`
	// example:
	//
	// http://xxxxxx
	TensorboardUrl *string `json:"TensorboardUrl,omitempty" xml:"TensorboardUrl,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// lycxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// tensorboard.pai
	Username    *string `json:"Username,omitempty" xml:"Username,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Tensorboard) String() string {
	return dara.Prettify(s)
}

func (s Tensorboard) GoString() string {
	return s.String()
}

func (s *Tensorboard) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Tensorboard) GetCpu() *int64 {
	return s.Cpu
}

func (s *Tensorboard) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *Tensorboard) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *Tensorboard) GetDisplayName() *string {
	return s.DisplayName
}

func (s *Tensorboard) GetDuration() *string {
	return s.Duration
}

func (s *Tensorboard) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Tensorboard) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *Tensorboard) GetGmtModifyTime() *string {
	return s.GmtModifyTime
}

func (s *Tensorboard) GetJobId() *string {
	return s.JobId
}

func (s *Tensorboard) GetMaxRunningTimeMinutes() *int64 {
	return s.MaxRunningTimeMinutes
}

func (s *Tensorboard) GetMemory() *int64 {
	return s.Memory
}

func (s *Tensorboard) GetOptions() *string {
	return s.Options
}

func (s *Tensorboard) GetPriority() *string {
	return s.Priority
}

func (s *Tensorboard) GetQuotaId() *string {
	return s.QuotaId
}

func (s *Tensorboard) GetQuotaName() *string {
	return s.QuotaName
}

func (s *Tensorboard) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *Tensorboard) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *Tensorboard) GetRequestId() *string {
	return s.RequestId
}

func (s *Tensorboard) GetStatus() *string {
	return s.Status
}

func (s *Tensorboard) GetSummaryPath() *string {
	return s.SummaryPath
}

func (s *Tensorboard) GetSummaryRelativePath() *string {
	return s.SummaryRelativePath
}

func (s *Tensorboard) GetTensorboardDataSources() []*TensorboardDataSourceSpec {
	return s.TensorboardDataSources
}

func (s *Tensorboard) GetTensorboardId() *string {
	return s.TensorboardId
}

func (s *Tensorboard) GetTensorboardSpec() *TensorboardSpec {
	return s.TensorboardSpec
}

func (s *Tensorboard) GetTensorboardUrl() *string {
	return s.TensorboardUrl
}

func (s *Tensorboard) GetToken() *string {
	return s.Token
}

func (s *Tensorboard) GetUserId() *string {
	return s.UserId
}

func (s *Tensorboard) GetUsername() *string {
	return s.Username
}

func (s *Tensorboard) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Tensorboard) SetAccessibility(v string) *Tensorboard {
	s.Accessibility = &v
	return s
}

func (s *Tensorboard) SetCpu(v int64) *Tensorboard {
	s.Cpu = &v
	return s
}

func (s *Tensorboard) SetDataSourceId(v string) *Tensorboard {
	s.DataSourceId = &v
	return s
}

func (s *Tensorboard) SetDataSourceType(v string) *Tensorboard {
	s.DataSourceType = &v
	return s
}

func (s *Tensorboard) SetDisplayName(v string) *Tensorboard {
	s.DisplayName = &v
	return s
}

func (s *Tensorboard) SetDuration(v string) *Tensorboard {
	s.Duration = &v
	return s
}

func (s *Tensorboard) SetGmtCreateTime(v string) *Tensorboard {
	s.GmtCreateTime = &v
	return s
}

func (s *Tensorboard) SetGmtFinishTime(v string) *Tensorboard {
	s.GmtFinishTime = &v
	return s
}

func (s *Tensorboard) SetGmtModifyTime(v string) *Tensorboard {
	s.GmtModifyTime = &v
	return s
}

func (s *Tensorboard) SetJobId(v string) *Tensorboard {
	s.JobId = &v
	return s
}

func (s *Tensorboard) SetMaxRunningTimeMinutes(v int64) *Tensorboard {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *Tensorboard) SetMemory(v int64) *Tensorboard {
	s.Memory = &v
	return s
}

func (s *Tensorboard) SetOptions(v string) *Tensorboard {
	s.Options = &v
	return s
}

func (s *Tensorboard) SetPriority(v string) *Tensorboard {
	s.Priority = &v
	return s
}

func (s *Tensorboard) SetQuotaId(v string) *Tensorboard {
	s.QuotaId = &v
	return s
}

func (s *Tensorboard) SetQuotaName(v string) *Tensorboard {
	s.QuotaName = &v
	return s
}

func (s *Tensorboard) SetReasonCode(v string) *Tensorboard {
	s.ReasonCode = &v
	return s
}

func (s *Tensorboard) SetReasonMessage(v string) *Tensorboard {
	s.ReasonMessage = &v
	return s
}

func (s *Tensorboard) SetRequestId(v string) *Tensorboard {
	s.RequestId = &v
	return s
}

func (s *Tensorboard) SetStatus(v string) *Tensorboard {
	s.Status = &v
	return s
}

func (s *Tensorboard) SetSummaryPath(v string) *Tensorboard {
	s.SummaryPath = &v
	return s
}

func (s *Tensorboard) SetSummaryRelativePath(v string) *Tensorboard {
	s.SummaryRelativePath = &v
	return s
}

func (s *Tensorboard) SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *Tensorboard {
	s.TensorboardDataSources = v
	return s
}

func (s *Tensorboard) SetTensorboardId(v string) *Tensorboard {
	s.TensorboardId = &v
	return s
}

func (s *Tensorboard) SetTensorboardSpec(v *TensorboardSpec) *Tensorboard {
	s.TensorboardSpec = v
	return s
}

func (s *Tensorboard) SetTensorboardUrl(v string) *Tensorboard {
	s.TensorboardUrl = &v
	return s
}

func (s *Tensorboard) SetToken(v string) *Tensorboard {
	s.Token = &v
	return s
}

func (s *Tensorboard) SetUserId(v string) *Tensorboard {
	s.UserId = &v
	return s
}

func (s *Tensorboard) SetUsername(v string) *Tensorboard {
	s.Username = &v
	return s
}

func (s *Tensorboard) SetWorkspaceId(v string) *Tensorboard {
	s.WorkspaceId = &v
	return s
}

func (s *Tensorboard) Validate() error {
	if s.TensorboardDataSources != nil {
		for _, item := range s.TensorboardDataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TensorboardSpec != nil {
		if err := s.TensorboardSpec.Validate(); err != nil {
			return err
		}
	}
	return nil
}
