// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelWithApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelDTO) *ModelRouterQueryModelWithApiKeyResponseBody
	GetData() *ModelDTO
	SetErrCode(v string) *ModelRouterQueryModelWithApiKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelWithApiKeyResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelWithApiKeyResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryModelWithApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryModelWithApiKeyResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryModelWithApiKeyResponseBody struct {
	// example:
	//
	// []
	Data *ModelDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterQueryModelWithApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelWithApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetData() *ModelDTO {
	return s.Data
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetData(v *ModelDTO) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetErrCode(v string) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetErrMessage(v string) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetRequestId(v string) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) SetSuccess(v bool) *ModelRouterQueryModelWithApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelWithApiKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
