// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitWorkspaceSystemMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *InitWorkspaceSystemMcpServerResponseBody
	GetData() interface{}
	SetErrorCode(v string) *InitWorkspaceSystemMcpServerResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *InitWorkspaceSystemMcpServerResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *InitWorkspaceSystemMcpServerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InitWorkspaceSystemMcpServerResponseBody
	GetSuccess() *bool
}

type InitWorkspaceSystemMcpServerResponseBody struct {
	// Indicates whether all system MCP services are initialized successfully. Returns true even when no system MCP services are available.
	//
	// example:
	//
	// true
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The return code. The value success is returned for successful requests. An error code is returned for failed requests.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when a system-level request failure occurs.
	//
	// example:
	//
	// Failed to initialize workspace system MCP servers
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 550***544
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. Valid values:
	//
	// - true: Succeeded.
	//
	// - false: Failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InitWorkspaceSystemMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitWorkspaceSystemMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *InitWorkspaceSystemMcpServerResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InitWorkspaceSystemMcpServerResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InitWorkspaceSystemMcpServerResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *InitWorkspaceSystemMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitWorkspaceSystemMcpServerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InitWorkspaceSystemMcpServerResponseBody) SetData(v interface{}) *InitWorkspaceSystemMcpServerResponseBody {
	s.Data = v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponseBody) SetErrorCode(v string) *InitWorkspaceSystemMcpServerResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponseBody) SetErrorMessage(v string) *InitWorkspaceSystemMcpServerResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponseBody) SetRequestId(v string) *InitWorkspaceSystemMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponseBody) SetSuccess(v bool) *InitWorkspaceSystemMcpServerResponseBody {
	s.Success = &v
	return s
}

func (s *InitWorkspaceSystemMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
