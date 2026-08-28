// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatrolConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdatePatrolConfigResponseBodyData) *UpdatePatrolConfigResponseBody
	GetData() *UpdatePatrolConfigResponseBodyData
	SetErrorCode(v string) *UpdatePatrolConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdatePatrolConfigResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdatePatrolConfigResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdatePatrolConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdatePatrolConfigResponseBody
	GetSuccess() *bool
}

type UpdatePatrolConfigResponseBody struct {
	// The inspection configuration response data.
	Data *UpdatePatrolConfigResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is not empty when success is false. This parameter is empty when success is true.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code, which is always 200. Use the success parameter to determine whether the request was successful.
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

func (s UpdatePatrolConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigResponseBody) GetData() *UpdatePatrolConfigResponseBodyData {
	return s.Data
}

func (s *UpdatePatrolConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdatePatrolConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdatePatrolConfigResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdatePatrolConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePatrolConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePatrolConfigResponseBody) SetData(v *UpdatePatrolConfigResponseBodyData) *UpdatePatrolConfigResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePatrolConfigResponseBody) SetErrorCode(v string) *UpdatePatrolConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePatrolConfigResponseBody) SetErrorMessage(v string) *UpdatePatrolConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePatrolConfigResponseBody) SetHttpCode(v int32) *UpdatePatrolConfigResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdatePatrolConfigResponseBody) SetRequestId(v string) *UpdatePatrolConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePatrolConfigResponseBody) SetSuccess(v bool) *UpdatePatrolConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePatrolConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePatrolConfigResponseBodyData struct {
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
	// Specifies whether to enable the inspection.
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
	ScopeConfig *UpdatePatrolConfigResponseBodyDataScopeConfig `json:"scopeConfig,omitempty" xml:"scopeConfig,omitempty" type:"Struct"`
	// The inspection scope type.
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
	// a14bda1c4a****
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdatePatrolConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigResponseBodyData) GetConfigCreatedAt() *int64 {
	return s.ConfigCreatedAt
}

func (s *UpdatePatrolConfigResponseBodyData) GetConfigUpdatedAt() *int64 {
	return s.ConfigUpdatedAt
}

func (s *UpdatePatrolConfigResponseBodyData) GetCron() *string {
	return s.Cron
}

func (s *UpdatePatrolConfigResponseBodyData) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdatePatrolConfigResponseBodyData) GetNamespace() *string {
	return s.Namespace
}

func (s *UpdatePatrolConfigResponseBodyData) GetNextPatrolAt() *int64 {
	return s.NextPatrolAt
}

func (s *UpdatePatrolConfigResponseBodyData) GetScopeConfig() *UpdatePatrolConfigResponseBodyDataScopeConfig {
	return s.ScopeConfig
}

func (s *UpdatePatrolConfigResponseBodyData) GetScopeType() *string {
	return s.ScopeType
}

func (s *UpdatePatrolConfigResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *UpdatePatrolConfigResponseBodyData) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePatrolConfigResponseBodyData) SetConfigCreatedAt(v int64) *UpdatePatrolConfigResponseBodyData {
	s.ConfigCreatedAt = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetConfigUpdatedAt(v int64) *UpdatePatrolConfigResponseBodyData {
	s.ConfigUpdatedAt = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetCron(v string) *UpdatePatrolConfigResponseBodyData {
	s.Cron = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetEnabled(v bool) *UpdatePatrolConfigResponseBodyData {
	s.Enabled = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetNamespace(v string) *UpdatePatrolConfigResponseBodyData {
	s.Namespace = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetNextPatrolAt(v int64) *UpdatePatrolConfigResponseBodyData {
	s.NextPatrolAt = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetScopeConfig(v *UpdatePatrolConfigResponseBodyDataScopeConfig) *UpdatePatrolConfigResponseBodyData {
	s.ScopeConfig = v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetScopeType(v string) *UpdatePatrolConfigResponseBodyData {
	s.ScopeType = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetTimezone(v string) *UpdatePatrolConfigResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) SetWorkspace(v string) *UpdatePatrolConfigResponseBodyData {
	s.Workspace = &v
	return s
}

func (s *UpdatePatrolConfigResponseBodyData) Validate() error {
	if s.ScopeConfig != nil {
		if err := s.ScopeConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePatrolConfigResponseBodyDataScopeConfig struct {
	// The list of deployment IDs. This parameter is valid only when scopeType is set to DEPLOYMENTS.
	DeploymentIds []*string `json:"deploymentIds,omitempty" xml:"deploymentIds,omitempty" type:"Repeated"`
	// The tag mapping. This parameter is valid only when scopeType is set to TAGS. The key is the tag name, and the value is a list of tag values.
	Tags map[string][]*string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (s UpdatePatrolConfigResponseBodyDataScopeConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatrolConfigResponseBodyDataScopeConfig) GoString() string {
	return s.String()
}

func (s *UpdatePatrolConfigResponseBodyDataScopeConfig) GetDeploymentIds() []*string {
	return s.DeploymentIds
}

func (s *UpdatePatrolConfigResponseBodyDataScopeConfig) GetTags() map[string][]*string {
	return s.Tags
}

func (s *UpdatePatrolConfigResponseBodyDataScopeConfig) SetDeploymentIds(v []*string) *UpdatePatrolConfigResponseBodyDataScopeConfig {
	s.DeploymentIds = v
	return s
}

func (s *UpdatePatrolConfigResponseBodyDataScopeConfig) SetTags(v map[string][]*string) *UpdatePatrolConfigResponseBodyDataScopeConfig {
	s.Tags = v
	return s
}

func (s *UpdatePatrolConfigResponseBodyDataScopeConfig) Validate() error {
	return dara.Validate(s)
}
