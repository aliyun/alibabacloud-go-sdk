// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApiKeyDecryptedDTO) *ModelRouterQueryApiKeyResponseBody
	GetData() *ApiKeyDecryptedDTO
	SetErrCode(v string) *ModelRouterQueryApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryApiKeyResponseBody struct {
	// example:
	//
	// []
	Data *ApiKeyDecryptedDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyResponseBody) GetData() *ApiKeyDecryptedDTO {
	return s.Data
}

func (s *ModelRouterQueryApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryApiKeyResponseBody) SetData(v *ApiKeyDecryptedDTO) *ModelRouterQueryApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) SetErrCode(v string) *ModelRouterQueryApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) SetErrMessage(v string) *ModelRouterQueryApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) SetRequestId(v string) *ModelRouterQueryApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) SetSuccess(v bool) *ModelRouterQueryApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
