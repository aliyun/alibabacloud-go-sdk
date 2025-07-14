// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateNamespaceResponseBody
	GetCode() *string
	SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody
	GetData() *CreateNamespaceResponseBodyData
	SetErrorCode(v string) *CreateNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNamespaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateNamespaceResponseBody
	GetTraceId() *string
}

type CreateNamespaceResponseBody struct {
	// bucketPath
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about a namespace.
	Data *CreateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// http://sae_pop_pre/#vpc
	//
	// example:
	//
	// bucketName
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The description of the custom namespace.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// mountDir
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateNamespaceResponseBody) GetData() *CreateNamespaceResponseBodyData {
	return s.Data
}

func (s *CreateNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNamespaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateNamespaceResponseBody) SetCode(v string) *CreateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetData(v *CreateNamespaceResponseBodyData) *CreateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateNamespaceResponseBody) SetErrorCode(v string) *CreateNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetMessage(v string) *CreateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetRequestId(v string) *CreateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetSuccess(v bool) *CreateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNamespaceResponseBody) SetTraceId(v string) *CreateNamespaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateNamespaceResponseBodyData struct {
	// Indicates whether the SAE built-in registry is enabled:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// Indicates whether the namespace was created. Valid values:
	//
	// 	- **true**: The instance was created.
	//
	// 	- **false**: The call failed to be created.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The short ID of the namespace.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The error code returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// Null
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
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
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateNamespaceResponseBodyData) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *CreateNamespaceResponseBodyData) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *CreateNamespaceResponseBodyData) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *CreateNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *CreateNamespaceResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *CreateNamespaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateNamespaceResponseBodyData) SetEnableMicroRegistration(v bool) *CreateNamespaceResponseBodyData {
	s.EnableMicroRegistration = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNameSpaceShortId(v string) *CreateNamespaceResponseBodyData {
	s.NameSpaceShortId = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceDescription(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceId(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetNamespaceName(v string) *CreateNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) SetRegionId(v string) *CreateNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
