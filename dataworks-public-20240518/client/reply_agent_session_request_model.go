// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplyAgentSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ReplyAgentSessionRequest
	GetId() *string
	SetJsonrpc(v string) *ReplyAgentSessionRequest
	GetJsonrpc() *string
	SetParams(v *ReplyAgentSessionRequestParams) *ReplyAgentSessionRequest
	GetParams() *ReplyAgentSessionRequestParams
}

type ReplyAgentSessionRequest struct {
	// The JSON-RPC correlation ID for the current reply request. It is returned as-is in the response and is different from PermissionRequestId.
	//
	// This parameter is required.
	//
	// example:
	//
	// reply-rpc-001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The JSON-RPC protocol version. Fixed to 2.0.
	//
	// example:
	//
	// 2.0
	Jsonrpc *string `json:"Jsonrpc,omitempty" xml:"Jsonrpc,omitempty"`
	// The user interaction reply parameters.
	//
	// This parameter is required.
	Params *ReplyAgentSessionRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Struct"`
}

func (s ReplyAgentSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionRequest) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionRequest) GetId() *string {
	return s.Id
}

func (s *ReplyAgentSessionRequest) GetJsonrpc() *string {
	return s.Jsonrpc
}

func (s *ReplyAgentSessionRequest) GetParams() *ReplyAgentSessionRequestParams {
	return s.Params
}

func (s *ReplyAgentSessionRequest) SetId(v string) *ReplyAgentSessionRequest {
	s.Id = &v
	return s
}

func (s *ReplyAgentSessionRequest) SetJsonrpc(v string) *ReplyAgentSessionRequest {
	s.Jsonrpc = &v
	return s
}

func (s *ReplyAgentSessionRequest) SetParams(v *ReplyAgentSessionRequestParams) *ReplyAgentSessionRequest {
	s.Params = v
	return s
}

func (s *ReplyAgentSessionRequest) Validate() error {
	if s.Params != nil {
		if err := s.Params.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplyAgentSessionRequestParams struct {
	// The answers to ask_user_question. The key is a 0-based string index of the question, and the value is the answer text. Fill in each question one by one for multiple questions. Omit this parameter for general tool authorization or cancellation.
	//
	// example:
	//
	// {"0":"lakehouse_uat"}
	Answers map[string]*string `json:"Answers,omitempty" xml:"Answers,omitempty"`
	// The processing result of the user for the current interaction.
	//
	// This parameter is required.
	Outcome *ReplyAgentSessionRequestParamsOutcome `json:"Outcome,omitempty" xml:"Outcome,omitempty" type:"Struct"`
	// The identifier of the current permission_request. It is obtained from _qwen/notify.params.data.requestId of the original SSE, and is not ToolCallId, HTTP RequestId, or the current JSON-RPC Id. It cannot be . or ..
	//
	// This parameter is required.
	//
	// example:
	//
	// permission-001
	PermissionRequestId *string `json:"PermissionRequestId,omitempty" xml:"PermissionRequestId,omitempty"`
	// The LSP session ID. Use the SessionId returned when creating the session, not the internal session ID of the daemon.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsp-session-001
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s ReplyAgentSessionRequestParams) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionRequestParams) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionRequestParams) GetAnswers() map[string]*string {
	return s.Answers
}

func (s *ReplyAgentSessionRequestParams) GetOutcome() *ReplyAgentSessionRequestParamsOutcome {
	return s.Outcome
}

func (s *ReplyAgentSessionRequestParams) GetPermissionRequestId() *string {
	return s.PermissionRequestId
}

func (s *ReplyAgentSessionRequestParams) GetSessionId() *string {
	return s.SessionId
}

func (s *ReplyAgentSessionRequestParams) SetAnswers(v map[string]*string) *ReplyAgentSessionRequestParams {
	s.Answers = v
	return s
}

func (s *ReplyAgentSessionRequestParams) SetOutcome(v *ReplyAgentSessionRequestParamsOutcome) *ReplyAgentSessionRequestParams {
	s.Outcome = v
	return s
}

func (s *ReplyAgentSessionRequestParams) SetPermissionRequestId(v string) *ReplyAgentSessionRequestParams {
	s.PermissionRequestId = &v
	return s
}

func (s *ReplyAgentSessionRequestParams) SetSessionId(v string) *ReplyAgentSessionRequestParams {
	s.SessionId = &v
	return s
}

func (s *ReplyAgentSessionRequestParams) Validate() error {
	if s.Outcome != nil {
		if err := s.Outcome.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplyAgentSessionRequestParamsOutcome struct {
	// Required and cannot be empty when Outcome is selected. It takes the optionId of the actual option in the options of the current event (when submitting an answer, the option with kind=allow_once is usually selected). Omit this parameter when Outcome is cancelled.
	//
	// example:
	//
	// option-from-event
	OptionId *string `json:"OptionId,omitempty" xml:"OptionId,omitempty"`
	// The processing type. selected indicates that an option is selected, and cancelled indicates that the user explicitly cancels the interaction.
	//
	// This parameter is required.
	//
	// example:
	//
	// selected
	Outcome *string `json:"Outcome,omitempty" xml:"Outcome,omitempty"`
}

func (s ReplyAgentSessionRequestParamsOutcome) String() string {
	return dara.Prettify(s)
}

func (s ReplyAgentSessionRequestParamsOutcome) GoString() string {
	return s.String()
}

func (s *ReplyAgentSessionRequestParamsOutcome) GetOptionId() *string {
	return s.OptionId
}

func (s *ReplyAgentSessionRequestParamsOutcome) GetOutcome() *string {
	return s.Outcome
}

func (s *ReplyAgentSessionRequestParamsOutcome) SetOptionId(v string) *ReplyAgentSessionRequestParamsOutcome {
	s.OptionId = &v
	return s
}

func (s *ReplyAgentSessionRequestParamsOutcome) SetOutcome(v string) *ReplyAgentSessionRequestParamsOutcome {
	s.Outcome = &v
	return s
}

func (s *ReplyAgentSessionRequestParamsOutcome) Validate() error {
	return dara.Validate(s)
}
