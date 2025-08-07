// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAllMessageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAllMessageResponseBody
	GetCode() *string
	SetData(v bool) *DeleteAllMessageResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteAllMessageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAllMessageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAllMessageResponseBody
	GetSuccess() *bool
}

type DeleteAllMessageResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAllMessageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAllMessageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAllMessageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAllMessageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAllMessageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAllMessageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAllMessageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAllMessageResponseBody) SetCode(v string) *DeleteAllMessageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetData(v bool) *DeleteAllMessageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetMessage(v string) *DeleteAllMessageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetRequestId(v string) *DeleteAllMessageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAllMessageResponseBody) SetSuccess(v bool) *DeleteAllMessageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAllMessageResponseBody) Validate() error {
	return dara.Validate(s)
}
