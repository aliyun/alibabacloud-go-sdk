// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceListPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceListPageResponseBody
	GetCode() *int32
	SetData(v *GetServiceListPageResponseBodyData) *GetServiceListPageResponseBody
	GetData() *GetServiceListPageResponseBodyData
	SetMessage(v string) *GetServiceListPageResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetServiceListPageResponseBody
	GetSuccess() *bool
}

type GetServiceListPageResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetServiceListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned for the request.
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

func (s GetServiceListPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceListPageResponseBody) GetData() *GetServiceListPageResponseBodyData {
	return s.Data
}

func (s *GetServiceListPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceListPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceListPageResponseBody) SetCode(v int32) *GetServiceListPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceListPageResponseBody) SetData(v *GetServiceListPageResponseBodyData) *GetServiceListPageResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceListPageResponseBody) SetMessage(v string) *GetServiceListPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceListPageResponseBody) SetSuccess(v bool) *GetServiceListPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceListPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceListPageResponseBodyData struct {
	// The data array that is returned.
	Content []*GetServiceListPageResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 8
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 6
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s GetServiceListPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBodyData) GetContent() []*GetServiceListPageResponseBodyDataContent {
	return s.Content
}

func (s *GetServiceListPageResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceListPageResponseBodyData) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *GetServiceListPageResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetServiceListPageResponseBodyData) SetContent(v []*GetServiceListPageResponseBodyDataContent) *GetServiceListPageResponseBodyData {
	s.Content = v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetSize(v int32) *GetServiceListPageResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetTotalElements(v int32) *GetServiceListPageResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) SetTotalPages(v int32) *GetServiceListPageResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetServiceListPageResponseBodyData) Validate() error {
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

type GetServiceListPageResponseBodyDataContent struct {
	// The ID of the application.
	//
	// example:
	//
	// efbda488-7b33-432f-****-36530047****
	EdasAppId *string `json:"EdasAppId,omitempty" xml:"EdasAppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// k8s-lq-cartservice
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	// The service group.
	//
	// example:
	//
	// DUBBO
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// The number of instances.
	//
	// example:
	//
	// 1
	InstanceNum *int32 `json:"InstanceNum,omitempty" xml:"InstanceNum,omitempty"`
	// The type of the service registry.
	//
	// example:
	//
	// xx
	RegisterType *string `json:"RegisterType,omitempty" xml:"RegisterType,omitempty"`
	// The ID of the service.
	//
	// example:
	//
	// xx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// The name of the service.
	//
	// example:
	//
	// com.alibabacloud.hipstershop.CartService
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetServiceListPageResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetServiceListPageResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *GetServiceListPageResponseBodyDataContent) GetEdasAppId() *string {
	return s.EdasAppId
}

func (s *GetServiceListPageResponseBodyDataContent) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *GetServiceListPageResponseBodyDataContent) GetGroup() *string {
	return s.Group
}

func (s *GetServiceListPageResponseBodyDataContent) GetInstanceNum() *int32 {
	return s.InstanceNum
}

func (s *GetServiceListPageResponseBodyDataContent) GetRegisterType() *string {
	return s.RegisterType
}

func (s *GetServiceListPageResponseBodyDataContent) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceListPageResponseBodyDataContent) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetServiceListPageResponseBodyDataContent) GetVersion() *string {
	return s.Version
}

func (s *GetServiceListPageResponseBodyDataContent) SetEdasAppId(v string) *GetServiceListPageResponseBodyDataContent {
	s.EdasAppId = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetEdasAppName(v string) *GetServiceListPageResponseBodyDataContent {
	s.EdasAppName = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetGroup(v string) *GetServiceListPageResponseBodyDataContent {
	s.Group = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetInstanceNum(v int32) *GetServiceListPageResponseBodyDataContent {
	s.InstanceNum = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetRegisterType(v string) *GetServiceListPageResponseBodyDataContent {
	s.RegisterType = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetServiceId(v string) *GetServiceListPageResponseBodyDataContent {
	s.ServiceId = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetServiceName(v string) *GetServiceListPageResponseBodyDataContent {
	s.ServiceName = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) SetVersion(v string) *GetServiceListPageResponseBodyDataContent {
	s.Version = &v
	return s
}

func (s *GetServiceListPageResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
