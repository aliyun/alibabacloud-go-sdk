// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataAgentMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *ModifyDataAgentMcpResponseBody
	GetData() interface{}
	SetErrorCode(v string) *ModifyDataAgentMcpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ModifyDataAgentMcpResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ModifyDataAgentMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyDataAgentMcpResponseBody
	GetSuccess() *bool
}

type ModifyDataAgentMcpResponseBody struct {
	// The updated MCP information.
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-***-695C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDataAgentMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataAgentMcpResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataAgentMcpResponseBody) GetData() interface{} {
	return s.Data
}

func (s *ModifyDataAgentMcpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ModifyDataAgentMcpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ModifyDataAgentMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataAgentMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyDataAgentMcpResponseBody) SetData(v interface{}) *ModifyDataAgentMcpResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDataAgentMcpResponseBody) SetErrorCode(v string) *ModifyDataAgentMcpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDataAgentMcpResponseBody) SetErrorMessage(v string) *ModifyDataAgentMcpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDataAgentMcpResponseBody) SetRequestId(v string) *ModifyDataAgentMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataAgentMcpResponseBody) SetSuccess(v bool) *ModifyDataAgentMcpResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDataAgentMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
