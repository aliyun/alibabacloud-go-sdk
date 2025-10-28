// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceMethodPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceMethodPageResponseBody
	GetCode() *string
	SetData(v *GetServiceMethodPageResponseBodyData) *GetServiceMethodPageResponseBody
	GetData() *GetServiceMethodPageResponseBodyData
	SetHttpCode(v string) *GetServiceMethodPageResponseBody
	GetHttpCode() *string
	SetMessage(v string) *GetServiceMethodPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceMethodPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceMethodPageResponseBody
	GetSuccess() *bool
}

type GetServiceMethodPageResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *GetServiceMethodPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpCode *string `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BDC0C0FE-D63B-4FC8-****-4081C57E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceMethodPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceMethodPageResponseBody) GetData() *GetServiceMethodPageResponseBodyData {
	return s.Data
}

func (s *GetServiceMethodPageResponseBody) GetHttpCode() *string {
	return s.HttpCode
}

func (s *GetServiceMethodPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceMethodPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceMethodPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceMethodPageResponseBody) SetCode(v string) *GetServiceMethodPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceMethodPageResponseBody) SetData(v *GetServiceMethodPageResponseBodyData) *GetServiceMethodPageResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceMethodPageResponseBody) SetHttpCode(v string) *GetServiceMethodPageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetServiceMethodPageResponseBody) SetMessage(v string) *GetServiceMethodPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceMethodPageResponseBody) SetRequestId(v string) *GetServiceMethodPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceMethodPageResponseBody) SetSuccess(v bool) *GetServiceMethodPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceMethodPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceMethodPageResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The data about the method.
	Result []*GetServiceMethodPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 6
	TotalSize *int32 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s GetServiceMethodPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetServiceMethodPageResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetServiceMethodPageResponseBodyData) GetResult() []*GetServiceMethodPageResponseBodyDataResult {
	return s.Result
}

func (s *GetServiceMethodPageResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetServiceMethodPageResponseBodyData) SetPageNumber(v int32) *GetServiceMethodPageResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyData) SetPageSize(v int32) *GetServiceMethodPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyData) SetResult(v []*GetServiceMethodPageResponseBodyDataResult) *GetServiceMethodPageResponseBodyData {
	s.Result = v
	return s
}

func (s *GetServiceMethodPageResponseBodyData) SetTotalSize(v int32) *GetServiceMethodPageResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceMethodPageResponseBodyDataResult struct {
	// The method.
	//
	// example:
	//
	// com.aliware.edas.EchoController
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
	// name
	NameDetail *string `json:"NameDetail,omitempty" xml:"NameDetail,omitempty"`
	// The definition of the parameter.
	//
	// example:
	//
	// [{"description":"","name":"arg0","type":"java.lang.String"}]
	ParameterDefinitions *string `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty"`
	// The details of the parameters.
	//
	// example:
	//
	// {}
	ParameterDetails *string `json:"ParameterDetails,omitempty" xml:"ParameterDetails,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// echo
	ParameterNames *string `json:"ParameterNames,omitempty" xml:"ParameterNames,omitempty"`
	// The data type of the parameter.
	//
	// example:
	//
	// java.lang.String
	ParameterTypes *string `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty"`
	// The method path.
	//
	// example:
	//
	// /consumer/alive
	Paths *string `json:"Paths,omitempty" xml:"Paths,omitempty"`
	// The request method.
	//
	// example:
	//
	// GET
	RequestMethods *string `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty"`
	// The return value.
	ReturnDefinition *GetServiceMethodPageResponseBodyDataResultReturnDefinition `json:"ReturnDefinition,omitempty" xml:"ReturnDefinition,omitempty" type:"Struct"`
	// The details of the response.
	//
	// example:
	//
	// java.lang.String
	ReturnDetails *string `json:"ReturnDetails,omitempty" xml:"ReturnDetails,omitempty"`
	// The data format of the response.
	//
	// example:
	//
	// java.lang.String
	ReturnType *string `json:"ReturnType,omitempty" xml:"ReturnType,omitempty"`
}

func (s GetServiceMethodPageResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetMethodController() *string {
	return s.MethodController
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetNameDetail() *string {
	return s.NameDetail
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterDefinitions() *string {
	return s.ParameterDefinitions
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterDetails() *string {
	return s.ParameterDetails
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterNames() *string {
	return s.ParameterNames
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterTypes() *string {
	return s.ParameterTypes
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetPaths() *string {
	return s.Paths
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetRequestMethods() *string {
	return s.RequestMethods
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetReturnDefinition() *GetServiceMethodPageResponseBodyDataResultReturnDefinition {
	return s.ReturnDefinition
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetReturnDetails() *string {
	return s.ReturnDetails
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetReturnType() *string {
	return s.ReturnType
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetMethodController(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.MethodController = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetName(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetNameDetail(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.NameDetail = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterDefinitions(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterDefinitions = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterDetails(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterDetails = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterNames(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterNames = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterTypes(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterTypes = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetPaths(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.Paths = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetRequestMethods(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.RequestMethods = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetReturnDefinition(v *GetServiceMethodPageResponseBodyDataResultReturnDefinition) *GetServiceMethodPageResponseBodyDataResult {
	s.ReturnDefinition = v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetReturnDetails(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ReturnDetails = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetReturnType(v string) *GetServiceMethodPageResponseBodyDataResult {
	s.ReturnType = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) Validate() error {
	if s.ReturnDefinition != nil {
		if err := s.ReturnDefinition.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceMethodPageResponseBodyDataResultReturnDefinition struct {
	// The ID of the return value.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The data format of the response.
	//
	// example:
	//
	// java.lang.String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceMethodPageResponseBodyDataResultReturnDefinition) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBodyDataResultReturnDefinition) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBodyDataResultReturnDefinition) GetId() *string {
	return s.Id
}

func (s *GetServiceMethodPageResponseBodyDataResultReturnDefinition) GetType() *string {
	return s.Type
}

func (s *GetServiceMethodPageResponseBodyDataResultReturnDefinition) SetId(v string) *GetServiceMethodPageResponseBodyDataResultReturnDefinition {
	s.Id = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResultReturnDefinition) SetType(v string) *GetServiceMethodPageResponseBodyDataResultReturnDefinition {
	s.Type = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResultReturnDefinition) Validate() error {
	return dara.Validate(s)
}
