// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveContactByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *RemoveContactByIdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *RemoveContactByIdResponseBody
	GetCode() *string
	SetData(v string) *RemoveContactByIdResponseBody
	GetData() *string
	SetMessage(v string) *RemoveContactByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveContactByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveContactByIdResponseBody
	GetSuccess() *bool
}

type RemoveContactByIdResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RemoveContactByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveContactByIdResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveContactByIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *RemoveContactByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveContactByIdResponseBody) GetData() *string {
	return s.Data
}

func (s *RemoveContactByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveContactByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveContactByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveContactByIdResponseBody) SetAccessDeniedDetail(v string) *RemoveContactByIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetCode(v string) *RemoveContactByIdResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetData(v string) *RemoveContactByIdResponseBody {
	s.Data = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetMessage(v string) *RemoveContactByIdResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetRequestId(v string) *RemoveContactByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveContactByIdResponseBody) SetSuccess(v bool) *RemoveContactByIdResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveContactByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
