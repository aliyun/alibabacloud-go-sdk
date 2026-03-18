// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConversationDTO) *ModelRouterCreateConversationResponseBody
	GetData() *ConversationDTO
	SetErrCode(v string) *ModelRouterCreateConversationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateConversationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateConversationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateConversationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateConversationResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateConversationResponseBody struct {
	// example:
	//
	// []
	Data *ConversationDTO `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateConversationResponseBody) GetData() *ConversationDTO {
	return s.Data
}

func (s *ModelRouterCreateConversationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateConversationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateConversationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateConversationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateConversationResponseBody) SetData(v *ConversationDTO) *ModelRouterCreateConversationResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) SetErrCode(v string) *ModelRouterCreateConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) SetErrMessage(v string) *ModelRouterCreateConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) SetRequestId(v string) *ModelRouterCreateConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) SetSuccess(v bool) *ModelRouterCreateConversationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateConversationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
