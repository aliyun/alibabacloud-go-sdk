// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeComponentsResponseBody
	GetCode() *string
	SetData(v []*DescribeComponentsResponseBodyData) *DescribeComponentsResponseBody
	GetData() []*DescribeComponentsResponseBodyData
	SetErrorCode(v string) *DescribeComponentsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeComponentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeComponentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeComponentsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeComponentsResponseBody
	GetTraceId() *string
}

type DescribeComponentsResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the supported components.
	Data []*DescribeComponentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The status code. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned.
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
	// Indicates whether the component version was obtained. Valid values:
	//
	// 	- **true**: The applications were obtained.
	//
	// 	- **false**: The applications failed to be queried.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeComponentsResponseBody) GetData() []*DescribeComponentsResponseBodyData {
	return s.Data
}

func (s *DescribeComponentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeComponentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComponentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeComponentsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeComponentsResponseBody) SetCode(v string) *DescribeComponentsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetData(v []*DescribeComponentsResponseBodyData) *DescribeComponentsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeComponentsResponseBody) SetErrorCode(v string) *DescribeComponentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetMessage(v string) *DescribeComponentsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetRequestId(v string) *DescribeComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetSuccess(v bool) *DescribeComponentsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeComponentsResponseBody) SetTraceId(v string) *DescribeComponentsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeComponentsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeComponentsResponseBodyData struct {
	// The description of the component.
	//
	// example:
	//
	// Open JDK 8
	ComponentDescription *string `json:"ComponentDescription,omitempty" xml:"ComponentDescription,omitempty"`
	// The component ID.
	//
	// example:
	//
	// Open JDK 8
	ComponentKey *string `json:"ComponentKey,omitempty" xml:"ComponentKey,omitempty"`
	// Indicates whether the component is expired. Valid values:
	//
	// 	- **true**: The component is expired.
	//
	// 	- **false**: The component is not expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The type of the component.
	//
	// example:
	//
	// JDK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeComponentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeComponentsResponseBodyData) GetComponentDescription() *string {
	return s.ComponentDescription
}

func (s *DescribeComponentsResponseBodyData) GetComponentKey() *string {
	return s.ComponentKey
}

func (s *DescribeComponentsResponseBodyData) GetExpired() *bool {
	return s.Expired
}

func (s *DescribeComponentsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *DescribeComponentsResponseBodyData) SetComponentDescription(v string) *DescribeComponentsResponseBodyData {
	s.ComponentDescription = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetComponentKey(v string) *DescribeComponentsResponseBodyData {
	s.ComponentKey = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetExpired(v bool) *DescribeComponentsResponseBodyData {
	s.Expired = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) SetType(v string) *DescribeComponentsResponseBodyData {
	s.Type = &v
	return s
}

func (s *DescribeComponentsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
