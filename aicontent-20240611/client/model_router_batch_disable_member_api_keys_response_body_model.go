// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchDisableMemberApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchOpResultDTO) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetData() *BatchOpResultDTO
	SetErrCode(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchDisableMemberApiKeysResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchDisableMemberApiKeysResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *BatchOpResultDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault information code.
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

func (s ModelRouterBatchDisableMemberApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchDisableMemberApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetData() *BatchOpResultDTO {
	return s.Data
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetData(v *BatchOpResultDTO) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetErrCode(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetErrMessage(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetRequestId(v string) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) SetSuccess(v bool) *ModelRouterBatchDisableMemberApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchDisableMemberApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
