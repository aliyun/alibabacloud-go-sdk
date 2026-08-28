// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatrolConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetPatrolConfigResponseBodyData) *GetPatrolConfigResponseBody
	GetData() *GetPatrolConfigResponseBodyData
	SetErrorCode(v string) *GetPatrolConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetPatrolConfigResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetPatrolConfigResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetPatrolConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPatrolConfigResponseBody
	GetSuccess() *bool
}

type GetPatrolConfigResponseBody struct {
	// The inspection configuration response data.
	Data *GetPatrolConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. This field is not empty when success is false. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This field is not empty when success is false. This field is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. The value is always 200. Use the success field to determine whether the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetPatrolConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigResponseBody) GetData() *GetPatrolConfigResponseBodyData {
	return s.Data
}

func (s *GetPatrolConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetPatrolConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetPatrolConfigResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetPatrolConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPatrolConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPatrolConfigResponseBody) SetData(v *GetPatrolConfigResponseBodyData) *GetPatrolConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetPatrolConfigResponseBody) SetErrorCode(v string) *GetPatrolConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetPatrolConfigResponseBody) SetErrorMessage(v string) *GetPatrolConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetPatrolConfigResponseBody) SetHttpCode(v int32) *GetPatrolConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetPatrolConfigResponseBody) SetRequestId(v string) *GetPatrolConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPatrolConfigResponseBody) SetSuccess(v bool) *GetPatrolConfigResponseBody {
	s.Success = &v
	return s
}

func (s *GetPatrolConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPatrolConfigResponseBodyData struct {
	// The configuration creation time, in milliseconds (UNIX timestamp).
	//
	// example:
	//
	// 1756200000000
	ConfigCreatedAt *int64 `json:"configCreatedAt,omitempty" xml:"configCreatedAt,omitempty"`
	// The configuration update time, in milliseconds (UNIX timestamp).
	//
	// example:
	//
	// 1756250000000
	ConfigUpdatedAt *int64 `json:"configUpdatedAt,omitempty" xml:"configUpdatedAt,omitempty"`
	// The cron expression that defines the inspection scheduling time.
	//
	// example:
	//
	// 0 2 	- 	- *
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// Indicates whether inspection is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The namespace.
	//
	// example:
	//
	// default-namespace
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The next inspection time, in milliseconds (UNIX timestamp).
	//
	// example:
	//
	// 1756260000000
	NextPatrolAt *int64 `json:"nextPatrolAt,omitempty" xml:"nextPatrolAt,omitempty"`
	// The inspection scope configuration.
	ScopeConfig *GetPatrolConfigResponseBodyDataScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type. Valid values:
	//
	// - ALL: inspects all deployments.
	//
	// - TAGS: filters deployments by tag.
	//
	// - DEPLOYMENTS: inspects specified deployments.
	//
	// example:
	//
	// ALL
	ScopeType *string `json:"scopeType,omitempty" xml:"scopeType,omitempty"`
	// The time zone.
	//
	// example:
	//
	// Asia/Shanghai
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// edcef******b4f
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetPatrolConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigResponseBodyData) GetConfigCreatedAt() *int64 {
	return s.ConfigCreatedAt
}

func (s *GetPatrolConfigResponseBodyData) GetConfigUpdatedAt() *int64 {
	return s.ConfigUpdatedAt
}

func (s *GetPatrolConfigResponseBodyData) GetCron() *string {
	return s.Cron
}

func (s *GetPatrolConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetPatrolConfigResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *GetPatrolConfigResponseBodyData) GetNextPatrolAt() *int64 {
	return s.NextPatrolAt
}

func (s *GetPatrolConfigResponseBodyData) GetScopeConfig() *GetPatrolConfigResponseBodyDataScopeConfig {
	return s.ScopeConfig
}

func (s *GetPatrolConfigResponseBodyData) GetScopeType() *string {
	return s.ScopeType
}

func (s *GetPatrolConfigResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetPatrolConfigResponseBodyData) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetPatrolConfigResponseBodyData) SetConfigCreatedAt(v int64) *GetPatrolConfigResponseBodyData {
	s.ConfigCreatedAt = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetConfigUpdatedAt(v int64) *GetPatrolConfigResponseBodyData {
	s.ConfigUpdatedAt = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetCron(v string) *GetPatrolConfigResponseBodyData {
	s.Cron = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetEnabled(v bool) *GetPatrolConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetNamespace(v string) *GetPatrolConfigResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetNextPatrolAt(v int64) *GetPatrolConfigResponseBodyData {
	s.NextPatrolAt = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetScopeConfig(v *GetPatrolConfigResponseBodyDataScopeConfig) *GetPatrolConfigResponseBodyData {
	s.ScopeConfig = v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetScopeType(v string) *GetPatrolConfigResponseBodyData {
	s.ScopeType = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetTimezone(v string) *GetPatrolConfigResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) SetWorkspace(v string) *GetPatrolConfigResponseBodyData {
	s.Workspace = &v
	return s
}

func (s *GetPatrolConfigResponseBodyData) Validate() error {
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPatrolConfigResponseBodyDataScopeConfig struct {
	// The list of deployment IDs. This field is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This field is valid only when scopeType is set to TAGS. The key is the tag name, and the value is the list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s GetPatrolConfigResponseBodyDataScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s GetPatrolConfigResponseBodyDataScopeConfig) GoString() string {
	return s.String()
}

func (s *GetPatrolConfigResponseBodyDataScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *GetPatrolConfigResponseBodyDataScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *GetPatrolConfigResponseBodyDataScopeConfig) SetDeploymentIds(v []*string) *GetPatrolConfigResponseBodyDataScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *GetPatrolConfigResponseBodyDataScopeConfig) SetTags(v map[string][]*string) *GetPatrolConfigResponseBodyDataScopeConfig {
	s.Tags = v
	return s
}

func (s *GetPatrolConfigResponseBodyDataScopeConfig) Validate() error {
	return dara.Validate(s)
}
