// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ApiKeyDTO) *ModelRouterCreateApiKeyResponseBody
	GetData() *ApiKeyDTO
	SetErrCode(v string) *ModelRouterCreateApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateApiKeyResponseBody struct {
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

func (s ModelRouterCreateApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateApiKeyResponseBody) GetData() *ApiKeyDTO {
	return s.Data
}

func (s *ModelRouterCreateApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateApiKeyResponseBody) SetData(v *ApiKeyDTO) *ModelRouterCreateApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) SetErrCode(v string) *ModelRouterCreateApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) SetErrMessage(v string) *ModelRouterCreateApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) SetRequestId(v string) *ModelRouterCreateApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) SetSuccess(v bool) *ModelRouterCreateApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
