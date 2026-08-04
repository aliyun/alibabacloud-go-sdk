// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAsyncCreateAgAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelAsyncCreateAgAccountResponseBody
	GetCode() *string
	SetMessage(v string) *CancelAsyncCreateAgAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelAsyncCreateAgAccountResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CancelAsyncCreateAgAccountResponseBody
	GetSuccess() *string
}

type CancelAsyncCreateAgAccountResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelAsyncCreateAgAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelAsyncCreateAgAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CancelAsyncCreateAgAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelAsyncCreateAgAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelAsyncCreateAgAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelAsyncCreateAgAccountResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CancelAsyncCreateAgAccountResponseBody) SetCode(v string) *CancelAsyncCreateAgAccountResponseBody {
	s.Code = &v
	return s
}

func (s *CancelAsyncCreateAgAccountResponseBody) SetMessage(v string) *CancelAsyncCreateAgAccountResponseBody {
	s.Message = &v
	return s
}

func (s *CancelAsyncCreateAgAccountResponseBody) SetRequestId(v string) *CancelAsyncCreateAgAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelAsyncCreateAgAccountResponseBody) SetSuccess(v string) *CancelAsyncCreateAgAccountResponseBody {
	s.Success = &v
	return s
}

func (s *CancelAsyncCreateAgAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
