// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentAccessKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody
	GetAccessKey() *string
	SetCode(v int32) *DescribeMonitoringAgentAccessKeyResponseBody
	GetCode() *int32
	SetMessage(v string) *DescribeMonitoringAgentAccessKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeMonitoringAgentAccessKeyResponseBody
	GetRequestId() *string
	SetSecretKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody
	GetSecretKey() *string
	SetSuccess(v bool) *DescribeMonitoringAgentAccessKeyResponseBody
	GetSuccess() *bool
}

type DescribeMonitoringAgentAccessKeyResponseBody struct {
	// The AccessKey ID that is required to install the agent.
	//
	// example:
	//
	// E7A27f9****
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C0F655E9-D793-51E6-BD78-CABBCCFC1047
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The AccessKey secret that is required to install the agent.
	//
	// example:
	//
	// sBjNOUhLylrxMX-Xv1****
	SecretKey *string `json:"SecretKey,omitempty" xml:"SecretKey,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentAccessKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetAccessKey() *string {
	return s.AccessKey
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetSecretKey() *string {
	return s.SecretKey
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetAccessKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.AccessKey = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetCode(v int32) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetMessage(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetRequestId(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetSecretKey(v string) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.SecretKey = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentAccessKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringAgentAccessKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
