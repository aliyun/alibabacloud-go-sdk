// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMethodsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListMethodsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListMethodsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMethodsResponseBody
	GetRequestId() *string
	SetServiceMethodList(v *ListMethodsResponseBodyServiceMethodList) *ListMethodsResponseBody
	GetServiceMethodList() *ListMethodsResponseBodyServiceMethodList
}

type ListMethodsResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message that indicates whether the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about service methods.
	ServiceMethodList *ListMethodsResponseBodyServiceMethodList `json:"ServiceMethodList,omitempty" xml:"ServiceMethodList,omitempty" type:"Struct"`
}

func (s ListMethodsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListMethodsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMethodsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMethodsResponseBody) GetServiceMethodList() *ListMethodsResponseBodyServiceMethodList {
	return s.ServiceMethodList
}

func (s *ListMethodsResponseBody) SetCode(v int32) *ListMethodsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMethodsResponseBody) SetMessage(v string) *ListMethodsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMethodsResponseBody) SetRequestId(v string) *ListMethodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMethodsResponseBody) SetServiceMethodList(v *ListMethodsResponseBodyServiceMethodList) *ListMethodsResponseBody {
	s.ServiceMethodList = v
	return s
}

func (s *ListMethodsResponseBody) Validate() error {
	if s.ServiceMethodList != nil {
		if err := s.ServiceMethodList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMethodsResponseBodyServiceMethodList struct {
	ServiceMethod []*ListMethodsResponseBodyServiceMethodListServiceMethod `json:"ServiceMethod,omitempty" xml:"ServiceMethod,omitempty" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodList) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodList) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodList) GetServiceMethod() []*ListMethodsResponseBodyServiceMethodListServiceMethod {
	return s.ServiceMethod
}

func (s *ListMethodsResponseBodyServiceMethodList) SetServiceMethod(v []*ListMethodsResponseBodyServiceMethodListServiceMethod) *ListMethodsResponseBodyServiceMethodList {
	s.ServiceMethod = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodList) Validate() error {
	if s.ServiceMethod != nil {
		for _, item := range s.ServiceMethod {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMethodsResponseBodyServiceMethodListServiceMethod struct {
	// The name of the application.
	//
	// example:
	//
	// App
	AppName     *string                                                           `json:"AppName,omitempty" xml:"AppName,omitempty"`
	InputParams *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams `json:"InputParams,omitempty" xml:"InputParams,omitempty" type:"Struct"`
	// The name of the service method.
	//
	// example:
	//
	// echo
	MethodName *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	// The return type of the service method.
	//
	// example:
	//
	// java.lang.string
	Output     *string                                                          `json:"Output,omitempty" xml:"Output,omitempty"`
	ParamTypes *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes `json:"ParamTypes,omitempty" xml:"ParamTypes,omitempty" type:"Struct"`
	// The name of the service.
	//
	// example:
	//
	// com.alibaba.edas.demo.EchoService
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethod) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethod) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetAppName() *string {
	return s.AppName
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetInputParams() *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams {
	return s.InputParams
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetMethodName() *string {
	return s.MethodName
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetOutput() *string {
	return s.Output
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetParamTypes() *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes {
	return s.ParamTypes
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetAppName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.AppName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetInputParams(v *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.InputParams = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetMethodName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.MethodName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetOutput(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.Output = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetParamTypes(v *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.ParamTypes = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) SetServiceName(v string) *ListMethodsResponseBodyServiceMethodListServiceMethod {
	s.ServiceName = &v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethod) Validate() error {
	if s.InputParams != nil {
		if err := s.InputParams.Validate(); err != nil {
			return err
		}
	}
	if s.ParamTypes != nil {
		if err := s.ParamTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMethodsResponseBodyServiceMethodListServiceMethodInputParams struct {
	InputParam []*string `json:"InputParam,omitempty" xml:"InputParam,omitempty" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) GetInputParam() []*string {
	return s.InputParam
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) SetInputParam(v []*string) *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams {
	s.InputParam = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodInputParams) Validate() error {
	return dara.Validate(s)
}

type ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes struct {
	ParamType []*string `json:"ParamType,omitempty" xml:"ParamType,omitempty" type:"Repeated"`
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) String() string {
	return dara.Prettify(s)
}

func (s ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) GoString() string {
	return s.String()
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) GetParamType() []*string {
	return s.ParamType
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) SetParamType(v []*string) *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes {
	s.ParamType = v
	return s
}

func (s *ListMethodsResponseBodyServiceMethodListServiceMethodParamTypes) Validate() error {
	return dara.Validate(s)
}
