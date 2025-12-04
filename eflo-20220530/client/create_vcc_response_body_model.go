// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateVccResponseBody
	GetCode() *int32
	SetContent(v *CreateVccResponseBodyContent) *CreateVccResponseBody
	GetContent() *CreateVccResponseBodyContent
	SetMessage(v string) *CreateVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVccResponseBody
	GetRequestId() *string
}

type CreateVccResponseBody struct {
	// The detailed information about the failed permission verification.
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
	// The returned data.
	Content *CreateVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// response message, if the success request is
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 039C3C3A-3C37-5672-80D5-D8CD48C676D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVccResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateVccResponseBody) GetContent() *CreateVccResponseBodyContent {
	return s.Content
}

func (s *CreateVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVccResponseBody) SetAccessDeniedDetail(v string) *CreateVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateVccResponseBody) SetCode(v int32) *CreateVccResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVccResponseBody) SetContent(v *CreateVccResponseBodyContent) *CreateVccResponseBody {
	s.Content = v
	return s
}

func (s *CreateVccResponseBody) SetMessage(v string) *CreateVccResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVccResponseBody) SetRequestId(v string) *CreateVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVccResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateVccResponseBodyContent struct {
	// The ID of the Lingjun connection instance.
	//
	// example:
	//
	// vcc-cn-zvp2w222001
	VccId *string `json:"VccId,omitempty" xml:"VccId,omitempty"`
}

func (s CreateVccResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateVccResponseBodyContent) GetVccId() *string {
	return s.VccId
}

func (s *CreateVccResponseBodyContent) SetVccId(v string) *CreateVccResponseBodyContent {
	s.VccId = &v
	return s
}

func (s *CreateVccResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
