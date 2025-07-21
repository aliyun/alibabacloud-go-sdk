// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetachLabelsResponseBody
	GetCode() *string
	SetMessage(v string) *DetachLabelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetachLabelsResponseBody
	GetRequestId() *string
}

type DetachLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DetachLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetachLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetachLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachLabelsResponseBody) SetCode(v string) *DetachLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *DetachLabelsResponseBody) SetMessage(v string) *DetachLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *DetachLabelsResponseBody) SetRequestId(v string) *DetachLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
