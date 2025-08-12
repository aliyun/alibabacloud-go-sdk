// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGroupMonitoringAgentProcessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateGroupMonitoringAgentProcessResponseBody
	GetCode() *string
	SetMessage(v string) *CreateGroupMonitoringAgentProcessResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateGroupMonitoringAgentProcessResponseBody
	GetRequestId() *string
	SetResource(v *CreateGroupMonitoringAgentProcessResponseBodyResource) *CreateGroupMonitoringAgentProcessResponseBody
	GetResource() *CreateGroupMonitoringAgentProcessResponseBodyResource
	SetSuccess(v bool) *CreateGroupMonitoringAgentProcessResponseBody
	GetSuccess() *bool
}

type CreateGroupMonitoringAgentProcessResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the call is successful.
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
	// The group process information.
	Resource *CreateGroupMonitoringAgentProcessResponseBodyResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessResponseBody) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) GetResource() *CreateGroupMonitoringAgentProcessResponseBodyResource {
	return s.Resource
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetCode(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Code = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetMessage(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Message = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetRequestId(v string) *CreateGroupMonitoringAgentProcessResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetResource(v *CreateGroupMonitoringAgentProcessResponseBodyResource) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Resource = v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) SetSuccess(v bool) *CreateGroupMonitoringAgentProcessResponseBody {
	s.Success = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateGroupMonitoringAgentProcessResponseBodyResource struct {
	// The ID of the group process.
	//
	// example:
	//
	// 7F2B0024-4F21-48B9-A764-211CEC48****
	GroupProcessId *string `json:"GroupProcessId,omitempty" xml:"GroupProcessId,omitempty"`
}

func (s CreateGroupMonitoringAgentProcessResponseBodyResource) String() string {
	return dara.Prettify(s)
}

func (s CreateGroupMonitoringAgentProcessResponseBodyResource) GoString() string {
	return s.String()
}

func (s *CreateGroupMonitoringAgentProcessResponseBodyResource) GetGroupProcessId() *string {
	return s.GroupProcessId
}

func (s *CreateGroupMonitoringAgentProcessResponseBodyResource) SetGroupProcessId(v string) *CreateGroupMonitoringAgentProcessResponseBodyResource {
	s.GroupProcessId = &v
	return s
}

func (s *CreateGroupMonitoringAgentProcessResponseBodyResource) Validate() error {
	return dara.Validate(s)
}
