// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListMcpServerToolsResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *GetListMcpServerToolsResultRequest
	GetDMSUnit() *string
	SetMcpServerUuid(v string) *GetListMcpServerToolsResultRequest
	GetMcpServerUuid() *string
	SetSessionId(v string) *GetListMcpServerToolsResultRequest
	GetSessionId() *string
}

type GetListMcpServerToolsResultRequest struct {
	// The DMS unit identifier. This is typically the same as the DMSUnit used in the request to start tool detection.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The MCP server ID used when starting the detection. It must match the detection record corresponding to SessionId.
	//
	// example:
	//
	// 44lg***z65
	McpServerUuid *string `json:"McpServerUuid,omitempty" xml:"McpServerUuid,omitempty"`
	// The temporary session ID returned by StartListMcpServerTools, which is used to locate this connectivity detection task.
	//
	// example:
	//
	// 1vwe***6wr
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s GetListMcpServerToolsResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListMcpServerToolsResultRequest) GoString() string {
	return s.String()
}

func (s *GetListMcpServerToolsResultRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *GetListMcpServerToolsResultRequest) GetMcpServerUuid() *string {
	return s.McpServerUuid
}

func (s *GetListMcpServerToolsResultRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *GetListMcpServerToolsResultRequest) SetDMSUnit(v string) *GetListMcpServerToolsResultRequest {
	s.DMSUnit = &v
	return s
}

func (s *GetListMcpServerToolsResultRequest) SetMcpServerUuid(v string) *GetListMcpServerToolsResultRequest {
	s.McpServerUuid = &v
	return s
}

func (s *GetListMcpServerToolsResultRequest) SetSessionId(v string) *GetListMcpServerToolsResultRequest {
	s.SessionId = &v
	return s
}

func (s *GetListMcpServerToolsResultRequest) Validate() error {
	return dara.Validate(s)
}
