// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWebFetchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentName(v string) *WebFetchRequest
	GetAgentName() *string
	SetOutputFormat(v string) *WebFetchRequest
	GetOutputFormat() *string
	SetRegionId(v string) *WebFetchRequest
	GetRegionId() *string
	SetUrl(v string) *WebFetchRequest
	GetUrl() *string
}

type WebFetchRequest struct {
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The output format. Valid values:
	//
	// - **markdown**: Markdown format.
	//
	// - **html**: HTML format.
	//
	// - **text**: Plain text format.
	//
	// example:
	//
	// markdown
	OutputFormat *string `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The URL of the target web page to crawl.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx/projects/spring-boot
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s WebFetchRequest) String() string {
	return dara.Prettify(s)
}

func (s WebFetchRequest) GoString() string {
	return s.String()
}

func (s *WebFetchRequest) GetAgentName() *string {
	return s.AgentName
}

func (s *WebFetchRequest) GetOutputFormat() *string {
	return s.OutputFormat
}

func (s *WebFetchRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WebFetchRequest) GetUrl() *string {
	return s.Url
}

func (s *WebFetchRequest) SetAgentName(v string) *WebFetchRequest {
	s.AgentName = &v
	return s
}

func (s *WebFetchRequest) SetOutputFormat(v string) *WebFetchRequest {
	s.OutputFormat = &v
	return s
}

func (s *WebFetchRequest) SetRegionId(v string) *WebFetchRequest {
	s.RegionId = &v
	return s
}

func (s *WebFetchRequest) SetUrl(v string) *WebFetchRequest {
	s.Url = &v
	return s
}

func (s *WebFetchRequest) Validate() error {
	return dara.Validate(s)
}
