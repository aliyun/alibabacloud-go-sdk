// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterTransferToMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ClientBalanceDTO) *ModelRouterTransferToMemberResponseBody
	GetData() *ClientBalanceDTO
	SetErrCode(v string) *ModelRouterTransferToMemberResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterTransferToMemberResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterTransferToMemberResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterTransferToMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterTransferToMemberResponseBody
	GetSuccess() *bool
}

type ModelRouterTransferToMemberResponseBody struct {
	// The data object.
	//
	// example:
	//
	// {}
	Data *ClientBalanceDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The fault code.
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

func (s ModelRouterTransferToMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterTransferToMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterTransferToMemberResponseBody) GetData() *ClientBalanceDTO {
	return s.Data
}

func (s *ModelRouterTransferToMemberResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterTransferToMemberResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterTransferToMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterTransferToMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterTransferToMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterTransferToMemberResponseBody) SetData(v *ClientBalanceDTO) *ModelRouterTransferToMemberResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) SetErrCode(v string) *ModelRouterTransferToMemberResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) SetErrMessage(v string) *ModelRouterTransferToMemberResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) SetHttpStatusCode(v int32) *ModelRouterTransferToMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) SetRequestId(v string) *ModelRouterTransferToMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) SetSuccess(v bool) *ModelRouterTransferToMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterTransferToMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
