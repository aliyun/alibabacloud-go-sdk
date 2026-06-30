// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentResponseBody
	GetCode() *string
	SetData(v *AgentInfo) *GetAgentResponseBody
	GetData() *AgentInfo
	SetMessage(v string) *GetAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentResponseBody
	GetSuccess() *bool
}

type GetAgentResponseBody struct {
	// The response code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *AgentInfo `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 6987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - true: The call was successful.
	//
	// - false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentResponseBody) GetData() *AgentInfo {
	return s.Data
}

func (s *GetAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *AgentInfo) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetSuccess(v bool) *GetAgentResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
