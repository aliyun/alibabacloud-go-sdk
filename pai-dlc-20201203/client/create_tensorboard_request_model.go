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
	// The job visibility. Valid values:
	//
	// 	- PUBLIC: Visible to all members in the workspace.
	//
	// 	- PRIVATE: Visible only to you and the administrator of the workspace.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The number of vCPU cores.
	//
	// example:
	//
	// 1
	Cpu *int64 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// The dataset ID.
	//
	// example:
	//
	// d-xxxxxxxx
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The dataset type. Valid values:
	//
	// 	- OSS
	//
	// 	- NAS
	//
	// example:
	//
	// OSS
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The configurations of the data source.
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The TensorBoard name
	//
	// example:
	//
	// tensorboard
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The job ID. For more information about how to query the job ID, see [ListJobs](https://help.aliyun.com/document_detail/459676.html).
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
	// The extended fields of the dataset are in the JSON format. MountPath: the path to mount the dataset.
	//
	// example:
	//
	// {"mountpath":"/root/data/"}
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// The priority of the job. Default value: 1. Valid values: 1 to 9.
	//
	// 	- 1 is the lowest priority.
	//
	// 	- 9 is the highest priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The resource quota ID. This parameter is required when you create a TensorBoard job by using a resource quota.
	//
	// This feature is currently limited to whitelisted users. If you need to use this feature, contact us.
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
	// The directory of summary.
	//
	// example:
	//
	// /root/data/
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// The relative path of summary.
	//
	// example:
	//
	// /summary/
	SummaryRelativePath *string `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	// The configurations of datasets mounted with the TensorBoard job.
	TensorboardDataSources []*TensorboardDataSourceSpec `json:"TensorboardDataSources,omitempty" xml:"TensorboardDataSources,omitempty" type:"Repeated"`
	// The pay-as-you-go configuration of TensorBoard, which is used to create TensorBoard jobs that use pay-as-you-go resources.
	TensorboardSpec *TensorboardSpec `json:"TensorboardSpec,omitempty" xml:"TensorboardSpec,omitempty"`
	// The dataset URI:
	//
	// 	- Value format when DataSourceType is set to OSS: `oss://[oss-bucket].[endpoint]/[path]`.
	//
	// 	- Value format when DataSourceType is set to NAS:`nas://[nas-filesystem-id].[region]/[path]`.
	//
	// example:
	//
	// oss://.oss-cn-shanghai-finance-1.aliyuncs.com/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	// The workspace ID.
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
