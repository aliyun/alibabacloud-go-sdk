// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceMethodPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetServiceMethodPageResponseBodyData) *GetServiceMethodPageResponseBody
	GetData() *GetServiceMethodPageResponseBodyData
	SetMessage(v string) *GetServiceMethodPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceMethodPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceMethodPageResponseBody
	GetSuccess() *bool
}

type GetServiceMethodPageResponseBody struct {
	Data *GetServiceMethodPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2DD5A212-C77B-3XXF-9XXE-XXX9XXXE5XX1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceMethodPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBody) GetData() *GetServiceMethodPageResponseBodyData {
	return s.Data
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

func (s *GetServiceMethodPageResponseBody) SetData(v *GetServiceMethodPageResponseBodyData) *GetServiceMethodPageResponseBody {
	s.Data = v
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
	return dara.Validate(s)
}

type GetServiceMethodPageResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Result   []*GetServiceMethodPageResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 100
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
	return dara.Validate(s)
}

type GetServiceMethodPageResponseBodyDataResult struct {
	// example:
	//
	// com.alibabacloud.mse.demo.a.AController
	MethodController *string `json:"MethodController,omitempty" xml:"MethodController,omitempty"`
	// example:
	//
	// aMethod
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 示例接口描述信息。
	NameDetail           *string                                                           `json:"NameDetail,omitempty" xml:"NameDetail,omitempty"`
	ParameterDefinitions []*GetServiceMethodPageResponseBodyDataResultParameterDefinitions `json:"ParameterDefinitions,omitempty" xml:"ParameterDefinitions,omitempty" type:"Repeated"`
	ParameterDetails     []*string                                                         `json:"ParameterDetails,omitempty" xml:"ParameterDetails,omitempty" type:"Repeated"`
	ParameterTypes       []*string                                                         `json:"ParameterTypes,omitempty" xml:"ParameterTypes,omitempty" type:"Repeated"`
	Paths                []*string                                                         `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	RequestMethods       []*string                                                         `json:"RequestMethods,omitempty" xml:"RequestMethods,omitempty" type:"Repeated"`
	// example:
	//
	// java.lang.String
	ReturnDetails *string `json:"ReturnDetails,omitempty" xml:"ReturnDetails,omitempty"`
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

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterDefinitions() []*GetServiceMethodPageResponseBodyDataResultParameterDefinitions {
	return s.ParameterDefinitions
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterDetails() []*string {
	return s.ParameterDetails
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetParameterTypes() []*string {
	return s.ParameterTypes
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetPaths() []*string {
	return s.Paths
}

func (s *GetServiceMethodPageResponseBodyDataResult) GetRequestMethods() []*string {
	return s.RequestMethods
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

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterDefinitions(v []*GetServiceMethodPageResponseBodyDataResultParameterDefinitions) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterDefinitions = v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterDetails(v []*string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterDetails = v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetParameterTypes(v []*string) *GetServiceMethodPageResponseBodyDataResult {
	s.ParameterTypes = v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetPaths(v []*string) *GetServiceMethodPageResponseBodyDataResult {
	s.Paths = v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResult) SetRequestMethods(v []*string) *GetServiceMethodPageResponseBodyDataResult {
	s.RequestMethods = v
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
	return dara.Validate(s)
}

type GetServiceMethodPageResponseBodyDataResultParameterDefinitions struct {
	// example:
	//
	// 参数描述示例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// aParam
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// java.lang.String
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetServiceMethodPageResponseBodyDataResultParameterDefinitions) String() string {
	return dara.Prettify(s)
}

func (s GetServiceMethodPageResponseBodyDataResultParameterDefinitions) GoString() string {
	return s.String()
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) GetDescription() *string {
	return s.Description
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) GetName() *string {
	return s.Name
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) GetType() *string {
	return s.Type
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) SetDescription(v string) *GetServiceMethodPageResponseBodyDataResultParameterDefinitions {
	s.Description = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) SetName(v string) *GetServiceMethodPageResponseBodyDataResultParameterDefinitions {
	s.Name = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) SetType(v string) *GetServiceMethodPageResponseBodyDataResultParameterDefinitions {
	s.Type = &v
	return s
}

func (s *GetServiceMethodPageResponseBodyDataResultParameterDefinitions) Validate() error {
	return dara.Validate(s)
}
