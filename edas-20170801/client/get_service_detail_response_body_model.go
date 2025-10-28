// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceDetailResponseBody
	GetCode() *int32
	SetData(v *GetServiceDetailResponseBodyData) *GetServiceDetailResponseBody
	GetData() *GetServiceDetailResponseBodyData
	SetMessage(v string) *GetServiceDetailResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetServiceDetailResponseBody
	GetSuccess() *bool
}

type GetServiceDetailResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data *GetServiceDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceDetailResponseBody) GetData() *GetServiceDetailResponseBodyData {
	return s.Data
}

func (s *GetServiceDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceDetailResponseBody) SetCode(v int32) *GetServiceDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetData(v *GetServiceDetailResponseBodyData) *GetServiceDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceDetailResponseBody) SetMessage(v string) *GetServiceDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceDetailResponseBody) SetSuccess(v bool) *GetServiceDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceDetailResponseBodyData struct {
	// The name of the Dubbo application.
	//
	// example:
	//
	// cartservice
	DubboApplicationName *string `json:"DubboApplicationName,omitempty" xml:"DubboApplicationName,omitempty"`
	// The name of the Enterprise Distributed Application Service (EDAS) application.
	//
	// example:
	//
	// test123
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
	Metadata *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The methods.
	Methods []*GetServiceDetailResponseBodyDataMethods `json:"Methods,omitempty" xml:"Methods,omitempty" type:"Repeated"`
	// The type of the service registry.
	//
	// example:
	//
	// agent
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// edas.service.consumer
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The type of the service.
	//
	// example:
	//
	// springCloud
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// The name of the Spring application.
	//
	// example:
	//
	// edas.service.consumer
	SpringApplicationName *string `json:"SpringApplicationName,omitempty" xml:"SpringApplicationName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyData) GetDubboApplicationName() *string {
	return s.DubboApplicationName
}

func (s *GetServiceDetailResponseBodyData) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *GetServiceDetailResponseBodyData) GetGroup() *string {
	return s.Group
}

func (s *GetServiceDetailResponseBodyData) GetMetadata() *string {
	return s.Metadata
}

func (s *GetServiceDetailResponseBodyData) GetMethods() []*GetServiceDetailResponseBodyDataMethods {
	return s.Methods
}

func (s *GetServiceDetailResponseBodyData) GetRegistryType() *string {
	return s.RegistryType
}

func (s *GetServiceDetailResponseBodyData) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceDetailResponseBodyData) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetServiceDetailResponseBodyData) GetSpringApplicationName() *string {
	return s.SpringApplicationName
}

func (s *GetServiceDetailResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *GetServiceDetailResponseBodyData) SetDubboApplicationName(v string) *GetServiceDetailResponseBodyData {
	s.DubboApplicationName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetEdasAppName(v string) *GetServiceDetailResponseBodyData {
	s.EdasAppName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetGroup(v string) *GetServiceDetailResponseBodyData {
	s.Group = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetMetadata(v string) *GetServiceDetailResponseBodyData {
	s.Metadata = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetMethods(v []*GetServiceDetailResponseBodyDataMethods) *GetServiceDetailResponseBodyData {
	s.Methods = v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetRegistryType(v string) *GetServiceDetailResponseBodyData {
	s.RegistryType = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceName(v string) *GetServiceDetailResponseBodyData {
	s.ServiceName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetServiceType(v string) *GetServiceDetailResponseBodyData {
	s.ServiceType = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetSpringApplicationName(v string) *GetServiceDetailResponseBodyData {
	s.SpringApplicationName = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) SetVersion(v string) *GetServiceDetailResponseBodyData {
	s.Version = &v
	return s
}

func (s *GetServiceDetailResponseBodyData) Validate() error {
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

type GetServiceDetailResponseBodyDataMethods struct {
	// The controllers.
	//
	// example:
	//
	// com.aliware.edas.DemoController
	MethodController *string `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// feign2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The specific name.
	//
	// example:
	//
	// test
	NameDetail *string `json:"NameDetail,omitempty" xml:"NameDetail,omitempty"`
	// The parameter definitions.
	//
	// example:
	//
	// [{"description":"","name":"arg0","type":"java.lang.String"}]
	ParameterDefinitions *string `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty"`
	// The parameter details.
	//
	// example:
	//
	// {}
	ParameterDetails *string `json:"ParameterDetails,omitempty" xml:"ParameterDetails,omitempty"`
	// The parameter names.
	//
	// example:
	//
	// test
	ParameterNames *string `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty"`
	// The parameter types.
	//
	// example:
	//
	// ["java.lang.String"]
	ParameterTypes *string `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty"`
	// The method paths.
	//
	// example:
	//
	// ["/consumer-echo/feign/{str}"]
	Paths *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	// The request methods.
	//
	// example:
	//
	// GET
	RequestMethods *string `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty"`
	// The definition of the value returned by the method.
	ReturnDefinition *GetServiceDetailResponseBodyDataMethodsReturnDefinition `json:"ReturnDefinition,omitempty" xml:"ReturnDefinition,omitempty" type:"Struct"`
	// The response details.
	//
	// example:
	//
	// test
	ReturnDetails *string `json:"ReturnDetails,omitempty" xml:"ReturnDetails,omitempty"`
	// The type of the response.
	//
	// example:
	//
	// java.lang.String
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GetServiceDetailResponseBodyDataMethods) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataMethods) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataMethods) GetMethodController() *string {
	return s.MethodController
}

