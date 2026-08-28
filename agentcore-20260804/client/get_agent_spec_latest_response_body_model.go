// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecLatestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgentSpecLatestResponseBodyData) *GetAgentSpecLatestResponseBody
	GetData() *GetAgentSpecLatestResponseBodyData
	SetRequestId(v string) *GetAgentSpecLatestResponseBody
	GetRequestId() *string
}

type GetAgentSpecLatestResponseBody struct {
	// The returned data.
	Data *GetAgentSpecLatestResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAgentSpecLatestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestResponseBody) GetData() *GetAgentSpecLatestResponseBodyData {
	return s.Data
}

func (s *GetAgentSpecLatestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSpecLatestResponseBody) SetData(v *GetAgentSpecLatestResponseBodyData) *GetAgentSpecLatestResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentSpecLatestResponseBody) SetRequestId(v string) *GetAgentSpecLatestResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSpecLatestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSpecLatestResponseBodyData struct {
	// The business tags.
	//
	// example:
	//
	// Sample property value
	BizTags *string `json:"bizTags,omitempty" xml:"bizTags,omitempty"`
	// The content.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The download count.
	//
	// example:
	//
	// 10
	DownloadCount *int64 `json:"downloadCount,omitempty" xml:"downloadCount,omitempty"`
	// Indicates whether the AgentSpec is enabled.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The list of MCP server references.
	McpServers []*GetAgentSpecLatestResponseBodyDataMcpServers `json:"mcpServers,omitempty" xml:"mcpServers,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The resource file mapping.
	Resource map[string]*DataResourceValue `json:"resource,omitempty" xml:"resource,omitempty"`
	// The visibility scope.
	//
	// example:
	//
	// PRIVATE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The list of skill references.
	Skills []*GetAgentSpecLatestResponseBodyDataSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The update time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetAgentSpecLatestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestResponseBodyData) GetBizTags() *string {
	return s.BizTags
}

func (s *GetAgentSpecLatestResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetAgentSpecLatestResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentSpecLatestResponseBodyData) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetAgentSpecLatestResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetAgentSpecLatestResponseBodyData) GetMcpServers() []*GetAgentSpecLatestResponseBodyDataMcpServers {
	return s.McpServers
}

func (s *GetAgentSpecLatestResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAgentSpecLatestResponseBodyData) GetResource() map[string]*DataResourceValue {
	return s.Resource
}

func (s *GetAgentSpecLatestResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *GetAgentSpecLatestResponseBodyData) GetSkills() []*GetAgentSpecLatestResponseBodyDataSkills {
	return s.Skills
}

func (s *GetAgentSpecLatestResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAgentSpecLatestResponseBodyData) SetBizTags(v string) *GetAgentSpecLatestResponseBodyData {
	s.BizTags = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetContent(v string) *GetAgentSpecLatestResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetDescription(v string) *GetAgentSpecLatestResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetDownloadCount(v int64) *GetAgentSpecLatestResponseBodyData {
	s.DownloadCount = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetEnable(v bool) *GetAgentSpecLatestResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetMcpServers(v []*GetAgentSpecLatestResponseBodyDataMcpServers) *GetAgentSpecLatestResponseBodyData {
	s.McpServers = v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetName(v string) *GetAgentSpecLatestResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetResource(v map[string]*DataResourceValue) *GetAgentSpecLatestResponseBodyData {
	s.Resource = v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetScope(v string) *GetAgentSpecLatestResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetSkills(v []*GetAgentSpecLatestResponseBodyDataSkills) *GetAgentSpecLatestResponseBodyData {
	s.Skills = v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) SetUpdateTime(v int64) *GetAgentSpecLatestResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyData) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentSpecLatestResponseBodyDataMcpServers struct {
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAgentSpecLatestResponseBodyDataMcpServers) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestResponseBodyDataMcpServers) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestResponseBodyDataMcpServers) GetName() *string {
	return s.Name
}

func (s *GetAgentSpecLatestResponseBodyDataMcpServers) SetName(v string) *GetAgentSpecLatestResponseBodyDataMcpServers {
	s.Name = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyDataMcpServers) Validate() error {
	return dara.Validate(s)
}

type GetAgentSpecLatestResponseBodyDataSkills struct {
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GetAgentSpecLatestResponseBodyDataSkills) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecLatestResponseBodyDataSkills) GoString() string {
	return s.String()
}

func (s *GetAgentSpecLatestResponseBodyDataSkills) GetName() *string {
	return s.Name
}

func (s *GetAgentSpecLatestResponseBodyDataSkills) SetName(v string) *GetAgentSpecLatestResponseBodyDataSkills {
	s.Name = &v
	return s
}

func (s *GetAgentSpecLatestResponseBodyDataSkills) Validate() error {
	return dara.Validate(s)
}
