// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchRenewMemberApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchOpResultDTO) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetData() *BatchOpResultDTO
	SetErrCode(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchRenewMemberApiKeysResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchRenewMemberApiKeysResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *BatchOpResultDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault message code.
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterBatchRenewMemberApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchRenewMemberApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetData() *BatchOpResultDTO {
	return s.Data
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetData(v *BatchOpResultDTO) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetErrCode(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetErrMessage(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetRequestId(v string) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) SetSuccess(v bool) *ModelRouterBatchRenewMemberApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchRenewMemberApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
