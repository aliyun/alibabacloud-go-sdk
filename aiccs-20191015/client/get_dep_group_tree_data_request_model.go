// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDepGroupTreeDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *GetDepGroupTreeDataRequest
	GetAgentId() *int64
	SetInstanceId(v string) *GetDepGroupTreeDataRequest
	GetInstanceId() *string
}

type GetDepGroupTreeDataRequest struct {
	// The agent ID.
	//
	// You can invoke the [GetAgent](https://help.aliyun.com/zh/aiccs/developer-reference/api-aiccs-2019-10-15-getagent) API and view the **AgentId*	- parameter in the response to obtain the agent ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The Artificial Intelligence Cloud Call Service (AICCS) instance ID.
	//
	// You can obtain it from **Instance Management*	- in the left-side navigation pane of the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview).
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDepGroupTreeDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDepGroupTreeDataRequest) GoString() string {
	return s.String()
}

func (s *GetDepGroupTreeDataRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *GetDepGroupTreeDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDepGroupTreeDataRequest) SetAgentId(v int64) *GetDepGroupTreeDataRequest {
	s.AgentId = &v
	return s
}

func (s *GetDepGroupTreeDataRequest) SetInstanceId(v string) *GetDepGroupTreeDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDepGroupTreeDataRequest) Validate() error {
	return dara.Validate(s)
}
