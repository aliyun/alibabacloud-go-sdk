// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpenSubProductStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetOpenSubProductStatusResponseBody
	GetCode() *string
	SetData(v string) *GetOpenSubProductStatusResponseBody
	GetData() *string
	SetMessage(v string) *GetOpenSubProductStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetOpenSubProductStatusResponseBody
	GetRequestId() *string
}

type GetOpenSubProductStatusResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpenSubProductStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpenSubProductStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpenSubProductStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetOpenSubProductStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetOpenSubProductStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetOpenSubProductStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpenSubProductStatusResponseBody) SetCode(v string) *GetOpenSubProductStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetOpenSubProductStatusResponseBody) SetData(v string) *GetOpenSubProductStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetOpenSubProductStatusResponseBody) SetMessage(v string) *GetOpenSubProductStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetOpenSubProductStatusResponseBody) SetRequestId(v string) *GetOpenSubProductStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpenSubProductStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
