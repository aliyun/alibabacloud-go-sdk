// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEngineConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *GetEngineConfigResponseBody
	GetConfigValue() *string
	SetDescription(v string) *GetEngineConfigResponseBody
	GetDescription() *string
	SetEnvironment(v string) *GetEngineConfigResponseBody
	GetEnvironment() *string
	SetGmtCreateTime(v string) *GetEngineConfigResponseBody
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *GetEngineConfigResponseBody
	GetGmtModifiedTime() *string
	SetGmtReleasedTime(v string) *GetEngineConfigResponseBody
	GetGmtReleasedTime() *string
	SetName(v string) *GetEngineConfigResponseBody
	GetName() *string
	SetRequestId(v string) *GetEngineConfigResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetEngineConfigResponseBody
	GetStatus() *string
}

type GetEngineConfigResponseBody struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2024-01-03T02:28:00.000Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2024-08-27T12:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2024-01-03 02:28:00
	GmtReleasedTime *string `json:"GmtReleasedTime,omitempty" xml:"GmtReleasedTime,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 59CE7EC6-F268-5D71-9215-32922CC50D72
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetEngineConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEngineConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetEngineConfigResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *GetEngineConfigResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEngineConfigResponseBody) GetEnvironment() *string {
	return s.Environment
}

func (s *GetEngineConfigResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetEngineConfigResponseBody) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetEngineConfigResponseBody) GetGmtReleasedTime() *string {
	return s.GmtReleasedTime
}

func (s *GetEngineConfigResponseBody) GetName() *string {
	return s.Name
}

func (s *GetEngineConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEngineConfigResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEngineConfigResponseBody) SetConfigValue(v string) *GetEngineConfigResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetDescription(v string) *GetEngineConfigResponseBody {
	s.Description = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetEnvironment(v string) *GetEngineConfigResponseBody {
	s.Environment = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtCreateTime(v string) *GetEngineConfigResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtModifiedTime(v string) *GetEngineConfigResponseBody {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetGmtReleasedTime(v string) *GetEngineConfigResponseBody {
	s.GmtReleasedTime = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetName(v string) *GetEngineConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetRequestId(v string) *GetEngineConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEngineConfigResponseBody) SetStatus(v string) *GetEngineConfigResponseBody {
	s.Status = &v
	return s
}

func (s *GetEngineConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
