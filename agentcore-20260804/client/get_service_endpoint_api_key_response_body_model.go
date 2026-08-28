// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceEndpointApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceEndpointApiKeyResponseBody
	GetCode() *string
	SetData(v *GetServiceEndpointApiKeyResponseBodyData) *GetServiceEndpointApiKeyResponseBody
	GetData() *GetServiceEndpointApiKeyResponseBodyData
	SetHttpStatusCode(v int32) *GetServiceEndpointApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetServiceEndpointApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceEndpointApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceEndpointApiKeyResponseBody
	GetSuccess() *bool
}

type GetServiceEndpointApiKeyResponseBody struct {
	// The response code. The value is SUCCESS when the request succeeds.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The currently active API Key information for the service endpoint.
	Data *GetServiceEndpointApiKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code. The value is 200 when the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message. The value is success when the request succeeds.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, used for troubleshooting and tracing.
	//
	// example:
	//
	// req-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. The value is true when the request succeeds.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetServiceEndpointApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceEndpointApiKeyResponseBody) GetData() *GetServiceEndpointApiKeyResponseBodyData {
	return s.Data
}

func (s *GetServiceEndpointApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetServiceEndpointApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceEndpointApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceEndpointApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceEndpointApiKeyResponseBody) SetCode(v string) *GetServiceEndpointApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) SetData(v *GetServiceEndpointApiKeyResponseBodyData) *GetServiceEndpointApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) SetHttpStatusCode(v int32) *GetServiceEndpointApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) SetMessage(v string) *GetServiceEndpointApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) SetRequestId(v string) *GetServiceEndpointApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) SetSuccess(v bool) *GetServiceEndpointApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceEndpointApiKeyResponseBodyData struct {
	// The currently active API Key for the service endpoint. The service reads this value from the gateway consumer in real time. AgentCore does not persist the plaintext. When calling the service endpoint, include this value in the request header specified by apiKeyName. Do not log this value or expose it in public configurations.
	//
	// example:
	//
	// example-api-key-value
	ApiKey *string `json:"apiKey,omitempty" xml:"apiKey,omitempty"`
	// The API Key fingerprint, which consists of the first 12 lowercase hexadecimal characters of the SHA-256 digest of the API Key. It can be used to identify the key version but cannot replace the API Key for authentication.
	//
	// example:
	//
	// b2520bf19231
	ApiKeyFingerprint *string `json:"apiKeyFingerprint,omitempty" xml:"apiKeyFingerprint,omitempty"`
	// The name of the HTTP request header used to pass the API Key. The value is currently fixed to x-api-key.
	//
	// example:
	//
	// x-api-key
	ApiKeyName *string `json:"apiKeyName,omitempty" xml:"apiKeyName,omitempty"`
	// The location where the API Key is passed. The value is currently fixed to Header, indicating that the API Key is passed through an HTTP request header.
	//
	// example:
	//
	// Header
	ApiKeySource *string `json:"apiKeySource,omitempty" xml:"apiKeySource,omitempty"`
	// The authentication type of the service endpoint. Valid values:
	//
	// - NONE: Authentication is not enabled.
	//
	// - API_KEY: API Key authentication is used.
	//
	// This operation succeeds only when the authentication type is API_KEY. Therefore, the value API_KEY is always returned in a successful response.
	//
	// example:
	//
	// API_KEY
	AuthenticationType *string `json:"authenticationType,omitempty" xml:"authenticationType,omitempty"`
	// The service endpoint ID.
	//
	// example:
	//
	// se-1
	ServiceEndpointId *string `json:"serviceEndpointId,omitempty" xml:"serviceEndpointId,omitempty"`
	// The ID of the workspace to which the service endpoint belongs.
	//
	// example:
	//
	// ws-1
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetServiceEndpointApiKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceEndpointApiKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetApiKey() *string {
	return s.ApiKey
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetApiKeyFingerprint() *string {
	return s.ApiKeyFingerprint
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetApiKeySource() *string {
	return s.ApiKeySource
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetAuthenticationType() *string {
	return s.AuthenticationType
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetServiceEndpointId() *string {
	return s.ServiceEndpointId
}

func (s *GetServiceEndpointApiKeyResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetApiKey(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.ApiKey = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetApiKeyFingerprint(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.ApiKeyFingerprint = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetApiKeyName(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.ApiKeyName = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetApiKeySource(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.ApiKeySource = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetAuthenticationType(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.AuthenticationType = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetServiceEndpointId(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.ServiceEndpointId = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) SetWorkspaceId(v string) *GetServiceEndpointApiKeyResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetServiceEndpointApiKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
