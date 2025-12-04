// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateSubnetResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateSubnetResponseBody
	GetCode() *int32
	SetContent(v *UpdateSubnetResponseBodyContent) *UpdateSubnetResponseBody
	GetContent() *UpdateSubnetResponseBodyContent
	SetMessage(v string) *UpdateSubnetResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateSubnetResponseBody
	GetRequestId() *string
}

type UpdateSubnetResponseBody struct {
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
	// The response content.
	Content *UpdateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3D9D6E7B-365B-5200-BFA6-9B79E269058C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateSubnetResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateSubnetResponseBody) GetContent() *UpdateSubnetResponseBodyContent {
	return s.Content
}

func (s *UpdateSubnetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSubnetResponseBody) SetAccessDeniedDetail(v string) *UpdateSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateSubnetResponseBody) SetCode(v int32) *UpdateSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateSubnetResponseBody) SetContent(v *UpdateSubnetResponseBodyContent) *UpdateSubnetResponseBody {
	s.Content = v
	return s
}

func (s *UpdateSubnetResponseBody) SetMessage(v string) *UpdateSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateSubnetResponseBody) SetRequestId(v string) *UpdateSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSubnetResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSubnetResponseBodyContent struct {
	// The subnet instance ID.
	//
	// example:
	//
	// subnet-yuvn29bn
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s UpdateSubnetResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *UpdateSubnetResponseBodyContent) GetSubnetId() *string {
	return s.SubnetId
}

func (s *UpdateSubnetResponseBodyContent) SetSubnetId(v string) *UpdateSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

func (s *UpdateSubnetResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
