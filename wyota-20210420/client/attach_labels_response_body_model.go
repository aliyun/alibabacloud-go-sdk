// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttachLabelsResponseBody
	GetCode() *string
	SetMessage(v string) *AttachLabelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachLabelsResponseBody
	GetRequestId() *string
}

type AttachLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachLabelsResponseBody) SetCode(v string) *AttachLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AttachLabelsResponseBody) SetMessage(v string) *AttachLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AttachLabelsResponseBody) SetRequestId(v string) *AttachLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
