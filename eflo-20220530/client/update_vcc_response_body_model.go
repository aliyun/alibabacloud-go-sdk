// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateVccResponseBody
	GetCode() *int32
	SetContent(v *UpdateVccResponseBodyContent) *UpdateVccResponseBody
	GetContent() *UpdateVccResponseBodyContent
	SetMessage(v string) *UpdateVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateVccResponseBody
	GetRequestId() *string
}

type UpdateVccResponseBody struct {
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
	// The response parameters.
	Content *UpdateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F906C4D3-7444-58E2-9819-E3D8563571A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVccResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateVccResponseBody) GetContent() *UpdateVccResponseBodyContent {
	return s.Content
}

func (s *UpdateVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVccResponseBody) SetAccessDeniedDetail(v string) *UpdateVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateVccResponseBody) SetCode(v int32) *UpdateVccResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateVccResponseBody) SetContent(v *UpdateVccResponseBodyContent) *UpdateVccResponseBody {
	s.Content = v
	return s
}

func (s *UpdateVccResponseBody) SetMessage(v string) *UpdateVccResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateVccResponseBody) SetRequestId(v string) *UpdateVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVccResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateVccResponseBodyContent struct {
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-2r42v22cn03
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s UpdateVccResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateVccResponseBodyContent) GetVccId() *string {
	return s.VccId
}

func (s *UpdateVccResponseBodyContent) SetVccId(v string) *UpdateVccResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *UpdateVccResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
