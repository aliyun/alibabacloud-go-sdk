// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteLabelResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteLabelResponseBody
	GetRequestId() *string
}

type DeleteLabelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLabelResponseBody) SetCode(v string) *DeleteLabelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteLabelResponseBody) SetMessage(v string) *DeleteLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteLabelResponseBody) SetRequestId(v string) *DeleteLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