func (s *GetServiceDetailResponseBodyDataMethods) GetName() *string {
	return s.Name
}

func (s *GetServiceDetailResponseBodyDataMethods) GetNameDetail() *string {
	return s.NameDetail
}

func (s *GetServiceDetailResponseBodyDataMethods) GetParameterDefinitions() *string {
	return s.ParameterDefinitions
}

func (s *GetServiceDetailResponseBodyDataMethods) GetParameterDetails() *string {
	return s.ParameterDetails
}

func (s *GetServiceDetailResponseBodyDataMethods) GetParameterNames() *string {
	return s.ParameterNames
}

func (s *GetServiceDetailResponseBodyDataMethods) GetParameterTypes() *string {
	return s.ParameterTypes
}

func (s *GetServiceDetailResponseBodyDataMethods) GetPaths() *string {
	return s.Paths
}

func (s *GetServiceDetailResponseBodyDataMethods) GetRequestMethods() *string {
	return s.RequestMethods
}

func (s *GetServiceDetailResponseBodyDataMethods) GetReturnDefinition() *GetServiceDetailResponseBodyDataMethodsReturnDefinition {
	return s.ReturnDefinition
}

func (s *GetServiceDetailResponseBodyDataMethods) GetReturnDetails() *string {
	return s.ReturnDetails
}

func (s *GetServiceDetailResponseBodyDataMethods) GetReturnType() *string {
	return s.ReturnType
}

func (s *GetServiceDetailResponseBodyDataMethods) SetMethodController(v string) *GetServiceDetailResponseBodyDataMethods {
	s.MethodController = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetName(v string) *GetServiceDetailResponseBodyDataMethods {
	s.Name = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetNameDetail(v string) *GetServiceDetailResponseBodyDataMethods {
	s.NameDetail = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetParameterDefinitions(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ParameterDefinitions = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetParameterDetails(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ParameterDetails = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetParameterNames(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ParameterNames = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetParameterTypes(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ParameterTypes = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetPaths(v string) *GetServiceDetailResponseBodyDataMethods {
	s.Paths = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetRequestMethods(v string) *GetServiceDetailResponseBodyDataMethods {
	s.RequestMethods = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetReturnDefinition(v *GetServiceDetailResponseBodyDataMethodsReturnDefinition) *GetServiceDetailResponseBodyDataMethods {
	s.ReturnDefinition = v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetReturnDetails(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ReturnDetails = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) SetReturnType(v string) *GetServiceDetailResponseBodyDataMethods {
	s.ReturnType = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethods) Validate() error {
	if s.ReturnDefinition != nil {
		if err := s.ReturnDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceDetailResponseBodyDataMethodsReturnDefinition struct {
	// The ID of the return value.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the response.
	//
	// example:
	//
	// foo
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceDetailResponseBodyDataMethodsReturnDefinition) String() string {
	return dara.Prettify(s)
}

func (s GetServiceDetailResponseBodyDataMethodsReturnDefinition) GoString() string {
	return s.String()
}

func (s *GetServiceDetailResponseBodyDataMethodsReturnDefinition) GetId() *string {
	return s.Id
}

func (s *GetServiceDetailResponseBodyDataMethodsReturnDefinition) GetType() *string {
	return s.Type
}

func (s *GetServiceDetailResponseBodyDataMethodsReturnDefinition) SetId(v string) *GetServiceDetailResponseBodyDataMethodsReturnDefinition {
	s.Id = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethodsReturnDefinition) SetType(v string) *GetServiceDetailResponseBodyDataMethodsReturnDefinition {
	s.Type = &v
	return s
}

func (s *GetServiceDetailResponseBodyDataMethodsReturnDefinition) Validate() error {
	return dara.Validate(s)
}
