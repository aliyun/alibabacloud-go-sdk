// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPredefinedModelProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPredefinedModelProvidersResponseBody
	GetCode() *string
	SetData(v []*ListPredefinedModelProvidersResponseBodyData) *ListPredefinedModelProvidersResponseBody
	GetData() []*ListPredefinedModelProvidersResponseBodyData
	SetHttpStatusCode(v int32) *ListPredefinedModelProvidersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListPredefinedModelProvidersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPredefinedModelProvidersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPredefinedModelProvidersResponseBody
	GetSuccess() *bool
}

type ListPredefinedModelProvidersResponseBody struct {
	// The business status code. The value is SUCCESS if the request is successful.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of predefined model providers.
	Data []*ListPredefinedModelProvidersResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code. The value is 200 if the request is successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The message about the request processing result.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListPredefinedModelProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPredefinedModelProvidersResponseBody) GetData() []*ListPredefinedModelProvidersResponseBodyData {
	return s.Data
}

func (s *ListPredefinedModelProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListPredefinedModelProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPredefinedModelProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPredefinedModelProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPredefinedModelProvidersResponseBody) SetCode(v string) *ListPredefinedModelProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) SetData(v []*ListPredefinedModelProvidersResponseBodyData) *ListPredefinedModelProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) SetHttpStatusCode(v int32) *ListPredefinedModelProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) SetMessage(v string) *ListPredefinedModelProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) SetRequestId(v string) *ListPredefinedModelProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) SetSuccess(v bool) *ListPredefinedModelProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBody) Validate() error {
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

type ListPredefinedModelProvidersResponseBodyData struct {
	// The default model service endpoint of the provider.
	//
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	DefaultEndpoint *string `json:"defaultEndpoint,omitempty" xml:"defaultEndpoint,omitempty"`
	// The default model invocation protocol of the provider.
	//
	// example:
	//
	// OpenAI/v1
	DefaultProtocol *string `json:"defaultProtocol,omitempty" xml:"defaultProtocol,omitempty"`
	// The display name of the model provider.
	//
	// example:
	//
	// Alibaba Cloud Model Studio
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The model provider type.
	//
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
}

func (s ListPredefinedModelProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPredefinedModelProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPredefinedModelProvidersResponseBodyData) GetDefaultEndpoint() *string {
	return s.DefaultEndpoint
}

func (s *ListPredefinedModelProvidersResponseBodyData) GetDefaultProtocol() *string {
	return s.DefaultProtocol
}

func (s *ListPredefinedModelProvidersResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListPredefinedModelProvidersResponseBodyData) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListPredefinedModelProvidersResponseBodyData) SetDefaultEndpoint(v string) *ListPredefinedModelProvidersResponseBodyData {
	s.DefaultEndpoint = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBodyData) SetDefaultProtocol(v string) *ListPredefinedModelProvidersResponseBodyData {
	s.DefaultProtocol = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBodyData) SetDisplayName(v string) *ListPredefinedModelProvidersResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBodyData) SetProviderType(v string) *ListPredefinedModelProvidersResponseBodyData {
	s.ProviderType = &v
	return s
}

func (s *ListPredefinedModelProvidersResponseBodyData) Validate() error {
	return dara.Validate(s)
}
