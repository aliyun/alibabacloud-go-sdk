// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddChatGroupResponseBody
	GetCode() *string
	SetMessage(v string) *AddChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddChatGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddChatGroupResponseBody
	GetSuccess() *bool
	SetUniqueCode(v string) *AddChatGroupResponseBody
	GetUniqueCode() *string
}

type AddChatGroupResponseBody struct {
	// Details of the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code. Valid values:
	//
	// - `OK`: The request was successful.
	//
	// - For other error codes, see the [error code list](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier for the request.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates if the API call succeeded. Valid values:
	//
	// - **true**: The API call was successful.
	//
	// - **false**: The API call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The unique request code.
	//
	// example:
	//
	// 8938****
	UniqueCode *string `json:"UniqueCode,omitempty" xml:"UniqueCode,omitempty"`
}

func (s AddChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddChatGroupResponseBody) GetUniqueCode() *string {
	return s.UniqueCode
}

func (s *AddChatGroupResponseBody) SetAccessDeniedDetail(v string) *AddChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddChatGroupResponseBody) SetCode(v string) *AddChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *AddChatGroupResponseBody) SetMessage(v string) *AddChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *AddChatGroupResponseBody) SetRequestId(v string) *AddChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddChatGroupResponseBody) SetSuccess(v bool) *AddChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *AddChatGroupResponseBody) SetUniqueCode(v string) *AddChatGroupResponseBody {
	s.UniqueCode = &v
	return s
}

func (s *AddChatGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
