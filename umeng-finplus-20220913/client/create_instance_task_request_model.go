// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *CreateInstanceTaskRequest
	GetAppId() *int64
	SetCalbackUrl(v string) *CreateInstanceTaskRequest
	GetCalbackUrl() *string
	SetClientId(v int64) *CreateInstanceTaskRequest
	GetClientId() *int64
	SetDatasetIds(v string) *CreateInstanceTaskRequest
	GetDatasetIds() *string
	SetMonitorType(v string) *CreateInstanceTaskRequest
	GetMonitorType() *string
	SetName(v string) *CreateInstanceTaskRequest
	GetName() *string
	SetOutputConfig(v string) *CreateInstanceTaskRequest
	GetOutputConfig() *string
	SetRequestId(v string) *CreateInstanceTaskRequest
	GetRequestId() *string
	SetScoreStrategyConfig(v string) *CreateInstanceTaskRequest
	GetScoreStrategyConfig() *string
}

type CreateInstanceTaskRequest struct {
	// This parameter is required.
	AppId      *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CalbackUrl *string `json:"CalbackUrl,omitempty" xml:"CalbackUrl,omitempty"`
	ClientId   *int64  `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// This parameter is required.
	DatasetIds  *string `json:"DatasetIds,omitempty" xml:"DatasetIds,omitempty"`
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// This parameter is required.
	Name                *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OutputConfig        *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScoreStrategyConfig *string `json:"ScoreStrategyConfig,omitempty" xml:"ScoreStrategyConfig,omitempty"`
}

func (s CreateInstanceTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceTaskRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *CreateInstanceTaskRequest) GetCalbackUrl() *string {
	return s.CalbackUrl
}

func (s *CreateInstanceTaskRequest) GetClientId() *int64 {
	return s.ClientId
}

func (s *CreateInstanceTaskRequest) GetDatasetIds() *string {
	return s.DatasetIds
}

func (s *CreateInstanceTaskRequest) GetMonitorType() *string {
	return s.MonitorType
}

func (s *CreateInstanceTaskRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceTaskRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *CreateInstanceTaskRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceTaskRequest) GetScoreStrategyConfig() *string {
	return s.ScoreStrategyConfig
}

func (s *CreateInstanceTaskRequest) SetAppId(v int64) *CreateInstanceTaskRequest {
	s.AppId = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetCalbackUrl(v string) *CreateInstanceTaskRequest {
	s.CalbackUrl = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetClientId(v int64) *CreateInstanceTaskRequest {
	s.ClientId = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetDatasetIds(v string) *CreateInstanceTaskRequest {
	s.DatasetIds = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetMonitorType(v string) *CreateInstanceTaskRequest {
	s.MonitorType = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetName(v string) *CreateInstanceTaskRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetOutputConfig(v string) *CreateInstanceTaskRequest {
	s.OutputConfig = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetRequestId(v string) *CreateInstanceTaskRequest {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceTaskRequest) SetScoreStrategyConfig(v string) *CreateInstanceTaskRequest {
	s.ScoreStrategyConfig = &v
	return s
}

func (s *CreateInstanceTaskRequest) Validate() error {
	return dara.Validate(s)
}
