// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInterruptForArbitrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InterruptForArbitrationResponseBody
	GetCode() *string
	SetMessage(v string) *InterruptForArbitrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *InterruptForArbitrationResponseBody
	GetRequestId() *string
}

type InterruptForArbitrationResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InterruptForArbitrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InterruptForArbitrationResponseBody) GoString() string {
	return s.String()
}

func (s *InterruptForArbitrationResponseBody) GetCode() *string {
	return s.Code
}

func (s *InterruptForArbitrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InterruptForArbitrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InterruptForArbitrationResponseBody) SetCode(v string) *InterruptForArbitrationResponseBody {
	s.Code = &v
	return s
}

func (s *InterruptForArbitrationResponseBody) SetMessage(v string) *InterruptForArbitrationResponseBody {
	s.Message = &v
	return s
}

func (s *InterruptForArbitrationResponseBody) SetRequestId(v string) *InterruptForArbitrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InterruptForArbitrationResponseBody) Validate() error {
	return dara.Validate(s)
}
