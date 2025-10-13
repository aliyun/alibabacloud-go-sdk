// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublishedServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPublishedServicesResponseBody
	GetCode() *string
	SetData(v []*ListPublishedServicesResponseBodyData) *ListPublishedServicesResponseBody
	GetData() []*ListPublishedServicesResponseBodyData
	SetErrorCode(v string) *ListPublishedServicesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListPublishedServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPublishedServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPublishedServicesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListPublishedServicesResponseBody
	GetTraceId() *string
}

type ListPublishedServicesResponseBody struct {
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
	Data []*ListPublishedServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s ListPublishedServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPublishedServicesResponseBody) GetData() []*ListPublishedServicesResponseBodyData {
	return s.Data
}

func (s *ListPublishedServicesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPublishedServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPublishedServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublishedServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPublishedServicesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListPublishedServicesResponseBody) SetCode(v string) *ListPublishedServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetData(v []*ListPublishedServicesResponseBodyData) *ListPublishedServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListPublishedServicesResponseBody) SetErrorCode(v string) *ListPublishedServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetMessage(v string) *ListPublishedServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetRequestId(v string) *ListPublishedServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetSuccess(v bool) *ListPublishedServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPublishedServicesResponseBody) SetTraceId(v string) *ListPublishedServicesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListPublishedServicesResponseBody) Validate() error {
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

type ListPublishedServicesResponseBodyData struct {
	// The ID of the application.
	//
	// example:
	//
	// b2a8a925-477a-4ed7-b825-d5e22500****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The reserved parameter. This parameter does not take effect.
	//
	// example:
	//
	// {}
	Group2Ip *string `json:"Group2Ip,omitempty" xml:"Group2Ip,omitempty"`
	// The service group that corresponds to the consumed service.
	Groups []*string `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The addresses where services can be subscribed to.
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
	// The version of the published services.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListPublishedServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPublishedServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPublishedServicesResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListPublishedServicesResponseBodyData) GetGroup2Ip() *string {
	return s.Group2Ip
}

func (s *ListPublishedServicesResponseBodyData) GetGroups() []*string {
	return s.Groups
}

func (s *ListPublishedServicesResponseBodyData) GetIps() []*string {
	return s.Ips
}

func (s *ListPublishedServicesResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListPublishedServicesResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListPublishedServicesResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListPublishedServicesResponseBodyData) SetAppId(v string) *ListPublishedServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetGroup2Ip(v string) *ListPublishedServicesResponseBodyData {
	s.Group2Ip = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetGroups(v []*string) *ListPublishedServicesResponseBodyData {
	s.Groups = v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetIps(v []*string) *ListPublishedServicesResponseBodyData {
	s.Ips = v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetName(v string) *ListPublishedServicesResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetType(v string) *ListPublishedServicesResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) SetVersion(v string) *ListPublishedServicesResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListPublishedServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
