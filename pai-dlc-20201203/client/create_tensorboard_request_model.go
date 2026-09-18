// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTensorboardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *CreateTensorboardRequest
	GetAccessibility() *string
	SetCpu(v int64) *CreateTensorboardRequest
	GetCpu() *int64
	SetDataSourceId(v string) *CreateTensorboardRequest
	GetDataSourceId() *string
	SetDataSourceType(v string) *CreateTensorboardRequest
	GetDataSourceType() *string
	SetDataSources(v []*DataSourceItem) *CreateTensorboardRequest
	GetDataSources() []*DataSourceItem
	SetDisplayName(v string) *CreateTensorboardRequest
	GetDisplayName() *string
	SetJobId(v string) *CreateTensorboardRequest
	GetJobId() *string
	SetMaxRunningTimeMinutes(v int64) *CreateTensorboardRequest
	GetMaxRunningTimeMinutes() *int64
	SetMemory(v int64) *CreateTensorboardRequest
	GetMemory() *int64
	SetOptions(v string) *CreateTensorboardRequest
	GetOptions() *string
	SetPriority(v string) *CreateTensorboardRequest
	GetPriority() *string
	SetQuotaId(v string) *CreateTensorboardRequest
	GetQuotaId() *string
	SetSourceId(v string) *CreateTensorboardRequest
	GetSourceId() *string
	SetSourceType(v string) *CreateTensorboardRequest
	GetSourceType() *string
	SetSummaryPath(v string) *CreateTensorboardRequest
	GetSummaryPath() *string
	SetSummaryRelativePath(v string) *CreateTensorboardRequest
	GetSummaryRelativePath() *string
	SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *CreateTensorboardRequest
	GetTensorboardDataSources() []*TensorboardDataSourceSpec
	SetTensorboardSpec(v *TensorboardSpec) *CreateTensorboardRequest
	GetTensorboardSpec() *TensorboardSpec
	SetUri(v string) *CreateTensorboardRequest
	GetUri() *string
	SetWorkspaceId(v string) *CreateTensorboardRequest
	GetWorkspaceId() *string
}

type CreateTensorboardRequest struct {
	// The visibility of the job. Valid values:
	//
	// - PUBLIC: Visible to all members in the workspace.
	//
	// - PRIVATE: Visible only to you and administrators in the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The number of CPU cores.
	//
	// example:
	//
	// 1
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The dataset ID. <props="china">For information about how to obtain the dataset ID, see [ListDatasets](https://help.aliyun.com/document_detail/457222.html).
	//
	// example:
	//
	// d-xxxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The dataset type. Valid values:
	//
	// - OSS
	//
	// - NAS
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The datasource configurations.
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The TensorBoard name.
	//
	// example:
	//
	// tensorboard
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The job ID. For information about how to obtain the job ID, see [ListJobs](https://help.aliyun.com/document_detail/459676.html).
	//
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The maximum running duration. Unit: minutes.
	//
	// example:
	//
	// 240
	MaxRunningTimeMinutes *int64 `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	// The memory size. Unit: GB.
	//
	// example:
	//
	// 1000
	Memory *int64 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The extended field of custom dataset in JSON format. Currently, MountPath is supported, which specifies the custom mount path of custom dataset.
	//
	// example:
	//
	// {"mountpath":"/root/data/"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. This is an optional parameter. Default value: 1. Valid values: 1 to 9.
	//
	// - 1: the lowest priority.
	//
	// - 9: the highest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota ID. This parameter is required when you create a TensorBoard job by using resources from a resource quota. <props="china">For information about how to obtain the resource quota ID, see [ListQuotas](https://help.aliyun.com/document_detail/2628071.html).
	//
	// <props="china">Published only on China site.
	//
	// Currently, only whitelisted users can create TensorBoard jobs by using resource quota resources. To use this feature, contact us.
	//
	// example:
	//
	// quota12345
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// The source ID.
	//
	// example:
	//
	// dlc-xxxxxx
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type.
	//
	// example:
	//
	// job
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The summary directory.
	//
	// example:
	//
	// /root/data/
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// The summary relative directory.
	//
	// example:
	//
	// /summary/
	SummaryRelativePath *string `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	// The list of dataset configurations mounted to the TensorBoard job.
	TensorboardDataSources []*TensorboardDataSourceSpec `json:"TensorboardDataSources,omitempty" xml:"TensorboardDataSources,omitempty" type:"Repeated"`
	// The pay-as-you-go configuration for TensorBoard, which is used to create a TensorBoard job that uses pay-as-you-go resources.
	TensorboardSpec *TensorboardSpec `json:"TensorboardSpec,omitempty" xml:"TensorboardSpec,omitempty"`
	// The URI of the dataset:
	//
	// - If DataSourceType is set to OSS, the format is `oss://[oss-bucket].[endpoint]/[path]`.
	//
	// - If DataSourceType is set to NAS, the format is `nas://[nas-filesystem-id].[region]/[path]`.
	//
	// example:
	//
	// oss://.oss-cn-shanghai-finance-1.aliyuncs.com/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The workspace ID. <props="china">For information about how to obtain the workspace ID, see [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html).
	//
	// example:
	//
	// 123***
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTensorboardRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *CreateTensorboardRequest) GetAccessibility() *string {
	return s.Accessibility
}

func (s *CreateTensorboardRequest) GetCpu() *int64 {
	return s.Cpu
}

func (s *CreateTensorboardRequest) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *CreateTensorboardRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *CreateTensorboardRequest) GetDataSources() []*DataSourceItem {
	return s.DataSources
}

func (s *CreateTensorboardRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateTensorboardRequest) GetJobId() *string {
	return s.JobId
}

func (s *CreateTensorboardRequest) GetMaxRunningTimeMinutes() *int64 {
	return s.MaxRunningTimeMinutes
}

func (s *CreateTensorboardRequest) GetMemory() *int64 {
	return s.Memory
}

func (s *CreateTensorboardRequest) GetOptions() *string {
	return s.Options
}

func (s *CreateTensorboardRequest) GetPriority() *string {
	return s.Priority
}

func (s *CreateTensorboardRequest) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CreateTensorboardRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateTensorboardRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateTensorboardRequest) GetSummaryPath() *string {
	return s.SummaryPath
}

