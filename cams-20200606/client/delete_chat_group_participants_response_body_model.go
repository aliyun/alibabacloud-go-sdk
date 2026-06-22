// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupParticipantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatGroupParticipantsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatGroupParticipantsResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatGroupParticipantsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatGroupParticipantsResponseBody
	GetRequestId() *string
	SetResult(v int64) *DeleteChatGroupParticipantsResponseBody
	GetResult() *int64
	SetSuccess(v bool) *DeleteChatGroupParticipantsResponseBody
	GetSuccess() *bool
}

type DeleteChatGroupParticipantsResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. Valid values:
	//
	// - OK: The request is successful.
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
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of rows affected.
	//
	// example:
	//
	// 20
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChatGroupParticipantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupParticipantsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupParticipantsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatGroupParticipantsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatGroupParticipantsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatGroupParticipantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatGroupParticipantsResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *DeleteChatGroupParticipantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatGroupParticipantsResponseBody) SetAccessDeniedDetail(v string) *DeleteChatGroupParticipantsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) SetCode(v string) *DeleteChatGroupParticipantsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) SetMessage(v string) *DeleteChatGroupParticipantsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) SetRequestId(v string) *DeleteChatGroupParticipantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) SetResult(v int64) *DeleteChatGroupParticipantsResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) SetSuccess(v bool) *DeleteChatGroupParticipantsResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatGroupParticipantsResponseBody) Validate() error {
	return dara.Validate(s)
}
