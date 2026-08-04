// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteContacterResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteContacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteContacterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteContacterResponseBody
	GetSuccess() *bool
}

type DeleteContacterResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteContacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContacterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContacterResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteContacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteContacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContacterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteContacterResponseBody) SetCode(v string) *DeleteContacterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteContacterResponseBody) SetMessage(v string) *DeleteContacterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteContacterResponseBody) SetRequestId(v string) *DeleteContacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContacterResponseBody) SetSuccess(v bool) *DeleteContacterResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteContacterResponseBody) Validate() error {
	return dara.Validate(s)
}
