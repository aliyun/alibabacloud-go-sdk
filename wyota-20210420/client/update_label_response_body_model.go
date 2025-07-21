// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLabelResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLabelResponseBody
	GetRequestId() *string
}

type UpdateLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLabelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLabelResponseBody) SetCode(v string) *UpdateLabelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLabelResponseBody) SetMessage(v string) *UpdateLabelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLabelResponseBody) SetRequestId(v string) *UpdateLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
