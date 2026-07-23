// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataAgentResponseBodyData) *DeleteDataAgentResponseBody
	GetData() *DeleteDataAgentResponseBodyData
	SetErrorCode(v string) *DeleteDataAgentResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataAgentResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAgentResponseBody
	GetSuccess() *bool
}

type DeleteDataAgentResponseBody struct {
	// The response struct.
	Data *DeleteDataAgentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// InvalidTid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentResponseBody) GetData() *DeleteDataAgentResponseBodyData {
	return s.Data
}

func (s *DeleteDataAgentResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataAgentResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAgentResponseBody) SetData(v *DeleteDataAgentResponseBodyData) *DeleteDataAgentResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataAgentResponseBody) SetErrorCode(v string) *DeleteDataAgentResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataAgentResponseBody) SetErrorMessage(v string) *DeleteDataAgentResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataAgentResponseBody) SetRequestId(v string) *DeleteDataAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAgentResponseBody) SetSuccess(v bool) *DeleteDataAgentResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteDataAgentResponseBodyData struct {
	// Agent Id
	//
	// example:
	//
	// cu0cs*******mf
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The Agent status.
	//
	// example:
	//
	// RUNNING
	AgentStatus *string `json:"AgentStatus,omitempty" xml:"AgentStatus,omitempty"`
}

func (s DeleteDataAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataAgentResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *DeleteDataAgentResponseBodyData) GetAgentStatus() *string {
	return s.AgentStatus
}

func (s *DeleteDataAgentResponseBodyData) SetAgentId(v string) *DeleteDataAgentResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *DeleteDataAgentResponseBodyData) SetAgentStatus(v string) *DeleteDataAgentResponseBodyData {
	s.AgentStatus = &v
	return s
}

func (s *DeleteDataAgentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
