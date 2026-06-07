// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServer(v *GetMcpServerResponseBodyMcpServer) *GetMcpServerResponseBody
	GetMcpServer() *GetMcpServerResponseBodyMcpServer
	SetRequestId(v string) *GetMcpServerResponseBody
	GetRequestId() *string
}

type GetMcpServerResponseBody struct {
	McpServer *GetMcpServerResponseBodyMcpServer `json:"McpServer,omitempty" xml:"McpServer,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBody) GetMcpServer() *GetMcpServerResponseBodyMcpServer {
	return s.McpServer
}

func (s *GetMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpServerResponseBody) SetMcpServer(v *GetMcpServerResponseBodyMcpServer) *GetMcpServerResponseBody {
	s.McpServer = v
	return s
}

func (s *GetMcpServerResponseBody) SetRequestId(v string) *GetMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpServerResponseBody) Validate() error {
	if s.McpServer != nil {
		if err := s.McpServer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServerResponseBodyMcpServer struct {
	Config *GetMcpServerResponseBodyMcpServerConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// example:
	//
	// 123456
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 123456
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// my-mcp-server
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetMcpServerResponseBodyMcpServer) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyMcpServer) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyMcpServer) GetConfig() *GetMcpServerResponseBodyMcpServerConfig {
	return s.Config
}

func (s *GetMcpServerResponseBodyMcpServer) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetMcpServerResponseBodyMcpServer) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetMcpServerResponseBodyMcpServer) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *GetMcpServerResponseBodyMcpServer) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetMcpServerResponseBodyMcpServer) GetName() *string {
	return s.Name
}

func (s *GetMcpServerResponseBodyMcpServer) GetVisibility() *string {
	return s.Visibility
}

func (s *GetMcpServerResponseBodyMcpServer) SetConfig(v *GetMcpServerResponseBodyMcpServerConfig) *GetMcpServerResponseBodyMcpServer {
	s.Config = v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetCreatorId(v string) *GetMcpServerResponseBodyMcpServer {
	s.CreatorId = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetGmtCreateTime(v string) *GetMcpServerResponseBodyMcpServer {
	s.GmtCreateTime = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetGmtModifiedTime(v string) *GetMcpServerResponseBodyMcpServer {
	s.GmtModifiedTime = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetModifierId(v string) *GetMcpServerResponseBodyMcpServer {
	s.ModifierId = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetName(v string) *GetMcpServerResponseBodyMcpServer {
	s.Name = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) SetVisibility(v string) *GetMcpServerResponseBodyMcpServer {
	s.Visibility = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServer) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpServerResponseBodyMcpServerConfig struct {
	CustomHeaders map[string]interface{} `json:"CustomHeaders,omitempty" xml:"CustomHeaders,omitempty"`
	// example:
	//
	// SSE
	Transport *string `json:"Transport,omitempty" xml:"Transport,omitempty"`
	// example:
	//
	// https://example.com/mcp/sse
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetMcpServerResponseBodyMcpServerConfig) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServerResponseBodyMcpServerConfig) GoString() string {
	return s.String()
}

func (s *GetMcpServerResponseBodyMcpServerConfig) GetCustomHeaders() map[string]interface{} {
	return s.CustomHeaders
}

func (s *GetMcpServerResponseBodyMcpServerConfig) GetTransport() *string {
	return s.Transport
}

func (s *GetMcpServerResponseBodyMcpServerConfig) GetUrl() *string {
	return s.Url
}

func (s *GetMcpServerResponseBodyMcpServerConfig) SetCustomHeaders(v map[string]interface{}) *GetMcpServerResponseBodyMcpServerConfig {
	s.CustomHeaders = v
	return s
}

func (s *GetMcpServerResponseBodyMcpServerConfig) SetTransport(v string) *GetMcpServerResponseBodyMcpServerConfig {
	s.Transport = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServerConfig) SetUrl(v string) *GetMcpServerResponseBodyMcpServerConfig {
	s.Url = &v
	return s
}

func (s *GetMcpServerResponseBodyMcpServerConfig) Validate() error {
	return dara.Validate(s)
}
