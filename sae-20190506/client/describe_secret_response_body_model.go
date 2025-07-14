// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecretResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSecretResponseBody
	GetCode() *string
	SetData(v *DescribeSecretResponseBodyData) *DescribeSecretResponseBody
	GetData() *DescribeSecretResponseBodyData
	SetErrorCode(v string) *DescribeSecretResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeSecretResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSecretResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSecretResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeSecretResponseBody
	GetTraceId() *string
}

type DescribeSecretResponseBody struct {
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
	// The response.
	Data *DescribeSecretResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Valid values:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the details of the Secret instance are successfully queried. Valid values:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The image failed to be found.
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

func (s DescribeSecretResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSecretResponseBody) GetData() *DescribeSecretResponseBodyData {
	return s.Data
}

func (s *DescribeSecretResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSecretResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSecretResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecretResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSecretResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeSecretResponseBody) SetCode(v string) *DescribeSecretResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSecretResponseBody) SetData(v *DescribeSecretResponseBodyData) *DescribeSecretResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecretResponseBody) SetErrorCode(v string) *DescribeSecretResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSecretResponseBody) SetMessage(v string) *DescribeSecretResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecretResponseBody) SetRequestId(v string) *DescribeSecretResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecretResponseBody) SetSuccess(v bool) *DescribeSecretResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecretResponseBody) SetTraceId(v string) *DescribeSecretResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeSecretResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecretResponseBodyData struct {
	// The time when the task was created.
	//
	// example:
	//
	// 1593746835111
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The associated application.
	RelateApps []*DescribeSecretResponseBodyDataRelateApps `json:"RelateApps,omitempty" xml:"RelateApps,omitempty" type:"Repeated"`
	// Secret key-value pair data.
	SecretData map[string]*string `json:"SecretData,omitempty" xml:"SecretData,omitempty"`
	// The ID of the Secret instance.
	//
	// example:
	//
	// 16
	SecretId *int64 `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The name of the Secret instance.
	//
	// example:
	//
	// registry-auth
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The type of the Secret instance.
	//
	// example:
	//
	// kubernetes.io/dockerconfigjson
	SecretType *string `json:"SecretType,omitempty" xml:"SecretType,omitempty"`
	// The time when the task was updated.
	//
	// example:
	//
	// 1593746835111
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSecretResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeSecretResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeSecretResponseBodyData) GetRelateApps() []*DescribeSecretResponseBodyDataRelateApps {
	return s.RelateApps
}

func (s *DescribeSecretResponseBodyData) GetSecretData() map[string]*string {
	return s.SecretData
}

func (s *DescribeSecretResponseBodyData) GetSecretId() *int64 {
	return s.SecretId
}

func (s *DescribeSecretResponseBodyData) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeSecretResponseBodyData) GetSecretType() *string {
	return s.SecretType
}

func (s *DescribeSecretResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeSecretResponseBodyData) SetCreateTime(v int64) *DescribeSecretResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecretResponseBodyData) SetNamespaceId(v string) *DescribeSecretResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeSecretResponseBodyData) SetRelateApps(v []*DescribeSecretResponseBodyDataRelateApps) *DescribeSecretResponseBodyData {
	s.RelateApps = v
	return s
}

func (s *DescribeSecretResponseBodyData) SetSecretData(v map[string]*string) *DescribeSecretResponseBodyData {
	s.SecretData = v
	return s
}

func (s *DescribeSecretResponseBodyData) SetSecretId(v int64) *DescribeSecretResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *DescribeSecretResponseBodyData) SetSecretName(v string) *DescribeSecretResponseBodyData {
	s.SecretName = &v
	return s
}

func (s *DescribeSecretResponseBodyData) SetSecretType(v string) *DescribeSecretResponseBodyData {
	s.SecretType = &v
	return s
}

func (s *DescribeSecretResponseBodyData) SetUpdateTime(v int64) *DescribeSecretResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *DescribeSecretResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeSecretResponseBodyDataRelateApps struct {
	// The application ID.
	//
	// example:
	//
	// f16b4000-9058-4c22-a49d-49a28f0b****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
}

func (s DescribeSecretResponseBodyDataRelateApps) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecretResponseBodyDataRelateApps) GoString() string {
	return s.String()
}

func (s *DescribeSecretResponseBodyDataRelateApps) GetAppId() *string {
	return s.AppId
}

func (s *DescribeSecretResponseBodyDataRelateApps) GetAppName() *string {
	return s.AppName
}

func (s *DescribeSecretResponseBodyDataRelateApps) SetAppId(v string) *DescribeSecretResponseBodyDataRelateApps {
	s.AppId = &v
	return s
}

func (s *DescribeSecretResponseBodyDataRelateApps) SetAppName(v string) *DescribeSecretResponseBodyDataRelateApps {
	s.AppName = &v
	return s
}

func (s *DescribeSecretResponseBodyDataRelateApps) Validate() error {
	return dara.Validate(s)
}
