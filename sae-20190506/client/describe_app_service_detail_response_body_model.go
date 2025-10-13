// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppServiceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAppServiceDetailResponseBody
	GetCode() *string
	SetData(v *DescribeAppServiceDetailResponseBodyData) *DescribeAppServiceDetailResponseBody
	GetData() *DescribeAppServiceDetailResponseBodyData
	SetErrorCode(v string) *DescribeAppServiceDetailResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeAppServiceDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAppServiceDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAppServiceDetailResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeAppServiceDetailResponseBody
	GetTraceId() *string
}

type DescribeAppServiceDetailResponseBody struct {
	// The HTTP status code. Valid values:
	//
	// - **2xx**: indicates that the call was successful.
	//
	// - **3xx**: indicates that the call was redirected.
	//
	// - **4xx**: indicates that the call failed.
	//
	// - **5xx**: indicates that a server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *DescribeAppServiceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error code. Valid values:
	//
	// - If the call is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The returned information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B2C7874F-F109-5B34-8618-2C10BBA2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the meta data was obtained. Valid values:
	//
	// 	- **true**: The metadata was obtained.
	//
	// 	- **false**: The metadata failed to be obtained.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0b16399316402420740034918e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeAppServiceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAppServiceDetailResponseBody) GetData() *DescribeAppServiceDetailResponseBodyData {
	return s.Data
}

