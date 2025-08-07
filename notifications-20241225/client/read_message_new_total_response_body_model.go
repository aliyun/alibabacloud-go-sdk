// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadMessageNewTotalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReadMessageNewTotalResponseBody
	GetCode() *string
	SetData(v int64) *ReadMessageNewTotalResponseBody
	GetData() *int64
	SetMessage(v string) *ReadMessageNewTotalResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadMessageNewTotalResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadMessageNewTotalResponseBody
	GetSuccess() *bool
}

type ReadMessageNewTotalResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadMessageNewTotalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadMessageNewTotalResponseBody) GoString() string {
	return s.String()
}

func (s *ReadMessageNewTotalResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReadMessageNewTotalResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ReadMessageNewTotalResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadMessageNewTotalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadMessageNewTotalResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadMessageNewTotalResponseBody) SetCode(v string) *ReadMessageNewTotalResponseBody {
	s.Code = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetData(v int64) *ReadMessageNewTotalResponseBody {
	s.Data = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetMessage(v string) *ReadMessageNewTotalResponseBody {
	s.Message = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetRequestId(v string) *ReadMessageNewTotalResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) SetSuccess(v bool) *ReadMessageNewTotalResponseBody {
	s.Success = &v
	return s
}

func (s *ReadMessageNewTotalResponseBody) Validate() error {
	return dara.Validate(s)
}
