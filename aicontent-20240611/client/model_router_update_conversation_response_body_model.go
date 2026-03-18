// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterUpdateConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConversationDTO) *ModelRouterUpdateConversationResponseBody
	GetData() *ConversationDTO
	SetErrCode(v string) *ModelRouterUpdateConversationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterUpdateConversationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterUpdateConversationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterUpdateConversationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterUpdateConversationResponseBody
	GetSuccess() *bool
}

type ModelRouterUpdateConversationResponseBody struct {
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

func (s ModelRouterUpdateConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterUpdateConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterUpdateConversationResponseBody) GetData() *ConversationDTO {
	return s.Data
}

func (s *ModelRouterUpdateConversationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterUpdateConversationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterUpdateConversationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterUpdateConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterUpdateConversationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterUpdateConversationResponseBody) SetData(v *ConversationDTO) *ModelRouterUpdateConversationResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) SetErrCode(v string) *ModelRouterUpdateConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) SetErrMessage(v string) *ModelRouterUpdateConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) SetHttpStatusCode(v int32) *ModelRouterUpdateConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) SetRequestId(v string) *ModelRouterUpdateConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) SetSuccess(v bool) *ModelRouterUpdateConversationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterUpdateConversationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
