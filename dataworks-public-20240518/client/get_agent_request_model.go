// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentVersion(v int32) *GetAgentRequest
	GetAgentVersion() *int32
	SetName(v string) *GetAgentRequest
	GetName() *string
}

type GetAgentRequest struct {
	// 省略时读取当前定义；指定已发布版本时读取对应快照及固定的官方 Skill、套件制品版本；指定当前未发布版本时仍返回草稿，不返回固定制品版本。不是 POP 系统参数 Version。
	//
	// example:
	//
	// 2
	AgentVersion *int32 `json:"AgentVersion,omitempty" xml:"AgentVersion,omitempty"`
	// The name of the skill to retrieve.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-skill
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgentRequest) GoString() string {
	return s.String()
}

func (s *GetAgentRequest) GetAgentVersion() *int32 {
	return s.AgentVersion
}

func (s *GetAgentRequest) GetName() *string {
	return s.Name
}

func (s *GetAgentRequest) SetAgentVersion(v int32) *GetAgentRequest {
	s.AgentVersion = &v
	return s
}

func (s *GetAgentRequest) SetName(v string) *GetAgentRequest {
	s.Name = &v
	return s
}

func (s *GetAgentRequest) Validate() error {
	return dara.Validate(s)
}
