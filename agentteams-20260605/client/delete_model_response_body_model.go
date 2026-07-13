// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteModelResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteModelResponseBody
	GetSuccess() *bool
}

type DeleteModelResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteModelResponseBody) SetCode(v string) *DeleteModelResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteModelResponseBody) SetMessage(v string) *DeleteModelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteModelResponseBody) SetRequestId(v string) *DeleteModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteModelResponseBody) SetSuccess(v bool) *DeleteModelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteModelResponseBody) Validate() error {
	return dara.Validate(s)
}
