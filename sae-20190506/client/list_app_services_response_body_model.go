// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAppServicesResponseBody
	GetCode() *string
	SetData(v []*ListAppServicesResponseBodyData) *ListAppServicesResponseBody
	GetData() []*ListAppServicesResponseBodyData
	SetErrorCode(v string) *ListAppServicesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ListAppServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAppServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAppServicesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ListAppServicesResponseBody
	GetTraceId() *string
}

type ListAppServicesResponseBody struct {
	// The HTTP status code that is returned. Valid values:
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
	// The details of the microservice.
	Data []*ListAppServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The status code. Valid values:
	//
	// 	- If the request was successful, the **ErrorCode*	- parameter is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. Valid values:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- If the request failed, an error message is returned.
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
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
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

func (s ListAppServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAppServicesResponseBody) GetData() []*ListAppServicesResponseBodyData {
	return s.Data
}

func (s *ListAppServicesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListAppServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAppServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAppServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAppServicesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ListAppServicesResponseBody) SetCode(v string) *ListAppServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListAppServicesResponseBody) SetData(v []*ListAppServicesResponseBodyData) *ListAppServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListAppServicesResponseBody) SetErrorCode(v string) *ListAppServicesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListAppServicesResponseBody) SetMessage(v string) *ListAppServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListAppServicesResponseBody) SetRequestId(v string) *ListAppServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppServicesResponseBody) SetSuccess(v bool) *ListAppServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ListAppServicesResponseBody) SetTraceId(v string) *ListAppServicesResponseBody {
	s.TraceId = &v
	return s
}

func (s *ListAppServicesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListAppServicesResponseBodyData struct {
	// The application ID.
	//
	// example:
	//
	// 0099b7be-5f5b-4512-a7fc-56049ef1****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// demo-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The number of instances of the microservice.
	//
	// example:
	//
	// 1
	InstanceCount *string `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The ID of the namespace to which the application belongs.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
	// The registry type. Valid values:
	//
	// 	- **0**：SAE Nacos
	//
	// 	- **1**: SAE built-in Nacos
	//
	// 	- **2**: MSE Nacos
	//
	// 	- **9**: SAE Kubernets service
	//
	// example:
	//
	// 0
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The IDs of the security groups.
	//
	// example:
	//
	// sg-wz969ngg2e49q5i4****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The group to which the microservice belongs.
	//
	// example:
	//
	// DEFAULT_GROUP
	ServiceGroup *string `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	// The name of the microservice.
	//
	// example:
	//
	// frontend
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The ports and protocols.
	ServicePortAndProtocol map[string]*string `json:"ServicePortAndProtocol,omitempty" xml:"ServicePortAndProtocol,omitempty"`
	// The list of ports.
	ServicePorts []*int32 `json:"ServicePorts,omitempty" xml:"ServicePorts,omitempty" type:"Repeated"`
	// The protocol used by the microservice.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The type of the microservice. Valid values:
	//
	// 	- **dubbo**
	//
	// 	- **springCloud**
	//
	// 	- **hsf**
	//
	// 	- **k8sService**
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The version of the microservice.
	//
	// example:
	//
	// 1.0.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s ListAppServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAppServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAppServicesResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *ListAppServicesResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *ListAppServicesResponseBodyData) GetInstanceCount() *string {
	return s.InstanceCount
}

func (s *ListAppServicesResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListAppServicesResponseBodyData) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *ListAppServicesResponseBodyData) GetRegistryType() *string {
	return s.RegistryType
}

func (s *ListAppServicesResponseBodyData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListAppServicesResponseBodyData) GetServiceGroup() *string {
	return s.ServiceGroup
}

func (s *ListAppServicesResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListAppServicesResponseBodyData) GetServicePortAndProtocol() map[string]*string {
	return s.ServicePortAndProtocol
}

func (s *ListAppServicesResponseBodyData) GetServicePorts() []*int32 {
	return s.ServicePorts
}

func (s *ListAppServicesResponseBodyData) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *ListAppServicesResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *ListAppServicesResponseBodyData) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *ListAppServicesResponseBodyData) SetAppId(v string) *ListAppServicesResponseBodyData {
	s.AppId = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetAppName(v string) *ListAppServicesResponseBodyData {
	s.AppName = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetInstanceCount(v string) *ListAppServicesResponseBodyData {
	s.InstanceCount = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetNamespaceId(v string) *ListAppServicesResponseBodyData {
	s.NamespaceId = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetNamespaceName(v string) *ListAppServicesResponseBodyData {
	s.NamespaceName = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetRegistryType(v string) *ListAppServicesResponseBodyData {
	s.RegistryType = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetSecurityGroupId(v string) *ListAppServicesResponseBodyData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServiceGroup(v string) *ListAppServicesResponseBodyData {
	s.ServiceGroup = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServiceName(v string) *ListAppServicesResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServicePortAndProtocol(v map[string]*string) *ListAppServicesResponseBodyData {
	s.ServicePortAndProtocol = v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServicePorts(v []*int32) *ListAppServicesResponseBodyData {
	s.ServicePorts = v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServiceProtocol(v string) *ListAppServicesResponseBodyData {
	s.ServiceProtocol = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServiceType(v string) *ListAppServicesResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *ListAppServicesResponseBodyData) SetServiceVersion(v string) *ListAppServicesResponseBodyData {
	s.ServiceVersion = &v
	return s
}

func (s *ListAppServicesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
