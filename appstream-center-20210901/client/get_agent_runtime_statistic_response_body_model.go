// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRuntimeStatisticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetAgentRuntimeStatisticResponseBodyData) *GetAgentRuntimeStatisticResponseBody
	GetData() []*GetAgentRuntimeStatisticResponseBodyData
	SetRequestId(v string) *GetAgentRuntimeStatisticResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetAgentRuntimeStatisticResponseBody
	GetTotalCount() *int32
}

type GetAgentRuntimeStatisticResponseBody struct {
	// The list of runtime statistics grouped by agent platform and agent provider. If AgentProvider is specified but no matching runtimes are found, a record with all statistical values set to 0 is returned.
	Data []*GetAgentRuntimeStatisticResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The request ID, which is used to uniquely identify the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned statistical records, which is the length of the Data list.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetAgentRuntimeStatisticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeStatisticResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeStatisticResponseBody) GetData() []*GetAgentRuntimeStatisticResponseBodyData {
	return s.Data
}

func (s *GetAgentRuntimeStatisticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentRuntimeStatisticResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetAgentRuntimeStatisticResponseBody) SetData(v []*GetAgentRuntimeStatisticResponseBodyData) *GetAgentRuntimeStatisticResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBody) SetRequestId(v string) *GetAgentRuntimeStatisticResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBody) SetTotalCount(v int32) *GetAgentRuntimeStatisticResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentRuntimeStatisticResponseBodyData struct {
	// The display name of the agent, which is determined by both the agent platform and the agent provider. For example, OpenClaw under the ENTERPRISE platform is displayed as OpenClaw, and OpenClaw under the JVS platform is displayed as JVS Claw.
	//
	// example:
	//
	// Agentic Computer
	AgentDisplayName *string `json:"AgentDisplayName,omitempty" xml:"AgentDisplayName,omitempty"`
	// The agent platform, such as ENTERPRISE_AGENTIC_COMPUTER.
	//
	// example:
	//
	// ENTERPRISE_AGENTIC_COMPUTER
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The agent provider, such as AgenticComputer.
	//
	// example:
	//
	// AgenticComputer
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The number of expired runtime resources in the group. A runtime is counted as expired if its resource status is Expired or Ceased, case-insensitive.
	//
	// example:
	//
	// 0
	ExpiredCount *int32 `json:"ExpiredCount,omitempty" xml:"ExpiredCount,omitempty"`
	// The number of runtime resources without channel configurations in the group, which is the total number of runtime resources minus the number of runtime resources with channel configurations.
	//
	// example:
	//
	// 3
	NoChannelConfiguredRuntimeCount *int32 `json:"NoChannelConfiguredRuntimeCount,omitempty" xml:"NoChannelConfiguredRuntimeCount,omitempty"`
	// The number of runtime resources without model configurations in the group. A runtime is considered to have a model configured if a model is configured for the runtime itself or its associated resource group.
	//
	// example:
	//
	// 2
	NoModelConfiguredRuntimeCount *int32 `json:"NoModelConfiguredRuntimeCount,omitempty" xml:"NoModelConfiguredRuntimeCount,omitempty"`
	// The number of risky runtime resources in the group, deduplicated by runtime ID. A runtime is counted as risky if it has any of the following risks: model configuration modified, third-party channel configuration modified, or agent being uninstalled.
	//
	// example:
	//
	// 1
	RiskyRuntimeCount *int32 `json:"RiskyRuntimeCount,omitempty" xml:"RiskyRuntimeCount,omitempty"`
	// The total number of runtime resources in the group, deduplicated by runtime ID.
	//
	// example:
	//
	// 10
	TotalRuntimeCount *int32 `json:"TotalRuntimeCount,omitempty" xml:"TotalRuntimeCount,omitempty"`
}

func (s GetAgentRuntimeStatisticResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRuntimeStatisticResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetAgentDisplayName() *string {
	return s.AgentDisplayName
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetExpiredCount() *int32 {
	return s.ExpiredCount
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetNoChannelConfiguredRuntimeCount() *int32 {
	return s.NoChannelConfiguredRuntimeCount
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetNoModelConfiguredRuntimeCount() *int32 {
	return s.NoModelConfiguredRuntimeCount
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetRiskyRuntimeCount() *int32 {
	return s.RiskyRuntimeCount
}

func (s *GetAgentRuntimeStatisticResponseBodyData) GetTotalRuntimeCount() *int32 {
	return s.TotalRuntimeCount
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetAgentDisplayName(v string) *GetAgentRuntimeStatisticResponseBodyData {
	s.AgentDisplayName = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetAgentPlatform(v string) *GetAgentRuntimeStatisticResponseBodyData {
	s.AgentPlatform = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetAgentProvider(v string) *GetAgentRuntimeStatisticResponseBodyData {
	s.AgentProvider = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetExpiredCount(v int32) *GetAgentRuntimeStatisticResponseBodyData {
	s.ExpiredCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetNoChannelConfiguredRuntimeCount(v int32) *GetAgentRuntimeStatisticResponseBodyData {
	s.NoChannelConfiguredRuntimeCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetNoModelConfiguredRuntimeCount(v int32) *GetAgentRuntimeStatisticResponseBodyData {
	s.NoModelConfiguredRuntimeCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetRiskyRuntimeCount(v int32) *GetAgentRuntimeStatisticResponseBodyData {
	s.RiskyRuntimeCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) SetTotalRuntimeCount(v int32) *GetAgentRuntimeStatisticResponseBodyData {
	s.TotalRuntimeCount = &v
	return s
}

func (s *GetAgentRuntimeStatisticResponseBodyData) Validate() error {
	return dara.Validate(s)
}
