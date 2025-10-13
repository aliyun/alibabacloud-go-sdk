// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEdasContainersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeEdasContainersResponseBody
	GetCode() *string
	SetData(v []*DescribeEdasContainersResponseBodyData) *DescribeEdasContainersResponseBody
	GetData() []*DescribeEdasContainersResponseBodyData
	SetErrorCode(v string) *DescribeEdasContainersResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeEdasContainersResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEdasContainersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEdasContainersResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeEdasContainersResponseBody
	GetTraceId() *string
}

type DescribeEdasContainersResponseBody struct {
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
	// The components.
	Data []*DescribeEdasContainersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message.
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
	// Indicates whether the list of container components of a microservices application was obtained. Valid values:
	//
	// 	- **true**: The list was obtained.
	//
	// 	- **false**: The list failed to be obtained.
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

func (s DescribeEdasContainersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdasContainersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeEdasContainersResponseBody) GetData() []*DescribeEdasContainersResponseBodyData {
	return s.Data
}

func (s *DescribeEdasContainersResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeEdasContainersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEdasContainersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEdasContainersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEdasContainersResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeEdasContainersResponseBody) SetCode(v string) *DescribeEdasContainersResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetData(v []*DescribeEdasContainersResponseBodyData) *DescribeEdasContainersResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetErrorCode(v string) *DescribeEdasContainersResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetMessage(v string) *DescribeEdasContainersResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetRequestId(v string) *DescribeEdasContainersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetSuccess(v bool) *DescribeEdasContainersResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) SetTraceId(v string) *DescribeEdasContainersResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeEdasContainersResponseBody) Validate() error {
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

type DescribeEdasContainersResponseBodyData struct {
	// Indicates whether the component is disabled. Valid values:
	//
	// 	- **true**: The component is disabled.
	//
	// 	- **false**: The component is not disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The version of the container, such as Ali-Tomcat, in which an application that is developed based on High-speed Service Framework (HSF) is deployed.
	//
	// example:
	//
	// 3.5.3
	EdasContainerVersion *string `json:"EdasContainerVersion,omitempty" xml:"EdasContainerVersion,omitempty"`
}

func (s DescribeEdasContainersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEdasContainersResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEdasContainersResponseBodyData) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeEdasContainersResponseBodyData) GetEdasContainerVersion() *string {
	return s.EdasContainerVersion
}

func (s *DescribeEdasContainersResponseBodyData) SetDisabled(v bool) *DescribeEdasContainersResponseBodyData {
	s.Disabled = &v
	return s
}

func (s *DescribeEdasContainersResponseBodyData) SetEdasContainerVersion(v string) *DescribeEdasContainersResponseBodyData {
	s.EdasContainerVersion = &v
	return s
}

func (s *DescribeEdasContainersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
