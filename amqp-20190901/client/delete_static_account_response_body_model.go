// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStaticAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteStaticAccountResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteStaticAccountResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteStaticAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteStaticAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStaticAccountResponseBody
	GetSuccess() *bool
}

type DeleteStaticAccountResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStaticAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStaticAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStaticAccountResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteStaticAccountResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteStaticAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteStaticAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStaticAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStaticAccountResponseBody) SetCode(v int32) *DeleteStaticAccountResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStaticAccountResponseBody) SetData(v bool) *DeleteStaticAccountResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteStaticAccountResponseBody) SetMessage(v string) *DeleteStaticAccountResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStaticAccountResponseBody) SetRequestId(v string) *DeleteStaticAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStaticAccountResponseBody) SetSuccess(v bool) *DeleteStaticAccountResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStaticAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
