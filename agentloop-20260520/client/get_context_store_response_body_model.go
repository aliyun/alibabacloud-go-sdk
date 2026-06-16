// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentSpace(v string) *GetContextStoreResponseBody
	GetAgentSpace() *string
	SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody
	GetConfig() *GetContextStoreResponseBodyConfig
	SetContextStoreName(v string) *GetContextStoreResponseBody
	GetContextStoreName() *string
	SetContextType(v string) *GetContextStoreResponseBody
	GetContextType() *string
	SetCreateTime(v string) *GetContextStoreResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetContextStoreResponseBody
	GetDescription() *string
	SetRegionId(v string) *GetContextStoreResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetContextStoreResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetContextStoreResponseBody
	GetStatus() *string
	SetUpdateTime(v string) *GetContextStoreResponseBody
	GetUpdateTime() *string
}

type GetContextStoreResponseBody struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string                            `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	Config     *GetContextStoreResponseBodyConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	// example:
	//
	// my-context-store
	ContextStoreName *string `json:"contextStoreName,omitempty" xml:"contextStoreName,omitempty"`
	// example:
	//
	// experience
	ContextType *string `json:"contextType,omitempty" xml:"contextType,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 我的上下文库
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetContextStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBody) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreResponseBody) GetConfig() *GetContextStoreResponseBodyConfig {
	return s.Config
}

func (s *GetContextStoreResponseBody) GetContextStoreName() *string {
	return s.ContextStoreName
}

func (s *GetContextStoreResponseBody) GetContextType() *string {
	return s.ContextType
}

func (s *GetContextStoreResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetContextStoreResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetContextStoreResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetContextStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContextStoreResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetContextStoreResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetContextStoreResponseBody) SetAgentSpace(v string) *GetContextStoreResponseBody {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBody) SetConfig(v *GetContextStoreResponseBodyConfig) *GetContextStoreResponseBody {
	s.Config = v
	return s
}

func (s *GetContextStoreResponseBody) SetContextStoreName(v string) *GetContextStoreResponseBody {
	s.ContextStoreName = &v
	return s
}

func (s *GetContextStoreResponseBody) SetContextType(v string) *GetContextStoreResponseBody {
	s.ContextType = &v
	return s
}

func (s *GetContextStoreResponseBody) SetCreateTime(v string) *GetContextStoreResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) SetDescription(v string) *GetContextStoreResponseBody {
	s.Description = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRegionId(v string) *GetContextStoreResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetRequestId(v string) *GetContextStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContextStoreResponseBody) SetStatus(v string) *GetContextStoreResponseBody {
	s.Status = &v
	return s
}

func (s *GetContextStoreResponseBody) SetUpdateTime(v string) *GetContextStoreResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetContextStoreResponseBody) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfig struct {
	// example:
	//
	// {"userId":"user_id","sessionId":"session_id"}
	MetadataField map[string]*string `json:"metadataField,omitempty" xml:"metadataField,omitempty"`
	// example:
	//
	// 1d
	MiningInterval *string `json:"miningInterval,omitempty" xml:"miningInterval,omitempty"`
	// example:
	//
	// ["order-service","payment-service"]
	ServiceNames []*string                                `json:"serviceNames,omitempty" xml:"serviceNames,omitempty" type:"Repeated"`
	Source       *GetContextStoreResponseBodyConfigSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s GetContextStoreResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfig) GetMetadataField() map[string]*string {
	return s.MetadataField
}

func (s *GetContextStoreResponseBodyConfig) GetMiningInterval() *string {
	return s.MiningInterval
}

func (s *GetContextStoreResponseBodyConfig) GetServiceNames() []*string {
	return s.ServiceNames
}

func (s *GetContextStoreResponseBodyConfig) GetSource() *GetContextStoreResponseBodyConfigSource {
	return s.Source
}

func (s *GetContextStoreResponseBodyConfig) SetMetadataField(v map[string]*string) *GetContextStoreResponseBodyConfig {
	s.MetadataField = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetMiningInterval(v string) *GetContextStoreResponseBodyConfig {
	s.MiningInterval = &v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetServiceNames(v []*string) *GetContextStoreResponseBodyConfig {
	s.ServiceNames = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) SetSource(v *GetContextStoreResponseBodyConfigSource) *GetContextStoreResponseBodyConfig {
	s.Source = v
	return s
}

func (s *GetContextStoreResponseBodyConfig) Validate() error {
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContextStoreResponseBodyConfigSource struct {
	// example:
	//
	// my-agent-space
	AgentSpace *string `json:"agentSpace,omitempty" xml:"agentSpace,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetContextStoreResponseBodyConfigSource) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponseBodyConfigSource) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponseBodyConfigSource) GetAgentSpace() *string {
	return s.AgentSpace
}

func (s *GetContextStoreResponseBodyConfigSource) GetStartTime() *string {
	return s.StartTime
}

func (s *GetContextStoreResponseBodyConfigSource) SetAgentSpace(v string) *GetContextStoreResponseBodyConfigSource {
	s.AgentSpace = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) SetStartTime(v string) *GetContextStoreResponseBodyConfigSource {
	s.StartTime = &v
	return s
}

func (s *GetContextStoreResponseBodyConfigSource) Validate() error {
	return dara.Validate(s)
}
