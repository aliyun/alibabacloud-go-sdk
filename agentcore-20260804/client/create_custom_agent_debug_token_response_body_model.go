// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomAgentDebugTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCustomAgentDebugTokenResponseBody
	GetCode() *string
	SetData(v *CreateCustomAgentDebugTokenResponseBodyData) *CreateCustomAgentDebugTokenResponseBody
	GetData() *CreateCustomAgentDebugTokenResponseBodyData
	SetHttpStatusCode(v int32) *CreateCustomAgentDebugTokenResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCustomAgentDebugTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCustomAgentDebugTokenResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCustomAgentDebugTokenResponseBody
	GetSuccess() *bool
}

type CreateCustomAgentDebugTokenResponseBody struct {
	// The business status code. SUCCESS is returned if the call is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The debug token issued in this request, the token fingerprint, and the corresponding resource and network information.
	Data *CreateCustomAgentDebugTokenResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the request.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The description of the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The unique identifier of the request, which can be used for troubleshooting.
	//
	// example:
	//
	// F2B17684-703A-4F29-9D14-30C334A85162
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateCustomAgentDebugTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentDebugTokenResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetData() *CreateCustomAgentDebugTokenResponseBodyData {
	return s.Data
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomAgentDebugTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetCode(v string) *CreateCustomAgentDebugTokenResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetData(v *CreateCustomAgentDebugTokenResponseBodyData) *CreateCustomAgentDebugTokenResponseBody {
	s.Data = v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetHttpStatusCode(v int32) *CreateCustomAgentDebugTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetMessage(v string) *CreateCustomAgentDebugTokenResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetRequestId(v string) *CreateCustomAgentDebugTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) SetSuccess(v bool) *CreateCustomAgentDebugTokenResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCustomAgentDebugTokenResponseBodyData struct {
	// The ID of the high-code agent corresponding to the token.
	//
	// example:
	//
	// agent-0123456789abcdef0123456789abcdef
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The token used by the AgentCore SDK for local debugging. Keep the token secure. Do not write it to logs or commit it to code repositories.
	//
	// example:
	//
	// eyJwcm9kdWN0****
	BootstrapToken *string `json:"bootstrapToken,omitempty" xml:"bootstrapToken,omitempty"`
	// The network type used for the token, which is the same as the networkType parameter in the request.
	//
	// example:
	//
	// INTERNET
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The fingerprint used to identify the issued token, which can be used to troubleshoot issuance records.
	//
	// example:
	//
	// c4f5b8813a0d
	TokenFingerprint *string `json:"tokenFingerprint,omitempty" xml:"tokenFingerprint,omitempty"`
	// The ID of the workspace corresponding to the token.
	//
	// example:
	//
	// ws-0123456789abcdef01234
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateCustomAgentDebugTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomAgentDebugTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) GetAgentId() *string {
	return s.AgentId
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) GetBootstrapToken() *string {
	return s.BootstrapToken
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) GetTokenFingerprint() *string {
	return s.TokenFingerprint
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) SetAgentId(v string) *CreateCustomAgentDebugTokenResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) SetBootstrapToken(v string) *CreateCustomAgentDebugTokenResponseBodyData {
	s.BootstrapToken = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) SetNetworkType(v string) *CreateCustomAgentDebugTokenResponseBodyData {
	s.NetworkType = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) SetTokenFingerprint(v string) *CreateCustomAgentDebugTokenResponseBodyData {
	s.TokenFingerprint = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) SetWorkspaceId(v string) *CreateCustomAgentDebugTokenResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateCustomAgentDebugTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
