// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteApiKeyResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApiKeyResponseBody
	GetRequestId() *string
}

type DeleteApiKeyResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiKeyResponseBody) SetCode(v string) *DeleteApiKeyResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetMessage(v string) *DeleteApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetRequestId(v string) *DeleteApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
