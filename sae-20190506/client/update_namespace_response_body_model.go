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
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The namespace information.
	Data *UpdateNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. The following rules apply:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidNamespaceName.Format
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information about the request result.
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
	// Indicates whether the namespace information is updated. Valid values:
	//
	// - **true**: The update is successful.
	//
	// - **false**: The update fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID. You can use this ID to precisely query invocation information.
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
	// Indicates whether the built-in registry of SAE is enabled. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	EnableMicroRegistration *bool `json:"EnableMicroRegistration,omitempty" xml:"EnableMicroRegistration,omitempty"`
	// The short namespace ID.
	//
	// example:
	//
	// test
	NameSpaceShortId *string `json:"NameSpaceShortId,omitempty" xml:"NameSpaceShortId,omitempty"`
	// The namespace description.
	//
	// example:
	//
	// desc
	NamespaceDescription *string `json:"NamespaceDescription,omitempty" xml:"NamespaceDescription,omitempty"`
	// The namespace ID.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The namespace name.
	//
	// example:
	//
	// name
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The region where the namespace is located.
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
