// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleClearPretendStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ConsoleClearPretendStatusResponseBody
	GetCode() *int32
	SetData(v bool) *ConsoleClearPretendStatusResponseBody
	GetData() *bool
	SetMessage(v string) *ConsoleClearPretendStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConsoleClearPretendStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConsoleClearPretendStatusResponseBody
	GetSuccess() *bool
}

type ConsoleClearPretendStatusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConsoleClearPretendStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConsoleClearPretendStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleClearPretendStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ConsoleClearPretendStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *ConsoleClearPretendStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConsoleClearPretendStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleClearPretendStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConsoleClearPretendStatusResponseBody) SetCode(v int32) *ConsoleClearPretendStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ConsoleClearPretendStatusResponseBody) SetData(v bool) *ConsoleClearPretendStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ConsoleClearPretendStatusResponseBody) SetMessage(v string) *ConsoleClearPretendStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ConsoleClearPretendStatusResponseBody) SetRequestId(v string) *ConsoleClearPretendStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleClearPretendStatusResponseBody) SetSuccess(v bool) *ConsoleClearPretendStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ConsoleClearPretendStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
