// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIdentityProviderSCIMResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetIdentityProviderSCIMResponseBody
	GetCode() *string
	SetData(v *GetIdentityProviderSCIMResponseBodyData) *GetIdentityProviderSCIMResponseBody
	GetData() *GetIdentityProviderSCIMResponseBodyData
	SetHttpStatusCode(v int32) *GetIdentityProviderSCIMResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetIdentityProviderSCIMResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetIdentityProviderSCIMResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetIdentityProviderSCIMResponseBody
	GetSuccess() *bool
}

type GetIdentityProviderSCIMResponseBody struct {
	// If the request fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data *GetIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The processing status of this request, which is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request acceptance or error message. For asynchronous tasks, verify the task result by using the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Used to associate this call with troubleshooting. Retain this ID if the call fails.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// A value of true indicates that the current request is processed successfully. For asynchronous writes, you must continue to query the task results.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetIdentityProviderSCIMResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSCIMResponseBody) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSCIMResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetIdentityProviderSCIMResponseBody) GetData() *GetIdentityProviderSCIMResponseBodyData {
	return s.Data
}

func (s *GetIdentityProviderSCIMResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetIdentityProviderSCIMResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetIdentityProviderSCIMResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIdentityProviderSCIMResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetIdentityProviderSCIMResponseBody) SetCode(v string) *GetIdentityProviderSCIMResponseBody {
	s.Code = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) SetData(v *GetIdentityProviderSCIMResponseBodyData) *GetIdentityProviderSCIMResponseBody {
	s.Data = v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) SetHttpStatusCode(v int32) *GetIdentityProviderSCIMResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) SetMessage(v string) *GetIdentityProviderSCIMResponseBody {
	s.Message = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) SetRequestId(v string) *GetIdentityProviderSCIMResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) SetSuccess(v bool) *GetIdentityProviderSCIMResponseBody {
	s.Success = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSCIMResponseBodyData struct {
	// The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual receiving status and URL.
	Configuration *GetIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The original status of the most recent independent SCIM operation task. An empty string ("") is returned if no historical tasks exist.
	//
	// example:
	//
	// SUCCESS
	OperationStatus *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
	// The SCIM independent configuration status.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Do not pass true during CustomSSO initialization. Subsequently, it can only be updated through UpdateIdentityProviderSCIM, and the result is determined in combination with operationStatus.
	SyncEnabled *bool `json:"syncEnabled,omitempty" xml:"syncEnabled,omitempty"`
}

func (s GetIdentityProviderSCIMResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSCIMResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSCIMResponseBodyData) GetConfiguration() *GetIdentityProviderSCIMResponseBodyDataConfiguration {
	return s.Configuration
}

func (s *GetIdentityProviderSCIMResponseBodyData) GetOperationStatus() *string {
	return s.OperationStatus
}

func (s *GetIdentityProviderSCIMResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderSCIMResponseBodyData) GetSyncEnabled() *bool {
	return s.SyncEnabled
}

func (s *GetIdentityProviderSCIMResponseBodyData) SetConfiguration(v *GetIdentityProviderSCIMResponseBodyDataConfiguration) *GetIdentityProviderSCIMResponseBodyData {
	s.Configuration = v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyData) SetOperationStatus(v string) *GetIdentityProviderSCIMResponseBodyData {
	s.OperationStatus = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyData) SetStatus(v string) *GetIdentityProviderSCIMResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyData) SetSyncEnabled(v bool) *GetIdentityProviderSCIMResponseBodyData {
	s.SyncEnabled = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyData) Validate() error {
	if s.Configuration != nil {
		if err := s.Configuration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIdentityProviderSCIMResponseBodyDataConfiguration struct {
	// The URL configured for the external provisioning system. The presence of a URL does not mean that the external configuration is completed or AgentCore synchronization is enabled.
	//
	// example:
	//
	// https://example.com/scim/v2
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// Enabled indicates that SCIM receiving is enabled, and Disabled indicates that it is disabled. This is not equivalent to the expected value of syncEnabled in AgentCore.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetIdentityProviderSCIMResponseBodyDataConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetIdentityProviderSCIMResponseBodyDataConfiguration) GoString() string {
	return s.String()
}

func (s *GetIdentityProviderSCIMResponseBodyDataConfiguration) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *GetIdentityProviderSCIMResponseBodyDataConfiguration) GetStatus() *string {
	return s.Status
}

func (s *GetIdentityProviderSCIMResponseBodyDataConfiguration) SetBaseUrl(v string) *GetIdentityProviderSCIMResponseBodyDataConfiguration {
	s.BaseUrl = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyDataConfiguration) SetStatus(v string) *GetIdentityProviderSCIMResponseBodyDataConfiguration {
	s.Status = &v
	return s
}

func (s *GetIdentityProviderSCIMResponseBodyDataConfiguration) Validate() error {
	return dara.Validate(s)
}
