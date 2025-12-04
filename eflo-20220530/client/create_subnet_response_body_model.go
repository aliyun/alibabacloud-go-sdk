// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubnetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateSubnetResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *CreateSubnetResponseBody
	GetCode() *int32
	SetContent(v *CreateSubnetResponseBodyContent) *CreateSubnetResponseBody
	GetContent() *CreateSubnetResponseBodyContent
	SetMessage(v string) *CreateSubnetResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSubnetResponseBody
	GetRequestId() *string
}

type CreateSubnetResponseBody struct {
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
	// The response parameters.
	Content *CreateSubnetResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// A88DFED5-24B7-5A3E-87DE-380BF06F3C90
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubnetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateSubnetResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSubnetResponseBody) GetContent() *CreateSubnetResponseBodyContent {
	return s.Content
}

func (s *CreateSubnetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSubnetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubnetResponseBody) SetAccessDeniedDetail(v string) *CreateSubnetResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateSubnetResponseBody) SetCode(v int32) *CreateSubnetResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSubnetResponseBody) SetContent(v *CreateSubnetResponseBodyContent) *CreateSubnetResponseBody {
	s.Content = v
	return s
}

func (s *CreateSubnetResponseBody) SetMessage(v string) *CreateSubnetResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSubnetResponseBody) SetRequestId(v string) *CreateSubnetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubnetResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSubnetResponseBodyContent struct {
	// Lingjun subnet instance ID
	//
	// example:
	//
	// subnet-yuvn29bn
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
}

func (s CreateSubnetResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s CreateSubnetResponseBodyContent) GoString() string {
	return s.String()
}

func (s *CreateSubnetResponseBodyContent) GetSubnetId() *string {
	return s.SubnetId
}

func (s *CreateSubnetResponseBodyContent) SetSubnetId(v string) *CreateSubnetResponseBodyContent {
	s.SubnetId = &v
	return s
}

func (s *CreateSubnetResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
