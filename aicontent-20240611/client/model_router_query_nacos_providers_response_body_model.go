// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ModelRouterQueryNacosProvidersResponseBodyData) *ModelRouterQueryNacosProvidersResponseBody
	GetData() []*ModelRouterQueryNacosProvidersResponseBodyData
	SetErrCode(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryNacosProvidersResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryNacosProvidersResponseBody struct {
	// example:
	//
	// []
	Data []*ModelRouterQueryNacosProvidersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryNacosProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetData() []*ModelRouterQueryNacosProvidersResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetData(v []*ModelRouterQueryNacosProvidersResponseBodyData) *ModelRouterQueryNacosProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetErrCode(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetErrMessage(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetRequestId(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetSuccess(v bool) *ModelRouterQueryNacosProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) Validate() error {
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

type ModelRouterQueryNacosProvidersResponseBodyData struct {
	BaseUrl *string                                                 `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	Models  []*ModelRouterQueryNacosProvidersResponseBodyDataModels `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	Name    *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Symbol  *string                                                 `json:"symbol,omitempty" xml:"symbol,omitempty"`
}

func (s ModelRouterQueryNacosProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) GetModels() []*ModelRouterQueryNacosProvidersResponseBodyDataModels {
	return s.Models
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) GetSymbol() *string {
	return s.Symbol
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) SetBaseUrl(v string) *ModelRouterQueryNacosProvidersResponseBodyData {
	s.BaseUrl = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) SetModels(v []*ModelRouterQueryNacosProvidersResponseBodyDataModels) *ModelRouterQueryNacosProvidersResponseBodyData {
	s.Models = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) SetName(v string) *ModelRouterQueryNacosProvidersResponseBodyData {
	s.Name = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) SetSymbol(v string) *ModelRouterQueryNacosProvidersResponseBodyData {
	s.Symbol = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyData) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModelRouterQueryNacosProvidersResponseBodyDataModels struct {
	Extensions *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions `json:"extensions,omitempty" xml:"extensions,omitempty" type:"Struct"`
	Identifier *string                                                         `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// text
	InOut       *string `json:"inOut,omitempty" xml:"inOut,omitempty"`
	InputToken  *string `json:"inputToken,omitempty" xml:"inputToken,omitempty"`
	OutputToken *string `json:"outputToken,omitempty" xml:"outputToken,omitempty"`
	// example:
	//
	// Chat
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ModelRouterQueryNacosProvidersResponseBodyDataModels) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponseBodyDataModels) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetExtensions() *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions {
	return s.Extensions
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetIdentifier() *string {
	return s.Identifier
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetInOut() *string {
	return s.InOut
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetInputToken() *string {
	return s.InputToken
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetOutputToken() *string {
	return s.OutputToken
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) GetType() *string {
	return s.Type
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetExtensions(v *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.Extensions = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetIdentifier(v string) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.Identifier = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetInOut(v string) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.InOut = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetInputToken(v string) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.InputToken = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetOutputToken(v string) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.OutputToken = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) SetType(v string) *ModelRouterQueryNacosProvidersResponseBodyDataModels {
	s.Type = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModels) Validate() error {
	if s.Extensions != nil {
		if err := s.Extensions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions struct {
	Async *bool `json:"async,omitempty" xml:"async,omitempty"`
}

func (s ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) GetAsync() *bool {
	return s.Async
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) SetAsync(v bool) *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions {
	s.Async = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBodyDataModelsExtensions) Validate() error {
	return dara.Validate(s)
}
