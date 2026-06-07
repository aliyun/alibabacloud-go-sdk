// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMcpServer(v *UpdateMcpServerResponseBodyMcpServer) *UpdateMcpServerResponseBody
	GetMcpServer() *UpdateMcpServerResponseBodyMcpServer
	SetRequestId(v string) *UpdateMcpServerResponseBody
	GetRequestId() *string
}

type UpdateMcpServerResponseBody struct {
	McpServer *UpdateMcpServerResponseBodyMcpServer `json:"McpServer,omitempty" xml:"McpServer,omitempty" type:"Struct"`
	// example:
	//
	// 824F80BA-1778-5D8A-BAFF-668A4D9C4CC7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerResponseBody) GetMcpServer() *UpdateMcpServerResponseBodyMcpServer {
	return s.McpServer
}

func (s *UpdateMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMcpServerResponseBody) SetMcpServer(v *UpdateMcpServerResponseBodyMcpServer) *UpdateMcpServerResponseBody {
	s.McpServer = v
	return s
}

func (s *UpdateMcpServerResponseBody) SetRequestId(v string) *UpdateMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMcpServerResponseBody) Validate() error {
	if s.McpServer != nil {
		if err := s.McpServer.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMcpServerResponseBodyMcpServer struct {
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
	// my-mcp-server
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateMcpServerResponseBodyMcpServer) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcpServerResponseBodyMcpServer) GoString() string {
	return s.String()
}

func (s *UpdateMcpServerResponseBodyMcpServer) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *UpdateMcpServerResponseBodyMcpServer) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *UpdateMcpServerResponseBodyMcpServer) GetName() *string {
	return s.Name
}

func (s *UpdateMcpServerResponseBodyMcpServer) SetGmtCreateTime(v string) *UpdateMcpServerResponseBodyMcpServer {
	s.GmtCreateTime = &v
	return s
}

func (s *UpdateMcpServerResponseBodyMcpServer) SetGmtModifiedTime(v string) *UpdateMcpServerResponseBodyMcpServer {
	s.GmtModifiedTime = &v
	return s
}

func (s *UpdateMcpServerResponseBodyMcpServer) SetName(v string) *UpdateMcpServerResponseBodyMcpServer {
	s.Name = &v
	return s
}

func (s *UpdateMcpServerResponseBodyMcpServer) Validate() error {
	return dara.Validate(s)
}
