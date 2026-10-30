// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindCustomAgentImagePreparationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FindCustomAgentImagePreparationResponseBody
	GetCode() *string
	SetData(v *FindCustomAgentImagePreparationResponseBodyData) *FindCustomAgentImagePreparationResponseBody
	GetData() *FindCustomAgentImagePreparationResponseBodyData
	SetHttpStatusCode(v int32) *FindCustomAgentImagePreparationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FindCustomAgentImagePreparationResponseBody
	GetMessage() *string
	SetRequestId(v string) *FindCustomAgentImagePreparationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FindCustomAgentImagePreparationResponseBody
	GetSuccess() *bool
}

type FindCustomAgentImagePreparationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The preparation query result. For the asynchronous execution status, refer to the phase.
	Data *FindCustomAgentImagePreparationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FindCustomAgentImagePreparationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FindCustomAgentImagePreparationResponseBody) GoString() string {
	return s.String()
}

func (s *FindCustomAgentImagePreparationResponseBody) GetCode() *string {
	return s.Code
}

func (s *FindCustomAgentImagePreparationResponseBody) GetData() *FindCustomAgentImagePreparationResponseBodyData {
	return s.Data
}

func (s *FindCustomAgentImagePreparationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FindCustomAgentImagePreparationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FindCustomAgentImagePreparationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FindCustomAgentImagePreparationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FindCustomAgentImagePreparationResponseBody) SetCode(v string) *FindCustomAgentImagePreparationResponseBody {
	s.Code = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) SetData(v *FindCustomAgentImagePreparationResponseBodyData) *FindCustomAgentImagePreparationResponseBody {
	s.Data = v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) SetHttpStatusCode(v int32) *FindCustomAgentImagePreparationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) SetMessage(v string) *FindCustomAgentImagePreparationResponseBody {
	s.Message = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) SetRequestId(v string) *FindCustomAgentImagePreparationResponseBody {
	s.RequestId = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) SetSuccess(v bool) *FindCustomAgentImagePreparationResponseBody {
	s.Success = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindCustomAgentImagePreparationResponseBodyData struct {
	// The ID of the target agent.
	//
	// example:
	//
	// agent-example
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The security error. This parameter is returned only when the task fails. Otherwise, the value is null.
	Error *FindCustomAgentImagePreparationResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The phase where the failure occurred. This parameter is returned only when the phase is FAILED. Otherwise, the value is null.
	//
	// example:
	//
	// COPYING
	FailedPhase *string `json:"failedPhase,omitempty" xml:"failedPhase,omitempty"`
	// The current phase. Only SUCCEEDED indicates that the target version is deployed successfully.
	//
	// example:
	//
	// FAILED
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The ID of the task submitted in the console.
	//
	// example:
	//
	// agent-example:1
	PreparationId *string `json:"preparationId,omitempty" xml:"preparationId,omitempty"`
}

func (s FindCustomAgentImagePreparationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FindCustomAgentImagePreparationResponseBodyData) GoString() string {
	return s.String()
}

func (s *FindCustomAgentImagePreparationResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *FindCustomAgentImagePreparationResponseBodyData) GetError() *FindCustomAgentImagePreparationResponseBodyDataError {
	return s.Error
}

func (s *FindCustomAgentImagePreparationResponseBodyData) GetFailedPhase() *string {
	return s.FailedPhase
}

func (s *FindCustomAgentImagePreparationResponseBodyData) GetPhase() *string {
	return s.Phase
}

func (s *FindCustomAgentImagePreparationResponseBodyData) GetPreparationId() *string {
	return s.PreparationId
}

func (s *FindCustomAgentImagePreparationResponseBodyData) SetAgentId(v string) *FindCustomAgentImagePreparationResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyData) SetError(v *FindCustomAgentImagePreparationResponseBodyDataError) *FindCustomAgentImagePreparationResponseBodyData {
	s.Error = v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyData) SetFailedPhase(v string) *FindCustomAgentImagePreparationResponseBodyData {
	s.FailedPhase = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyData) SetPhase(v string) *FindCustomAgentImagePreparationResponseBodyData {
	s.Phase = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyData) SetPreparationId(v string) *FindCustomAgentImagePreparationResponseBodyData {
	s.PreparationId = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FindCustomAgentImagePreparationResponseBodyDataError struct {
	// The business error code for image preparation. This code is used to identify the failure type. For the specific cause, refer to the message.
	//
	// example:
	//
	// ImageSyncFailed
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The specific cause of the failure. This may include underlying errors such as image not found, authentication failure, network timeout, or certificate verification error. Passwords, tokens, or authentication headers are not returned. If the underlying system does not provide a cause, a general description is returned.
	//
	// example:
	//
	// tls: failed to verify certificate: x509: certificate signed by unknown authority
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s FindCustomAgentImagePreparationResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s FindCustomAgentImagePreparationResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *FindCustomAgentImagePreparationResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *FindCustomAgentImagePreparationResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *FindCustomAgentImagePreparationResponseBodyDataError) SetCode(v string) *FindCustomAgentImagePreparationResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyDataError) SetMessage(v string) *FindCustomAgentImagePreparationResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *FindCustomAgentImagePreparationResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
