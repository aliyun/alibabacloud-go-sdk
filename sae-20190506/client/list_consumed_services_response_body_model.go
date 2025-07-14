// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumedServicesResponseBody
	GetCode() *string
	SetData(v []*ListConsumedServicesResponseBodyData) *ListConsumedServicesResponseBody
	GetData() []*ListConsumedServicesResponseBodyData
	SetErrorCode(v string) *ListConsumedServicesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListConsumedServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumedServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListConsumedServicesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListConsumedServicesResponseBody
	GetTraceId() *string
}

type ListConsumedServicesResponseBody struct {
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
	// The details of the microservices.
	Data []*ListConsumedServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code. Valid values:
	//
	// 	- If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
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
	// Indicates whether the list of microservices was queried. Valid values:
	//
	// 	- **true**: The list was queried.
	//
	// 	- **false**: The list failed to be queried.
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

func (s ListConsumedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumedServicesResponseBody) GetData() []*ListConsumedServicesResponseBodyData {
	return s.Data
}

func (s *ListConsumedServicesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListConsumedServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumedServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListConsumedServicesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListConsumedServicesResponseBody) SetCode(v string) *ListConsumedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetData(v []*ListConsumedServicesResponseBodyData) *ListConsumedServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumedServicesResponseBody) SetErrorCode(v string) *ListConsumedServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetMessage(v string) *ListConsumedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetRequestId(v string) *ListConsumedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetSuccess(v bool) *ListConsumedServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListConsumedServicesResponseBody) SetTraceId(v string) *ListConsumedServicesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListConsumedServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListConsumedServicesResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// {}
	Group2Ip *string `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	// The service groups that corresponds to the consumed services.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The addresses where the services can be subscribed to.
	Ips []*string `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The name of the published service.
	//
	// example:
	//
	// com.alibaba.nodejs.ItemService
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the published service.
	//
	// example:
	//
	// RPC
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version of the published service.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListConsumedServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListConsumedServicesResponseBodyData) GetGroup2Ip() *string {
	return s.Group2Ip
}

func (s *ListConsumedServicesResponseBodyData) GetGroups() []*string {
	return s.Groups
}

func (s *ListConsumedServicesResponseBodyData) GetIps() []*string {
	return s.Ips
}

func (s *ListConsumedServicesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListConsumedServicesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListConsumedServicesResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListConsumedServicesResponseBodyData) SetAppId(v string) *ListConsumedServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetGroup2Ip(v string) *ListConsumedServicesResponseBodyData {
	s.Group2Ip = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetGroups(v []*string) *ListConsumedServicesResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetIps(v []*string) *ListConsumedServicesResponseBodyData {
	s.Ips = v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetName(v string) *ListConsumedServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetType(v string) *ListConsumedServicesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) SetVersion(v string) *ListConsumedServicesResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListConsumedServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
