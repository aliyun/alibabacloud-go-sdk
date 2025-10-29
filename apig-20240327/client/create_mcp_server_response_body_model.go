// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateMcpServerResponseBody
	GetCode() *string
	SetData(v *CreateMcpServerResponseBodyData) *CreateMcpServerResponseBody
	GetData() *CreateMcpServerResponseBodyData
	SetMessage(v string) *CreateMcpServerResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateMcpServerResponseBody
	GetRequestId() *string
}

type CreateMcpServerResponseBody struct {
	// example:
	//
	// Ok
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateMcpServerResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 464F9EA0-1052-51BD-8187-D292AA2D8D24
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMcpServerResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateMcpServerResponseBody) GetData() *CreateMcpServerResponseBodyData {
	return s.Data
}

func (s *CreateMcpServerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMcpServerResponseBody) SetCode(v string) *CreateMcpServerResponseBody {
	s.Code = &v
	return s
}

func (s *CreateMcpServerResponseBody) SetData(v *CreateMcpServerResponseBodyData) *CreateMcpServerResponseBody {
	s.Data = v
	return s
}

func (s *CreateMcpServerResponseBody) SetMessage(v string) *CreateMcpServerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateMcpServerResponseBody) SetRequestId(v string) *CreateMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMcpServerResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateMcpServerResponseBodyData struct {
	// MCP Server ID
	//
	// example:
	//
	// mcp-afegaijoijaoji24a
	McpServerId *string `json:"mcpServerId,omitempty" xml:"mcpServerId,omitempty"`
	// example:
	//
	// test-mcp
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateMcpServerResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateMcpServerResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateMcpServerResponseBodyData) GetMcpServerId() *string {
	return s.McpServerId
}

func (s *CreateMcpServerResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateMcpServerResponseBodyData) SetMcpServerId(v string) *CreateMcpServerResponseBodyData {
	s.McpServerId = &v
	return s
}

func (s *CreateMcpServerResponseBodyData) SetName(v string) *CreateMcpServerResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateMcpServerResponseBodyData) Validate() error {
	return dara.Validate(s)
}
