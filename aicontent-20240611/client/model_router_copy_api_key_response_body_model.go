// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCopyApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApiKeyDTO) *ModelRouterCopyApiKeyResponseBody
	GetData() *ApiKeyDTO
	SetErrCode(v string) *ModelRouterCopyApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCopyApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCopyApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCopyApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCopyApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterCopyApiKeyResponseBody struct {
	// example:
	//
	// []
	Data *ApiKeyDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterCopyApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCopyApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCopyApiKeyResponseBody) GetData() *ApiKeyDTO {
	return s.Data
}

func (s *ModelRouterCopyApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCopyApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCopyApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCopyApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCopyApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCopyApiKeyResponseBody) SetData(v *ApiKeyDTO) *ModelRouterCopyApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) SetErrCode(v string) *ModelRouterCopyApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) SetErrMessage(v string) *ModelRouterCopyApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterCopyApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) SetRequestId(v string) *ModelRouterCopyApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) SetSuccess(v bool) *ModelRouterCopyApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCopyApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
