// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEngineConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEngineConfigs(v []*ListEngineConfigsResponseBodyEngineConfigs) *ListEngineConfigsResponseBody
	GetEngineConfigs() []*ListEngineConfigsResponseBodyEngineConfigs
	SetRequestId(v string) *ListEngineConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListEngineConfigsResponseBody
	GetTotalCount() *int64
}

type ListEngineConfigsResponseBody struct {
	EngineConfigs []*ListEngineConfigsResponseBodyEngineConfigs `json:"EngineConfigs,omitempty" xml:"EngineConfigs,omitempty" type:"Repeated"`
	// example:
	//
	// 74D958EF-3598-56FA-8296-FF1575CE43DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEngineConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEngineConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponseBody) GetEngineConfigs() []*ListEngineConfigsResponseBodyEngineConfigs {
	return s.EngineConfigs
}

func (s *ListEngineConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEngineConfigsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListEngineConfigsResponseBody) SetEngineConfigs(v []*ListEngineConfigsResponseBodyEngineConfigs) *ListEngineConfigsResponseBody {
	s.EngineConfigs = v
	return s
}

func (s *ListEngineConfigsResponseBody) SetRequestId(v string) *ListEngineConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEngineConfigsResponseBody) SetTotalCount(v int64) *ListEngineConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListEngineConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEngineConfigsResponseBodyEngineConfigs struct {
	// example:
	//
	// {}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2
	EngineConfigId *string `json:"EngineConfigId,omitempty" xml:"EngineConfigId,omitempty"`
	// example:
	//
	// Pre
	Environment *string `json:"Environment,omitempty" xml:"Environment,omitempty"`
	// example:
	//
	// 2023-08-07T01:43:42Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2023-08-27T12:00:00Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2023-08-29 12:00:00
	GmtReleasedTime *string `json:"GmtReleasedTime,omitempty" xml:"GmtReleasedTime,omitempty"`
	// example:
	//
	// engine_config_v1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Released
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 20230509161300
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListEngineConfigsResponseBodyEngineConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListEngineConfigsResponseBodyEngineConfigs) GoString() string {
	return s.String()
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetDescription() *string {
	return s.Description
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetEngineConfigId() *string {
	return s.EngineConfigId
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetEnvironment() *string {
	return s.Environment
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetGmtReleasedTime() *string {
	return s.GmtReleasedTime
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetName() *string {
	return s.Name
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetStatus() *string {
	return s.Status
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) GetVersion() *string {
	return s.Version
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetConfigValue(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.ConfigValue = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetDescription(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Description = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetEngineConfigId(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.EngineConfigId = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetEnvironment(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Environment = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtCreateTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtCreateTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtModifiedTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtModifiedTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetGmtReleasedTime(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.GmtReleasedTime = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetName(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Name = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetStatus(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Status = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) SetVersion(v string) *ListEngineConfigsResponseBodyEngineConfigs {
	s.Version = &v
	return s
}

func (s *ListEngineConfigsResponseBodyEngineConfigs) Validate() error {
	return dara.Validate(s)
}
