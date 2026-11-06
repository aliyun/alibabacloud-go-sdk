// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataProcessJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDataProcessJobRequest
	GetDescription() *string
	SetDestinationConfig(v string) *CreateDataProcessJobRequest
	GetDestinationConfig() *string
	SetDestinationId(v string) *CreateDataProcessJobRequest
	GetDestinationId() *string
	SetDestinationPath(v string) *CreateDataProcessJobRequest
	GetDestinationPath() *string
	SetDestinationType(v string) *CreateDataProcessJobRequest
	GetDestinationType() *string
	SetModelName(v string) *CreateDataProcessJobRequest
	GetModelName() *string
	SetProcessingMode(v string) *CreateDataProcessJobRequest
	GetProcessingMode() *string
	SetRegionId(v string) *CreateDataProcessJobRequest
	GetRegionId() *string
	SetSourceConfig(v string) *CreateDataProcessJobRequest
	GetSourceConfig() *string
	SetSourceId(v string) *CreateDataProcessJobRequest
	GetSourceId() *string
	SetSourcePath(v string) *CreateDataProcessJobRequest
	GetSourcePath() *string
	SetSourceType(v string) *CreateDataProcessJobRequest
	GetSourceType() *string
	SetTemplateId(v string) *CreateDataProcessJobRequest
	GetTemplateId() *string
}

type CreateDataProcessJobRequest struct {
	// The description of the data processing task.
	//
	// example:
	//
	// Image data processing task
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The JSON string for the destination database connection configuration, including Account and Password.
	//
	// example:
	//
	// {"Account":"example_user","Password":"example_password"}
	DestinationConfig *string `json:"DestinationConfig,omitempty" xml:"DestinationConfig,omitempty"`
	// The ID of the destination PolarDB cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// pc-test123456
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The name of the destination database.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdb
	DestinationPath *string `json:"DestinationPath,omitempty" xml:"DestinationPath,omitempty"`
	// The type of the destination. Only PolarDB is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarDB
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The name of the model used for feature extraction.
	//
	// example:
	//
	// model-v1
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The processing mode. Valid values: Full, Incremental, and Both.
	//
	// This parameter is required.
	//
	// example:
	//
	// Full
	ProcessingMode *string `json:"ProcessingMode,omitempty" xml:"ProcessingMode,omitempty"`
	// The ID of the region where the task resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The additional configuration for the data source in JSON string format. This parameter can be left empty.
	//
	// example:
	//
	// {}
	SourceConfig *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	// The ID of the PolarFS instance that serves as the data source.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-123456
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source data path in PolarFS. An empty string indicates the default path.
	//
	// This parameter is required.
	//
	// example:
	//
	// /images
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The type of the data source. Only PolarFS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// PolarFS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The ID of the task template used to process data.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s CreateDataProcessJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataProcessJobRequest) GoString() string {
	return s.String()
}

func (s *CreateDataProcessJobRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataProcessJobRequest) GetDestinationConfig() *string {
	return s.DestinationConfig
}

func (s *CreateDataProcessJobRequest) GetDestinationId() *string {
	return s.DestinationId
}

func (s *CreateDataProcessJobRequest) GetDestinationPath() *string {
	return s.DestinationPath
}

func (s *CreateDataProcessJobRequest) GetDestinationType() *string {
	return s.DestinationType
}

func (s *CreateDataProcessJobRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateDataProcessJobRequest) GetProcessingMode() *string {
	return s.ProcessingMode
}

func (s *CreateDataProcessJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDataProcessJobRequest) GetSourceConfig() *string {
	return s.SourceConfig
}

func (s *CreateDataProcessJobRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *CreateDataProcessJobRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *CreateDataProcessJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateDataProcessJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateDataProcessJobRequest) SetDescription(v string) *CreateDataProcessJobRequest {
	s.Description = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetDestinationConfig(v string) *CreateDataProcessJobRequest {
	s.DestinationConfig = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetDestinationId(v string) *CreateDataProcessJobRequest {
	s.DestinationId = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetDestinationPath(v string) *CreateDataProcessJobRequest {
	s.DestinationPath = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetDestinationType(v string) *CreateDataProcessJobRequest {
	s.DestinationType = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetModelName(v string) *CreateDataProcessJobRequest {
	s.ModelName = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetProcessingMode(v string) *CreateDataProcessJobRequest {
	s.ProcessingMode = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetRegionId(v string) *CreateDataProcessJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetSourceConfig(v string) *CreateDataProcessJobRequest {
	s.SourceConfig = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetSourceId(v string) *CreateDataProcessJobRequest {
	s.SourceId = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetSourcePath(v string) *CreateDataProcessJobRequest {
	s.SourcePath = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetSourceType(v string) *CreateDataProcessJobRequest {
	s.SourceType = &v
	return s
}

func (s *CreateDataProcessJobRequest) SetTemplateId(v string) *CreateDataProcessJobRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateDataProcessJobRequest) Validate() error {
	return dara.Validate(s)
}
