// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateVpdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateVpdResponseBody
	GetCode() *int32
	SetContent(v *UpdateVpdResponseBodyContent) *UpdateVpdResponseBody
	GetContent() *UpdateVpdResponseBodyContent
	SetMessage(v string) *UpdateVpdResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVpdResponseBody
	GetRequestId() *string
}

type UpdateVpdResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *UpdateVpdResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVpdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateVpdResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateVpdResponseBody) GetContent() *UpdateVpdResponseBodyContent {
	return s.Content
}

func (s *UpdateVpdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVpdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVpdResponseBody) SetAccessDeniedDetail(v string) *UpdateVpdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateVpdResponseBody) SetCode(v int32) *UpdateVpdResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVpdResponseBody) SetContent(v *UpdateVpdResponseBodyContent) *UpdateVpdResponseBody {
	s.Content = v
	return s
}

func (s *UpdateVpdResponseBody) SetMessage(v string) *UpdateVpdResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVpdResponseBody) SetRequestId(v string) *UpdateVpdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVpdResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVpdResponseBodyContent struct {
	// The ID of the VPD instance.
	//
	// example:
	//
	// vpd-lg4dppgi
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s UpdateVpdResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpdResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateVpdResponseBodyContent) GetVpdId() *string {
	return s.VpdId
}

func (s *UpdateVpdResponseBodyContent) SetVpdId(v string) *UpdateVpdResponseBodyContent {
	s.VpdId = &v
	return s
}

func (s *UpdateVpdResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
