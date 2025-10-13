// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateNamespaceResponseBody
	GetCode() *string
	SetData(v *UpdateNamespaceResponseBodyData) *UpdateNamespaceResponseBody
	GetData() *UpdateNamespaceResponseBodyData
	SetErrorCode(v string) *UpdateNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpdateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNamespaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpdateNamespaceResponseBody
	GetTraceId() *string
}

type UpdateNamespaceResponseBody struct {
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
	// The information about a namespace.
	Data *UpdateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned for the operation.
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
	// Indicates whether the information about the namespace was updated. Valid values:
	//
	// 	- **true**: The instance was updated.
	//
	// 	- **false**: The instance failed to be updated.
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

func (s UpdateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateNamespaceResponseBody) GetData() *UpdateNamespaceResponseBodyData {
	return s.Data
}

func (s *UpdateNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNamespaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpdateNamespaceResponseBody) SetCode(v string) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetData(v *UpdateNamespaceResponseBodyData) *UpdateNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateNamespaceResponseBody) SetErrorCode(v string) *UpdateNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetMessage(v string) *UpdateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetSuccess(v bool) *UpdateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetTraceId(v string) *UpdateNamespaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateNamespaceResponseBodyData struct {
	// Indicates whether to enable SAE built-in registry:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The short ID of the namespace.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The description of the namespace.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The region where the namespace resides.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBodyData) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *UpdateNamespaceResponseBodyData) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *UpdateNamespaceResponseBodyData) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *UpdateNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UpdateNamespaceResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *UpdateNamespaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateNamespaceResponseBodyData) SetEnableMicroRegistration(v bool) *UpdateNamespaceResponseBodyData {
	s.EnableMicroRegistration = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNameSpaceShortId(v string) *UpdateNamespaceResponseBodyData {
	s.NameSpaceShortId = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceDescription(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceId(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetNamespaceName(v string) *UpdateNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) SetRegionId(v string) *UpdateNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
