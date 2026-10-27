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
	// The business result code. If the request fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result. For the result structure, see the response fields of this API operation. For asynchronous write API operations, the response indicates that the request is accepted and does not guarantee that the background task is completed.
	Data *PrepareIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the current request. This code indicates the processing status of the request and is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The result message, which indicates that the request is accepted or contains an error message. For asynchronous tasks, check the task result by calling the corresponding query API operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. It is used to associate the current call with troubleshooting. Retain this ID if the request fails.
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
	// The AgentIdentity receiving configuration. The actual SAML configuration is read on demand only by calling GetIdentityProvider. SCIM query and prepare operations return the actual receiving status and URL.
	Configuration *PrepareIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The current credential summary, including credentialId, status, createTime, and expireTime. It does not include the repeatedly retrievable token.
	Credential *PrepareIdentityProviderSCIMResponseBodyDataCredential `json:"credential,omitempty" xml:"credential,omitempty" type:"Struct"`
	// The plaintext token returned only when a new credential is created. If the prepare operation reuses a valid credential, this value is null. The token cannot be retrieved again from the list. If the response is lost, query the list first. The token is not automatically reissued. Do not write the token to logs or persistent frontend caches.
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
	// The URL configured for the external provisioning system. The presence of a URL does not indicate that the external configuration is completed or that AC synchronization is enabled.
	//
	// example:
	//
	// https://example.com/scim/v2
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// The AgentIdentity receiving status. Enabled indicates that SCIM receiving is enabled, and Disabled indicates that it is disabled. This status is not equivalent to the expected value of syncEnabled in AC.
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
	// The credential ID. This parameter is required and cannot be empty when you delete a credential. Use the ID returned in the credential list, not the plaintext token.
	//
	// example:
	//
	// credential-example
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The expiration time returned by AgentIdentity. Only credentials that are in the Enabled state and have not expired are considered available.
	//
	// example:
	//
	// 2027-10-26T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The AgentIdentity credential status. Only credentials that are in the Enabled state and have not expired are considered available. This field is not the operationStatus of a SCIM task.
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
