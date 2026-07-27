// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMOUsageDetailExportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateMOUsageDetailExportRequest
	GetApiKey() *string
	SetEndTime(v string) *CreateMOUsageDetailExportRequest
	GetEndTime() *string
	SetInstanceId(v string) *CreateMOUsageDetailExportRequest
	GetInstanceId() *string
	SetModel(v string) *CreateMOUsageDetailExportRequest
	GetModel() *string
	SetStartTime(v string) *CreateMOUsageDetailExportRequest
	GetStartTime() *string
	SetUsageType(v string) *CreateMOUsageDetailExportRequest
	GetUsageType() *string
}

type CreateMOUsageDetailExportRequest struct {
	// example:
	//
	// sk-rds-*****
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// 结束时间，UTC 0 时区 ISO8601 字符串，格式 yyyy-MM-ddTHH:mm:ssZ；与 StartTime 跨度不超过 30 天
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-03-10T02:02:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rds_copilot***_public_cn-*********6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// qwen-flash
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// 起始时间，UTC 0 时区 ISO8601 字符串，格式 yyyy-MM-ddTHH:mm:ssZ
	//
	// This parameter is required.
	//
	// example:
	//
	// 2026-03-05T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// text
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s CreateMOUsageDetailExportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMOUsageDetailExportRequest) GoString() string {
	return s.String()
}

func (s *CreateMOUsageDetailExportRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateMOUsageDetailExportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateMOUsageDetailExportRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateMOUsageDetailExportRequest) GetModel() *string {
	return s.Model
}

func (s *CreateMOUsageDetailExportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateMOUsageDetailExportRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *CreateMOUsageDetailExportRequest) SetApiKey(v string) *CreateMOUsageDetailExportRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) SetEndTime(v string) *CreateMOUsageDetailExportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) SetInstanceId(v string) *CreateMOUsageDetailExportRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) SetModel(v string) *CreateMOUsageDetailExportRequest {
	s.Model = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) SetStartTime(v string) *CreateMOUsageDetailExportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) SetUsageType(v string) *CreateMOUsageDetailExportRequest {
	s.UsageType = &v
	return s
}

func (s *CreateMOUsageDetailExportRequest) Validate() error {
	return dara.Validate(s)
}