func (s *CreateTensorboardRequest) GetSummaryRelativePath() *string {
	return s.SummaryRelativePath
}

func (s *CreateTensorboardRequest) GetTensorboardDataSources() []*TensorboardDataSourceSpec {
	return s.TensorboardDataSources
}

func (s *CreateTensorboardRequest) GetTensorboardSpec() *TensorboardSpec {
	return s.TensorboardSpec
}

func (s *CreateTensorboardRequest) GetUri() *string {
	return s.Uri
}

func (s *CreateTensorboardRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateTensorboardRequest) SetAccessibility(v string) *CreateTensorboardRequest {
	s.Accessibility = &v
	return s
}

func (s *CreateTensorboardRequest) SetCpu(v int64) *CreateTensorboardRequest {
	s.Cpu = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceId(v string) *CreateTensorboardRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceType(v string) *CreateTensorboardRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSources(v []*DataSourceItem) *CreateTensorboardRequest {
	s.DataSources = v
	return s
}

func (s *CreateTensorboardRequest) SetDisplayName(v string) *CreateTensorboardRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTensorboardRequest) SetJobId(v string) *CreateTensorboardRequest {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *CreateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *CreateTensorboardRequest) SetMemory(v int64) *CreateTensorboardRequest {
	s.Memory = &v
	return s
}

func (s *CreateTensorboardRequest) SetOptions(v string) *CreateTensorboardRequest {
	s.Options = &v
	return s
}

func (s *CreateTensorboardRequest) SetPriority(v string) *CreateTensorboardRequest {
	s.Priority = &v
	return s
}

func (s *CreateTensorboardRequest) SetQuotaId(v string) *CreateTensorboardRequest {
	s.QuotaId = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceId(v string) *CreateTensorboardRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceType(v string) *CreateTensorboardRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryPath(v string) *CreateTensorboardRequest {
	s.SummaryPath = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryRelativePath(v string) *CreateTensorboardRequest {
	s.SummaryRelativePath = &v
	return s
}

func (s *CreateTensorboardRequest) SetTensorboardDataSources(v []*TensorboardDataSourceSpec) *CreateTensorboardRequest {
	s.TensorboardDataSources = v
	return s
}

func (s *CreateTensorboardRequest) SetTensorboardSpec(v *TensorboardSpec) *CreateTensorboardRequest {
	s.TensorboardSpec = v
	return s
}

func (s *CreateTensorboardRequest) SetUri(v string) *CreateTensorboardRequest {
	s.Uri = &v
	return s
}

func (s *CreateTensorboardRequest) SetWorkspaceId(v string) *CreateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTensorboardRequest) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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
