// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateBalanceTransactionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterCreateBalanceTransactionResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterCreateBalanceTransactionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateBalanceTransactionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateBalanceTransactionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateBalanceTransactionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateBalanceTransactionResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateBalanceTransactionResponseBody struct {
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

func (s ModelRouterCreateBalanceTransactionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateBalanceTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterCreateBalanceTransactionResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetErrCode(v string) *ModelRouterCreateBalanceTransactionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetErrMessage(v string) *ModelRouterCreateBalanceTransactionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateBalanceTransactionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetRequestId(v string) *ModelRouterCreateBalanceTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) SetSuccess(v bool) *ModelRouterCreateBalanceTransactionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateBalanceTransactionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
