// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceConsumersPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetServiceConsumersPageResponseBody
	GetCode() *int32
	SetData(v *GetServiceConsumersPageResponseBodyData) *GetServiceConsumersPageResponseBody
	GetData() *GetServiceConsumersPageResponseBodyData
	SetMessage(v string) *GetServiceConsumersPageResponseBody
	GetMessage() *string
	SetSuccess(v bool) *GetServiceConsumersPageResponseBody
	GetSuccess() *bool
}

type GetServiceConsumersPageResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *GetServiceConsumersPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s GetServiceConsumersPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConsumersPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceConsumersPageResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetServiceConsumersPageResponseBody) GetData() *GetServiceConsumersPageResponseBodyData {
	return s.Data
}

func (s *GetServiceConsumersPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceConsumersPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceConsumersPageResponseBody) SetCode(v int32) *GetServiceConsumersPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceConsumersPageResponseBody) SetData(v *GetServiceConsumersPageResponseBodyData) *GetServiceConsumersPageResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceConsumersPageResponseBody) SetMessage(v string) *GetServiceConsumersPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceConsumersPageResponseBody) SetSuccess(v bool) *GetServiceConsumersPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceConsumersPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceConsumersPageResponseBodyData struct {
	// The data array that is returned.
	Content []*GetServiceConsumersPageResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 5
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 3
	TotalElements *int32 `json:"TotalElements,omitempty" xml:"TotalElements,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s GetServiceConsumersPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConsumersPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceConsumersPageResponseBodyData) GetContent() []*GetServiceConsumersPageResponseBodyDataContent {
	return s.Content
}

func (s *GetServiceConsumersPageResponseBodyData) GetSize() *int32 {
	return s.Size
}

func (s *GetServiceConsumersPageResponseBodyData) GetTotalElements() *int32 {
	return s.TotalElements
}

func (s *GetServiceConsumersPageResponseBodyData) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetServiceConsumersPageResponseBodyData) SetContent(v []*GetServiceConsumersPageResponseBodyDataContent) *GetServiceConsumersPageResponseBodyData {
	s.Content = v
	return s
}

func (s *GetServiceConsumersPageResponseBodyData) SetSize(v int32) *GetServiceConsumersPageResponseBodyData {
	s.Size = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyData) SetTotalElements(v int32) *GetServiceConsumersPageResponseBodyData {
	s.TotalElements = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyData) SetTotalPages(v int32) *GetServiceConsumersPageResponseBodyData {
	s.TotalPages = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyData) Validate() error {
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

type GetServiceConsumersPageResponseBodyDataContent struct {
	// The name of the service consumer.
	//
	// example:
	//
	// k8s-lq-cartservice
	EdasAppName *string `json:"EdasAppName,omitempty" xml:"EdasAppName,omitempty"`
	// The ID of the service consumer.
	//
	// example:
	//
	// efbda488-7b33-432f-a40d-****0047****
	EdassAppId *string `json:"EdassAppId,omitempty" xml:"EdassAppId,omitempty"`
	// The IP address of the service consumer.
	//
	// example:
	//
	// 10.20.x.xx
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s GetServiceConsumersPageResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetServiceConsumersPageResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *GetServiceConsumersPageResponseBodyDataContent) GetEdasAppName() *string {
	return s.EdasAppName
}

func (s *GetServiceConsumersPageResponseBodyDataContent) GetEdassAppId() *string {
	return s.EdassAppId
}

func (s *GetServiceConsumersPageResponseBodyDataContent) GetIp() *string {
	return s.Ip
}

func (s *GetServiceConsumersPageResponseBodyDataContent) SetEdasAppName(v string) *GetServiceConsumersPageResponseBodyDataContent {
	s.EdasAppName = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyDataContent) SetEdassAppId(v string) *GetServiceConsumersPageResponseBodyDataContent {
	s.EdassAppId = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyDataContent) SetIp(v string) *GetServiceConsumersPageResponseBodyDataContent {
	s.Ip = &v
	return s
}

func (s *GetServiceConsumersPageResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
