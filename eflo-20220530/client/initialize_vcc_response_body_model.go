// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeVccResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *InitializeVccResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *InitializeVccResponseBody
	GetCode() *int32
	SetContent(v *InitializeVccResponseBodyContent) *InitializeVccResponseBody
	GetContent() *InitializeVccResponseBodyContent
	SetMessage(v string) *InitializeVccResponseBody
	GetMessage() *string
	SetRequestId(v string) *InitializeVccResponseBody
	GetRequestId() *string
}

type InitializeVccResponseBody struct {
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
	Content *InitializeVccResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E30DA7CB-03D0-51EB-8F18-856B99987E18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeVccResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeVccResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeVccResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *InitializeVccResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InitializeVccResponseBody) GetContent() *InitializeVccResponseBodyContent {
	return s.Content
}

func (s *InitializeVccResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InitializeVccResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeVccResponseBody) SetAccessDeniedDetail(v string) *InitializeVccResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *InitializeVccResponseBody) SetCode(v int32) *InitializeVccResponseBody {
	s.Code = &v
	return s
}

func (s *InitializeVccResponseBody) SetContent(v *InitializeVccResponseBodyContent) *InitializeVccResponseBody {
	s.Content = v
	return s
}

func (s *InitializeVccResponseBody) SetMessage(v string) *InitializeVccResponseBody {
	s.Message = &v
	return s
}

func (s *InitializeVccResponseBody) SetRequestId(v string) *InitializeVccResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeVccResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitializeVccResponseBodyContent struct {
	// The request ID.
	//
	// example:
	//
	// E30DA7CB-03D0-51EB-8F18-856B99987E18
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Linked Role of Lingjun Connection Instance (AliyunServiceRoleForEfloVcc)
	//
	// example:
	//
	// CloudConnectionOperationRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s InitializeVccResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s InitializeVccResponseBodyContent) GoString() string {
	return s.String()
}

func (s *InitializeVccResponseBodyContent) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeVccResponseBodyContent) GetRoleName() *string {
	return s.RoleName
}

func (s *InitializeVccResponseBodyContent) SetRequestId(v string) *InitializeVccResponseBodyContent {
	s.RequestId = &v
	return s
}

func (s *InitializeVccResponseBodyContent) SetRoleName(v string) *InitializeVccResponseBodyContent {
	s.RoleName = &v
	return s
}

func (s *InitializeVccResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
