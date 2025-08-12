// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteGroupMonitoringAgentProcessResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteGroupMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGroupMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGroupMonitoringAgentProcessResponseBody
	GetSuccess() *bool
}

type DeleteGroupMonitoringAgentProcessResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3F6150F9-45C7-43F9-9578-A58B2E726C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s DeleteGroupMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetCode(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DeleteGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGroupMonitoringAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
