// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchResetMemberAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchOpResultDTO) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetData() *BatchOpResultDTO
	SetErrCode(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchResetMemberAuthorizationResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchResetMemberAuthorizationResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *BatchOpResultDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The error message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterBatchResetMemberAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchResetMemberAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetData() *BatchOpResultDTO {
	return s.Data
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetData(v *BatchOpResultDTO) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetErrCode(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetErrMessage(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetRequestId(v string) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) SetSuccess(v bool) *ModelRouterBatchResetMemberAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchResetMemberAuthorizationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