func (s *DescribeAppServiceDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeAppServiceDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAppServiceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppServiceDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAppServiceDetailResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeAppServiceDetailResponseBody) SetCode(v string) *DescribeAppServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetData(v *DescribeAppServiceDetailResponseBodyData) *DescribeAppServiceDetailResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetErrorCode(v string) *DescribeAppServiceDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetMessage(v string) *DescribeAppServiceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetRequestId(v string) *DescribeAppServiceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetSuccess(v bool) *DescribeAppServiceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) SetTraceId(v string) *DescribeAppServiceDetailResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAppServiceDetailResponseBodyData struct {
	// The name of the Dubbo application.
	//
	// example:
	//
	// service-consumer
	DubboApplicationName *string `json:"DubboApplicationName,omitempty" xml:"DubboApplicationName,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// cn-zhangjiakou-micro-service-******
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	// The group to which the service belongs. You can create a custom group.
	//
	// example:
	//
	// springCloud
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The metadata. Example: `{side: "provider", port: "18081", preserved: {register: {source: "SPRING_CLOUD"}},…}`.
	//
	// example:
	//
	// {side: "provider", port: "18081", preserved: {register: {source: "SPRING_CLOUD"}},…}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The methods.
	Methods []*DescribeAppServiceDetailResponseBodyDataMethods `json:"Methods,omitempty" xml:"Methods,omitempty" type:"Repeated"`
	// The name of the service.
	//
	// example:
	//
	// service-provider
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The port used by the service.
	ServicePorts []*int64 `json:"ServicePorts,omitempty" xml:"ServicePorts,omitempty" type:"Repeated"`
	// The protocol used by the service.
	//
	// example:
	//
	// HTTP
	ServiceProtocol *string `json:"ServiceProtocol,omitempty" xml:"ServiceProtocol,omitempty"`
	// The tag of the service.
	ServiceTags []*string `json:"ServiceTags,omitempty" xml:"ServiceTags,omitempty" type:"Repeated"`
	// The type of the service. Valid values:
	//
	// 	- **dubbo**
	//
	// 	- **springCloud**
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The name of the Spring Cloud application.
	//
	// example:
	//
	// service-provider
	SpringApplicationName *string `json:"SpringApplicationName,omitempty" xml:"SpringApplicationName,omitempty"`
	// The version of the service. You can create a custom version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyData) GetDubboApplicationName() *string {
	return s.DubboApplicationName
}

func (s *DescribeAppServiceDetailResponseBodyData) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *DescribeAppServiceDetailResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *DescribeAppServiceDetailResponseBodyData) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *DescribeAppServiceDetailResponseBodyData) GetMethods() []*DescribeAppServiceDetailResponseBodyDataMethods {
	return s.Methods
}

func (s *DescribeAppServiceDetailResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *DescribeAppServiceDetailResponseBodyData) GetServicePorts() []*int64 {
	return s.ServicePorts
}

func (s *DescribeAppServiceDetailResponseBodyData) GetServiceProtocol() *string {
	return s.ServiceProtocol
}

func (s *DescribeAppServiceDetailResponseBodyData) GetServiceTags() []*string {
	return s.ServiceTags
}

func (s *DescribeAppServiceDetailResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *DescribeAppServiceDetailResponseBodyData) GetSpringApplicationName() *string {
	return s.SpringApplicationName
}

func (s *DescribeAppServiceDetailResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *DescribeAppServiceDetailResponseBodyData) SetDubboApplicationName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.DubboApplicationName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetEdasAppName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.EdasAppName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetGroup(v string) *DescribeAppServiceDetailResponseBodyData {
	s.Group = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetMetadata(v map[string]interface{}) *DescribeAppServiceDetailResponseBodyData {
	s.Metadata = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetMethods(v []*DescribeAppServiceDetailResponseBodyDataMethods) *DescribeAppServiceDetailResponseBodyData {
	s.Methods = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServicePorts(v []*int64) *DescribeAppServiceDetailResponseBodyData {
	s.ServicePorts = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceProtocol(v string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceProtocol = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceTags(v []*string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceTags = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetServiceType(v string) *DescribeAppServiceDetailResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetSpringApplicationName(v string) *DescribeAppServiceDetailResponseBodyData {
	s.SpringApplicationName = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) SetVersion(v string) *DescribeAppServiceDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyData) Validate() error {
	if s.Methods != nil {
		for _, item := range s.Methods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppServiceDetailResponseBodyDataMethods struct {
	// The class to which the method belongs.
	//
	// example:
	//
	// com.serverless.sae.controller.EchoController
	MethodController *string `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	// The name of the method.
	//
	// example:
	//
	// echo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The details of the method.
	//
	// example:
	//
	// description
	NameDetail *string `json:"NameDetail,omitempty" xml:"NameDetail,omitempty"`
	// The definition of the parameter.
	ParameterDefinitions []*DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	// The details of the parameters.
	ParameterDetails []*string `json:"ParameterDetails,omitempty" xml:"ParameterDetails,omitempty" type:"Repeated"`
	// The types of the parameters.
	ParameterTypes []*string `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty" type:"Repeated"`
	// The request paths. Format:
	//
	// `/path`
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The request methods. Valid values:
	//
	// 	- **GET**
	//
	// 	- **ALL**
	RequestMethods []*string `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty" type:"Repeated"`
	// The details of the response.
	//
	// example:
	//
	// test
	ReturnDetails *string `json:"ReturnDetails,omitempty" xml:"ReturnDetails,omitempty"`
	// The data format of the response.
	//
	// example:
	//
	// java.lang.String
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyDataMethods) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyDataMethods) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetMethodController() *string {
	return s.MethodController
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetName() *string {
	return s.Name
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetNameDetail() *string {
	return s.NameDetail
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetParameterDefinitions() []*DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	return s.ParameterDefinitions
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetParameterDetails() []*string {
	return s.ParameterDetails
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetParameterTypes() []*string {
	return s.ParameterTypes
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetPaths() []*string {
	return s.Paths
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetRequestMethods() []*string {
	return s.RequestMethods
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetReturnDetails() *string {
	return s.ReturnDetails
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) GetReturnType() *string {
	return s.ReturnType
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetMethodController(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.MethodController = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetName(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.Name = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetNameDetail(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.NameDetail = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterDefinitions(v []*DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterDefinitions = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterDetails(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterDetails = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetParameterTypes(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ParameterTypes = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetPaths(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.Paths = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetRequestMethods(v []*string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.RequestMethods = v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetReturnDetails(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ReturnDetails = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) SetReturnType(v string) *DescribeAppServiceDetailResponseBodyDataMethods {
	s.ReturnType = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethods) Validate() error {
	if s.ParameterDefinitions != nil {
		for _, item := range s.ParameterDefinitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions struct {
	// The description of the parameter.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// arg0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the parameter.
	//
	// example:
	//
	// java.lang.String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) GoString() string {
	return s.String()
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) GetDescription() *string {
	return s.Description
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) GetName() *string {
	return s.Name
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) GetType() *string {
	return s.Type
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetDescription(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Description = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetName(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Name = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) SetType(v string) *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions {
	s.Type = &v
	return s
}

func (s *DescribeAppServiceDetailResponseBodyDataMethodsParameterDefinitions) Validate() error {
	return dara.Validate(s)
}
