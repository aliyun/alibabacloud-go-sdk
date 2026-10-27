// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProviderSCIMCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIdentityProviderSCIMCredentialsResponseBody
	GetCode() *string
	SetData(v []*ListIdentityProviderSCIMCredentialsResponseBodyData) *ListIdentityProviderSCIMCredentialsResponseBody
	GetData() []*ListIdentityProviderSCIMCredentialsResponseBodyData
	SetHttpStatusCode(v int32) *ListIdentityProviderSCIMCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListIdentityProviderSCIMCredentialsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIdentityProviderSCIMCredentialsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIdentityProviderSCIMCredentialsResponseBody
	GetSuccess() *bool
}

type ListIdentityProviderSCIMCredentialsResponseBody struct {
	// The business result code. If the request fails, handle the error based on the HTTP status code and the message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result. For the result structure, see the response fields of this operation. For asynchronous write operations, the response indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data []*ListIdentityProviderSCIMCredentialsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code of the current request. This code indicates the processing status of the request and is not equivalent to the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Task result message, which indicates that the request is accepted or contains an error message. For task results of asynchronous tasks, check the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to associate the current call with troubleshooting. Retain this ID if the request fails.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the current request is successful. A value of true indicates that the request is processed successfully. For asynchronous write operations, you still need to query the task results.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListIdentityProviderSCIMCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderSCIMCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetData() []*ListIdentityProviderSCIMCredentialsResponseBodyData {
	return s.Data
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetCode(v string) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetData(v []*ListIdentityProviderSCIMCredentialsResponseBodyData) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetHttpStatusCode(v int32) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetMessage(v string) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetRequestId(v string) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) SetSuccess(v bool) *ListIdentityProviderSCIMCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProviderSCIMCredentialsResponseBodyData struct {
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
	// The expiration time returned by AgentIdentity. Only credentials with a status of Enabled that have not expired are considered available.
	//
	// example:
	//
	// 2027-10-26T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The credential status of AgentIdentity. Only credentials with a status of Enabled that have not expired are considered available. This field is not the operationStatus of a SCIM task.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListIdentityProviderSCIMCredentialsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProviderSCIMCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) SetCreateTime(v string) *ListIdentityProviderSCIMCredentialsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) SetCredentialId(v string) *ListIdentityProviderSCIMCredentialsResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) SetExpireTime(v string) *ListIdentityProviderSCIMCredentialsResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) SetStatus(v string) *ListIdentityProviderSCIMCredentialsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListIdentityProviderSCIMCredentialsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
