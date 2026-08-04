// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncModifyLoginEmailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAsyncModifyLoginEmailResponseBody
	GetCode() *string
	SetMessage(v string) *CancelAsyncModifyLoginEmailResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAsyncModifyLoginEmailResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CancelAsyncModifyLoginEmailResponseBody
	GetSuccess() *string
}

type CancelAsyncModifyLoginEmailResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAsyncModifyLoginEmailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncModifyLoginEmailResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAsyncModifyLoginEmailResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAsyncModifyLoginEmailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAsyncModifyLoginEmailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAsyncModifyLoginEmailResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CancelAsyncModifyLoginEmailResponseBody) SetCode(v string) *CancelAsyncModifyLoginEmailResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponseBody) SetMessage(v string) *CancelAsyncModifyLoginEmailResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponseBody) SetRequestId(v string) *CancelAsyncModifyLoginEmailResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponseBody) SetSuccess(v string) *CancelAsyncModifyLoginEmailResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAsyncModifyLoginEmailResponseBody) Validate() error {
	return dara.Validate(s)
}
