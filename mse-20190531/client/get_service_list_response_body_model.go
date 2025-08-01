// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceListResponseBody
	GetCode() *int32
	SetData(v []*GetServiceListResponseBodyData) *GetServiceListResponseBody
	GetData() []*GetServiceListResponseBodyData
	SetHttpStatusCode(v int32) *GetServiceListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetServiceListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceListResponseBody
	GetSuccess() *bool
}

type GetServiceListResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data entries returned.
	Data []*GetServiceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// xxx-xxx-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceListResponseBody) GetData() []*GetServiceListResponseBodyData {
	return s.Data
}

func (s *GetServiceListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetServiceListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceListResponseBody) SetCode(v int32) *GetServiceListResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceListResponseBody) SetData(v []*GetServiceListResponseBodyData) *GetServiceListResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceListResponseBody) SetHttpStatusCode(v int32) *GetServiceListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetServiceListResponseBody) SetMessage(v string) *GetServiceListResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceListResponseBody) SetRequestId(v string) *GetServiceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceListResponseBody) SetSuccess(v bool) *GetServiceListResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetServiceListResponseBodyData struct {
	// The name of the Dubbo application.
	//
	// example:
	//
	// dubbo-application
	DubboApplicationName *string `json:"DubboApplicationName,omitempty" xml:"DubboApplicationName,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// test
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	// The group.
	//
	// example:
	//
	// dubbo
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The metadata.
	//
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The methods.
	Methods []*GetServiceListResponseBodyDataMethods `json:"Methods,omitempty" xml:"Methods,omitempty" type:"Repeated"`
	// The type of the service registry.
	//
	// example:
	//
	// nacos
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// com.alibaba.xxx
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// dubbo
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The name of the Spring application.
	//
	// example:
	//
	// spring-application
	SpringApplicationName *string `json:"SpringApplicationName,omitempty" xml:"SpringApplicationName,omitempty"`
	// The version information.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceListResponseBodyData) GetDubboApplicationName() *string {
	return s.DubboApplicationName
}

func (s *GetServiceListResponseBodyData) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *GetServiceListResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *GetServiceListResponseBodyData) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *GetServiceListResponseBodyData) GetMethods() []*GetServiceListResponseBodyDataMethods {
	return s.Methods
}

func (s *GetServiceListResponseBodyData) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceListResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceListResponseBodyData) GetSpringApplicationName() *string {
	return s.SpringApplicationName
}

func (s *GetServiceListResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetServiceListResponseBodyData) SetDubboApplicationName(v string) *GetServiceListResponseBodyData {
	s.DubboApplicationName = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetEdasAppName(v string) *GetServiceListResponseBodyData {
	s.EdasAppName = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetGroup(v string) *GetServiceListResponseBodyData {
	s.Group = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetMetadata(v map[string]interface{}) *GetServiceListResponseBodyData {
	s.Metadata = v
	return s
}

func (s *GetServiceListResponseBodyData) SetMethods(v []*GetServiceListResponseBodyDataMethods) *GetServiceListResponseBodyData {
	s.Methods = v
	return s
}

func (s *GetServiceListResponseBodyData) SetRegistryType(v string) *GetServiceListResponseBodyData {
	s.RegistryType = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetServiceName(v string) *GetServiceListResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetServiceType(v string) *GetServiceListResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetSpringApplicationName(v string) *GetServiceListResponseBodyData {
	s.SpringApplicationName = &v
	return s
}

func (s *GetServiceListResponseBodyData) SetVersion(v string) *GetServiceListResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetServiceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetServiceListResponseBodyDataMethods struct {
	// The controller of the method.
	//
	// example:
	//
	// com.alibaba.SayHelloController
	MethodController *string `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	// The name of the method.
	//
	// example:
	//
	// sayHello
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The data types of the parameters.
	ParameterTypes []*string `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty" type:"Repeated"`
	// The paths.
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The methods.
	RequestMethods []*string `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty" type:"Repeated"`
	// The type of the return value.
	//
	// example:
	//
	// int
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GetServiceListResponseBodyDataMethods) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListResponseBodyDataMethods) GoString() string {
	return s.String()
}

func (s *GetServiceListResponseBodyDataMethods) GetMethodController() *string {
	return s.MethodController
}

func (s *GetServiceListResponseBodyDataMethods) GetName() *string {
	return s.Name
}

func (s *GetServiceListResponseBodyDataMethods) GetParameterTypes() []*string {
	return s.ParameterTypes
}

func (s *GetServiceListResponseBodyDataMethods) GetPaths() []*string {
	return s.Paths
}

func (s *GetServiceListResponseBodyDataMethods) GetRequestMethods() []*string {
	return s.RequestMethods
}

func (s *GetServiceListResponseBodyDataMethods) GetReturnType() *string {
	return s.ReturnType
}

func (s *GetServiceListResponseBodyDataMethods) SetMethodController(v string) *GetServiceListResponseBodyDataMethods {
	s.MethodController = &v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) SetName(v string) *GetServiceListResponseBodyDataMethods {
	s.Name = &v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) SetParameterTypes(v []*string) *GetServiceListResponseBodyDataMethods {
	s.ParameterTypes = v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) SetPaths(v []*string) *GetServiceListResponseBodyDataMethods {
	s.Paths = v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) SetRequestMethods(v []*string) *GetServiceListResponseBodyDataMethods {
	s.RequestMethods = v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) SetReturnType(v string) *GetServiceListResponseBodyDataMethods {
	s.ReturnType = &v
	return s
}

func (s *GetServiceListResponseBodyDataMethods) Validate() error {
	return dara.Validate(s)
}
