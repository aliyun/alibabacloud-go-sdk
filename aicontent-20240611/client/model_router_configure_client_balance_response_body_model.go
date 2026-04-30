// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterConfigureClientBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterConfigureClientBalanceResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterConfigureClientBalanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterConfigureClientBalanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterConfigureClientBalanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterConfigureClientBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterConfigureClientBalanceResponseBody
	GetSuccess() *bool
}

type ModelRouterConfigureClientBalanceResponseBody struct {
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

func (s ModelRouterConfigureClientBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterConfigureClientBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterConfigureClientBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterConfigureClientBalanceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetErrCode(v string) *ModelRouterConfigureClientBalanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetErrMessage(v string) *ModelRouterConfigureClientBalanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetHttpStatusCode(v int32) *ModelRouterConfigureClientBalanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetRequestId(v string) *ModelRouterConfigureClientBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) SetSuccess(v bool) *ModelRouterConfigureClientBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterConfigureClientBalanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
