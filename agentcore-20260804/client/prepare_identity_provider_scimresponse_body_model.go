// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrepareIdentityProviderSCIMResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PrepareIdentityProviderSCIMResponseBody
	GetCode() *string
	SetData(v *PrepareIdentityProviderSCIMResponseBodyData) *PrepareIdentityProviderSCIMResponseBody
	GetData() *PrepareIdentityProviderSCIMResponseBodyData
	SetHttpStatusCode(v int32) *PrepareIdentityProviderSCIMResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *PrepareIdentityProviderSCIMResponseBody
	GetMessage() *string
	SetRequestId(v string) *PrepareIdentityProviderSCIMResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PrepareIdentityProviderSCIMResponseBody
	GetSuccess() *bool
}

type PrepareIdentityProviderSCIMResponseBody struct {
	// The business result code. If the call fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result. For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data *PrepareIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The processing status of the current request. This status does not indicate the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request acceptance or error message. For asynchronous tasks, you must check the task result by calling the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. You can use this ID to associate the current call with troubleshooting. Keep the ID if the call fails.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous write operations, you must continue to query the task result.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s PrepareIdentityProviderSCIMResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMResponseBody) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetCode() *string {
	return s.Code
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetData() *PrepareIdentityProviderSCIMResponseBodyData {
	return s.Data
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrepareIdentityProviderSCIMResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetCode(v string) *PrepareIdentityProviderSCIMResponseBody {
	s.Code = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetData(v *PrepareIdentityProviderSCIMResponseBodyData) *PrepareIdentityProviderSCIMResponseBody {
	s.Data = v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetHttpStatusCode(v int32) *PrepareIdentityProviderSCIMResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetMessage(v string) *PrepareIdentityProviderSCIMResponseBody {
	s.Message = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetRequestId(v string) *PrepareIdentityProviderSCIMResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) SetSuccess(v bool) *PrepareIdentityProviderSCIMResponseBody {
	s.Success = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PrepareIdentityProviderSCIMResponseBodyData struct {
	// The actual SAML configuration is read on demand only by calling GetIdentityProvider. SCIM query and preparation operations return the actual receiving status and URL.
	Configuration *PrepareIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The current credential summary, including credentialId, status, createTime, and expireTime. The token is not included because it cannot be repeatedly obtained.
	Credential *PrepareIdentityProviderSCIMResponseBodyDataCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The plaintext of the newly issued SCIM token. The plaintext is returned only in the current response. If an existing valid credential is reused, an empty string ("") is returned. In this case, use the saved token.
	//
	// example:
	//
	// example-scim-token
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s PrepareIdentityProviderSCIMResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMResponseBodyData) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) GetConfiguration() *PrepareIdentityProviderSCIMResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) GetCredential() *PrepareIdentityProviderSCIMResponseBodyDataCredential {
	return s.Credential
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) SetConfiguration(v *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) *PrepareIdentityProviderSCIMResponseBodyData {
	s.Configuration = v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) SetCredential(v *PrepareIdentityProviderSCIMResponseBodyDataCredential) *PrepareIdentityProviderSCIMResponseBodyData {
	s.Credential = v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) SetToken(v string) *PrepareIdentityProviderSCIMResponseBodyData {
	s.Token = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyData) Validate() error {
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

type PrepareIdentityProviderSCIMResponseBodyDataConfiguration struct {
	// The URL configured for the external provisioning system. The presence of a URL does not indicate that the external configuration is completed or AgentCore synchronization is enabled.
	//
	// example:
	//
	// https://example.com/scim/v2
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// The receiving status of AgentIdentity. A value of Enabled indicates that SCIM receiving is enabled. A value of Disabled indicates that SCIM receiving is disabled. This status is not equivalent to the expected value of syncEnabled in AgentCore.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PrepareIdentityProviderSCIMResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) GetStatus() *string {
	return s.Status
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) SetBaseUrl(v string) *PrepareIdentityProviderSCIMResponseBodyDataConfiguration {
	s.BaseUrl = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) SetStatus(v string) *PrepareIdentityProviderSCIMResponseBodyDataConfiguration {
	s.Status = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}

type PrepareIdentityProviderSCIMResponseBodyDataCredential struct {
	// The creation time returned by AgentIdentity.
	//
	// example:
	//
	// 2026-10-26T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The credential ID. This parameter is required when you delete a credential. Use the ID returned in the credential list, not the plaintext of the token.
	//
	// example:
	//
	// credential-example
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The expiration time returned by AgentIdentity. Only credentials with the status of Enabled that have not expired are considered available.
	//
	// example:
	//
	// 2027-10-26T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The credential status of AgentIdentity. Only credentials with the status of Enabled that have not expired are considered available. This field is not the operationStatus of the SCIM task.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PrepareIdentityProviderSCIMResponseBodyDataCredential) String() string {
	return dara.Prettify(s)
}

func (s PrepareIdentityProviderSCIMResponseBodyDataCredential) GoString() string {
	return s.String()
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) GetCredentialId() *string {
	return s.CredentialId
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) GetStatus() *string {
	return s.Status
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) SetCreateTime(v string) *PrepareIdentityProviderSCIMResponseBodyDataCredential {
	s.CreateTime = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) SetCredentialId(v string) *PrepareIdentityProviderSCIMResponseBodyDataCredential {
	s.CredentialId = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) SetExpireTime(v string) *PrepareIdentityProviderSCIMResponseBodyDataCredential {
	s.ExpireTime = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) SetStatus(v string) *PrepareIdentityProviderSCIMResponseBodyDataCredential {
	s.Status = &v
	return s
}

func (s *PrepareIdentityProviderSCIMResponseBodyDataCredential) Validate() error {
	return dara.Validate(s)
}
