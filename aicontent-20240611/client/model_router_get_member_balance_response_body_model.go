// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterGetMemberBalanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterGetMemberBalanceResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterGetMemberBalanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterGetMemberBalanceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterGetMemberBalanceResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterGetMemberBalanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterGetMemberBalanceResponseBody
	GetSuccess() *bool
}

type ModelRouterGetMemberBalanceResponseBody struct {
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

func (s ModelRouterGetMemberBalanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterGetMemberBalanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterGetMemberBalanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterGetMemberBalanceResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetErrCode(v string) *ModelRouterGetMemberBalanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetErrMessage(v string) *ModelRouterGetMemberBalanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetHttpStatusCode(v int32) *ModelRouterGetMemberBalanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetRequestId(v string) *ModelRouterGetMemberBalanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) SetSuccess(v bool) *ModelRouterGetMemberBalanceResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterGetMemberBalanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
