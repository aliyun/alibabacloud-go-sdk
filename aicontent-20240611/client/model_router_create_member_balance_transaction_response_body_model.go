// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateMemberBalanceTransactionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateMemberBalanceTransactionResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateMemberBalanceTransactionResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *ClientBalanceDTO `json:"data,omitempty" xml:"data,omitempty"`
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

func (s ModelRouterCreateMemberBalanceTransactionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateMemberBalanceTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetErrCode(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetErrMessage(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetRequestId(v string) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) SetSuccess(v bool) *ModelRouterCreateMemberBalanceTransactionResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateMemberBalanceTransactionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
