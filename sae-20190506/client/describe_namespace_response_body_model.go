// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeNamespaceResponseBody
	GetCode() *string
	SetData(v *DescribeNamespaceResponseBodyData) *DescribeNamespaceResponseBody
	GetData() *DescribeNamespaceResponseBodyData
	SetErrorCode(v string) *DescribeNamespaceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeNamespaceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeNamespaceResponseBody
	GetTraceId() *string
}

type DescribeNamespaceResponseBody struct {
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
	// The information about the namespace.
	Data *DescribeNamespaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned if the request fails. For more information, see the **Error codes*	- section in this topic.
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
	// Indicates whether the information about the namespace was queried successfully. Valid values:
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
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeNamespaceResponseBody) GetData() *DescribeNamespaceResponseBodyData {
	return s.Data
}

func (s *DescribeNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeNamespaceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeNamespaceResponseBody) SetCode(v string) *DescribeNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetData(v *DescribeNamespaceResponseBodyData) *DescribeNamespaceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeNamespaceResponseBody) SetErrorCode(v string) *DescribeNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetMessage(v string) *DescribeNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetRequestId(v string) *DescribeNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetSuccess(v bool) *DescribeNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeNamespaceResponseBody) SetTraceId(v string) *DescribeNamespaceResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeNamespaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNamespaceResponseBodyData struct {
	// Indicates whether the SAE built-in registry is enabled. Valid values:
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
	// The ID of the namespace. The information about the default namespace cannot be queried or modified. The default namespace cannot be deleted.
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
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeNamespaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeNamespaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeNamespaceResponseBodyData) GetEnableMicroRegistration() *bool {
	return s.EnableMicroRegistration
}

func (s *DescribeNamespaceResponseBodyData) GetNameSpaceShortId() *string {
	return s.NameSpaceShortId
}

func (s *DescribeNamespaceResponseBodyData) GetNamespaceDescription() *string {
	return s.NamespaceDescription
}

func (s *DescribeNamespaceResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *DescribeNamespaceResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DescribeNamespaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNamespaceResponseBodyData) SetEnableMicroRegistration(v bool) *DescribeNamespaceResponseBodyData {
	s.EnableMicroRegistration = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNameSpaceShortId(v string) *DescribeNamespaceResponseBodyData {
	s.NameSpaceShortId = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceDescription(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceDescription = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceId(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetNamespaceName(v string) *DescribeNamespaceResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) SetRegionId(v string) *DescribeNamespaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *DescribeNamespaceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
