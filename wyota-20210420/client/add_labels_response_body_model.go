// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddLabelsResponseBody
	GetCode() *string
	SetMessage(v string) *AddLabelsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddLabelsResponseBody
	GetRequestId() *string
}

type AddLabelsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *AddLabelsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddLabelsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLabelsResponseBody) SetCode(v string) *AddLabelsResponseBody {
	s.Code = &v
	return s
}

func (s *AddLabelsResponseBody) SetMessage(v string) *AddLabelsResponseBody {
	s.Message = &v
	return s
}

func (s *AddLabelsResponseBody) SetRequestId(v string) *AddLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
