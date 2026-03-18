// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConversationDTO) *ModelRouterQueryConversationResponseBody
	GetData() *ConversationDTO
	SetErrCode(v string) *ModelRouterQueryConversationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryConversationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryConversationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryConversationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryConversationResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryConversationResponseBody struct {
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

func (s ModelRouterQueryConversationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationResponseBody) GetData() *ConversationDTO {
	return s.Data
}

func (s *ModelRouterQueryConversationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryConversationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryConversationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryConversationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryConversationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryConversationResponseBody) SetData(v *ConversationDTO) *ModelRouterQueryConversationResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) SetErrCode(v string) *ModelRouterQueryConversationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) SetErrMessage(v string) *ModelRouterQueryConversationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryConversationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) SetRequestId(v string) *ModelRouterQueryConversationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) SetSuccess(v bool) *ModelRouterQueryConversationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryConversationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
