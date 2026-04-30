// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetClientBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterGetClientBalanceResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterGetClientBalanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetClientBalanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetClientBalanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetClientBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetClientBalanceResponseBody
	GetSuccess() *bool
}

type ModelRouterGetClientBalanceResponseBody struct {
	// example:
	//
	// {}
	Data *ClientBalanceDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterGetClientBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetClientBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetClientBalanceResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterGetClientBalanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetClientBalanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetClientBalanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetClientBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetClientBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetClientBalanceResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterGetClientBalanceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) SetErrCode(v string) *ModelRouterGetClientBalanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) SetErrMessage(v string) *ModelRouterGetClientBalanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetClientBalanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) SetRequestId(v string) *ModelRouterGetClientBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) SetSuccess(v bool) *ModelRouterGetClientBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetClientBalanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
