// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallDataAgentMcpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *InstallDataAgentMcpResponseBody
	GetData() interface{}
	SetErrorCode(v string) *InstallDataAgentMcpResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *InstallDataAgentMcpResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *InstallDataAgentMcpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallDataAgentMcpResponseBody
	GetSuccess() *bool
}

type InstallDataAgentMcpResponseBody struct {
	// The MCP ID created in DataAgent.
	//
	// example:
	//
	// {"uuid":"6126jk***h2"}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// DMS-DA-40114
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the request failed.
	//
	// example:
	//
	// Specified parameter is not valid.
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

func (s InstallDataAgentMcpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallDataAgentMcpResponseBody) GoString() string {
	return s.String()
}

func (s *InstallDataAgentMcpResponseBody) GetData() interface{} {
	return s.Data
}

func (s *InstallDataAgentMcpResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *InstallDataAgentMcpResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *InstallDataAgentMcpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallDataAgentMcpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallDataAgentMcpResponseBody) SetData(v interface{}) *InstallDataAgentMcpResponseBody {
	s.Data = v
	return s
}

func (s *InstallDataAgentMcpResponseBody) SetErrorCode(v string) *InstallDataAgentMcpResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *InstallDataAgentMcpResponseBody) SetErrorMessage(v string) *InstallDataAgentMcpResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *InstallDataAgentMcpResponseBody) SetRequestId(v string) *InstallDataAgentMcpResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallDataAgentMcpResponseBody) SetSuccess(v bool) *InstallDataAgentMcpResponseBody {
	s.Success = &v
	return s
}

func (s *InstallDataAgentMcpResponseBody) Validate() error {
	return dara.Validate(s)
}
