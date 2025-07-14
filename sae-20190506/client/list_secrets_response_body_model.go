// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSecretsResponseBody
	GetCode() *string
	SetData(v *ListSecretsResponseBodyData) *ListSecretsResponseBody
	GetData() *ListSecretsResponseBodyData
	SetErrorCode(v string) *ListSecretsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListSecretsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSecretsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSecretsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListSecretsResponseBody
	GetTraceId() *string
}

type ListSecretsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ListSecretsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Take note of the following rules:
	//
	// 	- If the call is successful, **success*	- is returned.
	//
	// 	- If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSecretsResponseBody) GetData() *ListSecretsResponseBodyData {
	return s.Data
}

func (s *ListSecretsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSecretsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSecretsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSecretsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListSecretsResponseBody) SetCode(v string) *ListSecretsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSecretsResponseBody) SetData(v *ListSecretsResponseBodyData) *ListSecretsResponseBody {
	s.Data = v
	return s
}

func (s *ListSecretsResponseBody) SetErrorCode(v string) *ListSecretsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSecretsResponseBody) SetMessage(v string) *ListSecretsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSecretsResponseBody) SetRequestId(v string) *ListSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecretsResponseBody) SetSuccess(v bool) *ListSecretsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSecretsResponseBody) SetTraceId(v string) *ListSecretsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListSecretsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSecretsResponseBodyData struct {
	// The Secrets.
	Secrets []*ListSecretsResponseBodyDataSecrets `json:"Secrets,omitempty" xml:"Secrets,omitempty" type:"Repeated"`
}

func (s ListSecretsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodyData) GetSecrets() []*ListSecretsResponseBodyDataSecrets {
	return s.Secrets
}

func (s *ListSecretsResponseBodyData) SetSecrets(v []*ListSecretsResponseBodyDataSecrets) *ListSecretsResponseBodyData {
	s.Secrets = v
	return s
}

func (s *ListSecretsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSecretsResponseBodyDataSecrets struct {
	// The time when the Secret was created.
	//
	// example:
	//
	// 1593760185111
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The associated applications.
	RelateApps []*ListSecretsResponseBodyDataSecretsRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	// The Secret ID.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The Secret name.
	//
	// example:
	//
	// registry-auth
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The Secret type.
	//
	// Set the value to **kubernetes.io/dockerconfigjson**. The value indicates the secret for the username and password of the image repository and is used for authentication when images are pulled during application deployment.
	//
	// example:
	//
	// kubernetes.io/dockerconfigjson
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The time when the Secret was updated.
	//
	// example:
	//
	// 1593760185111
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListSecretsResponseBodyDataSecrets) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodyDataSecrets) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodyDataSecrets) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListSecretsResponseBodyDataSecrets) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSecretsResponseBodyDataSecrets) GetRelateApps() []*ListSecretsResponseBodyDataSecretsRelateApps {
	return s.RelateApps
}

func (s *ListSecretsResponseBodyDataSecrets) GetSecretId() *int64 {
	return s.SecretId
}

func (s *ListSecretsResponseBodyDataSecrets) GetSecretName() *string {
	return s.SecretName
}

func (s *ListSecretsResponseBodyDataSecrets) GetSecretType() *string {
	return s.SecretType
}

func (s *ListSecretsResponseBodyDataSecrets) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListSecretsResponseBodyDataSecrets) SetCreateTime(v int64) *ListSecretsResponseBodyDataSecrets {
	s.CreateTime = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetNamespaceId(v string) *ListSecretsResponseBodyDataSecrets {
	s.NamespaceId = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetRelateApps(v []*ListSecretsResponseBodyDataSecretsRelateApps) *ListSecretsResponseBodyDataSecrets {
	s.RelateApps = v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetSecretId(v int64) *ListSecretsResponseBodyDataSecrets {
	s.SecretId = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetSecretName(v string) *ListSecretsResponseBodyDataSecrets {
	s.SecretName = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetSecretType(v string) *ListSecretsResponseBodyDataSecrets {
	s.SecretType = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) SetUpdateTime(v int64) *ListSecretsResponseBodyDataSecrets {
	s.UpdateTime = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecrets) Validate() error {
	return dara.Validate(s)
}

type ListSecretsResponseBodyDataSecretsRelateApps struct {
	// The application ID.
	//
	// example:
	//
	// f16b4000-9058-4c22-a49d-49a28f0b****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application name.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s ListSecretsResponseBodyDataSecretsRelateApps) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsResponseBodyDataSecretsRelateApps) GoString() string {
	return s.String()
}

func (s *ListSecretsResponseBodyDataSecretsRelateApps) GetAppId() *string {
	return s.AppId
}

func (s *ListSecretsResponseBodyDataSecretsRelateApps) GetAppName() *string {
	return s.AppName
}

func (s *ListSecretsResponseBodyDataSecretsRelateApps) SetAppId(v string) *ListSecretsResponseBodyDataSecretsRelateApps {
	s.AppId = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecretsRelateApps) SetAppName(v string) *ListSecretsResponseBodyDataSecretsRelateApps {
	s.AppName = &v
	return s
}

func (s *ListSecretsResponseBodyDataSecretsRelateApps) Validate() error {
	return dara.Validate(s)
}
