// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceProvidersPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceProvidersPageResponseBody
	GetCode() *int32
	SetData(v *GetServiceProvidersPageResponseBodyData) *GetServiceProvidersPageResponseBody
	GetData() *GetServiceProvidersPageResponseBodyData
	SetMessage(v string) *GetServiceProvidersPageResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetServiceProvidersPageResponseBody
	GetSuccess() *bool
}

type GetServiceProvidersPageResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data structure.
	Data *GetServiceProvidersPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceProvidersPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvidersPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceProvidersPageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceProvidersPageResponseBody) GetData() *GetServiceProvidersPageResponseBodyData {
	return s.Data
}

func (s *GetServiceProvidersPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceProvidersPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceProvidersPageResponseBody) SetCode(v int32) *GetServiceProvidersPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceProvidersPageResponseBody) SetData(v *GetServiceProvidersPageResponseBodyData) *GetServiceProvidersPageResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceProvidersPageResponseBody) SetMessage(v string) *GetServiceProvidersPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceProvidersPageResponseBody) SetSuccess(v bool) *GetServiceProvidersPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceProvidersPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceProvidersPageResponseBodyData struct {
	// The data array returned.
	Content []*GetServiceProvidersPageResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 5
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s GetServiceProvidersPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvidersPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceProvidersPageResponseBodyData) GetContent() []*GetServiceProvidersPageResponseBodyDataContent {
	return s.Content
}

func (s *GetServiceProvidersPageResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceProvidersPageResponseBodyData) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *GetServiceProvidersPageResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetServiceProvidersPageResponseBodyData) SetContent(v []*GetServiceProvidersPageResponseBodyDataContent) *GetServiceProvidersPageResponseBodyData {
	s.Content = v
	return s
}

func (s *GetServiceProvidersPageResponseBodyData) SetSize(v int32) *GetServiceProvidersPageResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyData) SetTotalElements(v int32) *GetServiceProvidersPageResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyData) SetTotalPages(v int32) *GetServiceProvidersPageResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetServiceProvidersPageResponseBodyDataContent struct {
	// The remarks of the service provider.
	//
	// example:
	//
	// 172.178.XX.XX
	Iannotations *string `json:"Iannotations,omitempty" xml:"Iannotations,omitempty"`
	// The IP address of the service provider.
	//
	// example:
	//
	// 10.20.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port number of the service provider.
	//
	// example:
	//
	// 12345
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The serialization type.
	//
	// example:
	//
	// hessian2
	SerializeType *string `json:"SerializeType,omitempty" xml:"SerializeType,omitempty"`
	// The service timeout period.
	//
	// example:
	//
	// 1000
	Timeout *string `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s GetServiceProvidersPageResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetServiceProvidersPageResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *GetServiceProvidersPageResponseBodyDataContent) GetIannotations() *string {
	return s.Iannotations
}

func (s *GetServiceProvidersPageResponseBodyDataContent) GetIp() *string {
	return s.Ip
}

func (s *GetServiceProvidersPageResponseBodyDataContent) GetPort() *string {
	return s.Port
}

func (s *GetServiceProvidersPageResponseBodyDataContent) GetSerializeType() *string {
	return s.SerializeType
}

func (s *GetServiceProvidersPageResponseBodyDataContent) GetTimeout() *string {
	return s.Timeout
}

func (s *GetServiceProvidersPageResponseBodyDataContent) SetIannotations(v string) *GetServiceProvidersPageResponseBodyDataContent {
	s.Iannotations = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyDataContent) SetIp(v string) *GetServiceProvidersPageResponseBodyDataContent {
	s.Ip = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyDataContent) SetPort(v string) *GetServiceProvidersPageResponseBodyDataContent {
	s.Port = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyDataContent) SetSerializeType(v string) *GetServiceProvidersPageResponseBodyDataContent {
	s.SerializeType = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyDataContent) SetTimeout(v string) *GetServiceProvidersPageResponseBodyDataContent {
	s.Timeout = &v
	return s
}

func (s *GetServiceProvidersPageResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
