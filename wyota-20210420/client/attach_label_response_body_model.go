// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AttachLabelResponseBody
	GetCode() *string
	SetMessage(v string) *AttachLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *AttachLabelResponseBody
	GetRequestId() *string
}

type AttachLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachLabelResponseBody) GoString() string {
	return s.String()
}

func (s *AttachLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *AttachLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AttachLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachLabelResponseBody) SetCode(v string) *AttachLabelResponseBody {
	s.Code = &v
	return s
}

func (s *AttachLabelResponseBody) SetMessage(v string) *AttachLabelResponseBody {
	s.Message = &v
	return s
}

func (s *AttachLabelResponseBody) SetRequestId(v string) *AttachLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
