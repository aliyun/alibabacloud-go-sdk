// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetDataAgentMcpResponseBody
	GetData() interface{}
	SetErrorCode(v string) *GetDataAgentMcpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentMcpResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAgentMcpResponseBody
	GetSuccess() *bool
}

type GetDataAgentMcpResponseBody struct {
	// The MCP Server details.
	//
	// example:
	//
	// {"uuid":"	44lg***z65","name":"mcp","workspaceUuid":"	atvx***xmz","region":"cn-hangzhou","netType":"public","transportType":"sse","state":"ready","enable":true}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The return code. The value success is returned if the request was successful. An error code is returned if the request failed.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that describes the reason for the failure.
	//
	// example:
	//
	// Resource Not exist,Mcp Server you provide is not exist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 550e***000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAgentMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentMcpResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentMcpResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetDataAgentMcpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentMcpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAgentMcpResponseBody) SetData(v interface{}) *GetDataAgentMcpResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentMcpResponseBody) SetErrorCode(v string) *GetDataAgentMcpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentMcpResponseBody) SetErrorMessage(v string) *GetDataAgentMcpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentMcpResponseBody) SetRequestId(v string) *GetDataAgentMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentMcpResponseBody) SetSuccess(v bool) *GetDataAgentMcpResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
