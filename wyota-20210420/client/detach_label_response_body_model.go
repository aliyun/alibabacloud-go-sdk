// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachLabelResponseBody
	GetCode() *string
	SetMessage(v string) *DetachLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachLabelResponseBody
	GetRequestId() *string
}

type DetachLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachLabelResponseBody) SetCode(v string) *DetachLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DetachLabelResponseBody) SetMessage(v string) *DetachLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DetachLabelResponseBody) SetRequestId(v string) *DetachLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
