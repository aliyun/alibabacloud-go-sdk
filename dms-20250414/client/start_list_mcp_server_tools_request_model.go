// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartListMcpServerToolsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDMSUnit(v string) *StartListMcpServerToolsRequest
	GetDMSUnit() *string
	SetLanguage(v string) *StartListMcpServerToolsRequest
	GetLanguage() *string
	SetMcpServerUuid(v string) *StartListMcpServerToolsRequest
	GetMcpServerUuid() *string
}

type StartListMcpServerToolsRequest struct {
	// The identifier of the Data Management unit that runs the Data Agent resources.
	//
	// example:
	//
	// cn-hangzhou
	DMSUnit *string `json:"DMSUnit,omitempty" xml:"DMSUnit,omitempty"`
	// The language used for the detection session.
	//
	// example:
	//
	// CHINESE
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The ID of the MCP Server for which to detect connectivity and query the tool list. Only the service creator can trigger the detection.
	//
	// example:
	//
	// 44lg***z65
	McpServerUuid *string `json:"McpServerUuid,omitempty" xml:"McpServerUuid,omitempty"`
}

func (s StartListMcpServerToolsRequest) String() string {
	return dara.Prettify(s)
}

func (s StartListMcpServerToolsRequest) GoString() string {
	return s.String()
}

func (s *StartListMcpServerToolsRequest) GetDMSUnit() *string {
	return s.DMSUnit
}

func (s *StartListMcpServerToolsRequest) GetLanguage() *string {
	return s.Language
}

func (s *StartListMcpServerToolsRequest) GetMcpServerUuid() *string {
	return s.McpServerUuid
}

func (s *StartListMcpServerToolsRequest) SetDMSUnit(v string) *StartListMcpServerToolsRequest {
	s.DMSUnit = &v
	return s
}

func (s *StartListMcpServerToolsRequest) SetLanguage(v string) *StartListMcpServerToolsRequest {
	s.Language = &v
	return s
}

func (s *StartListMcpServerToolsRequest) SetMcpServerUuid(v string) *StartListMcpServerToolsRequest {
	s.McpServerUuid = &v
	return s
}

func (s *StartListMcpServerToolsRequest) Validate() error {
	return dara.Validate(s)
}
