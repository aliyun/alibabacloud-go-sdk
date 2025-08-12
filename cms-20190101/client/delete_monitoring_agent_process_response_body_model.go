// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMonitoringAgentProcessResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMonitoringAgentProcessResponseBody
	GetSuccess() *bool
}

type DeleteMonitoringAgentProcessResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 971CC023-5A96-452A-BB7C-2483F948BCFD
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

func (s DeleteMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetCode(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetMessage(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetRequestId(v string) *DeleteMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) SetSuccess(v bool) *DeleteMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitoringAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}
