// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteDataAgentMcpResponseBody
	GetData() interface{}
	SetErrorCode(v string) *DeleteDataAgentMcpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentMcpResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentMcpResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentMcpResponseBody struct {
	// Indicates whether the deletion status is updated. A value of true indicates that at least one matching record has been logically deleted.
	//
	// example:
	//
	// true
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The return code. The value success is returned for a successful request. An error code is returned for a failed request.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when a system-level request failure occurs.
	//
	// example:
	//
	// Ready MCP Servers not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate this call.
	//
	// example:
	//
	// 550e84***44
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentMcpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentMcpResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteDataAgentMcpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentMcpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentMcpResponseBody) SetData(v interface{}) *DeleteDataAgentMcpResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataAgentMcpResponseBody) SetErrorCode(v string) *DeleteDataAgentMcpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentMcpResponseBody) SetErrorMessage(v string) *DeleteDataAgentMcpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentMcpResponseBody) SetRequestId(v string) *DeleteDataAgentMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentMcpResponseBody) SetSuccess(v bool) *DeleteDataAgentMcpResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
