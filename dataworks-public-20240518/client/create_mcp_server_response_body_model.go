// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServer(v *CreateMcpServerResponseBodyMcpServer) *CreateMcpServerResponseBody
	GetMcpServer() *CreateMcpServerResponseBodyMcpServer
	SetRequestId(v string) *CreateMcpServerResponseBody
	GetRequestId() *string
}

type CreateMcpServerResponseBody struct {
	McpServer *CreateMcpServerResponseBodyMcpServer `json:"McpServer,omitempty" xml:"McpServer,omitempty" type:"Struct"`
	// example:
	//
	// valueA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpServerResponseBody) GetMcpServer() *CreateMcpServerResponseBodyMcpServer {
	return s.McpServer
}

func (s *CreateMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpServerResponseBody) SetMcpServer(v *CreateMcpServerResponseBodyMcpServer) *CreateMcpServerResponseBody {
	s.McpServer = v
	return s
}

func (s *CreateMcpServerResponseBody) SetRequestId(v string) *CreateMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpServerResponseBody) Validate() error {
	if s.McpServer != nil {
		if err := s.McpServer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServerResponseBodyMcpServer struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 1780555634000
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMcpServerResponseBodyMcpServer) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerResponseBodyMcpServer) GoString() string {
	return s.String()
}

func (s *CreateMcpServerResponseBodyMcpServer) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *CreateMcpServerResponseBodyMcpServer) GetName() *string {
	return s.Name
}

func (s *CreateMcpServerResponseBodyMcpServer) SetGmtCreateTime(v string) *CreateMcpServerResponseBodyMcpServer {
	s.GmtCreateTime = &v
	return s
}

func (s *CreateMcpServerResponseBodyMcpServer) SetName(v string) *CreateMcpServerResponseBodyMcpServer {
	s.Name = &v
	return s
}

func (s *CreateMcpServerResponseBodyMcpServer) Validate() error {
	return dara.Validate(s)
}
