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
	// The business result code. Use this code along with the HTTP status code and message for error handling.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The SCIM actual configuration, expected synchronization switch, and independent operation status. For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, but does not guarantee that the background task is complete.
	Data *GetIdentityProviderSCIMResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code of the current request. This code indicates the processing status of the request and is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Task result message. It indicates that the request is accepted or provides an error message. For asynchronous tasks, you must check task results by calling the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. It is used to associate the current call with troubleshooting. Retain this ID when a failure occurs.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous write operations, you must continue to query the task results.
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
	// The actual receiving status and URL. The actual SAML configuration is read back on demand only in GetIdentityProvider. SCIM queries and preparations return the actual receiving status and URL.
	Configuration *GetIdentityProviderSCIMResponseBodyDataConfiguration `json:"configuration,omitempty" xml:"configuration,omitempty" type:"Struct"`
	// The status of the most recent SCIM operation task. This status is independent of the SAML logon binding status. Valid values:
	//
	// - PENDING: Waiting for execution.
	//
	// - RUNNING: Executing.
	//
	// - PAUSED: Paused.
	//
	// - SUCCESS: Succeeded.
	//
	// - FAILED: Failed.
	//
	// - CANCELED: Canceled.
	//
	// If there are no tasks, it returns null or is not returned. This status only represents the most recent SCIM switch and the initial full import task when enabled. It does not indicate credential issuance, external provisioning connection tests, or continuous incremental synchronization health. Clients poll this operation when the status is PENDING or RUNNING.
	//
	// example:
	//
	// SUCCESS
	OperationStatus *string `json:"operationStatus,omitempty" xml:"operationStatus,omitempty"`
	// The independent SCIM configuration status. Valid values:
	//
	// - UNCONFIGURED: Preparation is not complete.
	//
	// - WAITING_CONFIGURATION: Receiving is enabled, a valid HTTPS URL and at least one valid token are available, waiting for external configuration and enablement confirmation.
	//
	// - UPDATING: The enablement or disablement task is PENDING or RUNNING.
	//
	// - READY: Synchronization is enabled and the current enablement task succeeded, and the receiving end, URL, and token are still available.
	//
	// - DISABLED: The disablement task succeeded, and the configuration, credentials, and users are retained.
	//
	// - UPDATE_FAILED: The operation failed, was paused, was canceled, or the configuration and credentials are unavailable when enablement is expected.
	//
	// This status is independent of the SAML or OIDC status. READY does not prove that Matrix execution or the continuous incremental link has been accepted.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Specifies whether local synchronization is expected. You must determine whether it takes effect based on the actual configuration and task records. Do not set this parameter to true during CustomSSO initialization. You can only update it later by calling UpdateIdentityProviderSCIM, and determine the result based on operationStatus.
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
	// The URL configured for the external provisioning system. The presence of a URL does not indicate that external configuration is complete or that AC synchronization is enabled.
	//
	// example:
	//
	// https://example.com/scim/v2
	BaseUrl *string `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	// The AgentIdentity receiving status. Enabled indicates that SCIM receiving is enabled. Disabled indicates that it is disabled. This is not equivalent to the expected value of syncEnabled for AC.
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
