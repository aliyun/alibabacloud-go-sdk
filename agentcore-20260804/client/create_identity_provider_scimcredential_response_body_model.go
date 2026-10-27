// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentityProviderSCIMCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIdentityProviderSCIMCredentialResponseBody
	GetCode() *string
	SetData(v *CreateIdentityProviderSCIMCredentialResponseBodyData) *CreateIdentityProviderSCIMCredentialResponseBody
	GetData() *CreateIdentityProviderSCIMCredentialResponseBodyData
	SetHttpStatusCode(v int32) *CreateIdentityProviderSCIMCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateIdentityProviderSCIMCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIdentityProviderSCIMCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIdentityProviderSCIMCredentialResponseBody
	GetSuccess() *bool
}

type CreateIdentityProviderSCIMCredentialResponseBody struct {
	// The business result code. Combine it with the HTTP status and message for handling upon failure.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result. For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data *CreateIdentityProviderSCIMCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the current request. It is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Task result message, indicating request acceptance or an error. For asynchronous tasks, check task results by using the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. It is used to associate the current call with troubleshooting. Keep it when a failure occurs.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous writes, you still need to continue querying the task results.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateIdentityProviderSCIMCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetData() *CreateIdentityProviderSCIMCredentialResponseBodyData {
	return s.Data
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetCode(v string) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetData(v *CreateIdentityProviderSCIMCredentialResponseBodyData) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.Data = v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetHttpStatusCode(v int32) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetMessage(v string) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetRequestId(v string) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) SetSuccess(v bool) *CreateIdentityProviderSCIMCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderSCIMCredentialResponseBodyData struct {
	// The AgentIdentity reception configuration. The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual reception status and URL.
	Configuration *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The current credential summary, including credentialId, status, createTime, and expireTime. It does not include the token, which cannot be retrieved repeatedly.
	Credential *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The token. The plaintext is returned only when a new credential is created. It is null when prepare reuses a valid credential. It cannot be retrieved again from the list. If the response is lost, query the list first instead of automatically re-issuing it. Do not write it to logs or frontend persistent cache.
	//
	// example:
	//
	// example-scim-token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) GetConfiguration() *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) GetCredential() *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential {
	return s.Credential
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) SetConfiguration(v *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) *CreateIdentityProviderSCIMCredentialResponseBodyData {
	s.Configuration = v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) SetCredential(v *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) *CreateIdentityProviderSCIMCredentialResponseBodyData {
	s.Credential = v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) SetToken(v string) *CreateIdentityProviderSCIMCredentialResponseBodyData {
	s.Token = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	if s.Credential != nil {
		if err := s.Credential.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration struct {
	// The URL configured for the external provisioning system. The presence of an address does not mean that the external configuration is completed or AC synchronization is enabled.
	//
	// example:
	//
	// https://example.com/scim/v2
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// The AgentIdentity reception status. Enabled indicates that SCIM reception is enabled, and Disabled indicates that it is disabled. This is not equivalent to the expected value of syncEnabled in AC.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) GetStatus() *string {
	return s.Status
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) SetBaseUrl(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration {
	s.BaseUrl = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) SetStatus(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration {
	s.Status = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type CreateIdentityProviderSCIMCredentialResponseBodyDataCredential struct {
	// The creation time returned by AgentIdentity.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The credential ID. It must be non-empty when deleting a credential. Use the ID returned in the credential list, not the plaintext token.
	//
	// example:
	//
	// credential-example
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The expiration time returned by AgentIdentity. Only credentials that are Enabled and have not expired are considered available.
	//
	// example:
	//
	// 2027-10-26T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The AgentIdentity credential status. Only credentials that are Enabled and have not expired are considered available. This field is not the operationStatus of a SCIM task.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) GoString() string {
	return s.String()
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) GetStatus() *string {
	return s.Status
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) SetCreateTime(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential {
	s.CreateTime = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) SetCredentialId(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential {
	s.CredentialId = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) SetExpireTime(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential {
	s.ExpireTime = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) SetStatus(v string) *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential {
	s.Status = &v
	return s
}

func (s *CreateIdentityProviderSCIMCredentialResponseBodyDataCredential) Validate() error {
	return dara.Validate(s)
}
