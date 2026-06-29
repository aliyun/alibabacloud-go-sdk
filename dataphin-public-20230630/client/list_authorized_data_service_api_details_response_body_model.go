// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAuthorizedDataServiceApiDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetRequestId() *string
	SetResult(v *ListAuthorizedDataServiceApiDetailsResponseBodyResult) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetResult() *ListAuthorizedDataServiceApiDetailsResponseBodyResult
	SetSuccess(v bool) *ListAuthorizedDataServiceApiDetailsResponseBody
	GetSuccess() *bool
}

type ListAuthorizedDataServiceApiDetailsResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// 非法入参
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. It is a unique identifier for the request.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query result.
	Result *ListAuthorizedDataServiceApiDetailsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetResult() *ListAuthorizedDataServiceApiDetailsResponseBodyResult {
	return s.Result
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetCode(v string) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetHttpStatusCode(v int32) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetMessage(v string) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetRequestId(v string) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetResult(v *ListAuthorizedDataServiceApiDetailsResponseBodyResult) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.Result = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) SetSuccess(v bool) *ListAuthorizedDataServiceApiDetailsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAuthorizedDataServiceApiDetailsResponseBodyResult struct {
	// The list of authorized API details.
	Data []*ListAuthorizedDataServiceApiDetailsResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResult) GetData() []*ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	return s.Data
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResult) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResult) SetData(v []*ListAuthorizedDataServiceApiDetailsResponseBodyResultData) *ListAuthorizedDataServiceApiDetailsResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResult) SetTotalCount(v int64) *ListAuthorizedDataServiceApiDetailsResponseBodyResult {
	s.TotalCount = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResult) Validate() error {
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

type ListAuthorizedDataServiceApiDetailsResponseBodyResultData struct {
	// API_ID
	//
	// example:
	//
	// 12345
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// GetData
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 12345
	AppId *int64 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The permission type.
	//
	// example:
	//
	// 使用权限
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The list of fields authorized for the API in the development environment.
	AuthorizedDevReturnParameters []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters `json:"AuthorizedDevReturnParameters,omitempty" xml:"AuthorizedDevReturnParameters,omitempty" type:"Repeated"`
	// The list of fields authorized for the API in the production environment.
	AuthorizedProdReturnParameters []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters `json:"AuthorizedProdReturnParameters,omitempty" xml:"AuthorizedProdReturnParameters,omitempty" type:"Repeated"`
	// The API description.
	//
	// example:
	//
	// Description1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The validity period of the API permission in the development environment.
	//
	// example:
	//
	// 2035-12-31
	DevAuthPeriod *string `json:"DevAuthPeriod,omitempty" xml:"DevAuthPeriod,omitempty"`
	// The validity period of the API permission in the production environment.
	//
	// example:
	//
	// 2035-12-31
	ProdAuthPeriod *string `json:"ProdAuthPeriod,omitempty" xml:"ProdAuthPeriod,omitempty"`
	// The data service project ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetApiName() *string {
	return s.ApiName
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetAppId() *int64 {
	return s.AppId
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetAuthType() *string {
	return s.AuthType
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetAuthorizedDevReturnParameters() []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	return s.AuthorizedDevReturnParameters
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetAuthorizedProdReturnParameters() []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	return s.AuthorizedProdReturnParameters
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetDescription() *string {
	return s.Description
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetDevAuthPeriod() *string {
	return s.DevAuthPeriod
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetProdAuthPeriod() *string {
	return s.ProdAuthPeriod
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetApiId(v int64) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.ApiId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetApiName(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.ApiName = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetAppId(v int64) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.AppId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetAuthType(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.AuthType = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetAuthorizedDevReturnParameters(v []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.AuthorizedDevReturnParameters = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetAuthorizedProdReturnParameters(v []*ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.AuthorizedProdReturnParameters = v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetDescription(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.Description = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetDevAuthPeriod(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.DevAuthPeriod = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetProdAuthPeriod(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.ProdAuthPeriod = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) SetProjectId(v int64) *ListAuthorizedDataServiceApiDetailsResponseBodyResultData {
	s.ProjectId = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultData) Validate() error {
	if s.AuthorizedDevReturnParameters != nil {
		for _, item := range s.AuthorizedDevReturnParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AuthorizedProdReturnParameters != nil {
		for _, item := range s.AuthorizedProdReturnParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters struct {
	// The example value.
	//
	// example:
	//
	// example1
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The field authorization status.
	//
	// example:
	//
	// 1
	IsAuthorized *int32 `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// The parameter type.
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description1
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// param1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GetIsAuthorized() *int32 {
	return s.IsAuthorized
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) SetExampleValue(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) SetIsAuthorized(v int32) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	s.IsAuthorized = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) SetParameterDataType(v int32) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) SetParameterDescription(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) SetParameterName(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters {
	s.ParameterName = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedDevReturnParameters) Validate() error {
	return dara.Validate(s)
}

type ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters struct {
	// The example value.
	//
	// example:
	//
	// example1
	ExampleValue *string `json:"ExampleValue,omitempty" xml:"ExampleValue,omitempty"`
	// The field authorization status.
	//
	// example:
	//
	// 1
	IsAuthorized *int32 `json:"IsAuthorized,omitempty" xml:"IsAuthorized,omitempty"`
	// The parameter type.
	//
	// example:
	//
	// 0
	ParameterDataType *int32 `json:"ParameterDataType,omitempty" xml:"ParameterDataType,omitempty"`
	// The description.
	//
	// example:
	//
	// description1
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// param1
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) String() string {
	return dara.Prettify(s)
}

func (s ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GoString() string {
	return s.String()
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GetExampleValue() *string {
	return s.ExampleValue
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GetIsAuthorized() *int32 {
	return s.IsAuthorized
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GetParameterDataType() *int32 {
	return s.ParameterDataType
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) SetExampleValue(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	s.ExampleValue = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) SetIsAuthorized(v int32) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	s.IsAuthorized = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) SetParameterDataType(v int32) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	s.ParameterDataType = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) SetParameterDescription(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	s.ParameterDescription = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) SetParameterName(v string) *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters {
	s.ParameterName = &v
	return s
}

func (s *ListAuthorizedDataServiceApiDetailsResponseBodyResultDataAuthorizedProdReturnParameters) Validate() error {
	return dara.Validate(s)
}
