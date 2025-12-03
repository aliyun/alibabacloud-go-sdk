// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessages(v string) *CreateQosPolicyResponseBody
	GetErrorMessages() *string
	SetQosPolicyId(v string) *CreateQosPolicyResponseBody
	GetQosPolicyId() *string
	SetRequestId(v string) *CreateQosPolicyResponseBody
	GetRequestId() *string
}

type CreateQosPolicyResponseBody struct {
	ErrorMessages *string `json:"ErrorMessages,omitempty" xml:"ErrorMessages,omitempty"`
	QosPolicyId   *string `json:"QosPolicyId,omitempty" xml:"QosPolicyId,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQosPolicyResponseBody) GetErrorMessages() *string {
	return s.ErrorMessages
}

func (s *CreateQosPolicyResponseBody) GetQosPolicyId() *string {
	return s.QosPolicyId
}

func (s *CreateQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQosPolicyResponseBody) SetErrorMessages(v string) *CreateQosPolicyResponseBody {
	s.ErrorMessages = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetQosPolicyId(v string) *CreateQosPolicyResponseBody {
	s.QosPolicyId = &v
	return s
}

func (s *CreateQosPolicyResponseBody) SetRequestId(v string) *CreateQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
