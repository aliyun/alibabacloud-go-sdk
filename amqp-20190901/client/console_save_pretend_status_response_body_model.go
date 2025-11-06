// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConsoleSavePretendStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ConsoleSavePretendStatusResponseBody
	GetCode() *int32
	SetData(v bool) *ConsoleSavePretendStatusResponseBody
	GetData() *bool
	SetMessage(v string) *ConsoleSavePretendStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConsoleSavePretendStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ConsoleSavePretendStatusResponseBody
	GetSuccess() *bool
}

type ConsoleSavePretendStatusResponseBody struct {
	Code      *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConsoleSavePretendStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConsoleSavePretendStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ConsoleSavePretendStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ConsoleSavePretendStatusResponseBody) GetData() *bool {
	return s.Data
}

func (s *ConsoleSavePretendStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConsoleSavePretendStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConsoleSavePretendStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ConsoleSavePretendStatusResponseBody) SetCode(v int32) *ConsoleSavePretendStatusResponseBody {
	s.Code = &v
	return s
}

func (s *ConsoleSavePretendStatusResponseBody) SetData(v bool) *ConsoleSavePretendStatusResponseBody {
	s.Data = &v
	return s
}

func (s *ConsoleSavePretendStatusResponseBody) SetMessage(v string) *ConsoleSavePretendStatusResponseBody {
	s.Message = &v
	return s
}

func (s *ConsoleSavePretendStatusResponseBody) SetRequestId(v string) *ConsoleSavePretendStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConsoleSavePretendStatusResponseBody) SetSuccess(v bool) *ConsoleSavePretendStatusResponseBody {
	s.Success = &v
	return s
}

func (s *ConsoleSavePretendStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
