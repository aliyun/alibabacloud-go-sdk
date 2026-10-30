// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomAgentImagePreparationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCustomAgentImagePreparationResponseBody
	GetCode() *string
	SetData(v *GetCustomAgentImagePreparationResponseBodyData) *GetCustomAgentImagePreparationResponseBody
	GetData() *GetCustomAgentImagePreparationResponseBodyData
	SetHttpStatusCode(v int32) *GetCustomAgentImagePreparationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCustomAgentImagePreparationResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCustomAgentImagePreparationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomAgentImagePreparationResponseBody
	GetSuccess() *bool
}

type GetCustomAgentImagePreparationResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The preparation query result. Refer to the phase for the asynchronous execution status.
	Data *GetCustomAgentImagePreparationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetCustomAgentImagePreparationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentImagePreparationResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomAgentImagePreparationResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCustomAgentImagePreparationResponseBody) GetData() *GetCustomAgentImagePreparationResponseBodyData {
	return s.Data
}

func (s *GetCustomAgentImagePreparationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCustomAgentImagePreparationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCustomAgentImagePreparationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomAgentImagePreparationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomAgentImagePreparationResponseBody) SetCode(v string) *GetCustomAgentImagePreparationResponseBody {
	s.Code = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) SetData(v *GetCustomAgentImagePreparationResponseBodyData) *GetCustomAgentImagePreparationResponseBody {
	s.Data = v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) SetHttpStatusCode(v int32) *GetCustomAgentImagePreparationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) SetMessage(v string) *GetCustomAgentImagePreparationResponseBody {
	s.Message = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) SetRequestId(v string) *GetCustomAgentImagePreparationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) SetSuccess(v bool) *GetCustomAgentImagePreparationResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentImagePreparationResponseBodyData struct {
	// The target agent ID.
	//
	// example:
	//
	// agent-example
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The error information. This field is returned only when the task fails. Otherwise, the value is null.
	Error *GetCustomAgentImagePreparationResponseBodyDataError `json:"error,omitempty" xml:"error,omitempty" type:"Struct"`
	// The failed phase. This field is returned only when the status is FAILED. Otherwise, the value is null.
	//
	// example:
	//
	// COPYING
	FailedPhase *string `json:"failedPhase,omitempty" xml:"failedPhase,omitempty"`
	// The current phase. Only SUCCEEDED indicates that the deployment of the target version is successful.
	//
	// example:
	//
	// FAILED
	Phase *string `json:"phase,omitempty" xml:"phase,omitempty"`
	// The console submission task ID.
	//
	// example:
	//
	// agent-example:1
	PreparationId *string `json:"preparationId,omitempty" xml:"preparationId,omitempty"`
}

func (s GetCustomAgentImagePreparationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentImagePreparationResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCustomAgentImagePreparationResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *GetCustomAgentImagePreparationResponseBodyData) GetError() *GetCustomAgentImagePreparationResponseBodyDataError {
	return s.Error
}

func (s *GetCustomAgentImagePreparationResponseBodyData) GetFailedPhase() *string {
	return s.FailedPhase
}

func (s *GetCustomAgentImagePreparationResponseBodyData) GetPhase() *string {
	return s.Phase
}

func (s *GetCustomAgentImagePreparationResponseBodyData) GetPreparationId() *string {
	return s.PreparationId
}

func (s *GetCustomAgentImagePreparationResponseBodyData) SetAgentId(v string) *GetCustomAgentImagePreparationResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyData) SetError(v *GetCustomAgentImagePreparationResponseBodyDataError) *GetCustomAgentImagePreparationResponseBodyData {
	s.Error = v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyData) SetFailedPhase(v string) *GetCustomAgentImagePreparationResponseBodyData {
	s.FailedPhase = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyData) SetPhase(v string) *GetCustomAgentImagePreparationResponseBodyData {
	s.Phase = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyData) SetPreparationId(v string) *GetCustomAgentImagePreparationResponseBodyData {
	s.PreparationId = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyData) Validate() error {
	if s.Error != nil {
		if err := s.Error.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCustomAgentImagePreparationResponseBodyDataError struct {
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

func (s GetCustomAgentImagePreparationResponseBodyDataError) String() string {
	return dara.Prettify(s)
}

func (s GetCustomAgentImagePreparationResponseBodyDataError) GoString() string {
	return s.String()
}

func (s *GetCustomAgentImagePreparationResponseBodyDataError) GetCode() *string {
	return s.Code
}

func (s *GetCustomAgentImagePreparationResponseBodyDataError) GetMessage() *string {
	return s.Message
}

func (s *GetCustomAgentImagePreparationResponseBodyDataError) SetCode(v string) *GetCustomAgentImagePreparationResponseBodyDataError {
	s.Code = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyDataError) SetMessage(v string) *GetCustomAgentImagePreparationResponseBodyDataError {
	s.Message = &v
	return s
}

func (s *GetCustomAgentImagePreparationResponseBodyDataError) Validate() error {
	return dara.Validate(s)
}
