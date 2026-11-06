// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreationTime(v string) *DescribeDataProcessJobAttributeResponseBody
	GetCreationTime() *string
	SetDestinationConfig(v string) *DescribeDataProcessJobAttributeResponseBody
	GetDestinationConfig() *string
	SetDestinationId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetDestinationId() *string
	SetDestinationPath(v string) *DescribeDataProcessJobAttributeResponseBody
	GetDestinationPath() *string
	SetDestinationType(v string) *DescribeDataProcessJobAttributeResponseBody
	GetDestinationType() *string
	SetJobId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetJobId() *string
	SetModelName(v string) *DescribeDataProcessJobAttributeResponseBody
	GetModelName() *string
	SetProcessingMode(v string) *DescribeDataProcessJobAttributeResponseBody
	GetProcessingMode() *string
	SetRegionId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetRequestId() *string
	SetSourceId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetSourceId() *string
	SetSourcePath(v string) *DescribeDataProcessJobAttributeResponseBody
	GetSourcePath() *string
	SetSourceType(v string) *DescribeDataProcessJobAttributeResponseBody
	GetSourceType() *string
	SetStatus(v string) *DescribeDataProcessJobAttributeResponseBody
	GetStatus() *string
	SetTemplateId(v string) *DescribeDataProcessJobAttributeResponseBody
	GetTemplateId() *string
	SetTemplateType(v string) *DescribeDataProcessJobAttributeResponseBody
	GetTemplateType() *string
}

type DescribeDataProcessJobAttributeResponseBody struct {
	// The time when the task was created, in RFC 3339 format.
	//
	// example:
	//
	// 2025-01-01T00:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The JSON string of the destination database connection configuration.
	//
	// example:
	//
	// {"Account":"example_user","Password":"example_password"}
	DestinationConfig *string `json:"DestinationConfig,omitempty" xml:"DestinationConfig,omitempty"`
	// The ID of the destination PolarDB cluster.
	//
	// example:
	//
	// pc-test123456
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// The name of the destination database.
	//
	// example:
	//
	// testdb
	DestinationPath *string `json:"DestinationPath,omitempty" xml:"DestinationPath,omitempty"`
	// The type of the destination. Currently, only PolarDB is supported.
	//
	// example:
	//
	// PolarDB
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// The ID of the data processing task.
	//
	// example:
	//
	// job-123456
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the model used for feature extraction.
	//
	// example:
	//
	// model-v1
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The data processing mode. Valid values: Full, Incremental, and Both.
	//
	// example:
	//
	// Full
	ProcessingMode *string `json:"ProcessingMode,omitempty" xml:"ProcessingMode,omitempty"`
	// The region ID of the region where the data processing task is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the source PolarFS instance.
	//
	// example:
	//
	// pfs-123456
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source data path in PolarFS.
	//
	// example:
	//
	// /data/input
	SourcePath *string `json:"SourcePath,omitempty" xml:"SourcePath,omitempty"`
	// The type of the data source. Currently, only PolarFS is supported.
	//
	// example:
	//
	// PolarFS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The status of the data processing task.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the template used by the task.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The node templatetype. Currently, only image is supported.
	//
	// example:
	//
	// image
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s DescribeDataProcessJobAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetDestinationConfig() *string {
	return s.DestinationConfig
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetDestinationId() *string {
	return s.DestinationId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetDestinationPath() *string {
	return s.DestinationPath
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetDestinationType() *string {
	return s.DestinationType
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetModelName() *string {
	return s.ModelName
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetProcessingMode() *string {
	return s.ProcessingMode
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetSourceId() *string {
	return s.SourceId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetSourcePath() *string {
	return s.SourcePath
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeDataProcessJobAttributeResponseBody) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetCreationTime(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetDestinationConfig(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.DestinationConfig = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetDestinationId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.DestinationId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetDestinationPath(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.DestinationPath = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetDestinationType(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.DestinationType = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetJobId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetModelName(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.ModelName = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetProcessingMode(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.ProcessingMode = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetRegionId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetRequestId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetSourceId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.SourceId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetSourcePath(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.SourcePath = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetSourceType(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.SourceType = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetStatus(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetTemplateId(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) SetTemplateType(v string) *DescribeDataProcessJobAttributeResponseBody {
	s.TemplateType = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
