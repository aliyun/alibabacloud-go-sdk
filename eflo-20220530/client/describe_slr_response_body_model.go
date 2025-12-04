// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeSlrResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DescribeSlrResponseBody
	GetCode() *int32
	SetContent(v *DescribeSlrResponseBodyContent) *DescribeSlrResponseBody
	GetContent() *DescribeSlrResponseBodyContent
	SetMessage(v string) *DescribeSlrResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSlrResponseBody
	GetRequestId() *string
}

type DescribeSlrResponseBody struct {
	// The information about the request denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Content *DescribeSlrResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 9C50C9CD-E799-54DA-BA7A-1FAF3DF80857
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeSlrResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeSlrResponseBody) GetContent() *DescribeSlrResponseBodyContent {
	return s.Content
}

func (s *DescribeSlrResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSlrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlrResponseBody) SetAccessDeniedDetail(v string) *DescribeSlrResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeSlrResponseBody) SetCode(v int32) *DescribeSlrResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSlrResponseBody) SetContent(v *DescribeSlrResponseBodyContent) *DescribeSlrResponseBody {
	s.Content = v
	return s
}

func (s *DescribeSlrResponseBody) SetMessage(v string) *DescribeSlrResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSlrResponseBody) SetRequestId(v string) *DescribeSlrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlrResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSlrResponseBodyContent struct {
	// Whether the role exists
	//
	// example:
	//
	// true
	HasRole *bool `json:"HasRole,omitempty" xml:"HasRole,omitempty"`
}

func (s DescribeSlrResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponseBodyContent) GetHasRole() *bool {
	return s.HasRole
}

func (s *DescribeSlrResponseBodyContent) SetHasRole(v bool) *DescribeSlrResponseBodyContent {
	s.HasRole = &v
	return s
}

func (s *DescribeSlrResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
