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
	// The business result code. In case of failure, evaluate this code together with the HTTP status code and message.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operation result. For the result structure, see the response fields of this operation. The response of an asynchronous write operation indicates that the request is accepted, but does not guarantee that the background task is completed.
	Data []*ListIdentityProviderSCIMCredentialsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The processing status of the current request. This does not indicate the final result of the background task.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message indicating request acceptance or an error. For asynchronous task results, verify them by using the corresponding query operation.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID. Used to correlate this call for troubleshooting. Retain this ID in case of failure.
	//
	// example:
	//
	// request-example
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates that the request was processed successfully. For asynchronous writes, you must continue to query the task result.
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
	// The credential ID. This field is required and cannot be empty when deleting a credential. Use the ID returned in the credential list, not the plaintext token.
	//
	// example:
	//
	// credential-example
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// The expiration time returned by AgentIdentity. Only credentials with the Enabled status that have not expired are considered available.
	//
	// example:
	//
	// 2027-10-26T00:00:00Z
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// The credential status. Only credentials with the Enabled status that have not expired are considered available. This field is not the operationStatus of a SCIM task.
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
