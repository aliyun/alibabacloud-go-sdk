// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRelatedResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListRelatedResourceResponseBody
	GetCode() *string
	SetData(v *ListRelatedResourceResponseBodyData) *ListRelatedResourceResponseBody
	GetData() *ListRelatedResourceResponseBodyData
	SetMessage(v string) *ListRelatedResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListRelatedResourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRelatedResourceResponseBody
	GetSuccess() *bool
}

type ListRelatedResourceResponseBody struct {
	// The error code. This field is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	Data *ListRelatedResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the request processing result. The value `success` is typically returned when the request is successful.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates that the request is successful, and a value of false indicates that the request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRelatedResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRelatedResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ListRelatedResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListRelatedResourceResponseBody) GetData() *ListRelatedResourceResponseBodyData {
	return s.Data
}

func (s *ListRelatedResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRelatedResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRelatedResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRelatedResourceResponseBody) SetCode(v string) *ListRelatedResourceResponseBody {
	s.Code = &v
	return s
}

func (s *ListRelatedResourceResponseBody) SetData(v *ListRelatedResourceResponseBodyData) *ListRelatedResourceResponseBody {
	s.Data = v
	return s
}

func (s *ListRelatedResourceResponseBody) SetMessage(v string) *ListRelatedResourceResponseBody {
	s.Message = &v
	return s
}

func (s *ListRelatedResourceResponseBody) SetRequestId(v string) *ListRelatedResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRelatedResourceResponseBody) SetSuccess(v bool) *ListRelatedResourceResponseBody {
	s.Success = &v
	return s
}

func (s *ListRelatedResourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRelatedResourceResponseBodyData struct {
	// The number of cloud desktops associated with the image.
	//
	// example:
	//
	// 1
	DesktopCount *int32 `json:"DesktopCount,omitempty" xml:"DesktopCount,omitempty"`
	// The number of cloud desktop pools associated with the image.
	//
	// example:
	//
	// 0
	DesktopGroupCount *int32 `json:"DesktopGroupCount,omitempty" xml:"DesktopGroupCount,omitempty"`
	// The number of associated templates.
	//
	// example:
	//
	// 1
	TemplateCount *int32 `json:"TemplateCount,omitempty" xml:"TemplateCount,omitempty"`
}

func (s ListRelatedResourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListRelatedResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListRelatedResourceResponseBodyData) GetDesktopCount() *int32 {
	return s.DesktopCount
}

func (s *ListRelatedResourceResponseBodyData) GetDesktopGroupCount() *int32 {
	return s.DesktopGroupCount
}

func (s *ListRelatedResourceResponseBodyData) GetTemplateCount() *int32 {
	return s.TemplateCount
}

func (s *ListRelatedResourceResponseBodyData) SetDesktopCount(v int32) *ListRelatedResourceResponseBodyData {
	s.DesktopCount = &v
	return s
}

func (s *ListRelatedResourceResponseBodyData) SetDesktopGroupCount(v int32) *ListRelatedResourceResponseBodyData {
	s.DesktopGroupCount = &v
	return s
}

func (s *ListRelatedResourceResponseBodyData) SetTemplateCount(v int32) *ListRelatedResourceResponseBodyData {
	s.TemplateCount = &v
	return s
}

func (s *ListRelatedResourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
