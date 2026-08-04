// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBatchCreateMemberApiKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BatchOpResultDTO) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetData() *BatchOpResultDTO
	SetErrCode(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterBatchCreateMemberApiKeysResponseBody
	GetSuccess() *bool
}

type ModelRouterBatchCreateMemberApiKeysResponseBody struct {
	// example:
	//
	// {}
	Data *BatchOpResultDTO `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// Unknown error
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

func (s ModelRouterBatchCreateMemberApiKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBatchCreateMemberApiKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetData() *BatchOpResultDTO {
	return s.Data
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetData(v *BatchOpResultDTO) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetErrCode(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetErrMessage(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetHttpStatusCode(v int32) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetRequestId(v string) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) SetSuccess(v bool) *ModelRouterBatchCreateMemberApiKeysResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterBatchCreateMemberApiKeysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
